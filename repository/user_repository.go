package repository

import (
	"context"
	"go-mongo-docker/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetOwnProjects(string) (*entity.User, error)
	CreateNewUser(*entity.User) (*entity.User, error)
	UpdateProjectMembers(string, string) (*entity.User, error)
}

type userRepository struct {
	db *mongo.Client
}

// NewProjectRepository returns "ProjectRepository"
func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

/**
 * @summary: gets projects by user id
 * @return : projects, error
 */
func (ur *userRepository) GetOwnProjects(id string) (*entity.User, error) {

	// get user
	userCollection := ur.db.Database("taski").Collection("users")
	userFindResult := userCollection.FindOne(context.Background(), bson.M{"id": id})

	var user *entity.User
	err := userFindResult.Decode(&user)
	avoidPanic(err)

	// get other users
	otherUsers, err := userCollection.Find(context.Background(), bson.D{})
	avoidPanic(err)

	var projectMembers []*entity.User
	for otherUsers.Next(context.Background()) {
		var projectMember *entity.User
		err := otherUsers.Decode(&projectMember)
		avoidPanic(err)
		projectMembers = append(projectMembers, projectMember)
	}

	// get projects
	projectCollection := ur.db.Database("taski").Collection("projects")
	projectFilter := options.Find()
	projectFilter.SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	projectFindResult, err := projectCollection.Find(context.Background(), bson.D{}, projectFilter)
	avoidPanic(err)

	var projects []*entity.Project
	for projectFindResult.Next(context.Background()) {
		var project *entity.Project
		err := projectFindResult.Decode(&project)
		avoidPanic(err)
		projects = append(projects, project)
	}

	// get todos
	todoCollection := ur.db.Database("taski").Collection("todos")
	todoFindResult, err := todoCollection.Find(context.Background(), bson.D{})
	avoidPanic(err)

	var todos []*entity.Todo
	for todoFindResult.Next(context.Background()) {
		var todo *entity.Todo
		err := todoFindResult.Decode(&todo)
		avoidPanic(err)
		todos = append(todos, todo)
	}

	// include todos in project
	for _, project := range projects {
		for _, todo := range todos {
			if project.Id == todo.ProjectId {
				project.Todos = append(project.Todos, *todo)
			}
		}
	}

	// include project members in project
	for _, project := range projects {
		for _, projectMemberId := range project.UserIds {
			for _, member := range projectMembers {
				if projectMemberId == member.Id {
					project.User = append(project.User, *member)
				}
			}
		}
	}

	// include projects in user
	for _, project := range projects {
		for _, userId := range project.UserIds {
			if user.Id == userId {
				user.Project = append(user.Project, *project)
			}
		}
	}

	return user, nil
}

func (ur *userRepository) CreateNewUser(user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := ur.db.Database("taski").Collection("users")

	var userProject [0]primitive.ObjectID
	insert := bson.D{
		{Key: "project", Value: userProject},
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "image", Value: user.Image},
	}

	// duplication validation
	validationResult := collection.FindOne(ctx, insert)

	var duplicatedUser *entity.User
	decodedValidationResult := validationResult.Decode(&duplicatedUser)

	if decodedValidationResult != nil {
		// user insertion
		incompleteInsertion, err := collection.InsertOne(ctx, insert)
		avoidPanic(err)

		autoIncrementedId := incompleteInsertion.InsertedID.(primitive.ObjectID).Hex()
		filter := bson.M{"_id": convertToObjectId(autoIncrementedId)}
		update := bson.M{
			"$set": bson.M{
				"id": autoIncrementedId,
			}}

		_, err = collection.UpdateOne(ctx, filter, update)
		avoidPanic(err)

		updateResult := collection.FindOne(ctx, bson.M{"id": autoIncrementedId})

		var newUser *entity.User
		decodedUpdateResult := updateResult.Decode(&newUser)
		avoidPanic(decodedUpdateResult)

		return newUser, nil
	} else {
		return duplicatedUser, nil
	}
}

func (ur *userRepository) UpdateProjectMembers(projectId string, userEmail string) (*entity.User, error) {
	// get project
	projectCollection := ur.db.Database("taski").Collection("projects")
	projectFindResult := projectCollection.FindOne(context.Background(), bson.M{"id": projectId})

	var project *entity.Project
	projectErr := projectFindResult.Decode(&project)
	avoidPanic(projectErr)

	// get user
	userCollection := ur.db.Database("taski").Collection("users")
	userFindResult := userCollection.FindOne(context.Background(), bson.M{"email": userEmail})

	var user *entity.User
	userErr := userFindResult.Decode(&user)
	if userErr != nil {
		return nil, bson.ErrDecodeToNil
	}

	// duplication validation
	for _, eachUser := range project.UserIds {
		if user.Id == eachUser {
			return user, nil
		}
	}

	// update project member
	pushFileter := bson.M{"id": projectId}
	push := bson.M{"$push": bson.M{"userIds": user.Id}}
	update := bson.M{"$set": bson.M{"updatedAt": time.Now()}}

	_, pushErr := projectCollection.UpdateOne(context.Background(), pushFileter, push)
	avoidPanic(pushErr)

	_, updateErr := projectCollection.UpdateOne(context.Background(), pushFileter, update)
	avoidPanic(updateErr)

	return user, nil
}

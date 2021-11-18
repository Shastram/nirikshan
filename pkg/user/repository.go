package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/utils"
)

// Repository holds the mongo database implementation of the Service
type Repository interface {
	GetUser(uid string) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	FindUsersByUID(uid []string) (*[]entities.User, error)
	FindUserByUsername(username string) (*entities.User, error)
	CheckPasswordHash(hash, password string) error
	UpdatePassword(userPassword *entities.UserPassword, isAdminBeingReset bool, ctx context.Context) error
	CreateUser(user *entities.User, ctx context.Context) error
	UpdateUser(user *entities.User, ctx context.Context) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) GetUser(uid string) (*entities.User, error) {
	var result = entities.User{}
	findOneErr := r.Collection.FindOne(context.TODO(), bson.M{
		"_id": uid,
	}).Decode(&result)

	if findOneErr != nil {
		return nil, findOneErr
	}
	return &(*result.SanitizedUser()), nil
}

// GetUsers fetches all the users from the database
func (r repository) GetUsers() (*[]entities.User, error) {
	var Users []entities.User
	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user entities.User
		_ = cursor.Decode(&user)
		Users = append(Users, *user.SanitizedUser())
	}
	return &Users, nil
}

// FindUsersByUID fetches the user from database that matches the passed uids
func (r repository) FindUsersByUID(uid []string) (*[]entities.User, error) {
	cursor, err := r.Collection.Find(context.Background(),
		bson.D{
			{"_id", bson.D{
				{"$in", uid},
			}},
		})

	if err != nil {
		return nil, err
	}

	var Users []entities.User
	for cursor.Next(context.TODO()) {
		var user entities.User
		_ = cursor.Decode(&user)
		Users = append(Users, *user.SanitizedUser())
	}

	return &Users, nil
}

// FindUserByUsername finds and returns a user if it exists
func (r repository) FindUserByUsername(username string) (*entities.User, error) {
	var result = entities.User{}
	findOneErr := r.Collection.FindOne(context.TODO(), bson.M{
		"username": username,
	}).Decode(&result)

	if findOneErr != nil {
		return nil, findOneErr
	}
	return &result, nil
}

func (r repository) FindUserByOauthID(oauthID string) (*entities.User, error) {
	var user entities.User
	err := r.Collection.FindOne(context.TODO(), bson.M{
		"oauth_id": oauthID,
	}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckPasswordHash checks password hash and password from user input
func (r repository) CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return utils.ErrWrongPassword
	}

	return nil
}

// UpdatePassword helps to update the password of the user, it acts as a resetPassword when isAdminBeingReset is set to true
func (r repository) UpdatePassword(userPassword *entities.UserPassword, isAdminBeingReset bool, ctx context.Context) error {
	var result = entities.User{}
	result.UserName = userPassword.Username
	findOneErr := r.Collection.FindOne(ctx, bson.M{
		"username": result.UserName,
	}).Decode(&result)
	if findOneErr != nil {
		return findOneErr
	}
	if isAdminBeingReset {
		err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userPassword.OldPassword))
		if err != nil {
			return err
		}
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword.NewPassword), utils.PasswordEncryptionCost)
	_, err = r.Collection.UpdateOne(context.Background(), bson.M{"_id": result.ID}, bson.M{"$set": bson.M{
		"password": string(newHashedPassword),
	}})
	if err != nil {
		return err
	}

	return nil
}

// CreateUser creates a new user in the database
func (r repository) CreateUser(user *entities.User, ctx context.Context) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.ErrUserExists
		}
		return err
	}

	return nil
}

// UpdateUser updates user details in the database
func (r repository) UpdateUser(user *entities.User, ctx context.Context) error {
	//TODO: update user
	//_, err := r.Collection.UpdateOne(ctx,)

	//if err != nil {
	//	return err
	//}

	return nil
}

// NewRepo creates a new instance of this repository
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

package services

import (
	"context"
	"errors"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(u *mongo.Collection, ctx context.Context) UserServices {
	return &UserServiceImpl{
		usercollection: u,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(a *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, a)
	return err
}

func (u *UserServiceImpl) CreateUsers(users *[]models.User) error {

	var AllUsers []interface{}
	for _, user := range *users {
		AllUsers = append(AllUsers, user)
	}

	_, err := u.usercollection.InsertMany(u.ctx, AllUsers)

	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {

	var NewUser models.User
	filter := bson.M{"user_name": name}

	err := u.usercollection.FindOne(u.ctx, filter).Decode(&NewUser)
	if err != nil {
		return nil, err
	}
	return &NewUser, nil
}

func (u *UserServiceImpl) GetUsers() (*[]models.User, error) {

	filter := bson.M{}

	cursor, err := u.usercollection.Find(u.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)

	var alluser []models.User
	var a models.User
	for cursor.Next(u.ctx) {

		if err := cursor.Decode(&a); err != nil {
			return nil, err
		} else {
			alluser = append(alluser, a)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &alluser, nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {

	filter := bson.M{"user_name": *name}
	result, err := u.usercollection.DeleteOne(u.ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no matched document for update")
	}
	return nil
}

func (u *UserServiceImpl) UpdateUser(a *models.User) error {
	filter := bson.M{"user_name": a.Name}
	update := bson.M{
		"$set": bson.M{
			"user_name":    a.Name,
			"user_age":     a.Age,
			"user_address": a.Address,
		},
	}

	err, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if err.MatchedCount != 1 {
		return errors.New("no matched document for update")
	}
	return nil
}

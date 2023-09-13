package services

import (
	"context"
	"errors"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseServiceImpl struct {
	coursecollection *mongo.Collection
	ctx              context.Context
}

func NewCourseService(u *mongo.Collection, ctx context.Context) CourseService {
	return &CourseServiceImpl{
		coursecollection: u,
		ctx:              ctx,
	}
}

func (u *CourseServiceImpl) GetCourses() (*[]models.Course, error) {

	filter := bson.M{}

	cursor, err := u.coursecollection.Find(u.ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)

	var allcourses []models.Course
	var course models.Course
	for cursor.Next(u.ctx) {

		if err := cursor.Decode(&course); err != nil {
			return nil, err
		} else {
			allcourses = append(allcourses, course)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &allcourses, nil
}

func (u *CourseServiceImpl) CreateCourse(newCourse *models.Course) error {

	_, err := u.coursecollection.InsertOne(u.ctx, newCourse)
	return err
}

func (u *CourseServiceImpl) UpdateCourse(id *string, a *models.Course) error {

	filter := bson.M{"code": a.Code}
	update := bson.M{
		"$set": bson.M{
			"code":        a.Code,
			"course_name": a.Name,
			"credithours": a.CreditHours,
			"status":      a.Status,
		},
	}

	n, _ := u.coursecollection.UpdateOne(u.ctx, filter, update)
	if n.MatchedCount != 1 {
		return errors.New("no matched document for update")
	}
	return nil
}

func (u *CourseServiceImpl) DeleteCourse(id *string) error {

	course, err := u.GetCourse(id)

	if err != nil {
		return err
	}

	filter := bson.M{"code": course.Code}

	n, _ := u.coursecollection.DeleteOne(u.ctx, filter)
	if n.DeletedCount != 1 {
		return errors.New("no matched document for update")
	}
	return nil
}

func (u *CourseServiceImpl) AddCourses(courses *[]models.Course) error {
	var allCourses []interface{}
	for _, course := range *courses {
		allCourses = append(allCourses, course)
	}

	_, err := u.coursecollection.InsertMany(u.ctx, allCourses)

	return err
}

func (u *CourseServiceImpl) GetCourse(id *string) (*models.Course, error) {

	var course models.Course
	filter := bson.M{"code": id}

	if err := u.coursecollection.FindOne(u.ctx, filter).Decode(&course); err != nil {
		return nil, err
	}

	return &course, nil
}

package models

type Course struct {
	Code        string `json:"code" bson:"code"`
	Name        string `json:"name" bson:"course_name"`
	CreditHours int    `json:"credithours" bson:"credithours"`
	Status      string `json:"status" bson:"status"`
}

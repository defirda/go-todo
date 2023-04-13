package models

import (
	"time"
)

type Activity struct {
	activity_id int       `bson:"activity_id"`
	title       string    `bson:"title"`
	email       string    `bson:"email"`
	created_at  time.Time `bson:"created_at"`
	updated_at  time.Time `bson:"updated_at"`
}

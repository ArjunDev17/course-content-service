package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Resource represents attachment inside a lesson (video, doc, quiz metadata)
type Resource struct {
	Type     string `bson:"type" json:"type"`           // e.g., "video", "pdf", "quiz"
	URL      string `bson:"url" json:"url,omitempty"`   // cloud storage URL
	Meta     any    `bson:"meta,omitempty" json:"meta"` // free-form metadata like duration
	FileName string `bson:"file_name,omitempty" json:"file_name,omitempty"`
}

// Lesson inside a module
type Lesson struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Resources   []Resource         `bson:"resources,omitempty" json:"resources,omitempty"`
	DurationMin int                `bson:"duration_min,omitempty" json:"duration_min,omitempty"`
	Order       int                `bson:"order,omitempty" json:"order,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Module contains lessons
type Module struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Lessons     []Lesson           `bson:"lessons,omitempty" json:"lessons,omitempty"`
	Order       int                `bson:"order,omitempty" json:"order,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Course top-level
type Course struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	Category     string             `bson:"category,omitempty" json:"category,omitempty"`
	Level        string             `bson:"level,omitempty" json:"level,omitempty"` // beginner, intermediate, advanced
	Price        float64            `bson:"price,omitempty" json:"price,omitempty"`
	InstructorID primitive.ObjectID `bson:"instructor_id,omitempty" json:"instructor_id,omitempty"`
	Tags         []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	Modules      []Module           `bson:"modules,omitempty" json:"modules,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

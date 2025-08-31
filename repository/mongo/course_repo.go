package mongo_repo

import (
	"context"
	"errors"
	"time"

	"github.com/ArjunDev17/course-content-service/model"
	"github.com/ArjunDev17/course-content-service/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository interface {
	Create(ctx context.Context, c *model.Course) (*model.Course, error)
	GetByID(ctx context.Context, id string) (*model.Course, error)
	GetAll(ctx context.Context, filter map[string]interface{}, page, limit int64) ([]*model.Course, int64, error)
	Update(ctx context.Context, id string, update map[string]interface{}) (*model.Course, error)
	Delete(ctx context.Context, id string) error
}

type courseRepo struct {
	collectionName string
}

func NewCourseRepository() CourseRepository {
	return &courseRepo{
		collectionName: db.CoursesCollection().Name(),
	}
}

func (r *courseRepo) coll() *mongo.Collection {
	return db.Client.Database(db.Client.Database("").Name()).Collection(r.collectionName)
}

// Create inserts a new course
func (r *courseRepo) Create(ctx context.Context, c *model.Course) (*model.Course, error) {
	c.CreatedAt = time.Now().UTC()
	c.UpdatedAt = c.CreatedAt
	res, err := db.CoursesCollection().InsertOne(ctx, c)
	if err != nil {
		return nil, err
	}
	c.ID = res.InsertedID.(primitive.ObjectID)
	return c, nil
}

// GetByID returns a course
func (r *courseRepo) GetByID(ctx context.Context, id string) (*model.Course, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var course model.Course
	if err := db.CoursesCollection().FindOne(ctx, bson.M{"_id": oid}).Decode(&course); err != nil {
		return nil, err
	}
	return &course, nil
}

// GetAll with filters, pagination
func (r *courseRepo) GetAll(ctx context.Context, filter map[string]interface{}, page, limit int64) ([]*model.Course, int64, error) {
	bfilter := bson.M{}
	for k, v := range filter {
		bfilter[k] = v
	}
	findOpts := options.Find()
	if limit <= 0 {
		limit = 20
	}
	findOpts.SetLimit(limit)
	if page > 0 {
		findOpts.SetSkip((page - 1) * limit)
	}
	cur, err := db.CoursesCollection().Find(ctx, bfilter, findOpts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []*model.Course
	for cur.Next(ctx) {
		var c model.Course
		if err := cur.Decode(&c); err != nil {
			return nil, 0, err
		}
		out = append(out, &c)
	}

	total, err := db.CoursesCollection().CountDocuments(ctx, bfilter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

// Update partial update
func (r *courseRepo) Update(ctx context.Context, id string, update map[string]interface{}) (*model.Course, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update["updated_at"] = time.Now().UTC()
	updateDoc := bson.M{"$set": update}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated model.Course
	if err := db.CoursesCollection().FindOneAndUpdate(ctx, bson.M{"_id": oid}, updateDoc, opts).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete
func (r *courseRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := db.CoursesCollection().DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return nil
}

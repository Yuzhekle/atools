package main

import (
	"context"
	"fmt"
	"time"

	// "go.mongodb.org/mongo-driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Person struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string
	Age    int
	Adress string
	Date   string
}

func main() {
	repo := NewRepo()

	// data := &Person{
	// 	Name:   "Nguyen Van A",
	// 	Age:    20,
	// 	Adress: "Ha Noi",
	// 	Date:   "2024-01-08",
	// }
	// err := repo.insert(data)
	// if err != nil {
	// 	panic(err)
	// }

	repo.find()
	repo.findOne()
}

type Repo struct {
	conn *mongo.Collection
}

func NewRepo() *Repo {
	return &Repo{
		conn: conn(),
	}
}

func (r *Repo) insert(data interface{}) (err error) {
	ctx := context.Background()
	_, err = r.conn.InsertOne(ctx, data)
	return
}

func (r *Repo) update() (err error) {
	ctx := context.Background()
	_, err = r.conn.UpdateOne(ctx, bson.M{"name": "Nguyen Van A"}, bson.M{"$set": bson.M{"name": "Nguyen Van B"}})
	return
}

func (r *Repo) findOne() (err error) {
	ctx := context.Background()
	result := Person{}
	err = r.conn.FindOne(ctx, bson.M{"date": "2024-01-28"}).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("ffff", result)
	return
}

func (r *Repo) find() (err error) {
	ctx := context.Background()
	cursor, err := r.conn.Find(ctx, bson.M{
		"date": bson.M{
			"$gte": "2024-01-02",
			"$lte": "2024-01-05",
		},
	})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result Person
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	return
}

func conn() *mongo.Collection {
	connStr := "mongodb://root:123456@localhost:27017/tms?authSource=admin"
	clientOptions := options.Client().ApplyURI(connStr)
	clientOptions.SetReadPreference(readpref.SecondaryPreferred())
	clientOptions.SetMaxConnIdleTime(10 * time.Second)
	clientOptions.SetMaxPoolSize(500)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	return client.Database("tms").Collection("Agoda_Roomtype_Mapping")
}

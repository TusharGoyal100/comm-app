package connector

import (
	"communication-app/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DbName string = "Communication-App"
const Collection string = "Message"

var collection *mongo.Collection

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(models.DbURL))
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected with database")
	collection = client.Database(DbName).Collection(Collection)
}

func InsertOne(msg interface{}) error {
	res, err := collection.InsertOne(context.Background(), msg)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted msg with ID:-", res.InsertedID)
	return nil
}

func GetOne(msgID string) (models.Message, error) {
	var msg models.Message
	msgIDHex, _ := primitive.ObjectIDFromHex(msgID)
	filter := bson.M{"_id": msgIDHex}

	result := collection.FindOne(context.Background(), filter)
	err := result.Decode(&msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func GetAll() ([]models.Message, error) {
	var result []models.Message

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var msg models.Message
		err := cur.Decode(&msg)
		if err != nil {
			return nil, err
		}
		result = append(result, msg)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteOne(msgID string) error {
	msgIDHex, _ := primitive.ObjectIDFromHex(msgID)
	filter := bson.M{"_id": msgIDHex}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Succesfully deleted no. of msg :-", result.DeletedCount)
	return nil
}

func DeleteAll() error {

	result, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	fmt.Println("Succesfully deleted no. of msg :-", result.DeletedCount)
	return nil
}

func UpdateOne(msgID string, newMessage interface{}) error {
	msgIDHex, _ := primitive.ObjectIDFromHex(msgID)

	data, err := bson.Marshal(newMessage)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": msgIDHex}
	var temp bson.D
	err = bson.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	updateFilter := bson.D{{"$set", temp}}
	res, err := collection.UpdateOne(context.Background(), filter, updateFilter)
	if err != nil {
		return err
	}

	fmt.Println("Succesfully updated msg;-", res.MatchedCount, res.ModifiedCount, res.UpsertedCount, res.UpsertedID)
	return nil
}

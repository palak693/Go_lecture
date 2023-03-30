package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Setting the MongoDB Compass connection which is at host and port localhost:27017
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connecting to the MongoDB database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	/* checking for the connection, and if there is no error identified, meaning the error is nil
	then the connection is established successfuly */
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	/* From my database "Palak's Database", I will select the "People" collection.
	This collection contains the Name, Profession, Hobby, Location, Profession and Sno or Serial num of the Person */
	coll := client.Database("mydb").Collection("People")

	// Declaring context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)

	// Before using selectors, I will first display all the documents from the collection:
	// Since I am displaying all the documents, I will use Find() method, but if I wanted to display
	//only one, then I would have used FindOne() method

	//setting an iterator that will go through all the docs and store them
	// and a variable err to store any errors

	cursor, err := coll.Find(context.TODO(), bson.D{}) // finding bson docs
	if err != nil {
		log.Fatal(err)
	} //close the program if error exists
	defer cursor.Close(context.TODO())

	//iterating the iterator to the next document until all docs are over
	for cursor.Next(context.TODO()) {
		var result bson.M //variable result of type bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result) //print the result
	}

	// Query 1: Find the document where Name: Reet Bhatia

	//again setting an iterator to store vals and err to store errors
	//setting the selector where name is specified as Reet Bhatia
	cursor, err = coll.Find(context.TODO(), bson.M{"Name": "Reet Bhatia"})
	if err != nil {
		log.Fatal(err)
	} //close program if error detected
	defer cursor.Close(context.TODO())

	//go to next iteration to check for more Reet Bhatia name entries
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// Query 2: Find documents where Location is UK and Hobby is Hiking
	cursor, err = coll.Find(context.TODO(), bson.D{"Location": "UK", "Hobby": "Hiking"})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// Query 3: Count total number of documents in the collection
	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total documents: %v\n", count)

	// Query 4: Find Sno = 4, ID = ObjectID(641de04e908f9f41e5e572cf)
	objid, _ := primitive.ObjectIDFromHex("641de04e908f9f41e5e572cf")
	cursor, err := coll.Find(context.TODO(), bson.D{"Sno": 4, "_id": objid})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// Query 5:
	// Query 5. Find the document where Profession is "Software Developer" and Location is NOT "India"
	cursor, err = coll.Find(context.TODO(), bson.M{"Profession": "Software Developer", "Location": bson.M{"$ne": "India"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	// Additional: Document where Profession is AI Rearcher or Hobby is Podcasting
	//storing the query into a variable can also work the same way as directly providing the condition in the Find()
	// selque := bson.M{"$or":[]bson.M{{"Profession":"AI Researcher"}, {"Hobby":"Podcasting"}}}
	cursors, err := coll.Find(context.TODO(), bson.M{"$or": []bson.M{{"Profession": "AI Researcher"},
		{"Hobby": "Podcasting"}}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursors.Close(context.TODO())

	for cursors.Next(context.TODO()) {
		var result bson.M
		err := cursors.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

}

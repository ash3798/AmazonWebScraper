package task

import (
	"context"
	"log"

	"github.com/ash3798/AmazonWebScraper/persistence/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Name         string
	ImageURL     string
	Description  string
	Price        string
	TotalReviews int
}

type ProductInfo struct {
	Url     string
	Product Product
}

var (
	MongoDB *mongo.Client
)

func InitDatabaseClient() bool {
	client, ok := connect()
	if !ok {
		return false
	}

	MongoDB = client
	return true
}

func connect() (*mongo.Client, bool) {
	mongoUri := config.GetMongoURL()
	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting to db , Error: ", err.Error())
		return nil, false
	}

	//check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("error not connected to db, Error : ", err.Error())
		return nil, false
	}

	log.Println("Connected to MongoDB")
	return client, true
}

func PersistDataToDB(productInfo ProductInfo) error {
	collection := MongoDB.Database(config.Manager.MongoDBName).Collection(config.Manager.MongoCollectionName)
	res, err := collection.InsertOne(context.TODO(), productInfo)
	if err != nil {
		log.Println("error while inserting the record , Error :", err.Error())
		return err
	}

	log.Println("successfully inserted single document: ", res.InsertedID)
	return nil
}

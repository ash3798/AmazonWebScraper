package task

import (
	"context"
	"log"
	"time"

	"github.com/ash3798/AmazonWebScraper/persistence/config"
	"go.mongodb.org/mongo-driver/bson"
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
	Url        string
	Product    Product
	ScrapeTime time.Time
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

	productInfo.ScrapeTime = time.Now()
	log.Printf("product info , %+v", productInfo)
	if !updateIfPresent(collection, productInfo) {
		return InsertToDB(collection, productInfo)
	}
	return nil
}

func updateIfPresent(collection *mongo.Collection, productInfo ProductInfo) bool {
	res := collection.FindOneAndReplace(
		context.TODO(),
		bson.M{"url": productInfo.Url},
		productInfo,
	)

	if res.Err() != nil {
		log.Println("Could not find the document present already with this url in store")
		return false
	}
	log.Println("Successfully updated record with newer info in database")
	return true
}

func InsertToDB(collection *mongo.Collection, productInfo ProductInfo) error {
	res, err := collection.InsertOne(context.TODO(), productInfo)
	if err != nil {
		log.Println("error while inserting the record , Error :", err.Error())
		return err
	}

	log.Println("Successfully inserted single document: ", res.InsertedID)
	return nil
}

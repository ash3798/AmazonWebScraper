# AMAZON-PRODUCT-SCRAPER
Amazon-product-scraper is a Golang based scraper making use of Colly to scrape the information about product from Amazon's product page.

## Backend
Golang , Gin

## Database
MongoDB

## Prerequisites
  * Docker must be installed in your setup.
  * Docker compose should be installed.

## Getting Started
Once prerequisites are set up, follow below steps to get started

1. Clone the repository to your local system
```bash
$ git clone https://github.com/ash3798/AmazonWebScraper.git
```
3. go to the root directory of repository, you will see a file named docker-compose.yml
```bash
$ ls
docker-compose.yml  go.mod  go.sum  persistence/  README.md  scraper/
```
5. use docker compose command to bring up the containers
```bash
$ docker compose up
```
> this step will take some time to build/pull images and bring up the application with all the containers
6. Verify the containers are up and running using either docker command or docker UI
```bash
$ docker compose ps
NAME                             SERVICE             STATUS              PORTS
amazonwebscraper_mongo_1         mongo               running             0.0.0.0:27017->27017/tcp, :::27017->27017/tcp
amazonwebscraper_persistence_1   persistence         running             0.0.0.0:9092->9092/tcp, :::9092->9092/tcp    
amazonwebscraper_scraper_1       scraper             running             0.0.0.0:9091->9091/tcp, :::9091->9091/tcp    
```
7. Once everything is up , you can start using the application.

## API Usage Guide

1. Send Scrape request
```bash
POST  /url/scrape

Body :   
{
  "url" : "https://www.amazon.com/Sony-Alpha-a6400-Mirrorless-Camera/dp/B07MTWVN3M/ref=sr_1_1?dchild=1&keywords=a6400&qid=1627662358&sr=8-1"
}

Response :
{
    "Url": "https://www.amazon.com/Sony-Alpha-a6400-Mirrorless-Camera/dp/B07MTWVN3M/ref=sr_1_1?dchild=1&keywords=a6400&qid=1627662358&sr=8-1",
    "Product": {
        "Name": "Sony Alpha a6400 Mirrorless Camera: Compact APS-C Interchangeable Lens Digital Camera with Real-Time Eye Auto Focus, 4K Video & Flip Up",
        "ImageURL": "https://images-na.ssl-images-amazon.com/images/I/41-P7hZaf6L.__AC_SY300_SX300_QL70_ML2_.jpg",
        "Description": "Next Gen speed: experience the world’s fastest 0. 02 sec AF with real-time AF and object tracking.Enhanced subject capture: wide 425 Phase/ 425 contrast detection points over 84% of the sensor",
        "Price": "$898.00",
        "TotalReviews": 1084
    }
}
```
  > * Api has been made syncronous as if now for easy accessibility of scrape results. Thus upon request it will scrape, prepare result and respond back with result.
  > * For faster responses to api request, scraping can also be made asyncronous and OK status can be sent back to user on successful receive. 
  > * Once server is done with scraping, it will send the scrape results to the database (done asyncronously).

## Visualizing Data
Data can be seen by in MongoDB. You can access Mongo using two ways
 1. You can then access your mongoDB using "MongoDB Compass UI" to see the records inserted to it.
    >Mongo DB can be accessed at "localhost:27017" since port has been exposed by container

 2. You can access MongoDB by entering into the mongo container
    * exec into the mongo container
      ```bash
      docker exec -it amazonwebscraper_mongo_1 bash
      ```
    * enter the mongo shell
      ```bash
      $ mongosh
      Current Mongosh Log ID: 610549146c95e8fe2bb0622c
      Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000
      Using MongoDB:          5.0.1
      Using Mongosh:          1.0.1
      ```
    * show all DB's
      ```bash
      test> show dbs
      admin     41 kB
      config   111 kB
      local     41 kB
      product   41 kB
      test> use product
      switched to db product
      product>
      ```
      > You will see a "product" DB would have been created by our app. Documents will be stored in this DB
    * list collections
      ```bash
      product> show collections
      shopitems
      ```
      > shopitems collections is created by our app
    * List out the data in the collections, you will find the data of the product you just scraped for.
      ```bash
      product> db.shopitems.find().pretty()
      [
        {
          _id: ObjectId("610548929abf18b9c2b5e57c"),
          url: 'https://www.amazon.com/Sony-Alpha-a6400-Mirrorless-Camera/dp/B07MTWVN3M/ref=sr_1_1?dchild=1&keywords=a6400&qid=1627662358&sr=8-1',
          product: {
            name: 'Sony Alpha a6400 Mirrorless Camera: Compact APS-C Interchangeable Lens Digital Camera with Real-Time Eye Auto Focus, 4K Video',
            imageurl: 'https://images-na.ssl-images-amazon.com/images/I/41-P7hZaf6L.__AC_SY300_SX300_QL70_ML2_.jpg',
            description: 'Next Gen speed: experience the world’s fastest 0. 02 sec AF with real-time AF and object tracking.Enhanced subject capture: wide 425 Phase/ 425 contrast detection points over 84% of the sensor',
            price: '$898.00',
            totalreviews: 1083
          }
        }
      ]
      ```

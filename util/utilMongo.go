package util

import (
	"context"
	"errors"
	"github.com/akshay1192/restaurantBooking"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"strings"
	"time"
)

type available struct {
	OccupiedAt   time.Time
	OccupiedTill time.Time
}

type table struct {
	MinCapacity   int `bson:"min_occupency"`
	MaxCapacity   int `bson:"max_seating"`
	NextAvailable []available
}

type wait struct {
	MinPartySize int    `bson:"min_party_size"`
	MaxPartySize int    `bson:"max_party_size"`
	AvgTot       string `bson:"avg_tot"`
}

// Tables will contain the current table setup details
var Tables []table

// MongoClient is the mongo client handle for queries
var MongoClient *mongo.Client

// InitialiseMongoClient initialises mongoClient and populates Tables
func InitialiseMongoClient() error {

	var err error

	// creating context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// creating mongo client and then connecting the client to server
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(restaurantBooking.mongoHostURL))

	// checking the connectivity
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	err = PopulateTablesData()
	if err != nil {
		return errors.New("could not read configs from database, hence program is exiting")
	}

	return nil
}

// RefreshTablesData function is called on each request to reload latest table setup
func RefreshTablesData() error {

	previousBookedSlots := make([][]available, len(Tables))
	for i, table := range Tables {
		previousBookedSlots[i] = table.NextAvailable
	}

	// creating context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := fetchDataFromDatabase(ctx)
	if err != nil {
		return err
	}

	for i, curTable := range Tables {

		t := table{}

		minCollection := MongoClient.Database(restaurantBooking.bookingDatabase).Collection(restaurantBooking.mintableSeatingCollection)
		if minCollection == nil {
			return errors.New("error in getting collection handle")
		}

		minCollection.FindOne(ctx, bson.M{"seating_capacity": curTable.MaxCapacity}).Decode(&t)
		curTable.MinCapacity = t.MinCapacity
		curTable.NextAvailable = previousBookedSlots[i]
		Tables[i] = curTable
	}

	return nil
}

// PopulateTablesData function is called on service startup only
func PopulateTablesData() error {

	// creating context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := fetchDataFromDatabase(ctx)
	if err != nil {
		return err
	}

	for i, curTable := range Tables {

		t := table{}

		minCollection := MongoClient.Database(restaurantBooking.bookingDatabase).Collection(restaurantBooking.mintableSeatingCollection)
		if minCollection == nil {
			return errors.New("error in getting collection handle")
		}

		minCollection.FindOne(ctx, bson.M{"seating_capacity": curTable.MaxCapacity}).Decode(&t)
		curTable.MinCapacity = t.MinCapacity
		Tables[i] = curTable
	}

	return nil
}

// fetchDataFromDatabase fetches data from db and populates Tables
func fetchDataFromDatabase(ctx context.Context) error {

	Tables = []table{}

	// returns handle to collection
	maxCollection := MongoClient.Database(restaurantBooking.bookingDatabase).Collection(restaurantBooking.maxtableSeatingCollection)
	if maxCollection == nil {
		return errors.New("error in getting collection handle")
	}

	// reading docs
	cursor, err := maxCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var currTable table
		err := cursor.Decode(&currTable)
		if err != nil {
			return err
		}

		currTable.NextAvailable = nil
		Tables = append(Tables, currTable)
	}

	return cursor.Err()
}

// GetAvgTotForGivenSeats given average tot for given party size
func GetAvgTotForGivenSeats(partySize int) int {

	t := wait{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := MongoClient.Database(restaurantBooking.bookingDatabase).Collection(restaurantBooking.avgTotCollection)
	collection.FindOne(ctx, bson.M{"$and": []interface{}{bson.M{"min_party_size": bson.M{"$lte": partySize}}, bson.M{"max_party_size": bson.M{"$gte": partySize}}}}).Decode(&t)

	// if given party_size does not fit in any criteria
	if t.AvgTot == "" {
		t.AvgTot = restaurantBooking.maxWaitTime
	}

	strings := strings.Split(t.AvgTot, "m")
	waitTime, _ := strconv.Atoi(strings[0])

	return waitTime
}

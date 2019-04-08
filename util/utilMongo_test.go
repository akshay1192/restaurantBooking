package util

import (
	"github.com/akshay1192/restaurantBooking"
	"testing"
)

// start mongo server and check for connection
func TestInitialiseMongoClient(t *testing.T) {

	err := restaurantBooking.InitialiseMongoClient()
	if err != nil {
		t.Fail()
	}
}

// populating the table data
func TestPopulateTablesData(t *testing.T) {

	err := restaurantBooking.PopulateTablesData()
	if err != nil && len(restaurantBooking.Tables) != 0 {
		t.Fail()
	}
}

func TestRefreshTablesData(t *testing.T) {

	err := restaurantBooking.RefreshTablesData()
	if err != nil && len(restaurantBooking.Tables) != 0 {
		t.Fail()
	}

}

// party_size falling in any given criteria
func TestGetAvgTotForGivenSeats(t *testing.T) {

	wait := restaurantBooking.GetAvgTotForGivenSeats(1)
	wait2 := restaurantBooking.GetAvgTotForGivenSeats(3)
	wait3 := restaurantBooking.GetAvgTotForGivenSeats(8)

	if wait != 15 {
		t.Fail()
	}

	if wait2 != 15 {
		t.Fail()
	}

	if wait3 != 50 {
		t.Fail()
	}

}

// default time
func TestGetAvgTotForGivenSeats1(t *testing.T) {

	wait := restaurantBooking.GetAvgTotForGivenSeats(10)
	if wait != 90 {
		t.Fail()
	}

}

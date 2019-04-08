package util

import (
	"github.com/akshay1192/restaurantBooking"
	"testing"
	"time"
)

// incorrect date/time format
func TestInitialDateTimeChecks(t *testing.T) {

	expected := false

	timeFormatCheck, _ := restaurantBooking.InitialDateTimeChecks("2019-04-07", "15:010")
	dateFormatCheck, _ := restaurantBooking.InitialDateTimeChecks("201904-07", "15:00")

	if (timeFormatCheck && dateFormatCheck) != expected {
		t.Fail()
	}
}

// booking < 2 hr OR booking > 48 hr
func TestInitialDateTimeChecks2(t *testing.T) {

	expected := false

	currDate := time.Now().Format("2006-01-02")
	currTime := time.Now().Add(1 * time.Hour).Format("15:04")
	bookingUnder2HourCheck, _ := restaurantBooking.InitialDateTimeChecks(currDate, currTime)

	currDate = time.Now().Format("2006-01-02")
	currTime = time.Now().Add(49 * time.Hour).Format("15:04")
	bookingOver48HourCheck, _ := restaurantBooking.InitialDateTimeChecks(currDate, currTime)

	if (bookingUnder2HourCheck && bookingOver48HourCheck) != expected {
		t.Fail()
	}
}

// booking before/after restaurant opens/closes
func TestInitialDateTimeChecks3(t *testing.T) {

	expected := false

	currDate := time.Now().Format("2006-01-02")
	beforeRestaurantOpenCheck, _ := restaurantBooking.InitialDateTimeChecks(currDate, "12:00")
	afterRestaurantOpenCheck, _ := restaurantBooking.InitialDateTimeChecks(currDate, "22:00")

	if (beforeRestaurantOpenCheck && afterRestaurantOpenCheck) != expected {
		t.Fail()
	}
}

// booking when restaurant is open
func TestInitialDateTimeChecks4(t *testing.T) {

	expected := true

	currDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	actual, _ := restaurantBooking.InitialDateTimeChecks(currDate, "15:00")

	if actual != expected {
		t.Fail()
	}

}

func TestCheckFromCurrentlyAvailableTables(t *testing.T) {

	restaurantBooking.InitialiseMongoClient()

	currDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	index := restaurantBooking.CheckFromCurrentlyAvailableTables(restaurantBooking.Tables, 3, currDate, "19:00")
	if index == -1 {
		t.Fail()
	}
}

func TestCheckFromCurrentlyAvailableTables1(t *testing.T) {

	restaurantBooking.InitialiseMongoClient()

	currDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	booking1 := restaurantBooking.CheckFromCurrentlyAvailableTables(restaurantBooking.Tables, 3, currDate, "19:00")
	booking2 := restaurantBooking.CheckFromCurrentlyAvailableTables(restaurantBooking.Tables, 3, currDate, "19:00")

	if booking1 == -1 || booking2 == -1 {
		t.Fail()
	}

}

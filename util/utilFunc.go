package util

import (
	"errors"
	"github.com/akshay1192/restaurantBooking"
	"math"
	"time"
)

const (
	restaurantBookingStartTime = "13:00"
	restaurantBookingCloseTime = "21:00"
	spaceDelimeter             = " "
	dateTimeLayout             = "2006-01-02 15:04"
	maxWaitTime                = "90m"
	mongoHostURL               = "mongodb://localhost:27017"
	bookingDatabase            = "booking"
	maxtableSeatingCollection  = "maxTableSeating"
	mintableSeatingCollection  = "minTableSeating"
	avgTotCollection           = "averageTot"
)

// InitialDateTimeChecks function checks preliminary checks like restaurant open, time diff >=2 and <= 48
// return false on failure
func InitialDateTimeChecks(dateStr, timeStr string) (bool, error) {

	//booking date and time from request param
	dateTimeStr := dateStr + spaceDelimeter + timeStr

	//converting string to date format
	bookingTime, err := time.ParseInLocation(dateTimeLayout, dateTimeStr, time.Now().Location())
	if err != nil {
		return false, errors.New("please give date (YYYY-MM-DD) and time (HH:MM) in correct format : " + err.Error())

	}

	//restaurant open and difference between booking and current time in hrs
	diff := bookingTime.Sub(time.Now()).Hours()
	if (timeStr < restaurantBookingStartTime || timeStr > restaurantBookingCloseTime) || (diff < 2 || diff > 48) {
		return false, errors.New("restaurant currently not taking bookings for given date and time")
	}

	return true, nil
}

// CheckFromCurrentlyAvailableTables checks and returns index of currently available table
// in case no table is available, it return -1
func CheckFromCurrentlyAvailableTables(tables []restaurantBooking.table, number int, dateStr, timeStr string) int {
	currLoss := math.MaxInt32
	index := -1

	loc := time.Now().Location()
	dateTimeStr := dateStr + spaceDelimeter + timeStr

	bookingTime, err := time.ParseInLocation(dateTimeLayout, dateTimeStr, loc)
	if err != nil {
		return index
	}

	for i, table := range tables {

		tableLoss := math.MaxInt32
		tableIndex := -1

		if number >= table.MinCapacity && number <= table.MaxCapacity {

			for _, t := range table.NextAvailable {

				avgWaitTOT := time.Duration(restaurantBooking.GetAvgTotForGivenSeats(number)) * time.Minute

				if (bookingTime.Sub(t.OccupiedAt) >= 0 && bookingTime.Sub(t.OccupiedTill) < 0) || (bookingTime.Sub(t.OccupiedAt) <= 0 && bookingTime.Add(avgWaitTOT).Sub(t.OccupiedAt) > 0) {
					tableLoss = math.MaxInt32
					tableIndex = -1
					break
				}

				loss := table.MaxCapacity - number
				if loss < tableLoss {
					tableIndex = i
					tableLoss = loss
				}

			}

			if len(table.NextAvailable) == 0 {

				loss := table.MaxCapacity - number
				if loss < tableLoss {
					tableIndex = i
					tableLoss = loss
				}
			}
		}

		if tableLoss < currLoss {
			index = tableIndex
			currLoss = tableLoss
		}
	}

	if index != -1 {
		bookedTill := bookingTime.Add(time.Duration(restaurantBooking.GetAvgTotForGivenSeats(number)) * time.Minute)
		nextAvailable := restaurantBooking.available{bookingTime, bookedTill}
		tables[index].NextAvailable = append(tables[index].NextAvailable, nextAvailable)
	}

	return index
}

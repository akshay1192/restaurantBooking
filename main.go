package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jitinp/go-exercise-akshay/util"
	"io/ioutil"
	"net/http"
	"time"
)

type request struct {
	PartySize int    `json:"party_size"`
	Phone     int    `json:"phone"`
	Name      string `json:"name"`
}

type response struct {
	Success bool       `json:"success"`
	Result  *result    `json:"result,omitempty"`
	Errors  *errorType `json:"errors,omitempty"`
}

type result struct {
	TableID     int    `json:"table_id"`
	BookingDate string `json:"booking_date"`
	BookingTime string `json:"booking_time"`
	Seats       int    `json:"no_of_seats"`
}

type errorType struct {
	Reason string `json:"reason"`
}

const (
	bookingRoute    = "/booking/{date}/{time}"
	port            = ":8080"
	dateParam       = "date"
	timeParam       = "time"
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

func (res *response) buildErrorResponse(reason string) []byte {

	errorType := &errorType{}
	res.Success = false
	res.Errors = errorType

	errorType.Reason = reason

	response, _ := json.Marshal(res)
	return response
}

func (res *response) buildSuccessResponse(tableID, seats int, date, time string) []byte {

	result := &result{}
	res.Success = true
	res.Result = result

	result.BookingDate = date
	result.BookingTime = time
	result.TableID = tableID
	result.Seats = seats

	response, _ := json.Marshal(res)
	return response
}

func bookTableHandler(w http.ResponseWriter, r *http.Request) {

	dateStr := mux.Vars(r)[dateParam]
	timeStr := mux.Vars(r)[timeParam]

	w.Header().Set(contentType, applicationJSON)

	if ok, err := util.InitialDateTimeChecks(dateStr, timeStr); !ok {

		res := &response{}
		responseInByte := res.buildErrorResponse(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseInByte)
		return
	}

	req := request{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res := &response{}
		responseInByte := res.buildErrorResponse(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseInByte)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		res := &response{}
		responseInByte := res.buildErrorResponse(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseInByte)
		return
	}

	err = util.RefreshTablesData()
	if err != nil {
		fmt.Println("Could not read configs from database, hence program is exiting")
		panic(err)
	}

	index := util.CheckFromCurrentlyAvailableTables(util.Tables, req.PartySize, dateStr, timeStr)
	if index != -1 {

		res := &response{}
		responseInByte := res.buildSuccessResponse(index+1, req.PartySize, dateStr, timeStr)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseInByte)
		return
	}

	res := &response{}
	responseInByte := res.buildErrorResponse("Sorry, no tables available for booking !!")

	w.WriteHeader(http.StatusBadRequest)
	w.Write(responseInByte)

}

func main() {

	// creating context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := util.InitialiseMongoClient()
	if err != nil {
		fmt.Println("Could not connect to mongo database, hence program is exiting")
		panic(err)
	}
	defer util.MongoClient.Disconnect(ctx)

	router := mux.NewRouter()
	router.HandleFunc(bookingRoute, bookTableHandler).Methods(http.MethodPost)

	fmt.Println("----------- Server getting started for restaurant booking -----------")

	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("Server shutting down")
		panic(err)
	}
}

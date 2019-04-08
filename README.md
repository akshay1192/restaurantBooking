# The Problem:
McDoe Restaurant is a popular fine-dine restaurant in South Bangalore, but in the recent quarters the restaurant has incurred loss. Management team of McDoe found following reasons to for the consistent  loss of revenue
1. Revenue loss due to Over/Under booking
2. Revenue loss due to unoptimized utilization of seating capacity

McDoe management has asked for your help, to build the API backend for an online reservation system. The booking system will allow user to submit a table booking request from McDoe’s website. Upon receiving the request the system will reserve a table, if the following conditions are met 
+ Restaurant is open
+ Table is available for occupancy
+ Number of requested seats should be within the minimum and maximum of seating capacity of a table
+ Booking request should be made 2 to 48 hrs prior

The system should also ensure that the following criterias are fulfilled
+ At no point of time the restaurant is overbooked 
+ The table allocation should be done in a way, that restaurant can accommodate maximum possible dinners.

## Configurations and Preconditions:
1. The solution should be implemented as a REST API Endpoint
2. For now you need to consider only a single table can be booked by one booking request. If the system cannot find a suitable match, it will return a failure
3. The floor layout/ seating arrangement will be passed to the application as a JSON data. The configuration can be changed anytime by simply swapping the old configuration with new one.  Here is a sample  json from **layout.json** -
```js
[
  {"table_id":1 , "max_seating": 1},
  {"table_id":2 , "max_seating": 2},
  {"table_id":3 , "max_seating": 2},
  {"table_id":4 , "max_seating": 2},
  {"table_id":5 , "max_seating": 3},
  {"table_id":6 , "max_seating": 3},
  {"table_id":7 , "max_seating": 4},
  {"table_id":8 , "max_seating": 4},
  {"table_id":9 , "max_seating": 6},
  {"table_id":10 , "max_seating": 6},
  {"table_id":11 , "max_seating": 8},
  {"table_id":12 , "max_seating": 12}
]
```
4. The minimum allowed occupancy for a table will be decided by some predefined configuration json. These rules can be modified by the admin at anypoint of time. Here is a sample **seating.json** -
```js
[
  {"seating_capacity": 2 , "min_occupency":1},
  {"seating_capacity": 3 , "min_occupency":1},
  {"seating_capacity": 4 , "min_occupency":2},
  {"seating_capacity": 5 , "min_occupency":2},
  {"seating_capacity": 6 , "min_occupency":3},
  {"seating_capacity": 7 , "min_occupency":3},
  {"seating_capacity": 8 , "min_occupency":4},
  {"seating_capacity": 9 , "min_occupency":4},
  {"seating_capacity": 10 , "min_occupency":5},
  {"seating_capacity": 11 , "min_occupency":5},
  {"seating_capacity": 12 , "min_occupency":6}
]
```
5. The average table occupancy times (Avg.TOT) for dinners are pre-assumed based on research data. This will be passed to the application as a json configuration. The Avg.TOT is dependent on the party size (Number of seat reserved). Here is a sample **tot.json** -
```js
[
  {"min_party_size":1,"max_party_size":3,"avg_tot": "15m"},
  {"min_party_size":4,"max_party_size":5,"avg_tot": "35m"},
  {"min_party_size":6,"max_party_size":9,"avg_tot": "50m"},
  {"min_party_size":10,"avg_tot": "90m"}
]
```

## API Contract
You need to create a post API endpoint which accepts booking date, booking time and party size like below
```sh
 POST /booking/:date/:time
   content-type : application-json
   body : { party_size: Number , phone: Number , name : String }
 ```
   
The API should return a boolean success parameter to indicate, whether booking request is a successful or not. For a successful booking, it should return a result object, and in case of failure an error object will be expected. All the responses should have a relevant http status code, like - 200 for success , 4XX and 5XX for errors. The Response schema should be like following
```js
{
  success: Boolean,
  result: {  // optional
    table_id : Number,
    booking_date : String,
    booking_time : String,
    no_of_seats: Number,
  }
  errors : { //optional
    reason: String
  }
}
```
## General Instructions
1. Use Go 1.5 or latest
2. You can use any of these databases - MySQL/MariaDB, Postgresql, MongoDB
3. We appreciate multiple small commits
4. The code should have unit-test cases with maximum test coverage
5. The code should not have any golint warning
6. You need **follow the same naming convention for the configuration files and API request/response as mentioned in the document**. You work will be assessed by automated test-cases (BDD), which follows the naming convention mentioned in the document.
7. **Commit your code to this repository on master branch**

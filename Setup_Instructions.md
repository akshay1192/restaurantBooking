Assumptions :
    
    * Mongo db should be installed and running on system
        
    * Mongodb url : "mongodb://localhost:27017"  
    
    * Mongo database : booking
        
    * Mongo collections : maxTableSeating, minTableSeating, averageTot

    * Service running on localhost on port : 8080

    * API Endpoint : localhost:8080/booking/2019-04-08/17:50   (where :date is 2019-04-08  and :time is 17:50)

    * Restaurant starts taking booking from 13:00 and closes booking on 21:00
    
    
Prerequisites :

    
    1. Mongo server should be running at localhost on port :27017 (mongodbUrl : "mongodb://localhost:27017")



Mongo Dataset 1 :

    ------------------------ InsertData (Dataset 1) into mongo -------------------------


    1. Open mongo client in one terminal
    2. Type in these commands on terminal :

        * use booking


        * db.maxTableSeating.insertMany([
            {
              "table_id": 1,
              "max_seating": 4
            },
            {
              "table_id": 2,
              "max_seating": 7
            },
            {
              "table_id": 3,
              "max_seating": 15
            },
            {
              "table_id": 4,
              "max_seating": 5
            }
          ])


        * db.minTableSeating.insertMany([
            {
              "seating_capacity": 4,
              "min_occupency": 1
            },
            {
              "seating_capacity": 7,
              "min_occupency": 5
            },
            {
              "seating_capacity": 15,
              "min_occupency": 8
            },
            {
              "seating_capacity": 5,
              "min_occupency": 2
            }
          ])


         * db.averageTot.insertMany([
             {
               "min_party_size": 1,
               "max_party_size": 3,
               "avg_tot": "15m"
             },
             {
               "min_party_size": 4,
               "max_party_size": 5,
               "avg_tot": "35m"
             },
             {
               "min_party_size": 6,
               "max_party_size": 9,
               "avg_tot": "50m"
             },
             {
               "min_party_size": 10,
               "avg_tot": "90m"
             }
           ])
           
           
           
Mongo Dataset 2 :     

    ------------------------ InsertData (Dataset 2) into mongo -------------------------


    1. Open mongo client in one terminal
    2. Type in these commands on terminal :

            * db.dropDatabase()

            * use booking

            * db.maxTableSeating.insertMany([
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
              ])

            * db.minTableSeating.insertMany([
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
              ])

            * db.averageTot.insertMany([
                 {"min_party_size":1,"max_party_size":3,"avg_tot": "15m"},
                 {"min_party_size":4,"max_party_size":5,"avg_tot": "35m"},
                 {"min_party_size":6,"max_party_size":9,"avg_tot": "50m"},
                 {"min_party_size":10,"avg_tot": "90m"}
              ])


Run service :

    cd go-exercise-akshay
    go run main.go
    
    Postman collection : https://www.getpostman.com/collections/4711a4906c6aa790711c
    
    
Run Test cases :

    cd go-exercise-akshay/util
    go test
    
    
Test cases code coverage :

    cd go-exercise-akshay/util
    go test -cover
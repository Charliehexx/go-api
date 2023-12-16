<h1>CRUD REST API on Car management service using Gofr framework</h1>


 <h2>Requirements to run the project locally on system</h2>
 
 ## ⚡️ Quick Start Guide
    1- Go lang should be installed 
    2- Docker should be installed for container services

## <h2>Setting up the Project</h2>

 ```Clone the repo and exeucte these commands
    1-  docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
    2-  docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE cars (id INT AUTO_INCREMENT PRIMARY KEY,license_plate VARCHAR(20) NOT NULL, model VARCHAR(50) NOT NULL,color VARCHAR(20) NOT NULL,repair_status VARCHAR(50) DEFAULT NULL,entry_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP );"
   3-  go mod tidy
   4- go run main.go
```
<h1> Postman Collection Workspace </h1>
After successfully cloning the repository and configuring mysql database and running the project you can head to [Postman Workspace](https://www.postman.com/maintenance-participant-81084192/workspace/my-workspace/collection/21737598-d2e28910-2ce0-4ffd-832c-41d0095bf966?action=share&source=copy-link&creator=21737598) and test the APIs by hitting on endpoints.Ensure that postman is installed on your system.

<h2>Screenshot and Working of the APIs</h2>
   <h2>Working of POST Method.</h2>
   <h3>For adding car into the database Endpoint- ```http://localhost:9000/car/enter```</h3>
   <h3>POST Request Successfully Made in the database with the required Parameter.</h3>
   
![Screenshot (25)](https://github.com/Charliehexx/go-api/assets/86345323/901171c4-bcc0-4385-9ab3-c08ce67a62f9)

   <h3>POST  Request Rejected as the license_plate with same value already exists in database.</h3>

![Screenshot (26)](https://github.com/Charliehexx/go-api/assets/86345323/1c232393-7a38-4c39-99bf-fe53e0295d30)

<h3>POST Request rejected as required parameter here(repair_status) is missing.</h3>

![Screenshot (27)](https://github.com/Charliehexx/go-api/assets/86345323/7cf003db-2dd5-4c09-9647-fe69bef5ebe4)


 <h2>Working of GET Method.</h2>
  <h3>For gettign all the cars from the  database Endpoint- ```http://localhost:9000/car```</h3>
  <i>We get all the  information about cars and their status till now present in the database</i> 
    <h3>GET Request Successfully Made in the database with the required Parameter.</h3>

![Screenshot (28)](https://github.com/Charliehexx/go-api/assets/86345323/41753d8d-3f02-4023-a672-08cce5b0ac5c)

 <h2>Working of PUT Method.</h2>
    <h3>For updating repair_status of the car into the database Endpoint- ```http://localhost:9000/car/update/{id}```</h3>
    <i>Car reapir_status gets updated in the database on the id which is called.</i>
    <h3>PUT Request Successfully Made in the database updating the repair_status of the car.</h3>
    
![Screenshot (29)](https://github.com/Charliehexx/go-api/assets/86345323/39397069-9fb5-42f6-bf2b-f818c6e3196a)

  <h2>Working of DELETE Method.</h2>
   <h3>For updating repair_status of the car into the database Endpoint- ```http://localhost:9000/car/delete/{id}```</h3>
  <i>Car gets deleted successfully from the database while hitting a request with existing id in the database.</i>
    <h3>DELETE Request Successfully Made in the database updating the repair_status of the car.</h3>
    
![Screenshot (30)](https://github.com/Charliehexx/go-api/assets/86345323/a7d5c8a5-660e-4f90-b00b-ba779f75747a)
    
<h2>Sequence Diagram of the Proeject</h2>
```
+----------------+         +--------------------------+         +-----------------
--+
|     Client     |         |         App              |         |      Database     |
+----------------+         +--------------------------+         +-------------------+
       |                             |                                   |
       |   POST /car/enter            |                                   |
       |---------------------------> |                                   |
       |                             |                                   |
       |                             |                                   |
       |                             |   Execute SQL: INSERT INTO cars |
       |                             |--------------------------------->|
       |                             |                                   |
       |                             |                                   |
       |                             |           Return Response       |
       |                             |<----------------------------------|
       |                             |                                   |
       |                             |                                   |
       |   GET /car                  |                                   |
       |---------------------------> |                                   |
       |                             |                                   |
       |                             |   Execute SQL: SELECT * FROM cars|
       |                             |--------------------------------->|
       |                             |                                   |
       |                             |           Return Cars List      |
       |                             |<----------------------------------|
       |                             |                                   |
       |   PUT /car/update/{id}      |                                   |
       |---------------------------> |                                   |
       |                             |                                   |
       |                             | Decode JSON: repair_status      |
       |                             |--------------------------------->|
       |                             |                                   |
       |                             |   Execute SQL: UPDATE cars      |
       |                             |--------------------------------->|
       |                             |                                   |
       |                             |           Return Response       |
       |                             |<----------------------------------|
       |                             |                                   |
       |   DELETE /car/delete/{id}   |                                   |
       |---------------------------> |                                   |
       |                             |                                   |
       |                             |   Execute SQL: DELETE FROM cars |
       |                             |--------------------------------->|
       |                             |                                   |
       |                             |           Return Response       |
       |                             |<----------------------------------|       
```


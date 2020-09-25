# PassHash

## Overview

This is a simple service written in golang. The primary purpose of this service is to allow for a password to be POST'ed to an endpoint and hashed. After this password is hashed it should be accessible through a GET request. There is no persistence on the passwords, as they are only stored in memory. 

## Endpoints
All of the endpoints are defined in this Postman collection. You can import the collection, and use it as a base to interact with the service.

[Postman Collection](assets/PassHash.postman_collection.json)

### POST /hash
This is the primary endpoint for submitting a request to have a password be hashed. Once a password has been queued the endpoint will return a unique identifier that can be used to retrieve the hashed password once complete. 
#### Params
 **Password** - This field is required, and cannot exceed 30 characters in length. 

Example Request
```bash
$ curl -X POST -H"Content-Type:application/x-www-form-urlencoded" -d 'password=123' http://127.0.0.1:8880/hash
1
```

### GET /hash/:id
This endpoint can be used to retrieve your queued hash results, this value should be avlible within 5 seconds. 

#### Params
 **id** - This field is required, and must be numeric

Example Request
```bash
$ curl -X GET  http://127.0.0.1:8880/hash/1
PJkJr+wlNU1VHa4hWQuybjjVPyFzuNPcPu5MBH56scHri4UQPjvnumE7MbtcnDYhTcnxSkL9ei/bhIVrylxEwg==
```

### GET /stats
This endpoint can be used to retrieve server statistics on processing. 

Example Request
```bash
$ curl -X GET  http://127.0.0.1:8880/stats
{"total":1,"average":36}
```

### GET /shutdown
This endpoint will shutdown the service, warning this will reset all the data in the service as it is a in memory service.  

Example Request
```bash
$ $ curl -X GET  http://127.0.0.1:8880/shutdown
  OK
```

### Errors
Standard HTTP status codes are used. Here are some descriptions of the common ones in the service. 

Bad Request  400  You failed validation, there will me a message with details in the response. 
Bad Request  400  You failed validation, there will me a message with details in the response. 

## Starting the service
You can start the service with the following command
```bash
$ go build github.com/jonccrawley/passhash
$ ./passhash
```

The Server will start up at http://127.0.0.1:8880/ by default

## Test the codebase
You can run the unit tests with the following command from the projects root directory  
```bash
$ go test ./...
```

You can also use apache benchmarking to test out concurrent invocations. 
```bash
$ echo "password=1235456" >> /tmp/post.data
$ ab -n 10000 -c 5 -p /tmp/post.data -T application/x-www-form-urlencoded  -l http://127.0.0.1:8880/hash
```

## Directories

It might just be the java developer in me, but i like to have my code segmented out into logical packages. These are the intentions for each of the packages. 

### /assets

Files that can be linked in the README.md

### /backend

Backend and worker functions used to sha512 and Base64Encode the password

### /backend

Backend and worker functions used to sha512 and Base64Encode the password

### /definition
The primary place for interfaces to be defined for cross package access and reference. 

### /handler
Common folder for all endpoints

### /model
Location for co0mmon object definitions 

### /repository
In memory data persistence implementations

### /utils
Any helper functions or tools







 

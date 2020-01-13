# babe_api
## docker container run
```
docker-compose up -d
```
## go run
```
go run main.go
```

## directory structure
### commands
 Batch files
### configs
 Common variables setting
### controllers
 Contains handler functions for particular route to be called when an api is called
### db
 Contains migration files and db settings
### middlewares
 Middleware to be used for this project
### models
 Database tables to be used as model struct
### routes
 Url routing when a request comes
### services
 Contains all the core business logic including db connections
### tests
 Testing files
### .env
 Environment variables setting
### main.go
 Starting point

## rest api url design
- Register user
POST /register
- Login user
POST /login
- Logout user
POST /logout
- get common profile questions
GET /questions
- get your profiles
GET /profiles
- get user profiles
GET /users/{userId}/profiles
- save your profiles
POST /profiles
- get recommended users
GET /users/recommended
- save like or profile impression
POST /users/{userId}/impression
- get your reviews
GET /reviews
- get user review (comment, rating)
GET /users/{userId}/reviews
- save user review (comment, rating)
POST /users/{userId}/review
- block user review
POST /reviews/{reviewId}/block
- get chat rooms
GET /chatRoomConfigs
- save chat room, or matching
POST /chatRoomConfigs
- block chat room
POST /users/{userId}/chatRoomConfigs/block

## database schema design
### users
#### manage user info
id: autoincrement
username: string
password: string
image_path: string
gender: unsignedTinyInteger
age: unsignedInteger
country: unsignedInteger
job: unsignedInteger
prefecture: unsignedInteger
user_type: unsignedTinyInteger
created_at: timestamp
updated_at: timestamp
### questions
#### manage profile questions
id: auto increment
question_type: unsignedInteger
question: string
del_flg: unsignedTinyInteger
created_at: timestamp
updated_at: timestamp
### profiles
#### manage profile info
id: auto increment
user_id: unsignedInteger
(question_id: unsignedInteger)
self_introduction_comment: text
answer: unsignedInteger(or json)
protection: unsignedInteger
created_at: timestamp
updated_at: timestamp
### chat_room_configs
#### manage chat room info
id: auto increment
male_user_id: unsignedInteger
female_user_id: unsignedInteger
status: unsignedInteger
created_at: timestamp
updated_at: timestamp
### reviews
#### manage review
id: auto increment
user_id_from: unsignedInteger
user_id_to: unsignedInteger
rating: unsignedTinyInteger
comments: text
like: boolean
status: unsignedInteger
created_at: timestamp
updated_at: timestamp

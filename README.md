# Virtual Pets (Backend)

This is the back end CRUD server for my Virtual Pets Projects using Gorilla/Mux and MongoDB.
The front end project is [here](https://github.com/luycaslima/virtual-pets)

## What is Virtual Pets the objective of this project?

Virtual Pets is a Web-based game inspired by Neopets and Tamagotchi/Digimon.
Where you can create, feed, play and train your pets.

### Functional Fetaures at the moment
 - Create Account
   - Log in / Log off
   - JWT authentication
   - Register
 - Create Pet
   - link to the user account

## Setup

To run this project, you need to set up an .env with MONGOURI (URL to the MongoDB instance (online or local))
and a JWT_SECRET_KEY string to set up the JWT authentication.

### How to run it and Documentation

The Documentation is made with the Swagger bind of Golang: [Swaggo](https://github.com/swaggo/swag)
To test the Routes and see the Documentation, while running the project access the localhost:8080/documentation 
```
go run main.go
```


## TODO
 - [ ] Migrate documentation to another Swagger package
 - [ ] User Routes
   - [X] Log in/ Log off
   - [X] Register Account
   - [X] Create Pet
   - [ ] Feed Pet
   - [ ] Train Pet
   - [ ] Play with Pet
 - [ ] Pet Routes
   - [X] Create a pet from a Specie
 - [ ]  Admin Routes
   - [X]  Create Species
     - [ ] Documentation
   - [ ]  Auth
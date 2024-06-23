# Auth App
Simple Auth App Project

## Tech Stack
- Golang

## Project Structure

### **`cmd/`**
Contains the application's entry point.  
### **`internal/`**
Contains the core application logic, divided into subdirectories:  
-`entity/`: Defines the core business entities.  
-`usecase/`: Implements the business logic and application use cases.  
-`repository/`: Defines data access interfaces and implementations.  
-`handler/`: Handles HTTP requests and responses.  
-`utils/`: Utility functions or helper code.
### **`mocks/`**
Contains mocks for all interface in the project

## How to run
Go into your terminal/command/etc and type the command below  
`./script.sh run`  
This will build and run the app
  
After the apps running you will see  
`Server is running on port 8080`

## Other script
`./script.sh build`  
Build the app

`./script.sh test`  
Run all unit test

`./script.sh mock`  
Generate mocks for interfaces
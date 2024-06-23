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

## How to run

`./script.sh build`  
Build the app

`./script.sh run`  
Build and run the app

`./script.sh test`  
Run all unit test

`./script.sh mock`  
Generate mocks for interfaces
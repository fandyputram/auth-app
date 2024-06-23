#!/bin/bash

# Function to build the app
build_app() {
    echo "Building the project..."
    go build -o auth-app ./cmd
}

# Function to build and run the app
build_and_run_app() {
    build_app
    echo "Running the project..."
    ./auth-app
}

# Function to clean the app
clean_app() {
    echo "Cleaning up..."
    go clean
    rm -f auth-app
    echo "Cleanup complete."
}

# Function to test the app
test_app() {
    echo "Running tests..."
    go test ./...
}

# Function to generate mocks
generate_mock() {
    echo "Generating mocks..."
    # Generate Usecase mocks
    mockgen -source=internal/usecase/auth.go -destination=mocks/mock_usecase_auth.go -package=mocks

    # Generate Repository mocks
    mockgen -source=internal/interface/repository/user.go -destination=mocks/mock_repository_user.go -package=mocks
}

# Main script logic
case "$1" in
    build)
        build_app
        ;;
    run)
        build_and_run_app
        ;;
    test)
        test_app
        ;;
    mock)
        generate_mock
        ;;
    clean)
        clean_app
        ;;
    *)
        echo "Usage: $0 {build|run|test|mock|clean}"
        ;;
esac

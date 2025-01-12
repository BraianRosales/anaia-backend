GO-1 class

## Initializes the project with go mod init:

cd Anaia-backend
go mod init anaia-backend
touch main.go

run the app: go run main.go

-----------------------------------------------------------------------------

## Design patterns applied.

* Clean architecture: Layers that inherit from the lowest to the highest layer, can't get dependencies from higher layers.

    - Repository: Communicates with the database.
    - Service/Business logic: Business logic (e.g., constraints).
    - Presentation: API

    - Repository < Service < API 

* Domain Driven design

----------------------------------------------------------------------------

## Packages

* settings: package where I get all the application configuration, port where it will be running, database credentials configuration, etc.

* internal: Everything that is sensitive, golang as internal hidden security. Here are the layers.

---------------------------------------------------------------------------

## Libraries: 

* fx (Used to import dependencies into main in a more sensible and readable way).

---------------------------------------------------------------------------

## Project Structure

/anaia-backend
  /internal
    /api
      api.go // API logic (HTTP controllers, endpoints)
    /repository
      repository.go // Interaction with database or data repositories
    /service
      service.go // Core business logic
  /settings
    settings.go // Global project settings (such as credentials, configurations)
    settings.yaml // Configuration file (e.g. environment variables)
  main.go // Main application entry (initialization and execution)
  README.md // Project Overview

# netmonk-asset-management

Asset Management is my intern first mini project.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- Docker version 19.03.0
- Docker-compose version 1.24.1

### Installing

Run this command before run the app

For the first step, we need open the project folder then build the app image and database image using these commands :

```
docker build --target asset_management_app -t am_app .

docker build --target asset_management_db -t am_db .
```

### Running The Project

   1. Go to the project directory
   2. Open _scripts/docker-compose folder
   3. Run the docker container ```docker-compose up```
   4. Open the app on 'localhost:8080'

Thank you for last three months, enjoy.
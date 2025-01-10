# Cat APP Go Project

This project aims to fetch data from Booking.com and store it in a local PostgreSQL database. We will create a backend API using the Beego framework to serve the property data and design two front-end pages: property listing and property details.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Technologies Used](#technologies-used)
3. [Key Features](#key-features)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [Running the Project](#running-the-project)
6. [Running the Test](#running-the-test)
7. [API Endpoints](#api-endpoints)
8. [Project Structure](#project-structure)
9. [Contributing](#contributing)
10. [Authors](#authors)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (version 1.16 or later)
- Git
- Beego

## Technologies Used

- [Go](https://golang.org/) - The main programming language
- [Beego v2](https://github.com/beego/beego) - Web framework for building the API
- [TailwindCSS](https://tailwindcss.com/) - Front-end css framework


## Key Features

## Key Features
- **Data Fetching and Storage:**
  - Fetch data from Booking.com API.
  - Store the fetched data in a local PostgreSQL database.

- **Database Schema:**
  - **Location Table:** Stores location names details.
  - **Rental Property Table:** Provides basic information about properties such as name, type, bedroom, bathroom,  guest, location breadcrumb and amenities.
  - **Property Details Table:** Stores detailed information about properties, including images, city and descriptions.

- **Backend API using Beego Framework:**
  - **/v1/property/list:** Endpoint to provide a list of properties for a specific location.
  - **/v1/property/details:** Endpoint to provide detailed information about a specific property.

- **Front-end Pages:**
  - **Property Listing Page:**
    - Displays a list of properties for a given location.
    - Includes a search box for location-based searches.
    - Uses AJAX to load property tiles without page reloads.
  - **Property Details Page:**
    - Displays detailed information about a selected property.
    - Navigates from the property listing page by clicking on a property tile.

## Installation

### 1. Install Go

If you haven't installed Go, follow these steps:

1. Visit the official Go downloads page: https://golang.org/dl/
2. Download the appropriate installer for your operating system.
3. Follow the installation instructions for your OS:
   - Windows: Run the MSI installer and follow the prompts.
   - macOS: Open the package file and follow the prompts.
   - Linux: Extract the archive to `/usr/local`:
     ```
     tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
     ```
4. Add Go to your PATH:
   - For bash, add the following to your `~/.bashrc` or `~/.bash_profile`:
     ```
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - For other shells, add the equivalent to your shell's configuration file.
5. Verify the installation by opening a new terminal and running:
      ```
      go version
      ```
6. create src to keep your  all beego app here
      ```
      mkdir -p $GOPATH/src/
      ```

### 2. Clone the Repository

Clone this repository to your local machine in the default Go path:

```
cd $GOPATH/src/ 
git clone https://github.com/Khairun-Nahar-Munne/rent.git
cd rent
```

### 3. Install Dependencies

This project uses Beego v2 and other dependencies. Install them using the following commands:

```
go get github.com/beego/beego/v2
go get github.com/astaxie/beego/logs
go mod tidy
```


## Configuration

### 1. API Key

This project requires an API key from The Cat API. To obtain one:

1. Visit https://rapidapi.com/ntd119/api/booking-com18/playground/apiendpoint_ba764901-8751-43a1-9fb1-d882600b9cf8
2. Sign up for an account
3. Subscribe for the test
3. Ann API key will be generated in your dashboard
4. There is already API Key of mine. You can use it untill the api request rate limit is exceeded. You have to change api keys in the location controller and rental controller. 
### 2. Configuration File

Create a `conf/app.conf` file in the project root with the following content:

```ini
appname = rent-web
httpport = 8080
runmode = dev
api_key = "5871a5c278msh858d5b59a204dbbp1b44d5jsnd5c0bc2d25ee"
```
You can use your api_key if you blocked using my api_key
## Running the Project

To run the project, use the following command from the project root directory:

1. For using postgres and pgadmin:
```
docker-compose up
```
2. Open another terminal and run the project

```
bee run
```

If you don't have `bee` installed, you can install it with:

```
go get github.com/beego/bee/v2
```

3. To fetch the data from booking.com api:

- Run this `http://localhost:8080/api/locations/fetch` for fetching and storing locations data.
- Run this `http://localhost:8080/api/properties/fetch` for fetching and storing rental property and property details data.

4. To see the rent website: 
- Run this `http://localhost:8080/`.


## API Endpoints

The project provides the following API endpoints:

- `GET /api/locations/fetch`: Fetch and store location
- `GET /api/properties/fetch`: Fetch and store rental property and property details
- `GET /v1/property/list`: Feth data from location and rental property tables
- `GET /v1/property/details`: Fetch data from rental property and property details tables
- `GET /`: Render property lists
- `GET /property-details`: Render property details

## Project Structure

The project follows a standard Beego directory structure:

```
rent/
├── conf/
│   └── app.conf
├── controllers/
│   ├── location_controller.go
│   ├── rental_controller.go
│   ├── property_controller.go
│   └── default.go
├── routers/
│   └── router.go
├── models/
│   └── models.go
├── views/
│   ├── index.tpl
│   └── property-details.tpl
├── static/
│   ├── img
│   └── js
│       ├── listing.js
│       ├── details.js
│       └── reload.min.js
├── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

- `conf/`: Contains configuration files
- `controllers/`: Contains the logic for handling API requests (fetcing and store data from booking.com api and fetch the data from local database)
- `routers/`: Defines the routing for the application
- `models/`: Define structure of the data of database
- `main.go`: The entry point of the application


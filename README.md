# Cat APP Go Project

This project is a Go-based backend service that interacts with The Cat API to provide various cat-related functionalities. It uses the Beego web framework and offers endpoints for fetching random cat images, breed information, favoriting images, and voting on cats.

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
- [The Cat API](https://thecatapi.com/) - External API for cat-related data
- [TailwindCSS](https://tailwindcss.com/) - Front-end css framework


## Key Features

- Fetch random cat images
- Retrieve comprehensive breed information
- Get breed-specific images
- Favorite cat images
- Vote on cat images
- Manage favorite cat images (add, list, delete)
- Configuration management using Beego's app.conf
- Testing 

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
4. There is already API Key of mine. You can use it untill the api request rate limit is exceeded. You have to change api keys in the location controller and rental controller. (in 4 places)
### 2. Configuration File

Create a `conf/app.conf` file in the project root with the following content:

```ini
appname = rent
httpport = 8080
runmode = dev
cat_api_url = "https://api.thecatapi.com/v1"
api_key = "live_i5ikdgttExQhChEMpt7UuGateqcwS8IVYlMGGBwSO3Mp7CatoYAhl9VAUgZ76Pqa"
```

You can replace `api_key` with the API key you obtained from The Cat API.

## Running the Project

To run the project, use the following command from the project root directory:

```
bee run
```

If you don't have `bee` installed, you can install it with:

```
go get github.com/beego/bee/v2
```


The server will start, and you should see output indicating that it's running on `http://localhost:8080`.

## Running the Test

Run the test
```
go test -v ./tests 
```
Get Coverage
```
Step 1: go test -coverprofile coverage.out ./...
Step 2: go tool cover -func=coverage.out | grep total: | awk '{print $3}'
```
If you want to see coverage of each file in html
```
go tool cover -html coverage.out    
```

## API Endpoints

The project provides the following API endpoints:

- `GET /api/cat/fetch`: Fetch a random cat image
- `GET /api/breeds`: Get a list of all cat breeds
- `GET /api/breed/:id`: Get specific breeds
- `POST /api/favorites`: Add a cat image to favorites
- `POST /api/votes`: Vote on a cat image
- `GET /api/votes`: Get latest votes
- `GET /api/get_favorites`: Get a list of favorite cat images
- `DELETE /api/favourite/:id`: Delete a favorite cat image


## Project Structure

The project follows a standard Beego directory structure:

```
cat-app/
├── conf/
│   └── app.conf
├── controllers/
│   ├── location_controller.go
│   ├── rental_controller.go
│   ├── property_controller.go
│   ├── default.go
│   └── vote_controller.go
├── routers/
│   └── router.go
├── models/
│   └── models.go
├── views/
│   └── index.tpl
│   └── property-details.tpl
├── static/
│   ├── img
│   └── js
│       ├── listing.js
│       ├── details.js
│       └── reload.min.js
├── main.go
├── go.mod
├── go.sum
└── README.md
```

- `conf/`: Contains configuration files
- `controllers/`: Contains the logic for handling API requests (fetcing and store data from booking.com api and fetch the data from local database)
- `routers/`: Defines the routing for the application
- `models/`: Define structure of the data of database
- `tests/`: contains all tests file
- `main.go`: The entry point of the application

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Authors
This project is for Django admin panel practice created by me for assignment purposes during my internship days at w3 egnineers ltd. 
 
Khairun Nahar Munne  – khairunnaharmunne81@gmail.com
 
 You can find me here at:
[Github](https://github.com/Khairun-Nahar-Munne) 
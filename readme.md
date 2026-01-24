# Simple Category REST API (Go)

This project is a simple RESTful API built with **Go (net/http)** that
provides CRUD (Create, Read, Update, Delete) operations for categories.\
All data is stored **in memory**, making this project suitable for
learning, testing, and small demos.

------------------------------------------------------------------------

## Features

-   Health check endpoint
-   Get all categories
-   Get category by ID
-   Create a new category
-   Update an existing category
-   Delete a category
-   JSON-based request and response

------------------------------------------------------------------------

## Tech Stack

-   Go (Golang)
-   net/http
-   encoding/json

------------------------------------------------------------------------

## Data Structure

### Category

``` json
{
  "id": 1,
  "name": "Makanan",
  "description": "Kategori Makanan"
}
```

### Standard Response

``` json
{
  "status": true,
  "message": "API Running"
}
```

### Response With Data

``` json
{
  "status": true,
  "message": "Get Category",
  "data": {}
}
```

------------------------------------------------------------------------

## Getting Started

### Prerequisites

-   Go 1.21.0 or later

### Run the Server

``` bash
go run main.go
```

The server will run on:

    http://localhost:8080

------------------------------------------------------------------------

## API Endpoints

### Health Check

**GET** `/api/health`

**Response**

``` json
{
  "status": true,
  "message": "API Running"
}
```

------------------------------------------------------------------------

### Get All Categories

**GET** `/api/categories`

**Response**

``` json
{
  "status": true,
  "message": "Get Category",
  "data": [
    {
      "id": 1,
      "name": "Makanan",
      "description": "Kategori Makanan"
    }
  ]
}
```

------------------------------------------------------------------------

### Get Category By ID

**GET** `/api/categories/{id}`

**Response**

``` json
{
  "status": true,
  "message": "Get Category",
  "data": {
    "id": 1,
    "name": "Makanan",
    "description": "Kategori Makanan"
   }
}
```

------------------------------------------------------------------------

### Create Category

**POST** `/api/categories`

**Request Body**

``` json
{
  "name": "Dessert",
  "description": "Dessert Category"
}
```

**Response**

``` json
{
  "status": true,
  "message": "Create Category",
  "data": {
    "id": 4,
    "name": "Dessert",
    "description": "Dessert Category"
  }
}
```

------------------------------------------------------------------------

### Update Category

**PUT** `/api/categories/{id}`

**Request Body**

``` json
{
  "name": "Updated Name",
  "description": "Updated Description"
}
```

**Response**

``` json
{
  "status": true,
  "message": "Update Category",
  "data": {
    "id": 1,
    "name": "Updated Name",
    "description": "Updated Description"
  }
}
```

------------------------------------------------------------------------

### Delete Category

**DELETE** `/api/categories/{id}`

**Response**

``` json
{
  "status": true,
  "message": "Success delete category"
}
```

------------------------------------------------------------------------

## Error Response Format

``` json
{
  "status": false,
  "message": "Category not found"
}
```

Common error messages: - Invalid request - Invalid category ID -
Category not found

------------------------------------------------------------------------

## Notes

-   Data is stored in memory and will be reset when the server restarts.
-   No database is used.
-   No authentication or authorization implemented.
-   Suitable for learning REST API fundamentals in Go.

------------------------------------------------------------------------

## License

This project is open-source and free to use for learning and
experimentation.
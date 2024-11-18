# Author
Cliff Li
li.tuox@northeastern.edu

# Receipt Processor

A simple receipt processing API built in Go(Fetch take-home excercise).


## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Testing](#testing)
- [Example Requests](#example-requests)
- [Troubleshooting](#troubleshooting)

## Installation

### Prerequisites

- Go 1.20 or higher must be installed.

### Clone the Repository

```bash
git clone https://github.com/txli299/receipt-processor.git
cd receipt-processor
```

### Install Dependencies

```bash
go mod tidy
```

## Usage

### Running the Server
Start the server using:
```bash
go run main.go
```
The server will start at http://localhost:8080.


## Endpoints

### 1. **Submit Receipt**

- **Method**: `POST`
- **URL**: `/receipts/process`
- **Description**: Submits a receipt for processing and returns a unique ID.

#### Request

- **Headers**:
  - `Content-Type: application/json`
- **Body**:

  ```json
  {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
        },{
        "shortDescription": "Emils Cheese Pizza",
        "price": "12.25"
        },{
        "shortDescription": "Knorr Creamy Chicken",
        "price": "1.26"
        },{
        "shortDescription": "Doritos Nacho Cheese",
        "price": "3.35"
        },{
        "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
        "price": "12.00"
        }
    ],
    "total": "35.35"
    }
  ```

#### Response

- **Status Code**: `200 OK`
- **Body**:

  ```json
  {
    "id": "3d41851f-caac-40b7-b799-582fc8ae8a66"
  }
  ```

#### Error Responses

- **Status Code**: `400 Bad Request`
- **Body**:

  ```json
  {
    "error": "Invalid request body"
  }
  ```

---

### 2. **Get Points for a Receipt**

- **Method**: `GET`
- **URL**: `/receipts/{id}/points`
- **Description**: Returns the points awarded for a specific receipt.

#### Request

- **Path Parameter**:
  - `id` (string): The unique ID of the receipt.
- **Example URL**:

  ```
  http://localhost:8080/receipts/3d41851f-caac-40b7-b799-582fc8ae8a66/points
  ```

#### Response

- **Status Code**: `200 OK`
- **Body**:

  ```json
  {
    "points": 28
  }
  ```

#### Error Responses

- **Status Code**: `404 Not Found`
- **Body**:

  ```json
  {
    "error": "Receipt not found"
  }
  ```
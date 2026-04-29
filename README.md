# URL Shortener (url_shorterner)

A fast, lightweight, and scalable URL shortener service built with Go (Golang). It uses **MongoDB** as the primary database for persistent storage and **Redis** as an in-memory cache to ensure rapid redirections.

## Features

- **Shorten URLs:** Convert long URLs into a compact, 8-character short code.
- **Fast Redirections:** Utilizes Redis caching to serve recently accessed short URLs with minimal latency.
- **Persistent Storage:** Stores all URL mappings in MongoDB.
- **Cache Fallback Strategy:** On cache miss, retrieves the URL from MongoDB and repopulates the Redis cache.
- **Standard Library HTTP:** Built using Go's lightweight and robust `net/http` standard library.

## Architecture

1. **Client** makes a `POST` request to shorten a URL.
2. **Service** generates an 8-character short code.
3. The mapping `(shortCode -> longURL)` is saved into **MongoDB** and cached in **Redis**.
4. When a **Client** visits the short URL, the service first checks **Redis**. 
   - If found (**Cache Hit**), it redirects immediately.
   - If not found (**Cache Miss**), it fetches from **MongoDB**, updates the **Redis** cache, and then redirects.

## Prerequisites

Before running the application, ensure you have the following installed and running:

- **Go** (Version 1.25.1 or newer)
- **MongoDB** (Running on `localhost:27017`)
- **Redis** (Running on `localhost:6379`)

## Setup & Installation

1. **Clone the repository** (if applicable) or navigate to the project directory:
   ```bash
   cd path/to/url_shorterner_m
   ```

2. **Download dependencies:**
   ```bash
   go mod tidy
   ```

3. **Ensure your databases are running.**
   - MongoDB should be accessible at `mongodb://localhost:27017`.
   - Redis should be accessible at `localhost:6379`.

## Running the Application

Start the server by running:

```bash
go run main.go
```

The server will start and listen on port `8080`.

```text
✅ Connected to MongoDB
📌 Connected DB: url_shortener
📂 Using Collection: urls
Redis connected ✅
Server running on port 8080...
```

## API Endpoints

### 1. Shorten a URL

- **Endpoint:** `/shorten`
- **Method:** `POST`
- **Content-Type:** `application/json`

**Request:**
```json
{
  "long_url": "https://www.google.com/search?q=golang+web+development"
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/aBcDeFgH"
}
```

### 2. Redirect to Original URL

- **Endpoint:** `/{shortCode}`
- **Method:** `GET`

**Description:** Navigating to the generated `short_url` (e.g., `http://localhost:8080/aBcDeFgH`) in your browser will automatically redirect you to the original long URL with an HTTP `302 Found` status.

## Testing with Postman

You can easily test the API endpoints using Postman:

### 1. Test POST `/shorten`
1. Open Postman and create a new request.
2. Set the HTTP method to **POST**.
3. Set the URL to `http://localhost:8080/shorten`.
4. Go to the **Body** tab, select **raw**, and choose **JSON** from the dropdown menu.
5. Paste the following payload:
   ```json
   {
     "long_url": "https://www.google.com/search?q=golang+web+development"
   }
   ```
6. Click **Send**. You should receive a `200 OK` status with the shortened URL in the response body.

### 2. Test GET `/{shortCode}`
1. Create a new request in Postman.
2. Set the HTTP method to **GET**.
3. Paste the generated short URL from the previous step (e.g., `http://localhost:8080/aBcDeFgH`).
4. Click **Send**.
*(Note: Postman automatically follows redirects by default, so you will see the HTML response of the original long URL. You can disable "Automatically follow redirects" in the Postman settings for that request if you want to inspect the `302 Found` response instead.)*

## Project Structure

```text
url_shorterner_m/
├── handlers/      # HTTP handlers (Controller layer)
│   └── url_handler.go
├── models/        # Database models & structures
│   └── url_model.go
├── services/      # Business logic (Shortening, caching strategies)
│   └── url_service.go
├── storage/       # Database & Cache connection initializations
│   ├── mongodb.go
│   └── redis.go
├── utils/         # Helper functions (e.g., short code generator)
│   └── generator.go
├── go.mod         # Go module dependencies
├── go.sum         # Go module checksums
└── main.go        # Application entry point
```

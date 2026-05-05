# URL Shortner (url_shortner)

A fast, lightweight, and scalable URL shortener service built with Go (Golang). It uses **MongoDB** as the primary database for persistent storage and **Redis** as an in-memory cache to ensure rapid redirections.

## Features

- **Shorten URLs:** Convert long URLs into a compact, 8-character short code using Base62 encoding combined with random padding to ensure unique and varied characters.
- **Fast Redirections:** Utilizes Redis caching to serve recently accessed short URLs with minimal latency.
- **Persistent Storage:** Stores all URL mappings in MongoDB.
- **Cache Fallback Strategy:** On cache miss, retrieves the URL from MongoDB and repopulates the Redis cache.
- **Rate Limiting:** Protects endpoints from abuse using an in-memory Token Bucket algorithm (per IP address).
- **Standard Library HTTP:** Built using Go's lightweight and robust `net/http` standard library.

## Architecture

1. **Client** makes a `POST` request to shorten a URL.
2. **Service** generates a unique ID using a Redis counter and encodes it into an 8-character short code using Base62 encoding with random padding.
3. The mapping `(shortCode -> longURL)` is saved into **MongoDB** and cached in **Redis**.
4. When a **Client** visits the short URL, the service first checks **Redis**. 
   - If found (**Cache Hit**), it redirects immediately.
   - If not found (**Cache Miss**), it fetches from **MongoDB**, updates the **Redis** cache, and then redirects.

## Prerequisites & Installation

Before running the application, ensure you have **Go**, **MongoDB**, and **Redis** installed and running on your system. 

### For Mac (macOS)
The easiest way to install dependencies on a Mac is using [Homebrew](https://brew.sh/).

1. **Install Go:**
   ```bash
   brew install go
   ```
2. **Install & Start Redis:**
   ```bash
   brew install redis
   brew services start redis
   ```
3. **Install & Start MongoDB:**
   ```bash
   brew tap mongodb/brew
   brew install mongodb-community
   brew services start mongodb-community
   ```

### For Windows
1. **Install Go:** Download and install from the [official Go website](https://go.dev/dl/).
2. **Install MongoDB:** Download and install the [MongoDB Community Server](https://www.mongodb.com/try/download/community). Ensure the MongoDB service is started.
3. **Install Redis:** The recommended way to run Redis on Windows is via [WSL (Windows Subsystem for Linux)](https://learn.microsoft.com/en-us/windows/wsl/install). 
   - Open your WSL terminal (e.g., Ubuntu).
   - Run `sudo apt update && sudo apt install redis-server`.
   - Start the service: `sudo service redis-server start`.
   - *(Alternatively, you can use Docker or Memurai).*

## Step-by-Step Guide to Run the Project

Follow these steps clearly to get the project running locally:

### Step 1: Clone the repository
Open your terminal and clone the project:
```bash
git clone https://github.com/PriyankaPatil064/url_shortner.git
cd url_shortener
```

### Step 2: Download dependencies
Fetch all required Go modules:
```bash
go mod tidy
```

### Step 3: Verify Databases are Running
Ensure your local databases are active:
- **MongoDB:** Accessible at `mongodb://localhost:27017`
- **Redis:** Accessible at `localhost:6379`

### Step 4: Run the Application
Start the Go server:
```bash
go run main.go
```

The server will start and listen on port `8080`. You should see the following success output in your terminal:

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

### 3. Viewing Data in MongoDB Compass
To verify that your URLs are successfully stored in the database:
1. Open **MongoDB Compass** and connect to `mongodb://localhost:27017`.
2. Navigate to the `url_shortener` database and the `urls` collection.
3. **Note:** MongoDB Compass does not auto-refresh. After shortening a new URL, you must click the **Refresh button** (the circular arrow `↻`) in Compass to see the newly inserted document.

## Project Structure

```text
url_shortener/
├── handlers/      # HTTP handlers (Controller layer)
│   └── url_handler.go
├── middleware/    # HTTP middlewares (e.g., Rate Limiter)
│   └── rate_limiter.go
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

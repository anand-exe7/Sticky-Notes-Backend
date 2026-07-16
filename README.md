# Sticky Notes API - Go Backend

A RESTful API built in **Go (Golang)** to manage and persist Sticky Notes data. 

## 🚀 Why Go instead of Node.js?

While the frontend runs on JavaScript (TypeScript/Lit), the backend was intentionally built using Go to take advantage of its superior server performance:

* **Compiled vs. Interpreted:** Unlike JavaScript which is interpreted at runtime by Node.js, Go compiles directly down to raw machine code. This means the server executes instructions instantly without the overhead of an interpreter.
* **Blazing Fast Execution:** Go handles HTTP requests, database connections, and JSON serialization significantly faster than standard Express.js backends. 
* **Cold Starts:** Because it's a single compiled binary, the server starts up incredibly fast, which is absolutely vital for cloud-hosted environments like Render.

## 🛠️ Tech Stack

* **Language:** Go (1.21+)
* **Router:** `go-chi/chi` (Extremely lightweight and fast)
* **Database:** MongoDB Atlas (Cloud)
* **Driver:** Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`)
* **Hosting:** Render (Web Service)

## 📦 API Endpoints

### Notes
* `GET /api/notes` - Fetch all active notes (excludes trash)
* `POST /api/notes` - Create a new note
* `PUT /api/notes/{id}` - Update an existing note
* `PUT /api/notes/{id}/pin` - Toggle the pinned status of a note
* `DELETE /api/notes/{id}` - Move a note to the recycle bin (soft delete)

### Recycle Bin (Trash)
* `GET /api/notes/trash` - Fetch all soft-deleted notes
* `PUT /api/notes/{id}/restore` - Restore a note from the trash back to the board
* `DELETE /api/notes/trash/{id}` - Permanently delete a note from the database forever

## 💻 Local Development

1. **Clone the repository.**
2. **Set up environment variables:**
   Create a `.env` file in the root directory and add your MongoDB connection string:
   ```env
   MONGO_URI=mongodb+srv://<username>:<password>@cluster.mongodb.net/?appName=Lost-and-found
   ```
3. **Run the server:**
   ```bash
   go run cmd/main.go
   ```
   The server will default to `http://localhost:8080`.

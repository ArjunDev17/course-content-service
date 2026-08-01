````markdown
# Course Content Service

The **Course Content Service** is a microservice responsible for managing courses, modules, and lessons in the EdTech platform.  
It provides **REST** and **gRPC APIs** to create, update, retrieve, and manage study materials.  
This service uses **MongoDB** as its primary data store due to the flexible and schema-less nature of course content.

---

## 🚀 Features
- Manage Courses, Modules, and Lessons
- Store metadata for study materials (title, description, tags, difficulty, etc.)
- RESTful APIs (Swagger/OpenAPI support)
- gRPC endpoints for inter-service communication
- Structured logging and middleware support
- Unit tests, integration tests, and load tests
- Dockerized & ready for Kubernetes deployment

---

## 🛠️ Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gin / Fiber (REST APIs)
- **Database:** MongoDB
- **gRPC:** Protobuf for inter-service contracts
- **Logging:** Structured logging (Zap/Logrus)
- **Testing:** Go test, mocks, load tests
- **Containerization:** Docker, Helm charts (for deployment)

---

## 📂 Project Structure
```bash
course-content-service/
├── apis/                       # Swagger/OpenAPI specs
├── cmd/
│   └── server/                 # Entry point (main.go)
├── config/
│   ├── config.yaml             # Default configs
│   └── loader.go               # Loads env/config
├── env/                        # Env files (.env, dev.env, prod.env)
├── handler/                    # API layer (controllers)
│   ├── http/                   # REST controllers
│   └── grpc/                   # gRPC handlers
├── model/                      # Domain models (Course, Module, Lesson)
├── repository/                 # Database layer
│   ├── mongo/                  # MongoDB implementation
│   └── mock/                   # Mocks for testing
├── service/                    # Business logic layer
│   ├── course/                 # Course service
│   ├── lesson/                 # Lesson service
│   └── test/                   # Unit tests
├── server/                     # Server bootstrap
│   ├── api/                    # API setup (Gin/Fiber routes)
│   └── props/                  # Server props (ports, timeouts)
├── middleware/                 # Logging, tracing, auth, rate limiting
├── pkg/                        # Shared package (db conn, cache conn)
│   ├── db/
│   │   └── mongo.go            # Mongo connection
│   └── logger/                 # Structured logging setup
├── scripts/                    # Local scripts (migrations, seeding)
├── test/                       # Integration & load tests
│   ├── api/
│   └── load/
├── swagger/                    # Swagger spec + docs
├── Dockerfile
├── Makefile
└── go.mod
````

---

## ⚡ Getting Started

### Prerequisites

* [Go 1.21+](https://go.dev/dl/)
* [MongoDB](https://www.mongodb.com/docs/manual/installation/)
* [Docker](https://docs.docker.com/get-docker/)
* [Make](https://www.gnu.org/software/make/)

### 1️⃣ Clone Repository

```bash
git clone https://github.com/your-org/course-content-service.git
cd course-content-service
```

### 2️⃣ Setup Environment

Create a `.env` file or use existing `env/dev.env`:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=course_service
PORT=8080
```

### 3️⃣ Run Locally

```bash
go mod tidy
go run cmd/server/main.go
```

### 4️⃣ Run with Docker

```bash
docker build -t course-content-service:latest .
docker run -p 8080:8080 --env-file=env/dev.env course-content-service:latest
```

---

## 📜 API Documentation

* REST APIs documented with **Swagger** → `http://localhost:8080/swagger/index.html`
* gRPC APIs defined in `/apis/proto/course.proto`

---

## 🧪 Testing

```bash
make test
```

* **Unit Tests** → `service/test/`
* **Integration Tests** → `test/api/`
* **Load Tests** → `test/load/`

---

## 📦 Deployment

* Containerized using Docker
* Kubernetes-ready via Helm charts (`/helm/course-content/`)
* Supports CI/CD with GitHub Actions / Jenkins pipelines

---

## 📌 Future Enhancements

* Add content versioning (track updates to lessons/modules)
* Integrate with Elasticsearch for fast course search
* Add GraphQL API layer
* Caching with Redis for frequently accessed course data

---

## 👨‍💻 Maintainers

* Arjun Singh 
* Team/ArjunDev17


# Kafka Learning Roadmap

This repository is being enhanced step by step into a production-ready Event-Driven microservice.

## Phase 1 - Kafka Producer ✅
- Create Course API
- Store Course in MongoDB
- Publish `course.created` event to Kafka
- Understand Producers, Topics, and Event Contracts

## Phase 2
- Notification Service (Kafka Consumer)

## Phase 3
- Subscriber Management

## Phase 4
- Notification Worker Pool

## Phase 5
- Retry Topics

## Phase 6
- Dead Letter Topic (DLT)

## Phase 7
- Consumer Groups

## Phase 8
- Partitions

## Phase 9
- Ordering Guarantees

## Phase 10
- Production Deployment
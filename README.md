# Binar - Backend Developer Home Assignment Test

# Tech Stack

| Technology | Purpose                                   | Version | URL                                                        |
|------------|-------------------------------------------|---------|------------------------------------------------------------|
| Go         | Main Programming Language                 | 1.22.0  | [https://golang.org/](https://golang.org/)                 |
| PostgreSQL | Relational Database Management System     | 15      | [https://www.postgresql.org/](https://www.postgresql.org/) |
| RabbitMQ   | Message Queue for Asynchronous Processing | 3.12    | [https://www.rabbitmq.com/](https://www.rabbitmq.com/)     |

| Library/Framework | Purpose                 | Version  |
|-------------------|-------------------------|----------|
| Echo              | Web Framework           | v4.12.0  |
| GORM              | ORM Library             | v1.25.11 |
| Google Wire       | Dependency Injection    | v0.6.0   |
| Zap               | Logging                 | v1.27.0  |
| Cobra             | CLI Application Library | v1.8.0   |

# Architecture

![image](https://github.com/user-attachments/assets/ca5ced75-a22c-410f-ac8b-ccddba377c06)

### Explanation :

Before i start an explanation of how this project implement clean architecture,
i would inform that i take reference about clean architecture from :
- The Creator : https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- My Mentor : https://github.com/khannedy/golang-clean-architecture/tree/main
- Random Guy I Pick From Internet : https://github.com/bxcodec/go-clean-arch
- Bard AI : https://gemini.google.com/

Okay now i'll start explain what i think about clean architecture.
In my first opinion before fetch all reference just base on my experience
and random diagram on internet, i though clean architecture just a concept of separation
from MVC (Model -> View -> Controller). But after a long time search, i relize
it not just MVC Concept.

And then what i get about after find about clean architecture ?
and how my project applying them ? Here is it

Separation of Concerns :
I clearly separate different components into distinct layers
in this case `cmd -> internal -> pkg`

1. Cmd : Contains main entry points for different applications (API, CLI, worker)
- API For this test case is using HTTP Protocol
- CLI For this test case is using cobra and migration purpose
- Worker For this test case is using RabbitMQ as part of Broadcast Message Queue Processing

2. Internal : My core application logic
- This is the place where i hold the whole main business logic, and i'll explain it later

3. pkg : The place where i holds shared packages and utilities

Domain Centric :
I create domain free of any library or other, it is one of most important things
from clean architecture. Domain should be free and not depend on any dependency

Dependency Inversion : 
I implement interfaces (ex : on my repo and service)
for adapting the Dependency Inversion Principle, allowing for loose coupling between layers.

Clear Dependencies:
Yeah we know it just an MVC like on Laravel :)
but here is it, `delivery depends on service, service depends on repository, and repository implements domain interfaces`.

Separation of Infrastructure :
Yap, i not inject all infra into service or repo, just what they need

Clean API Layer:
The separation of HTTP handlers and payload structures in the delivery layer keeps the API clean and decoupled from the core logic.

Repository Pattern :
We can't far away from Repository Pattern, because my memorable lesson from PZN it's Repository Pattern :)
Repository Pattern is the part of Clean Architecture too, it implements database independence, and make it easy
to mock test.

### Flow (On Adapting Clean Architecture) Explanation
Generally, i separate this project into 4 Layer.
`Delivery Layer` `Service Layer` `Domain Layer` `Repository Layer`

1. Every request will come from delivery (For this case only HTTP, it suitable to gRPC and other)
as an external system. It will be wrapped into `payload -> request.go` and then i transform it into
main domain entity
2. Delivery call service, and then service will call repository.
3. Call between repo and service connected via domain
4. And then it will reverse `Repo -> Domain -> Service -> Domain -> DTO(Response Wrapper) -> Handler -> Response`

# Usage
Follow these steps carefully to set up and run the application:

1. **Set Docker Environment**
   ```
   make set-docker-env
   ```
   This updates the configuration to use Docker container names for service connections.

2. **Build and Run the Application**
   ```
   make build-app
   ```
   This builds the Docker images and starts the containers.

3. **Set Migration Environment**
   ```
   make set-migration-env
   ```
   This command updates the configuration to use localhost for database connections.

4. **Run Database Migrations**
   ```
   make migrate-up
   ```
   This applies all pending database migrations.

5. **Seed the Database**
   ```
   make seed
   ```
   This populates the database with initial data.

### Additional Commands

- **Clean Application (Use with Caution)**
  ```
  make clean-app
  ```
  **WARNING**: This command removes all Docker images, volumes, and prunes the system. Use only when necessary and be aware that it will delete all related Docker resources.

### Best Practices

1. Always run `set-migration-env` before performing database operations locally.
2. Ensure all migrations are up to date before seeding the database.
3. Switch to `set-docker-env` before building and running the application in Docker.
4. Regularly backup your data before running migrations or using `clean-app`.
5. Review the `app.yaml` file after environment switches to ensure correct configuration.
6. Use `clean-app` sparingly and only when you need a completely fresh environment.

# ERD
![image](https://github.com/user-attachments/assets/e7f37c1e-06ac-4690-bf0e-1bb784bdb750)

# Lack
- Import Naming: The project uses GoLand's default auto-import feature, which may result in numerically suffixed package aliases (e.g., domain1, domain2).
- Unit Tests: Due to time constraints, comprehensive unit tests are not included in the current version.
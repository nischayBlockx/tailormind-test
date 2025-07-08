This Go microservice enables the following functionality:

Add, update, delete student records

Generate a downloadable PDF report for a student by their ID

It connects to a PostgreSQL database and serves REST APIs using the Gin framework.


go-service/
├── api/handler            # HTTP handlers
├── internal/models        # Data structures
├── internal/repository    # Database interaction logic
├── internal/service       # Logic layer
├── pkg/storage/           # PostgreSQL client
├── scripts/migrations     # SQL migration files
├── cmd/main/main.go       # Entry point
├── go.mod / go.sum
├── .env                   # Environment variables

**Features**

#RESTful API to:

POST /api/v1/students - Add student

PUT /api/v1/students - Update student

DELETE /api/v1/students/:id - Delete student

GET /api/v1/students/:id/report - Generate and download student PDF report



Create your DB manually in pgAdmin or CLI:

CREATE DATABASE student_db;

Then apply migrations:

```bash
export DATABASE_URL=postgres://<user>@localhost:5432/student_db?sslmode=disable
migrate -path scripts/migrations -database "$DATABASE_URL" up
```



**Running the Service**
```bash
go run cmd/main/main.go

```
Server starts on port :8081.




**Sample cURL Commands**

*Add Student*

```bash
curl -X POST http://localhost:8081/api/v1/students \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "e4eaaaf2-d142-11e1-b3e4-080027620cdd",
    "name": "SkillTest",
    "email": "skilltest@example.com",
    "phone": "1234567890",
    "gender": "Male",
    "dob": "2000-01-01",
    "class": "10",
    "section": "A",
    "roll": "25",
    "fatherName": "Father SkillTest",
    "fatherPhone": "9876543210",
    "motherName": "Mother SkillTest",
    "motherPhone": "9876543211",
    "guardianName": "Uncle SkillTest",
    "guardianPhone": "8888888888",
    "relationOfGuardian": "Uncle",
    "currentAddress": "123 Main St",
    "permanentAddress": "123 Main St",
    "admissionDate": "2016-06-15",
    "reporterName": "Admin"
  }'
```

  **Get PDF Report**
```bash
  curl -X GET http://localhost:8081/api/v1/students/e4eaaaf2-d142-11e1-b3e4-080027620cdd/report --output student_report.pdf
```

  NOte: Student ID is passed manually now. You can use github.com/google/uuid to generate UUIDs in backend before insert to ensure stable and unique identifiers.

  
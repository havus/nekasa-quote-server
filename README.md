# nekasa-quote-server

## Getting Started
Welcome to the nekasa-quote-server project! This guide will help you get your development environment set up so you can start working on the project.

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations


### Setup

#### 1. Clone the repository

Start by cloning the repository to your local machine:
```bash
> git clone https://github.com/havus/nekasa-quote-server.git
> cd nekasa-quote-server
```

#### 2. Install dependencies

Navigate to the project directory and install the required Go modules:
```bash
> go mod download
```

#### Database

We are using [golang-migrate/migrate](https://github.com/golang-migrate/migrate) to handle migrations.

**Install on your local machine**: [Official installation link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)<br>
```bash
> go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

**Create MySQL database**:<br>
```bash
> mysql -u root -p
mysql> CREATE DATABASE nekasa_quotes;
```


**Run migration**:<br>
```bash
> $GOPATH/bin/migrate -database "mysql://root:password@tcp(localhost:3306)/nekasa_quotes" -path internal/infrastructure/database/migrations up
```

**Add migration**:<br>
```bash
> $GOPATH/bin/migrate create -ext sql -dir internal/infrastructure/database/migrations create_table_users
```

**Inspect schema migration**:<br>
```sql
SELECT
  TABLE_NAME, COLLATION_NAME,
  COLUMN_NAME AS COL_NAME, COLUMN_TYPE AS COL_TYPE, COLUMN_KEY AS COL_KEY, COLUMN_DEFAULT AS COL_DEFAULT,
  IS_NULLABLE, CHARACTER_SET_NAME AS CHAR_SET_NAME, CHARACTER_MAXIMUM_LENGTH AS CHAR_MAX_LEN,
  NUMERIC_PRECISION AS NUM_PRECISION, NUMERIC_SCALE AS NUM_SCALE
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_SCHEMA='nekasa_quotes' AND TABLE_NAME='users';
```

#### 3. Configure the environment
Create a `.env` file in the root directory of the project and add your environment-specific variables. An example `.env` file might look like `.env.template`.


#### 4. Run the application
Once everything is set up, you can run the application using the following command:
```bash
> go run cmd/nekasa-quote/main.go
```

#### 5. Running tests
To run the tests, use the following command:
```bash
> go test ./...
```
This will execute all the tests in the project and display the results.


### Definitions
DTO: Data Transfer Object


### Contributing

We welcome contributions to the nekasa-quote-server project. If you have an idea for a new feature or have found a bug, please open an issue or submit a pull request.

### License

This project is licensed under.....

# Blog Aggregator (Go)

A blog aggregator written in **Go** that collects and stores blog data using **PostgreSQL**.

## Prerequisites

Before you begin, make sure you have the following installed on your system:

* **Go toolchain** (latest stable version recommended)
* **PostgreSQL** (version 13+ recommended)


## Installation

### 1. Install the Go toolchain

Download and install Go from the official website:

```
https://go.dev/dl/
```

Verify installation:

```
go version
```

### 2. Install PostgreSQL

Install **Postgres** using the method appropriate for your operating system.

sudo apt install postgresql postgresql-contrib --- Linux
sudo apt install postgresql postgresql-contrib --- MacOS

After installation, ensure it is available in your PATH and can be started from the command line.

### 2. Install Goose

Goose is used to run the migration but you are free to use whatever tool 
you prefer

Install Goose : go install github.com/pressly/goose/v3/cmd/goose@latest

Run goose -version to make sure it's installed correctly.


### 3. Create the PostgreSQL database

Ensure PostgreSQL is running, then create a database:

```
Mac: psql postgres
Linux: sudo -u postgres psql
```

Inside the psql shell:

```
CREATE DATABASE gator;
```

Exit psql:

```
\q
```


### 4. Create the configuration file

Create a configuration file named **`.gatorconfig.json`** in your **home directory**:

```
~/.gatorconfig.json
```

This file is used to configure the PostgreSQL database connection and other application settings.

Example structure:

```
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

> Make sure the file exists and contains a valid database URL before running the application.



### 5. Clone the repository

Clone this project from GitHub:

```
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
```


### 6. Run database migrations

From the project root directory, run the migration files to initialize the database schema:

```
psql "<database-url>" -f migrations/001_init.sql
```


## Building and Installing

Install the application using Go:

```
go install
```

Or build a local binary:

```
go build
```

Run the compiled binary:

```
./<binary-name>
```

## Running the Project

If installed via `go install`, run:

```
<binary-name>
```

If running from source:

```
go run .
``
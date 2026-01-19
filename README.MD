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

**Postgres install commands**

Linux (Debian/Ubuntu):

```bash
sudo apt install postgresql postgresql-contrib
```

MacOS:

```bash
brew install postgresql@15
```

After installation, ensure it is available in your PATH and can be started from the command line.

### 2. Install Goose

Goose is used to run the migration but you are free to use whatever tool 
you prefer

Install Goose :

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

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
goose -dir ./sql/schema postgres "$DATABASE_URL" up
```

## Building and Installing

Install the application using Go:

```bash
go install github.com/flogit2161/BlogAggregator@latest
```

If installed via `go install`, run:

```
BlogAggregator register "username"
```

If running from source:

```
go run . register "username"
```

## How to use

Run register with your username and start adding feeds to follow, you can then browse your feed and aggregate them to get their RSS

## Cmds List

- register - creates and logs in user
- login - log in with user created beforehand when reusing program
- users - check all users available
- feed - get feeds for user
- addfeed - add a feed
- follow - follow a feed
- following - check followed feeds

- unfollow - unfollow feed
- agg - aggregates followed feeds
- browse - browse aggregated feeds

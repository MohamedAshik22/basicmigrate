# Basic Migrate Go App

This is a simple Go app that manages database migrations and includes a basic HTTP server. The app provides commands for applying migrations, rolling them back, undoing specific migrations, and running the application.

## Features
- **Database Migrations**: Apply, rollback, or undo specific migrations.
- **Migration History Recording**: Keep track of migration history.
- **Basic HTTP Server**: Start a simple server on `localhost:8080`.

## Commands

The application has four main commands that you can run:

### 1. Run All Migrations

Run the migrations by applying all pending migrations in sequence.

```bash
go run main.go --up
```

### 2. Rollback All Migrations

Undo all the applied migrations and bring the database back to its initial state.

```bash
go run main.go --down
```

### 3. Undo a Specific Migration

Undo a specific migration by providing its ID. For example, if you want to undo the migration `0001_create_users`:

```bash
go run main.go --undo 0001_create_users
```

### 4. Run the App (Start the Server)

If no migration-related flags are provided, the app will start an HTTP server on `localhost:8080`.

```bash
go run main.go
```

Once the server is running, you can visit [http://localhost:8080](http://localhost:8080) to see a simple "Hello, World!" response.


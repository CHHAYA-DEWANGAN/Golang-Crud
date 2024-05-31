# Golang-Crud

To Run GOlang Project
- go run main.go

Create a migration directory:

Create a directory within your project to store migration files. A common convention is db/migrations.
Generate a new migration:

Use the migrate create command to generate a new migration file with an appropriate version number:

- **migrate create -dir db/migrations -ext .sql NewTableName  # Replace "NewTableName" with your actual table name**

This command creates two files within the specified directory:
YYYYMMDD_HHMMSS_NewTableName_up.sql: Contains SQL statements to create the table (up migration).
YYYYMMDD_HHMMSS_NewTableName_down.sql: Contains SQL statements to drop the table (down migration).
Edit the up.sql file:

Open the up.sql file and write the SQL code to create your table with its columns, data types, constraints, and indexes:
SQL
CREATE TABLE IF NOT EXISTS your_table_name (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
Use code with caution.
content_copy
Replace your_table_name with your desired name.
Customize the columns, data types, constraints, and indexes as needed for your specific table structure.
(Optional) Edit the down.sql file:

The down.sql file is automatically generated to drop the table. You can modify it if you have specific requirements for rolling back the migration.
Run migrations:

Apply the migrations to your database:
Bash
**migrate -path db/migrations -database "your_database_url" up**
Use code with caution.
content_copy
Replace "your_database_url" with the connection string for your database (refer to your database driver's documentation for URL format).

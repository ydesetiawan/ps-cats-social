# ps-cats-social
Project Sprint for Cats Social

## 🔍Requirements

This application requires the following:

- Golang: https://go.dev/dl/
- PostgreSQL: https://www.postgresql.org/download/
- Golang Migrate: https://github.com/golang-migrate/migrate
- Project Requerement : https://openidea-projectsprint.notion.site/Cats-Social-9e7639a6a68748c38c67f81d9ab3c769

## 🛠️Installation

To install the boilerplate, follow these steps:

1. Make sure you have Golang, PostgreSQL, and Golang Migrate installed and configured on your system.

2. Clone this repository:

   ```bash
   git clone https://github.com/ydesetiawan/ps-cats-social.git
   ```

3. Navigate to the project directory:

   ```bash
   cd ps-cats-social
   ```

4. Run the following command to install dependencies:
   ```bash
   go mod download
   ```

## 🚀Usage

1. **Setting Up Environment Variables**

   Before starting the application, you need to set up the following environment variables:

    - `DB_NAME`: Name of your PostgreSQL database
    - `DB_PORT`: Port of your PostgreSQL database (default: 5432)
    - `DB_HOST`: Hostname or IP address of your PostgreSQL server
    - `DB_USERNAME`: Username for your PostgreSQL database
    - `DB_PASSWORD`: Password for your PostgreSQL database
    - `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)
    - `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)
    - `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)

2. **Database Migrations**

   Cats Social uses Golang Migrate to manage database schema changes. Here's how to work with migrations:

    - Apply migrations to the database:

      ```bash
      make migration_setup
      make migration_up
      ```

3. **Running the Application**

   Once you have set up the environment variables, you can start the application by running:

   ```bash
   go run cmd/api/main.go
   ```

   This will start the Cats Social application on the default port (usually 8080). You can access the application in your web browser at http://localhost:8080

## ⚙️Configuration

The application uses environment variables for configuration. You can configure the database connection, JWT secret, and bcrypt salt by setting the following environment variables:

- Refer to the [Usage](#usage) section for a detailed explanation of each environment variable.

## 💾Database Migration

Database migration must use golang-migrate as a tool to manage database migration

1. Direct your terminal to your project folder first
2. Create migration using MakeFIle

   ```bash
    make migration_setup
    make migration_up
    make migration_down
    make migration_fix
   ```




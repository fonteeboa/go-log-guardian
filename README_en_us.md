<h1 align="center">Log Guardian</h1>

<p align="center"><strong>Open-Source Logger Assistant</strong></p>

[![link to Go version](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-log-guardian)](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-log-guardian)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/fonteeBoa/go-log-guardian)
[![go report card](https://goreportcard.com/badge/github.com/fonteeBoa/go-log-guardian "go report card")](https://goreportcard.com/report/github.com/fonteeBoa/go-log-guardian)

README em Português: Acesse a versão em Português da documentação do projeto [aqui](https://github.com/fonteeboa/go-log-guardian/blob/master/README.md)

Log Guardian is a library developed to standardize and efficiently manage logs in systems, emphasizing the importance of translation for the end user. With the flexibility to handle different types of logs, this library provides a consistent structure for generating and managing general and specific logs, allowing integration with various systems.

<h2 align="center"><strong>Features</strong></h2>

🔹 Log Standardization: Log Guardian offers a unified framework for different types of logs, from function logs and database operation logs to request logs.

🔹 Flexible Configuration: Allows easy integration with different systems, enabling customization and configuration of logs according to specific needs.

🔹 Database Connection: In addition to log management, Log Guardian can integrate with different types of databases such as PostgreSQL, MySQL, SQLite, and MongoDB. Configuration is simple, using environment variables to specify connection details.

🔹 Automatic Log Insertion: When correctly configured with environment variables, Log Guardian can automatically insert logs into the specified database.

<h2 align="center"><strong>Usage</strong></h2>

Log Guardian is flexible and adapts to the configuration of the environment in which it is run. If the required environment variables are not configured, Log Guardian can still return the specific log template for manual insertion into the database.

If the environment variables are correctly set with the details of the desired database, Log Guardian can automatically connect to the specified database and insert logs directly into the corresponding table. It returns a boolean value indicating success or failure in inserting data into the database.

This flexibility allows easy integration and use of Log Guardian in different configuration scenarios, whether just providing log templates for manual insertion or performing automatic insertions into the configured database.

It is advisable to refer to the Configuration section for details on the necessary environment variables for a complete setup of Log Guardian.

<h2 align="center"><strong>Configuration</strong></h2>

Log Guardian uses environment variables to configure its database operations, including database connection settings and other essential configurations. Here is the list of available environment variables:

<h4 align="center"><strong>Relational Database</strong></h4>

#### PostgreSQL
```
POSTGRES_HOST: Defines the host address for PostgreSQL.
POSTGRES_EXTERNAL_PORT: Specifies the external port for PostgreSQL.
POSTGRES_USER: Username for authentication in PostgreSQL.
POSTGRES_PASSWORD: Password for authentication in PostgreSQL.
POSTGRES_DB: Name of the PostgreSQL database to be used.
```
#### MySQL
```
MYSQL_HOST: Defines the host address for MySQL.
MYSQL_PORT: Specifies the port for MySQL.
MYSQL_USER: Username for authentication in MySQL.
MYSQL_PASSWORD: Password for authentication in MySQL.
MYSQL_DBNAME: Name of the MySQL database to be used.
```
#### SQLite
```
SQLITE_PATH: Path to the SQLite file if it is the chosen database.
```

<h4 align="center"><strong>NoSQL Database</strong></h4>

#### MongoDB
```
MONGODB_URI: Defines the connection URI for MongoDB.
MONGODB_DBNAME: Name of the MongoDB database to be used.
```

<h4 align="center"><strong>General Configuration</strong></h4>

```
LOG_GUARDIAN_DATABASE_TYPE: Specifies the type of database to be used by Log Guardian (Values: sqlite, postgres, mysql, mongodb).
```

<h4 align="center"><strong>Notes</strong></h4>

To utilize the automatic functions of go-log-guardian, using the LOG_GUARDIAN_DATABASE_TYPE variable is mandatory, as some validations are performed based on this variable before calling the insertion routines.

Ensure to provide valid and correct values for each of these environment variables. This ensures proper connection and functioning of Log Guardian with the desired database.

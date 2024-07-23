# Setup Project

1. download go from here <https://go.dev/dl/go1.22.2.windows-amd64.msi>
2. Download database driver [dbeaver](https://dbeaver.io/download/) and database tools [jdbc](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads) and install it
3. pull this project from github and locate it in your go workspace where the src folder is located
4. open the project in your favorite IDE
5. open the terminal and run the following command to install the dependencies
    ```bash
    go mod download #same as npm install
    ```
    if you want to clean the package that not used in the project use the following command

    ```bash
    go mod tidy #same as npm install bur remove the unused package
    ```
6. try to install all package
    ```
    go install
    ```
7. create a new file named .env in the root of the project and copy the content of .env.example to .env
8. **IMPORTANT** If u're new on this project try to migrate the database first
    ```bash
    go run database/database.go -migrate
    ```

    if you want to drop all table use the following command
    ```bash
    go run database/database.go -migrate:fresh
    ```

    if you want to seed the database use the following command
    ```bash
    go run database/database.go -seed
    ```
9. run the project using the following command

- Run Project with CompileDaemon **(RECOMMENDED)**
    ```bash
    CompileDaemon -command="./skripsi-be"
    ```

- Run Project
    ```bash
    go run main.go 
    ```

# Package of this project

1. Framework: [Gin](https://gin-gonic.com/docs/quickstart/)
2. Golang ORM: [GORM Postgres](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
3. Database:
    - PostgreeSQL main: [postgree](https://www.postgresql.org/download/)
    - PostgreeSQL tools: [dbeaver](https://dbeaver.io/download/)
    - PostgreeSQL driver: [jdbc](https://jdbc.postgresql.org/download/)
4. env: [godotenv](https://github.com/joho/godotenv)
5. compiler: go with [CompileDaemon](https://github.com/githubnemo/CompileDaemon)

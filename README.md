# Life Saver Project
The intent of this project is for the assessment module provided by Life Saver.

The following APIs are provided to perform basic CRUD operations on user (Hosted at :8081)
Endpoint - /user

- Create User   GET
- Get User      POST    /:userId
- Update User   UPDATE  /:userId
- Delete User   DELETE  /userId

####

Stack :

Application : GoLang
Database : Postgres SQL

Continious Integration : Docker
Continious Deployment : GitHub Actions

####

External Dependencies Used :

Gin Gonic : Http Helper
Simplified way to write endpoints using Http Gin Context, very suitable for using JWT for Authorization
https://github.com/gin-gonic/gin

Postgres Connection - https://github.com/lib/pq

Flyway Migration : github.com/golang-migrate/migrate/v4

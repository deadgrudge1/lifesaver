# Life Saver Project
The intent of this project is for the assessment module provided by Life Saver.

The following APIs are provided to perform basic CRUD operations on user (Hosted at :8081)
Endpoint - /user

- Create User   
  POST    /user/:userId
  
- Get User      
  GET     /user
  
- Update User   
  UPDATE  /user/:userId
  
- Delete User   
  DELETE  /user/userId

####

Stack :

Application : GoLang
Database : Postgres SQL

Continious Integration : Docker
Continious Deployment : GitHub Actions

####

External Dependencies Used :

Gin Gonic : Http Helper
Simplified way to write endpoints using Http Gin Context, very suitable for using JWT for Authentication
https://github.com/gin-gonic/gin

Postgres Connection - https://github.com/lib/pq

Flyway Migration : github.com/golang-migrate/migrate/v4

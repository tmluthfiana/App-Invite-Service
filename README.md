
Once you run the API server, it will automatically create two tables in the database and add a default admin user (username: userPulse, password: qwerty1234)

1. **admin** table will contain all the admin credentials.
2. **client_tokens** table will contain all the invitation tokens.



Setup
--


1. Git clone this repo.
2. Create a mysql database 

    ```
    CREATE DATABASE invitation-app
    ```

3. Run the following in the terminal

    ```
    export mysql_users_password="your database password"
    export mysql_users_schema="your database name"
    export mysql_users_host="localhost"
    export mysql_users_username="your database username"
    ```

4. In the project root folder, run the following to start the API server

    ```
    go mod tidy
	go run .
    ```
    
Generating an admin token for authentication
--    

To generate a token, run the following in Postman

    ```
    POST http://localhost:8040/login
    ```
    
     {
        "username":"userPulse",
        "password":"qwerty1234"
     }


Generating an invite token
--    

To generate an invite token, run the following in Postman

```
GET http://localhost:8040/requesttoken
```

In Header, you will need to add the Authorization with Bearer the generated token

```
{
    "id": 22,
    "token": "9nd9TX",
    "status": "Active"
}
```

A new record will be added in client_token table. The generated token will be converted into md5 hash for security purposes.

Validating an invite token
--    

To validate an invite token, run the following in Postman

```
GET http://localhost:8040/validateToken
```

Here is a sample json request

```
{
    "token":"eAvBGO"
}
```

Listing all invite tokens
-- 
To list all the generated invite tokens, run the following in Postman

```
GET http://localhost:8040/getalltokens
```
In Header, you will need to add the Authorization with the generated token

Swagger
-- 

To view the Swagger documentation, open a browser and access this url

```
http://localhost:8040/swagger/index.html#/

```












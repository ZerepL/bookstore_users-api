# bookstore_users-api
 
A Golang microservice that will be used as an Users API to handle and store differents users.

## Architecture

![Arch](./miscs/arch.jpg)

## Requirements

### Standalone

* [MySQL](https://www.mysql.com/)

## Before running

This app collect some data from env, bellow you can find a list of all vars and their values:

|       Variable       |   Description   |
|:--------------------:|:---------------:|
| mysql_users_username |  Database User  |
| mysql_users_password |  Database Pass  |
|   mysql_users_host   |  Database URL   |
|  mysql_users_schema  | Database Schema |

Inside this folder you will find a file called *migration.sql*, run it in your database.

## Running


``` shell
go run *.go
```

## API

Inside docs you can find a swagger file with all details about API.

### /users/

| Function      | Method | Path  | Expected |
|:-------------:|:------:|:-----:|:--------:|
|  Get user     |  GET   | /{id} |  Int     |
|  Create new   |  POST  | /     | JSON     |
|  Update user  |  PUT   | /{id} | JSON     |
|  Patch user   |  Patch | /{id} | JSON     |
|  Delete user  | DELETE | /{id} | Int      |

#### Create user payload

``` json
{
    "first_name": "x",
    "last_name": "x",
    "email": "x@email.com",
    "password": "batatinha-1"
}
```

#### Update user payload

``` json
{
    "first_name": "x",
    "last_name": "x",
    "email": "x@email.com"
}
```

#### Patch user payload

It can be any field of update payload.

### /internal/

| Function      | Method | Path                          | Expected |
|:-------------:|:------:|:-----------------------------:|:--------:|
| Search status |  GET   | /users/search?status={status} |  String  |

#### Status available

* Active

## TODO

* Create Docker
* Add tests
* Create deployment for K8s

## Credits

This microservice is based in this [course](https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/)

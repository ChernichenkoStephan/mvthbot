# mvthbot

<img src="assets/robot.jpeg" alt="mvthbot header picture" width="200"/>

A service-bot with users, variables, and calculation history support to solve equations in text format, to help with calculations in chats. With an open API for interacting with other services and registering via telegram.

## Roadmap

- [ ] mvthbot MVP
    - [x] Equation solving
        - [x] Reverse Polish notation conversion
        - [x] Variables support
    - [ ] REST API (using Fiber)
        - [x] Middleware (logging, parsing, etc.)
        - [x] Auth (using JWT)
        - [ ] OpenAPI (using Swagger)
    - [x] Telegram bot (using tucnak/telebot)
        - [x] Broadcasting
        - [x] Middleware (logging, parsing etc.)
        - [x] Commands set
    - [x] error handling (using emperror.dev/emperror)
    - [x] logging (using uber.org/zap)
    - [ ] Database
        - [ ] PostgreSQL (using github.com/jackc/pgx)
        - [x] Migrations (using golang-migrate)
    - [ ] configuration (using spf13/viper)
    - [ ] Deploy (using Docker and Makefile)
- [ ] mvthbot 1.0
    - [ ] health checks (using AppsFlyer/go-sundheit)
    - [ ] graceful restart (using cloudflare/tableflip) and shutdown
    - [ ] metrics and tracing using Prometheus and Jaeger (via OpenCensus)
    - [ ] support for multiple server/daemon instances (using oklog/run)
- [ ] mvthbot 2.0 (rebuild to microservices)
    - [ ] Services
        - [ ] Solving
        - [ ] User
        - [ ] Bot
        - [ ] Gateway (Nginx & REST & grpc)
        - [ ] Auth
    - [ ] messaging (using ThreeDotsLabs/watermill)
- [ ] mvthbot 3.0
    - [ ] Proper testing
        - [ ] Tests refactoring
        - [ ] Coverage increase
        - [ ] Load testing
    - [ ] CI/CD
        - [ ] Docker-compose
        - [ ] Kubernetes
        - [ ] Ansible
        - [ ] Jenkins
        - [ ] GitHub Actions
    - [ ] Dashboard (using GoAdminGroup/go-admin)
    - [ ] Rewrite DB with codegen

## Project structure

```
📦 mvthbot
├─ api
│  └─ swagger.yaml
├─ cmd
│  └─ mvthbot
│     ├─ api.go
│     ├─ app.go
│     ├─ bot.go
│     ├─ main.go
│     └─ storage.go
├─ configs
├─ internal
└─ tools
```
©generated by [Project Tree Generator](https://woochanleee.github.io/project-tree-generator)

> Table with desctiprion

## Prerequisites

* Golang (version `unknown`)
* Docker (version `unknown`)
* PostgreSQL (version `unknown`)

## Installation

> Path

## Configuration

> Table with desctiprion

## API

### Available endpoints

#### Auth
| Type | Path | Variables | Description |
| --- | --- | --- | --- |
| POST | api/v1/auth/login/ | none | getting JWT to access |
| POST | api/v1/auth/logout/ | none | dismissing JWT |

> **User ID defines from JWT cookie**

#### Solve
| Type | Path | Variables | Description |
| --- | --- | --- | --- |
| POST | api/v1/solve/:equation/ | *:equation* equation to solve coded in LF | returns result, no variable created |
| POST | api/v1/solve/ | none | returns result of equations in body |

#### Variables
| Type | Path | Variables | Description |
| --- | --- | --- | --- |
| POST | api/v1/variables/:name/:equation | *:name*  user variable name to set *:equation* equation to solve | setting variable to result of equation |
| POST | api/v1/variables | none | Setting  user variables to results of equations in body, returns results list |
| GET | api/v1/variables/:name | *:name* variable name | returns  value of  user variable |
| GET | api/v1/variables | none | returns values of  user variables from body |
| DELETE | api/v1/variables/:name | *:name* | deleting user variable with name |
| DELETE | api/v1/variables | none | deletes all user variables for user, optional names ion body |

#### Auth
| Type | Path | Variables | Description |
| --- | --- | --- | --- |
| GET | api/v1/auth/history | none | getting all history of equations for user |
| DELETE | api/v1/auth/history | none | clearing user history |
   
### Auth

[POST] api/v1/auth/login/

    {
        "username": "default",
        "password": "password"
    }

[POST] api/v1/auth/logout/

    with JWT Cookie

### Solve

[POST] api/v1/solve/2%2B2/

		Empty
    
[POST] api/v1/solve/

		{
        "equations": ["2+2", "1+2+a"]
	    }

### Variables

[POST] api/v1/variables/a/"2+2"

		Empty

[POST] api/v1/variables

    {
       "statements": [
        {
            "names": ["a"],
            "equation": "2+2"
        },
        {
            "names": ["b"],
            "equation": "1+2+a"
        },
        {
            "names": ["c", "d"],
            "equation": "b"
        },
    }

[GET] api/v1/variables/a

		Empty

[GET] api/v1/variables

    {
        "names": ["a", "b"]
    }

[DELETE] api/v1/variables/a

		Empty

[POST] api/v1/variables

    {
       "names": ["a", "b"]
    }

[DELETE] api/v1/variables

		Empty
		
*Optional*

	{
		"names": ["a", "b"]
	}


### History

[GET] api/v1/history

		Empty

[DELETE] api/v1/history

		Empty

## Bot commands

> Change to list (with output examples)

* `/s` **Solve command.**

	*Simple solve:*
	
	    /s 1+2
	    
         **Output:**
         3
	    
	*Solve with user variables set:*
	
	    /s a = 2 + 2
	    
	     **Output:**
	     a = 4
	     
	    /s a = b = 1+1
	    
	     **Output:**
	     a = b = 2
	     
	    /s a = b
	    
	     **Output (b == 3):**
	     a = b = 3
	
	*Multiple solve:*
	
		/s a = 2+2
	    b = 3 + 3
	    c = 1 + 1
	    
	    **Output:**
	    a = 4
	    b = 6
	    c = 6
	
* `/get` **Get user variable value command;**

	*Simple*
	
	    /get a
	    
	    **Output (a == 4):**
	    4
	    
	
	*Multiple*
	
	    /get a b c
	    **Output (a == 4, b == 2, c == 1):**
	    a = 4
	    b = 2
	    c = 1

* `/getall` **Get all user variables values command;**
    
        /getall
 	     **Output**
	     a = 1
         
         b = 2

         c = 3
	     

* `/del` **Delete user variable command.**

	*Simple*
	
	    /del a
	    
	     **Output variants:**
	     Success
	     Fault <Not found>
	    
	
	*Multiple*
	
	    /del a b c
	    **Output variants:**
	     Success
	     Fault <Not found>

*  `/delall` **Delete all user variables command.**

 	     **Output variants:**
	     Success
	     Fault <Not found>

*  `/hist` **Delete user variable command.**

	*Output sample:*
	
			 **Output:**
		    1+2
		    a = 2+2
		    4-1
		    a = b = 2+3

* `/clear` 	**Clears user history and variables**

			**Output variants:**
			Success
			Fault <Not found>
			
* `/password` **Returns password for accessing through REST API**

		**Output:**
		eifjkvncqe;dow

* `/genpassword` **Revokes or generates password for accessing through REST API**

		**Output:**
		eifjkvncqe;dow


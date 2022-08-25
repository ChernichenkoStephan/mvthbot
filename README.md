# mvthbot

> WORK IN PROCESS...

<img src="assets/robot.jpeg" alt="mvthbot header picture" width="200"/>

A service-bot with users, variables, and calculation history support to solve equations in text format, to help with calculations in chats. With an open API for interacting with other services and registering via telegram.

## Roadmap

- [ ] mvthbot MVP
    - [x] Equation solving
        - [x] Reverse Polish notation conversion
        - [x] Variables support
    - [x] REST API (using Fiber)
        - [x] Middleware (logging, parsing, etc.)
        - [x] Auth (using JWT)
        - [x] OpenAPI (using swaggo/swag)
    - [x] Telegram bot (using tucnak/telebot)
        - [x] Broadcasting
        - [x] Middleware (logging, parsing etc.)
        - [x] Commands set
    - [x] error handling (using emperror.dev/emperror)
    - [x] logging (using uber.org/zap)
    - [x] Database
        - [x] PostgreSQL (using github.com/jackc/pgx)
        - [x] Migrations (using golang-migrate)
    - [x] Configuration (using env variables, spf13/viper and spf13/cobra)
    - [x] Graceful shutdown (using sync/errgroup and context)
    - [ ] Deploy (using Docker and Makefile)
- [ ] mvthbot 1.0
    - [ ] Health checks (using AppsFlyer/go-sundheit)
    - [ ] Metrics and tracing using Prometheus and Jaeger (via OpenCensus)
    - [ ] Proper testing
        - [ ] Tests refactoring
        - [ ] Coverage increase
        - [ ] Load testing

## Project structure

```
📦 mvthbot
├─ api/                          # OpenApi, etc.
├─ assets/                       # Files for github
├─ bin/                          # For output binaries
├─ cmd/                           
│  └─ mvthbot/                   # Setups and runs functions
├─ configs/                      # Configuration files
├─ db/                           
│  ├─ migration/                 # Migration sql scripts
│  └─ scripts/                   # Help sql scripts
├─ deployments/                  # Docker files, etc.
├─ docs/                         # Some documentation
├─ internal/                      
│  ├─ app/                       # App setup functions for: gracefull shutdown, etc.
│  ├─ auth/                      # JWT auth
│  ├─ bot/                       # Telegram bot module
│  ├─ converting/                # Converting to Reverse polish notation
│  ├─ fixing/                    # Fixing small misspells in equations
│  ├─ lexemes/                   # Lexemes for equation parsing 
│  ├─ logging/                   # zap.Logger setup
│  ├─ misc/                      # Helth API handler
│  ├─ solving/                   # Base equations solving
│  ├─ user/                      # User working logic
│  └─ utils/                     # Some helping functions
├─ logs/                         # Folder for log files
└─ scripts/                      # Some helping bash scripts
```
©generated by [Project Tree Generator](https://woochanleee.github.io/project-tree-generator)

## Prerequisites

* Golang (version 1.18+)
* Docker (version `unknown`)
* PostgreSQL (version ~14.5)
* Go migrate

## Installing

Install the Golang and GO environment

	https://golang.org/doc/install

Install Postgresql (if you want to run locally)

Install go migrate 

	$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

Clone repository

	git clone git@github.com:ChernichenkoStephan/mvthbot.git (go get)

## Build

	go mod tidy && make build

## Test

	./scripts/go.test.sh

## Run

1. Setup environment variables
2. Run DB ``` ./scripts/db.start.sh ```
3. Setup DB ``` make createdb && make migrateset ```
4. Run with ``` ./scripts/go.run.sh ```

## Configuration

### Available flags

| Flag |  Short  | Description |
| --- | --- | --- |
| log | l | log dir path |
| config | c | config file path |
| token | t | bot token |
| port | p | api listen port |

### config.yaml

> Example in folder

### Environment variables

| Name | Description |
| --- | --- |
| SECRET | Seed for API password generatioEnvironmentn |
| BOT_TOKEN | Token for telegram bot |
| DB_USER | PostgreSQL user name |
| DB_NAME | PostgreSQL database name |

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


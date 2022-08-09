# mvthbot
Math solving service

## Features

- [x] Equasion solving
    - [x] Reverse Polish notation conversion
    - [x] Variables support
- [ ] REST API (using Fiber)
    - [ ] Middleware
    - [ ] Auth (using JWT)
    - [ ] OpenAPI (using Swagger)
- [ ] Telegram bot (using tucnak/telebot)
    - [ ] Broadcasting
- [ ] error handling (using emperror.dev/emperror)
- [ ] logging (using zap)
- [ ] Database
    - [ ] PostgreSQL (using xxx)
    - [ ] Migrations (using Ansible)
- [ ] ES Deploy (using Docker and Makefile)
- [ ] metrics and tracing using Prometheus and Jaeger (via OpenCensus)
- [ ] health checks (using AppsFlyer/go-sundheit)
- [ ] graceful restart (using cloudflare/tableflip) and shutdown
- [ ] support for multiple server/daemon instances (using oklog/run)
- [ ] messaging (using ThreeDotsLabs/watermill)
- [ ] configuration (using spf13/viper)
- [ ] Dashboard (using GoAdminGroup/go-admin)
- [ ] Full Deploy (using Dockercompose)
- [ ] Advanced Testing
- [ ] CL/CI

## API

also in JWT and headers

### Auth

auth
    [POST] api/v1/solve/2%2B2/

### Solve

    [POST] api/v1/solve/2%2B2/

    [POST] api/v1/solve/
    {
        equations: ["2+2", "1+2+a"]
    }

### Variables

    [POST] api/v1/variables/a/"2+2"

    [POST] api/v1/variables
    {
        statements: [
        {
            names: ["a"], 
            equation: "2+2"
        },
        {
            names: ["b"], 
            equation: "1+2+a"
        },
        {
            names: ["c", "d"], 
            equation: "b"
        },
    }

    [GET] api/v1/variables/a

    [GET] api/v1/variables
    {
        names: ["a", "b"]
    }

    [DELETE] api/v1/variables/a

    [DELETE] api/v1/variables
    {
        names:["a", "b"]
    }

    [DELETE] api/v1/variables

### History

    [GET] api/v1/history

    [DELETE] api/v1/history

## Bot

    /key

    /slv 1+2

    /set a = 2 + 2

    /set a = 2+2
    b = 3 + 3
    c = 1 + 1

    /set a = b = 1+1

    /set a = b

    /get a

    /del a

    /del a b

    /clear

    /history

    **Output:**
    1+2
    a = 2+2
    4-1
    a = b = 2+3




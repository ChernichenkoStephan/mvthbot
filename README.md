# mvthbot
Math solving service

## API

also in JWT and headers

api/v1/<userkey>/solve/1+2+3/

api/v1/<userkey>/solve/

{
    [{"", "1+2"}, {"a", "2+2"}]
}

api/v1/<userkey>/set/a/1+2+3

api/v1/<userkey>/set/[a,v]/1,2,3

api/v1/<userkey>/set/a/b

api/v1/<userkey>/get/a/b

api/v1/<userkey>/get/a/b

api/v1/<userkey>/history

api/v1/<userkey>/delete/a/b

api/v1/<userkey>/clear

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

/api
OpenAPI/Swagger specs, JSON schema files, protocol definition files.

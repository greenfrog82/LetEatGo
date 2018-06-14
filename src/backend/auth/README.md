# go package install
```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/dgrijalva/jwt-go
```

# postgresql user 설정
참고 URL : http://brownbears.tistory.com/19

## postgresql user 정보
```
DBNAME : leteatgo
Username : admin
Password : qwer1234
```

# API doc
```
회원가입
[REQ]
POST  HTTP/1.1
Host: localhost:7878/signup
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: e6617146-76b7-b7d8-cf9e-ea7896818fab

{"username": "test444", "password": "1234"}

[RES]
{
    "success": true
}


로그인
[REQ]
POST /login HTTP/1.1
Host: localhost:7878
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: 9f80b9c8-640c-3729-c2de-34b7eeb5a615

{"username": "test", "password": "1234"}


[RES]
{
    "expire": "1528433488",
    "token": "eqAY4BozT_-1lCpK122LUr35_RIKjHI..."
}
```

```
회원 가입

POST  HTTP/1.1
Host: localhost:7878/signup
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: e6617146-76b7-b7d8-cf9e-ea7896818fab

{"username": "test444", "password": "1234"}


HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 05 Jun 2018 05:13:02 GMT
Content-Length: 19

{"success": "true"}

----


로그인

POST /login HTTP/1.1
Content-Type: application/json
cache-control: no-cache
Postman-Token: c16bde4b-fb55-45f1-b909-251d2a2609ff
User-Agent: PostmanRuntime/7.1.1
Accept: */*
Host: localhost:7878
cookie: csrftoken=cedf1c5854d236b672edf325da884197; sessionid=7812c18bbb3bdb63d02471fa4fa20f3b
accept-encoding: gzip, deflate
content-length: 40
Connection: keep-alive

{"username": "test", "password": "1234"}

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 05 Jun 2018 05:13:02 GMT
Content-Length: 179

{"expire":"1528434782","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mjg0MzQ3ODIsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoidGVzdCJ9.yOLZYgDWlDwonf19kH_sZmX4Tg3yb9HyG0pOfQ4NHCc"}

```

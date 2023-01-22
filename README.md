# web-project
Web Programming final project

## Installation
run `docker-compose up --build`
## Endpoints
### signup
```
{
   "name": "signup",
   "url": "http://localhost:3000/users/signup",
   "method": "POST",
   "body": {
    "type": "json",
    "raw": {
      "user_name" : "armin2",
      "password" : "armin21234"
    }
   }
}
```
### Login 
```
{
 "name": "login",
 "url": "http://localhost:3000/users/login",
 "method": "POST",
 "body": {
  "type": "json",
  "raw": {
   "user_name": "armin1",
   "password": "armin11234"
  }
 }
}
```
### Add url 
```
{
 "name": "add url",
 "url": "http://localhost:3000/urls",
 "method": "POST",
 "body": {
  "type": "json",
  "raw": {
   "user_name": "armin1",
   "address": "aut.ac.ir",
   "treshold": 2
  }
 },
 "auth": {
  "type": "bearer",
  "bearer": "your token"
 }
}
```
### Get urls for user 
```
  {
   "name": "get urls",
   "url": "http://localhost:3000/urls",
   "method": "GET",
   "auth": {
    "type": "bearer",
    "bearer": "your token"
   }
  }
```
### Get url detailes 
```
{
 "name": "get url",
 "url": "http://localhost:3000/urls/:url_id",
 "method": "GET",
 "auth": {
  "type": "bearer",
  "bearer": "your token"
 }
}
```
### Get Alerts for an url 
```
  {
   "name": "get alerts",
   "url": "http://localhost:3000/alerts/:url_id",
   "method": "GET",
   "auth": {
    "type": "bearer",
    "bearer": "your token"
   }
  }
```

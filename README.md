# CVWO Web Forum (API)

## Deployment
[https://cvwo-web-forum-api.onrender.com](https://cvwo-web-forum-api.onrender.com)

## Documentation
For information about the endpoints of the API refer to [https://documenter.getpostman.com/view/24839876/2s8ZDScQqM](https://documenter.getpostman.com/view/24839876/2s8ZDScQqM)

## Getting Started with your Own Local Version
### Requirements
* Go: known working version - 1.19.4
### Run the App in Development Mode
```
$ git clone https://github.com/vansh284/CVWOWebForumAPI
$ cd ./CVWOWebForumAPI
$ go get ./...
$ go run ./cmd/main/main.go
```
* Rename .env.example to .env. 
* Set up your MySQL database and modify the DSN in .env accordingly. Modify other values in .env according to needs.


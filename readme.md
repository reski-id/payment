

<h3 align="center">Rest API <br>
<h5 align="center" >Golang Echo Clean Architecture <h5>
<br>
</h4>
<p align="left">
<h2>
  Content <br></h2>
  • Key Features <br>
  • Installing Using Github<br>
  • Installing Using Docker<br>
  • End Point<br>
  • Technologi that i use<br>
  • Contact me<br>
</p>

##  Problem Solving
Solution `TASK1 - ALGO`

##  Screensoot Postman
Screenshoot `ScreenshootProgram`

## 📱 Features

* register
* login
* Top Up
* Payment
* Transfer


## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/news_restAPI.git

# Go into the repository
$ cd news_restAPI

# Install dependencies
$ go get

# load .env env 
# use bash cmd and type this..
$ source .env

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allready create database for this app see local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/portal

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name portal -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="portal" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/portal

# Run the app
$ docker logs portal
```

## 📜 End Point  

User
| Methode       | End Point      | used for            | Using JWT
| ------------- | -------------  | -----------         | -----------
| `POST`        | /register      | Register            | No
| `POST`        | /login         | Login               | No 
| `GET`         | /my            | Get my Profile      | Yes
| `PUT`         | /profile         | Update my Profile   | Yes

Transaction
| Methode       | End Point      | used for            | Using JWT
| ------------- | -------------  | -----------         | -----------
| `POST`        | /topup      | /topup            | Yes
| `POST`        | /payment      | /payment            | Yes
| `POST`        | /transfer      | /transfer            | Yes


## 🛠️ Technologi

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Echo Framework](https://echo.labstack.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Mockery Unit Testing](https://github.com/vektra/mockery)
- [Git Repository](https://github.com/reski-id)
- [Trunk Base Development](https://trunkbaseddevelopment.com/)
- [JSON Web Tokens (JWT)](https://jwt.io/)

## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>

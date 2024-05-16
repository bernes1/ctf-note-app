# Music cooking 


we cooking upp a lil app for saving links to good dj sets and songs accross platforms.
Do you also have the problem with not knowing what banger tunes you sould listen to when you have listened to bangers before.


# feature requests
get a nice TUI




# database plan??

### some toughts about the database
 we write some of that squeel 

#### artist table
  - id 
  - name

#### platform table 
   - id
   - name


#### djset table
  - id 
  - name
  - url
  - artist

##### beautiful sql

```sql
CREATE TABLE Platform (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE Artist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    UNIQUE(name)
);

CREATE TABLE DJset (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    platform_id INTEGER NOT NULL,
    artist_id INTEGER NOT NULL,
    FOREIGN KEY (platform_id) REFERENCES Platform(id),
    FOREIGN KEY (artist_id) REFERENCES Artist(id)
);
```


# Setup

## Prerequisites

Before running the program, make sure you have the following:

- Docker installed on your machine https://docs.docker.com/get-docker/
- Go programming language installed https://go.dev/doc/install



### Getting started

First make a folder called vars in the project folder
```shell
 cd music-cooking
 mkdir vars
```

Create a .env file in vars folder with the following contents

```
  // vars/.env
  DATABASE_URL=postgresql://postgres:postgres@localhost/setsdb
  POSTGRES_USER=postgres
  POSTGRES_PASSWORD=postgres
  POSTGRES_SERVER=localhost
  POSTGRES_PORT=5432
  POSTGRES_DB=setsdb
```
After you have added the .env file to the vars folder run this comand in the project root folder.

```shell
docker compose --env-file="vars/.env" up -d
```
this comand wil start the database.


To run the program run the following 
```shell
go run src/main.go
``` 





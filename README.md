# Music cooking 


we cooking upp a lil app for saving links to good dj sets and songs accross platforms.
Do you also have the problem with not knowing what banger tunes you sould listen to when you have listened to bangers before.



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
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Artist (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE DJset (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    platform_id INT NOT NULL,
    artist_id INT NOT NULL,
    FOREIGN KEY (platform_id) REFERENCES Platform(id),
    FOREIGN KEY (artist_id) REFERENCES Artist(id)
);
```


# Setup

Create a .env file with the following contents

```
  // .env
  DATABASE_URL = "postgresql://user:password@postgresserver/db"
  POSTGRES_USER = "postgres"
  POSTGRES_PASSWORD = "postgres"
  POSTGRES_SERVER = "localhost"
  POSTGRES_PORT = "5432"
  POSTGRES_DB = "setsdb"
```


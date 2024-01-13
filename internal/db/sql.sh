# create the database

CREATE TABLE users(
   id VARCHAR(40) PRIMARY KEY,
   username VARCHAR(20) NOT NULL ,
   password VARCHAR(100) NOT NULL,
   email VARCHAR(45) DEFAULT '',
   first_name VARCHAR(45) DEFAULT '',
   last_name VARCHAR(45) DEFAULT '',
   phone VARCHAR(15) DEFAULT '',
   github VARCHAR(45) DEFAULT '',
   medium VARCHAR(45) DEFAULT '',
   twitter VARCHAR(45) DEFAULT '',
   linkedin VARCHAR(45) DEFAULT '',
   objective VARCHAR(500) DEFAULT '',
   tagline VARCHAR(150) DEFAULT '',
   theme VARCHAR(15) DEFAULT '',
   skills TEXT DEFAULT '{
        "hard": [],
        "soft": [],
        "interest": []
    }',
   projects TEXT DEFAULT '[]'
   experience TEXT DEFAULT '[]'
);


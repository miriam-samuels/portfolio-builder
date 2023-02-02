# create the database

CREATE DATABASE portfolio;

CREATE TABLE users(
   username VARCHAR(30)
   email VARCHAR(45)
   password VARCHAR(100)
   first_name VARCHAR(45)
   last_name VARCHAR(45)
   phone VARCHAR(20)
   github VARCHAR(45)
   medium VARCHAR(45)
   twitter VARCHAR(45)
   linkedin VARCHAR(45)
   objective VARCHAR(400)
   tagline VARCHAR(150)
   skills JSON
   projects JSON
   theme VARCHAR(30)
);

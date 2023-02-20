# web-backend

## Usage

```go
go run main.go
```

## API call

```bash
localhost:8005/posts/create

json = {
    "Title": "cjdhuds",
    "Genre": "cjdhuds",
    "Director": "cjdhuds",
    "Release_year": 2008
}

expected result:
{
    "message": "successfully created"
}


```

## MySQL

since i create a mySQL db in my local computer, you can run the following command line to create another table with same schema.

```bash
$ mysql -u root -p

mysql> CREATE DATABASE movies;

mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| movies             |
| mysql              |
| performance_schema |
| QA                 |
+--------------------+


mysql> USE movies;

mysql> CREATE TABLE movies(title VARCHAR(50) NOT NULL,genre VARCHAR(30) NOT NULL,director VARCHAR(60) NOT NULL,release_year INT NOT NULL,PRIMARY KEY(title));

mysql> DESCRIBE movies;
+--------------+-------------+------+-----+---------+-------+
| Field        | Type        | Null | Key | Default | Extra |
+--------------+-------------+------+-----+---------+-------+
| title        | varchar(50) | NO   | PRI | NULL    |       |
| genre        | varchar(30) | NO   |     | NULL    |       |
| director     | varchar(60) | NO   |     | NULL    |       |
| release_year | int         | NO   |     | NULL    |       |
+--------------+-------------+------+-----+---------+-------+


mysql> INSERT INTO movies VALUE ("Joker", "psychological thriller", "Todd Phillips", 2019);

mysql> SELECT * FROM movies;

+-------+------------------------+---------------+--------------+
| title | genre                  | director      | release_year |
+-------+------------------------+---------------+--------------+
| Joker | psychological thriller | Todd Phillips |         2019 |
+-------+------------------------+---------------+--------------+

```

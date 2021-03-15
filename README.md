# Fuzzy Dollop (GitHub auto-generated name)

Test HTTP Application built with TDD start to finish

## Goals

* The API should support `CREATE`, `READ`, `UPDATE`, `DELETE` of the `posts` resource with validation.
* Support fetching multiple posts, a single post, and filtering posts by user.
* Use an in-memory database or native data structures of your choice for persistence. 
* Cover all the important behavior of the code with automated unit tests. 

## Extras

* Add authentication
* Use PostgreSQL
* Deploy to Heroku

## Next Steps

If I had more time, I'd:
* containerize and deploy to Heroku. 
* add more test cases for each unit test
* add further unit tests for stuff outside db model methods
* make error handling and error messages consistent
* set up logging to also pipe to persistent file
* determine the best way to make sure SQL `UPDATE`s do not overwrite data whether that would be by coordinating with the client-side or by making more use of SQL

## Instructions

To install dependencies, run:
```bash
go get
```

Set up a PostgreSQL with credentials that you can store in an `.env` file. Example `.env` file:

```
PSQL_HOST="localhost"
PSQL_PORT=5432
PSQL_USER="postgres"
PSQL_DBNAME="furrydollop"
PSQL_password="password"
PSQL_PW="password"
```

You can run the SQL scripts located in `furry-dollop/scripts` to create the `posts` and `users` tables.

To turn on the app, run:
```bash
go run .
```

To test model methods, run:
```bash
go run ./models/psql -v
```

## Sample Data

### Posts

``` javascript
[
  {
    "userId": 1,
    "id": 1,
    "title": "Node is awesome",
    "body": "Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine."
  },
  {
    "userId": 1,
    "id": 2,
    "title": "Spring Boot is cooler",
    "body": "Spring Boot makes it easy to create stand-alone, production-grade Spring based Applications that you can "just run"."
  },
  {
    "userId": 2,
    "id": 3,
    "title": "Go is faster",
    "body": "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
  },
  {
    "userId": 3,
    "id": 4,
    "title": "'What about me?' -Rails",
    "body": "Ruby on Rails makes it much easier and more fun. It includes everything you need to build fantastic applications, and you can learn it with the support of our large, friendly community."
  }
]
```

### Users

```javascript
[
  {
    "id": 1,
    "name": "Ryan Dahl",
    "email": "node4lyfe@example.com",
    "expertise": "Node"
  },
  {
    "id": 2,
    "name": "Rob Pike",
    "email": "gofarther@example.com",
    "expertise": "Go"
  },
  {
    "id": 3,
    "name": "DHH",
    "email": "magic@example.com",
    "expertise": "Rails"
  }
]
```
# Backend Exercise

**Please limit your time working on this exercise to four hours.**

Please create an application implementing a CRUD web service for `posts` using technology applicable to the position you are applying for.

- The API should support `CREATE`, `READ`, `UPDATE`, `DELETE` of the `posts` resource with validation.
- Support fetching multiple posts, a single post, and filtering posts by user.
- Use an in-memory database or native data structures of your choice for managing data. 
- Cover all the important behavior of the code with automated tests. Feel free to choose between integration tests and unit tests.

#### Additional topics

If you have extra time, we'd love to see what you create around any of the topics below:

- Authentication (Basic, OAuth2, JWT)
- Utilize a database for persistent storage (Feel free to choose an engine of your choice)
- Deploy the application to a cloud computing platform (Heroku, GCP, AWS, Azure, etc.)

When you are done, zip up the project and return it to your Big Nerd Ranch contact.

#### Sample Data

Below, you'll find some sample data to help structure your models and seed your database:

### Posts

```
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
  },
  {
    "userId": 4,
    "id": 5,
    "title": ".NET has the gravy",
    "body": ".NET enables engineers to develop blazing fast web services with ease, utilizing tools developed by Microsoft!"
  }
  
]
```

### Users

```
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
  },
  {
    "id": 4,
    "name": "John Watkins",
    "email": "jwats@example.com",
    "expertise": ".NET"
  }
]
```
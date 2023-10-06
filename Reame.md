
# Simple TechBlog Posts CRUD App

A basic CRUD application for managing tech blog posts. It provides a RESTful API to create, read, update, and delete posts.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Usage](#usage)

## Features

- In-memory database for storing blog posts.
- RESTful API endpoints for CRUD operations.
- Middleware for request logging.

## Requirements

- Go 1.16 or newer

## Usage

1. Start the server:

```bash
go run .
```

2. The server will start on port `8080`. You can use the API endpoints as mentioned below:

- **Create Post**: `POST /v1/posts`
- **Get All Posts**: `GET /v1/posts`
- **Get Post by ID**: `GET /v1/posts?id=1`
- **Update Post**: `PUT /v1/posts?id=1`
- **Delete Post**: `DELETE /v1/posts?id=1`

## Author

Daniel Cardenas
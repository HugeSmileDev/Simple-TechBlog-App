package main

import (
	"Simple-TechBlog-Posts-CRUD-App/handlers"
	"Simple-TechBlog-Posts-CRUD-App/middleware"
	"fmt"
	"net/http"
)

const apiVersion = "v1"

func main() {
	http.HandleFunc(fmt.Sprintf("/%s/posts", apiVersion), middleware.LogRequest(handlers.PostsEndpoint))
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

package handlers

import (
	"Simple-TechBlog-Posts-CRUD-App/models"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

var posts = []models.Post{
	{1, 1, "Node is awesome", "Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine."},
	{1, 2, "Spring Boot is cooler", "Spring Boot makes it easy to create stand-alone, production-grade Spring based Applications that you can \"just run\"."},
	{2, 3, "Go is faster", "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."},
	{3, 4, "'What about me?' -Rails", "Ruby on Rails makes it much easier and more fun. It includes everything you need to build fantastic applications, and you can learn it with the support of our large, friendly community."},
	{4, 5, ".NET has the gravy", ".NET enables engineers to develop blazing fast web services with ease, utilizing tools developed by Microsoft!"},
}

var postIDCounter = len(posts)
var postMutex = &sync.Mutex{}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId != "" {
		userPosts := make([]models.Post, 0)
		id, _ := strconv.Atoi(userId)
		for _, post := range posts {
			if post.UserId == id {
				userPosts = append(userPosts, post)
			}
		}
		json.NewEncoder(w).Encode(userPosts)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	for _, post := range posts {
		if post.ID == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid post data", http.StatusBadRequest)
		return
	}
	//Validation for each Post
	validationMsg := models.ValidatePost(&post)
	if validationMsg != "" {
		http.Error(w, validationMsg, http.StatusBadRequest)
		return
	}

	postMutex.Lock()
	postIDCounter++
	post.ID = postIDCounter
	posts = append(posts, post)
	postMutex.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully!"})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	var updatedPost models.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid post data", http.StatusBadRequest)
		return
	}

	//Validation for each updating Post
	validationMsg := models.ValidatePost(&updatedPost)
	if validationMsg != "" {
		http.Error(w, validationMsg, http.StatusBadRequest)
		return
	}

	postMutex.Lock()
	for i, post := range posts {
		if post.ID == id {
			updatedPost.ID = id
			posts[i] = updatedPost
			postMutex.Unlock()
			response := struct {
				Post    models.Post `json:"post"`
				Message string      `json:"message"`
			}{
				Post:    updatedPost,
				Message: "Post updated successfully!",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	postMutex.Unlock()
	http.Error(w, "Post not found", http.StatusNotFound)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	postMutex.Lock()
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			postMutex.Unlock()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted successfully!"})
			return
		}
	}
	postMutex.Unlock()
	http.Error(w, "Post not found", http.StatusNotFound)
}

func PostsEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Get("id") != "" {
			GetPostById(w, r)
		} else {
			GetPosts(w, r)
		}
	case http.MethodPost:
		CreatePost(w, r)
	case http.MethodPut:
		UpdatePost(w, r)
	case http.MethodDelete:
		DeletePost(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fredy-bambang/golearn/internal/driver"
	models "github.com/fredy-bambang/golearn/internal/models"
	repository "github.com/fredy-bambang/golearn/internal/repository"
	post "github.com/fredy-bambang/golearn/internal/repository/post"
	"github.com/go-chi/chi"
	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
	"github.com/payfazz/go-router/segment"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

// Fetch all post data
func (p *Post) Fetch(r *http.Request) *handler.Response {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	// respondwithJSON(w, http.StatusOK, payload)
	return defresponse.JSON(200, payload)
}

// Create a new post
func (p *Post) Create(r *http.Request) *handler.Response {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	fmt.Println(newID)
	if err != nil {
		defresponse.Text(500, "Server Error")
	}

	return defresponse.JSON(201, map[string]string{"message": "Successfully Created"})
}

// Update a post by id
func (p *Post) Update(r *http.Request) *handler.Response {
	// id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	id, _ := segment.Get(r, "id")
	intID, _ := strconv.Atoi(id)
	data := models.Post{ID: int64(intID)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		defresponse.Text(http.StatusInternalServerError, "Server Error")
	}

	return defresponse.JSON(201, payload)
}

// GetByID returns a post details
func (p *Post) GetByID(r *http.Request) *handler.Response {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		return defresponse.Text(http.StatusNoContent, "Content not found")
	}

	return defresponse.JSON(http.StatusOK, payload)
}

// Delete a post
func (p *Post) Delete(r *http.Request) *handler.Response {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		return defresponse.Text(http.StatusInternalServerError, "Server Error")
	}

	return defresponse.JSON(http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func get_books(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func create_book(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func book_by_id(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
		{

		}
	}
	return nil, errors.New("book not found")
}

func get_book_by_id(c *gin.Context) {
	id := c.Param("id")
	book, err := book_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "the book was not found"})
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkout_book(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "enter the query params"})
	}

	book, err := book_by_id(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "there is some error in finding the book"})
	}

}

func main() {
	router := gin.Default()
	router.GET("/books", get_books)
	router.POST("/create", create_book)
	router.GET("/getBook/:id", get_book_by_id)
	router.Run("localhost:3000")
}

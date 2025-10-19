package main

import (
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"slices"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	ISBN   string  `json:"isbn"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{ID: "1", Title: "Fundamental of Deep Learning in Practice", Author: "Nuttachot Promrit and Sajjaporn Waijanya", ISBN: "978-1234567890", Year: 2023, Price: 59.99},
	{ID: "2", Title: "Practical DevOps and Cloud Engineering", Author: "Nuttachot Promrit", ISBN: "978-0987654321", Year: 2024, Price: 49.99},
	{ID: "3", Title: "Mastering Golang for E-commerce Back End Development", Author: "Nuttachot Promrit", ISBN: "978-1111222233", Year: 2023, Price: 54.99},
}

func getBooks(c *gin.Context) {
	yearQuery := c.Query("year")

	if yearQuery != "" {
		year, err := strconv.Atoi(yearQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year format"})
			return
		}

		var filteredBooks []Book
		for _, book := range books {
			if book.Year == year {
				filteredBooks = append(filteredBooks, book)
			}
		}
		c.JSON(http.StatusOK, filteredBooks)
		return
	}
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if newBook.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if newBook.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author is required"})
		return
	}
	if newBook.Year < 1900 || newBook.Year > 2100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "year must be between 1900-2100"})
		return
	}
	if newBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	// Generate new ID
	newBook.ID = fmt.Sprintf("%d", len(books)+1)

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if updatedBook.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if updatedBook.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author is required"})
		return
	}
	if updatedBook.Year < 1900 || updatedBook.Year > 2100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "year must be between 1900-2100"})
		return
	}
	if updatedBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			books = slices.Delete(books, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "book-api"})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/books", getBooks)
		api.GET("/books/:id", getBook)
		api.POST("/books", createBook)
		api.PUT("/books/:id", updateBook)
		api.DELETE("/books/:id", deleteBook)
	}

	fmt.Println("Server is running on port 8080...")
	r.Run(":8080")
}
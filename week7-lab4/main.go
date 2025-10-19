package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB() {
	var err error
	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "postgres")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "5432")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, name)
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to open database:", err)  // เปลี่ยน Fatal -> Println
		return  // เพิ่ม return
	}
	
	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to database:", err)  // เปลี่ยน Fatal -> Println
		return  // เพิ่ม return
	}
	
	log.Println("Successfully connected to database")
}

func main() {
	initDB()
	defer db.Close()
	
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "unhealthy", 
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"message": "healthy"})
	})

	//api := r.Group("/api/v1")
	//{
	//	api.GET("/books", getBooks)
	//	api.GET("/books/:id", getBook)
	//	api.POST("/books", createBook)
	//	api.PUT("/books/:id", updateBook)
	//	api.DELETE("/books/:id", deleteBook)
	//}

	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}
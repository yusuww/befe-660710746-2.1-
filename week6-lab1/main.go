package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin" 
	"slices" //มันไม่มีคำสั่งลบเลยเพิ่อิมพอตนี้เข้ามา
)

// Student struct
type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

func getStudents(c *gin.Context) {
	yearQuery := c.Query("year")

	if yearQuery != "" {
		filter := []Student{}
		for _, student := range students {
			if fmt.Sprint(student.Year) == yearQuery {
				filter = append(filter, student)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)
}

func getStudent(c *gin.Context) {
	id := c.Param("id")

	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
}

func createStudent(c *gin.Context) { // แก้ชื่อ function จาก creatStudent
	var newStudent Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newStudent.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"}) // ลบ () ออก
		return
	}

	if newStudent.Year < 1 || newStudent.Year > 4 { // แก้ไข: เพิ่ม .Year
		c.JSON(http.StatusBadRequest, gin.H{"error": "year must be 1-4"}) // ลบ () ออก
		return // เพิ่ม return
	}

	newStudent.ID = fmt.Sprintf("%d", len(students)+1) // แก้ไข: เพิ่ม fmt.

	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent) // เปลี่ยนเป็น StatusCreated และ return newStudent แทน students
}

func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updateStudent Student

	// ตรวจสอบ JSON format
	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เพิ่มการ validate ข้อมูล (optional)
	if updateStudent.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if updateStudent.Year < 1 || updateStudent.Year > 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "year must be 1-4"})
		return
	}

	// หา student และอัปเดต
	for i, student := range students {
		if student.ID == id {
			updateStudent.ID = id
			students[i] = updateStudent
			c.JSON(http.StatusOK, updateStudent)
			return
		}
	}

	// หาไม่เจอ
	c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
}

func deleteStudent (c *gin.Context){
	id := c.Param("id")
	
	for i, student := range students {

		if student.ID == id{
			students = slices.Delete(students, i, i+1) //ถ้ามี 12345 จะไม่ลบ 5 // := เป็ฯการประกาศตัวแปรใหม่
			c.JSON(http.StatusOK, gin.H{"massage": "Student delete successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents)
		api.GET("/students/:id", getStudent)
		api.POST("/students", createStudent) // แก้ไข: ลบ /:id ออก เพราะ POST ไม่ต้องการ ID
		api.PUT("/students/:id", updateStudent)
		api.DELETE("/students/:id", deleteStudent)
	}

	r.Run(":8080")
}
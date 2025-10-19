package main //ทุกไฟล์ต้องมีแพ็คเกต

import (
	"github.com/gin-gonic/gin"
)

//นิยามสตัค
type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func main(){
	r := gin.Default() //Default เราต์ไปยัวฟังก์ที่ถูกต้อง

	r.GET("/users" , func(c *gin.Context) {//รัยรีเควสแบบget

		user :=[]User{{ID:"1", Name:"Sunita"}} //ส่งjson
		c.JSON(200, user)
	}) 

	r.Run(":8080")
}
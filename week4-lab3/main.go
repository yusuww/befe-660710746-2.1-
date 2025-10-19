package main /**/
import (/*การเรียกใช้*/
	"fmt"
	"errors"
)

type Student struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s*Student) IsHonor() bool { /* ซึ่งคืนค่าเป็น boleen  ,  "*"  คือ ประกาศว่าชี้ไปที่ Student  เมธอทคือ IsHonor()*/
	return s.GPA >= 3.50 /*ส่งค่าreturnไปเปรียบเทียบ*/
}

/*เช็ค error*/
func (s *Student) Validate() error { /*error(ชนิดของข้อมูล)*/
	if s.Name == ""{/*เช็คว่ามีช่องว่างในชื่อไหม*/
		return errors.New("name is required") /*การใช้เพ็คเกต errors(แพ็คเกต)*/ 
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil /* nil = ว่างไม่error */
}

func main(){
	//var st  Student = Student{ID:"1", Name:"Sunita", Email:"krittinnawong_s@silpakron.edu", Year:4, GPA:2.75} /*กำหนดค่า*/

	// st := Student = Student({ID:"1", Name:"Sunita", Email:"krittinnawong_s@silpakron.edu", Year:4, GPA:2.75}) อีกวิธี

	Student := []Student{// [] เก้บstudentหลายคน
		{ID:"1", Name:"Sunita", Email:"krittinnawong_s@silpakron.edu", Year:4, GPA:2.75}, // คั่นด้วย ","
		{ID:"2", Name:"Alice", Email:"Alice@silpakron.edu", Year:4, GPA:3.75},
	} 
	
	newStudent := Student{{ID:"2", Name:"trudy", Email:"trudy@silpakron.edu", Year:4, GPA:3.50}}
	Students = append(students, newStudent)

	for i, student := range student{ //ประกาศตัวแปรใหม่สองตัวซึ่งเริ่มต้นที่ศูนย์ _ไม่ต้องเอาตัวแปรไปรับจากที่ใช้ i
		fmt.Printf("%dHonor = %v\n", st.IsHonor())//ปริ้นเรียกใช้ fmt
		fmt.Printf("%dValidation = %v\n", st.Validate())
	}
}
/**/
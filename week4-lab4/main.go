package main /**/

import (
	"fmt"
	"errors"
)

func divide(a, b float64) (float64, error){
	if b == 0{
		return 0, errors.New("cannot divide by 0")
	}
	return a/b, nil
}

func main(){
	result, err:= divide(10,2) //ไปเช็ค error ที่ divide
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Result = ", result)
}

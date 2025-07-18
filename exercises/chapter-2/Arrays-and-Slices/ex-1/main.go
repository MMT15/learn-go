package main

import "fmt"

var grades [8]int = [8]int{90, 85, 92, 78, 95}

func main() {
	fmt.Println(grades)
	studentGrades:=grades
	studentGrades[0]=100
	fmt.Println(studentGrades, grades)
}

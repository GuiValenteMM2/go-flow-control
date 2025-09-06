package main

import (
	"fmt"
	"go-flow-control/internal/employee"
)

func main() {
	employee.RepoTest()

	job1 := employee.Job{JobName: "Dev", Salary: 3000, YearsReq: 2}
	employee1 := employee.Employee{Id: 2, Name: "Gui", Job: job1}

	fmt.Printf("Hi my name is %v, I am a %v and I earn %v \n", employee1.Name, employee1.Job.JobName, employee1.Job.Salary)

	total := 15 * 20 / 100

	fmt.Println(total)
}

package main

import (
	"fmt"
	"go-flow-control/internal/employee"
)

func main() {

	job1 := employee.Job{JobName: "Dev", Salary: 3000, YearsReq: 2}
	job2 := employee.Job{JobName: "PO", Salary: 8000, YearsReq: 3}

	employee1 := employee.Employee{Id: 2, Name: "Gui", Job: job1}
	employee2 := employee.Employee{Id: 1, Name: "Johnny Bodybuilder", Job: job2}

	employee.AddToRepo(employee1)
	employee.AddToRepo(employee2)

	employee.RemoveFromRepo(2)

	empList := employee.GetAllFromRepo()

	for i := range len(empList) {
		fmt.Println(employee.ToString(empList[i]))
	}

	fmt.Println(employee.GetFromRepo(1))

}

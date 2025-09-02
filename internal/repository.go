package internal

var employeeRepository []Employee = make([]Employee, 0)

func addToRepo(employee Employee) {
	employeeRepository = append(employeeRepository, employee)
}

func getFromRepo(employeeId uint16) (Employee, bool) {
	for _, employee := range employeeRepository {
		if employee.id != 0 && employee.id == employeeId {
			return employee, true
		}
	}
	return Employee{}, false
}

func getAllFromRepo() []Employee {
	return employeeRepository
}

// func removeFromRepo(employeeId uint16) bool {
// 	for _, employee := range employeeRepository {
// 		if employee.id != 0 && employee.id == employeeId {

// 		}
// 	}
// }

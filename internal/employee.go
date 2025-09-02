package internal

type Employee struct {
	id   uint16
	name string
	Job
}

type Job struct {
	jobName  string
	salary   uint16
	yearsReq uint8
}

func newEmployee(employee Employee) Employee {
	return employee
}

func newJob(job Job) Job {
	return job
}

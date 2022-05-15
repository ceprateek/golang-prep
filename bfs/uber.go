package bfs

import "fmt"

/*
We are given a list of entries which has information about employees in an organization.
The format of each entry is as follows.

employeeId, firstName, lastName, managerId

Here is a sample of the inputs.
Example:
1, Rob, Choi, 2
2, Jennifer, Grant
3, Andy, Zuckerman, 1
4, Mark, O’Donnell, 1

Notes:
Each employee only reports to one manager
If an employee does not have a managerId, that employee is the head of the organization

We want to print the hierarchy of the organization in a readable format as shown below in a sample.
Output
Jennifer Grant
 	Rob Choi
 		Andy Zuckerman
 		Mark O'Donnell
*/

//1. Map of all employees [managerId] -> employee
//2. start from root
//3. queue add root, iterate while q is empty
//keep adding reportees from map to q

type employee struct {
	employeeId, managerId int
	firstname, lastname   string
}

func RunPrintDirectoryOfEMployee() {
	//fmt.Println("Hello world")
	listOfExistingEmps := make([]*employee, 8)
	j := employee{2, 0, "Jennifer", "Grant"}
	jr := employee{5, 2, "James", "Bond"}
	jrR := employee{6, 5, "Batman", "Most"}
	r := employee{1, 2, "Rob", "Choi"}
	a := employee{3, 1, "Andy", "Zuckerman"}
	m := employee{4, 1, "Mark", "O'Donnell"}
	ar := employee{7, 3, "Jason", "Stettam"}
	jsr := employee{8, 7, "Prateek", "Garg"}
	listOfExistingEmps[0] = &j
	listOfExistingEmps[1] = &r
	listOfExistingEmps[2] = &a
	listOfExistingEmps[3] = &m
	listOfExistingEmps[4] = &jr
	listOfExistingEmps[5] = &jrR
	listOfExistingEmps[6] = &ar
	listOfExistingEmps[7] = &jsr
	printDirectoryOfEMployee(listOfExistingEmps)
}

func printDirectoryOfEMployee(employess []*employee) {
	//org head
	var rootEmployee *employee
	//Map of all employees [managerId] -> employee
	cacheOfEmployessToManagerId := make(map[int][]*employee)
	for _, emp := range employess {
		managerId := emp.managerId
		if managerId == 0 {
			rootEmployee = emp
		} else {
			if reportees, ok := cacheOfEmployessToManagerId[managerId]; ok {
				reportees = append(reportees, emp)
				cacheOfEmployessToManagerId[managerId] = reportees
			} else {
				reportees = make([]*employee, 0)
				reportees = append(reportees, emp)
				cacheOfEmployessToManagerId[managerId] = reportees
			}
		}
	}
	out := make([]string, 0)
	printDirectoryOfEMployeeHelper(cacheOfEmployessToManagerId, rootEmployee, "", &out)
	for _, val := range out {
		fmt.Println(val)
	}
}

func printDirectoryOfEMployeeHelper(cache map[int][]*employee, root *employee, level string, out *[]string) {
	*out = append(*out, level+root.firstname+" "+root.lastname)
	var listOfReportees []*employee
	if _, ok := cache[root.employeeId]; ok {
		listOfReportees = cache[root.employeeId]
	}
	for _, emp := range listOfReportees {
		printDirectoryOfEMployeeHelper(cache, emp, level+"  ", out)
	}
}

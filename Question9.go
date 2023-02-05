/*
Define structure EmpInfo that holds the following

Empno(5 digits) , date-of-join (dd, mm, yy) , locationcode (string of 4 characters)

Create an array of 10 elements of the structure type and initialize the structure, in random
order.

Implement the functions
	- printEmpByLocation() which takes pointer to EmpInfo structure & location code as arguments, and prints list of employees & their date of join.
	- printEmpByYearOfJoin() which takes pointer to EmpInfo structure & year as arguments, and prints list of employees, date of join & location.

The program should support the following command line agruments:
	-l locationcode 	--> 	should print employees of the specified location
	-y year 					--> 	should print employees in that particular year
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

type DOJ struct {
	dd, mm uint8
	yyyy   uint16
}

type Empinfo struct {
	Empno        uint32
	date_of_join DOJ
	locationcode string
}

func printEmpByLocation(employeeArray *[10]Empinfo, locationcode string) {
	for _, employee := range *employeeArray {
		if employee.locationcode == locationcode {
			fmt.Println("Employee ID =", employee.Empno, "&", "Date Of Join =", employee.date_of_join)
		}
	}
}

func printEmpByYearOfJoin(employeeArray *[10]Empinfo, year uint16) {
	for _, employee := range *employeeArray {
		if employee.date_of_join.yyyy == year {
			fmt.Println(employee)
		}
	}
}

func main() {
	employeeArray := [10]Empinfo{{15122, DOJ{12, 10, 2003}, "KOLK"}, {34522, DOJ{11, 02, 2010}, "HYDR"}, {11111, DOJ{01, 10, 2015}, "BIHA"},
		{32978, DOJ{12, 12, 2018}, "DELH"}, {25743, DOJ{12, 04, 2008}, "CHAN"}, {17194, DOJ{12, 10, 1998}, "DELH"}, {42694, DOJ{12, 10, 1998}, "BIHA"},
		{89145, DOJ{12, 10, 1998}, "HYDR"}, {30119, DOJ{12, 10, 1998}, "KOLK"}, {45433, DOJ{12, 10, 1998}, "BENG"}}
	if len(os.Args) != 3 {
		fmt.Println("Invalid Argument Types")
	} else if os.Args[1] == "-l" {
		printEmpByLocation(&employeeArray, os.Args[2])
	} else if os.Args[1] == "-y" {
		year, _ := strconv.ParseUint(os.Args[2], 10, 16)
		printEmpByYearOfJoin(&employeeArray, uint16(year))
	}

}

package main

import (
	"fmt"

	"nurgazinovd_golang1/employee"
	"nurgazinovd_golang1/engineer"
	"nurgazinovd_golang1/finance"
	"nurgazinovd_golang1/hr"
	"nurgazinovd_golang1/manager"
	"nurgazinovd_golang1/sales"
)

func main() {

	// Create employees
	mgr := manager.Manager{
		Position: "CEO",
		Salary:   100000,
		Address:  "Tole bi 59",
	}

	eng := engineer.Engineer{
		Position: "Software Engineer",
		Salary:   80000,
		Address:  "Satpaev 14",
	}

	// Print employee info
	printEmployee(&mgr)
	printEmployee(&eng)

	// Create Sales, HR, Finance employees
	salesRep := sales.Sales{
		Position: "Salesman",
		Salary:   50000,
		Address:  "Japanese town 111343",
	}

	hrManager := hr.HR{
		Position: "HR Manager",
		Salary:   55123,
		Address:  "Planet 808",
	}

	accountant := finance.Finance{
		Position: "Accountant Finance Manager",
		Salary:   63400,
		Address:  "888 Pushkina Kolotushkina",
	}

	printEmployee(&salesRep)
	printEmployee(&hrManager)
	printEmployee(&accountant)

}

func printEmployee(e employee.Employee) {
	fmt.Println(e.GetPosition())
	fmt.Println(e.GetSalary())
	fmt.Println(e.GetAddress())
}

package models

import (
	"database/sql"
	"fmt"

	"../db"
)

//Employee Define
type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json:"employee_salary"`
	Age    string `json:"employee_age"`
}

//Employees array Define
type Employees struct {
	Employees []Employee `json:"employee"`
}

var con *sql.DB

//GetEmployee get data of employee
func GetEmployee() Employees {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result

}

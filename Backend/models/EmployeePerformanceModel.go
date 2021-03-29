package models

import (
	"backend-nbs/helpers"

	"github.com/google/uuid"
)

type EmployeePerformance struct {
	Id          string `json:"id"`
	Employee_Id string `json:"employee_id"`
	Score       int    `json:"score"`
	Created_At  string `json:"createdAt"`
	Updated_At  string `json:"updateAt"`
}

func CreateEmployeePerformance(employee_id string) error {

	db, err := helpers.ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("insert into employees_performance (id,employee_id) VALUES(?, ?)", uuid.NewString(), employee_id)
	return err
}

func GetEmployeePerformanceById(employee_id string) EmployeePerformance {
	db, err := helpers.ConnectMySQL()
	helpers.LogFatal(err)
	defer db.Close()

	employeePerformance := &EmployeePerformance{}
	db.QueryRow("select id,employee_id,score,created_at,updated_at from employees_performance where employee_id = ?", employee_id).Scan(
		&employeePerformance.Id, &employeePerformance.Employee_Id, &employeePerformance.Score, &employeePerformance.Created_At, &employeePerformance.Updated_At)
	return *employeePerformance
}
func GetEmployeePerformance() EmployeePerformance {
	db, err := helpers.ConnectMySQL()
	helpers.LogFatal(err)
	defer db.Close()

	employeePerformance := &EmployeePerformance{}
	db.QueryRow("select id,employee_id,score,created_at,updated_at from employees_performance").Scan(
		&employeePerformance.Id, &employeePerformance.Employee_Id, &employeePerformance.Score, &employeePerformance.Created_At, &employeePerformance.Updated_At)
	return *employeePerformance
}

func UpdateScoreEmployeePerformance(score int, employee_id string, currentDate string) error {
	db, err := helpers.ConnectMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("update employees_performance set score = ?,updated_at = ? where employee_id = ?", score, currentDate, employee_id)
	return err

}

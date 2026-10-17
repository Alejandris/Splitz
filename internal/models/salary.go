package models

// AddSalaryRequest representa el salario neto ingresado por el cliente.
type AddSalaryRequest struct {
	NetSalary float64 `json:"net_salary"`
}

// AddSalaryResponse representa el resultado listo para el siguiente endpoint.
type AddSalaryResponse struct {
	NetSalary          float64 `json:"net_salary"`
	Salary             float64 `json:"salary"`
	ReadyForProcessing bool    `json:"ready_for_processing"`
}

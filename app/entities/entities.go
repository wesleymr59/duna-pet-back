package entities

type MessageHealthy struct {
	Message string `json:"message"`
}

type User struct {
	ID     int    `gorm:"primaryKey,autoIncrement:true;"`
	Nome   string `json:"nome"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Active int    `json:"active"`
}

type Login struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type Dog struct {
	ID   int    `gorm:"primaryKey,autoIncrement:true;"`
	Race string `json:"race"`
}

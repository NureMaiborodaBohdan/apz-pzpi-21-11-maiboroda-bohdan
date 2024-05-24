package AlcoSafe

type User struct {
	UserID     int    `json:"UserID" db:"UserID"`
	Username   string `json:"Username" binding:"required" db:"Username"`
	Email      string `json:"Email" binding:"required" db:"Email"`
	Password   string `json:"Password" binding:"required" db:"Password"`
	Role       string `json:"Role" db:"Role"`
	Name       string `json:"Name"`
	Surname    string `json:"Surname"`
	Patronymic string `json:"Patronymic"`
	CompanyID  int    `json:"CompanyID"`
	Sex        string `json:"Sex"`
}

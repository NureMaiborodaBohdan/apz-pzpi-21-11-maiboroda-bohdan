package AlcoSafe

type User struct {
	UserID     uint64 `json:"UserID"`
	Username   string `json:"Username"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Role       string `json:"Role"`
	Name       string `json:"Name"`
	Surname    string `json:"Surname"`
	Patronymic string `json:"Patronymic"`
	CompanyID  int    `json:"CompanyID"`
	Sex        string `json:"Sex"`
}

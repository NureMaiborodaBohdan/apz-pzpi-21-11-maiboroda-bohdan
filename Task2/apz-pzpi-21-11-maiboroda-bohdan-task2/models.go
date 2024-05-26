package AlcoSafe

type TestResult struct {
	TestID       int     `json:"TestID" db:"TestID"`
	UserID       int     `json:"UserID" db:"UserID" binding:"required"`
	TestTime     string  `json:"TestTime" db:"TestTime" binding:"required"`
	AlcoholLevel float64 `json:"AlcoholLevel" db:"AlcoholLevel" binding:"required"`
	Description  string  `json:"Description" db:"Description"`
}

type Location struct {
	LocationID int    `json:"LocationID" db:"LocationID"`
	Country    string `json:"Country" db:"Country" binding:"required"`
	City       string `json:"City" db:"City" binding:"required"`
	Address    string `json:"Address" db:"Address" binding:"required"`
	PostCode   int    `json:"PostCode" db:"PostCode" binding:"required"`
}

type Company struct {
	CompanyID   int     `json:"CompanyID" db:"CompanyID"`
	Name        string  `json:"Name" db:"Name" binding:"required"`
	Description string  `json:"Description" db:"Description"`
	LegalLimit  float64 `json:"LegalLimit" db:"LegalLimit" binding:"required"`
	LocationID  int     `json:"LocationID" db:"LocationID" binding:"required"`
}

type Notification struct {
	NotificationID int    `json:"NotificationID" db:"NotificationID"`
	Message        string `json:"Message" db:"Message" binding:"required"`
	SentTime       string `json:"SentTime" db:"SentTime" binding:"required"`
	UserID         int    `json:"UserID" db:"UserID" binding:"required"`
}

type ThresholdValues struct {
	ThresholdID int     `json:"ThresholdID" db:"ThresholdID"`
	LegalLimit  float64 `json:"LegalLimit" db:"LegalLimit" binding:"required"`
}

type AccessControl struct {
	AccessID   int    `json:"AccessID" db:"AccessID"`
	UserID     int    `json:"UserID" db:"UserID" binding:"required"`
	AccessTime string `json:"AccessTime" db:"AccessTime" binding:"required"`
}

package AlcoSafe

type TestResult struct {
	TestID       int     `json:"TestID"`
	UserID       int     `json:"UserID"`
	TestTime     string  `json:"TestTime"`
	AlcoholLevel float64 `json:"AlcoholLevel"`
	Description  string  `json:"Description"`
}

type Location struct {
	LocationID int    `json:"LocationID"`
	Country    string `json:"Country"`
	City       string `json:"City"`
	Address    string `json:"Address"`
	PostCode   int    `json:"PostCode"`
}

type Company struct {
	CompanyID   int    `json:"CompanyID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	LocationID  int    `json:"LocationID"`
	ThresholdID int    `json:"ThresholdID"`
}

type Notification struct {
	NotificationID int    `json:"NotificationID"`
	Message        string `json:"Message"`
	SentTime       string `json:"SentTime"`
	UserID         int    `json:"UserID"`
}

type ThresholdValues struct {
	ThresholdID int     `json:"ThresholdID"`
	LegalLimit  float64 `json:"LegalLimit"`
}

type AccessControl struct {
	AccessID   int    `json:"AccessID"`
	UserID     int    `json:"UserID"`
	AccessTime string `json:"AccessTime"`
}

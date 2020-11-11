package backend

// modal for user
type User struct {
	ID             string `json:"UserId`
	Firstname      string `json:"FirstName"`
	Lastname       string `json:"LastName"`
	Instructor     bool   `json:"Instructor"`
	Workshopsgiven int    `json:"workshopsGiven"`
}

// modal for uworkshop

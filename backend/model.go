package backend

// modal for user
type User struct {
	ID             string `json:"_id`
	Firstname      string `json:"firstName"`
	Lastname       string `json:"lastName"`
	Instructor     bool   `json:"instructor"`
	Workshopsgiven int    `json:"workshopsGiven"`
}

// modal for uworkshop

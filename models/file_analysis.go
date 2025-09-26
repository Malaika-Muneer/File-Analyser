package models

// Struct for storing file analysis results
type FileAnalysis struct {
	User_id      int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Vowels       int    `json:"vowels"`
	Consonants   int    `json:"consonants"`
	Digits       int    `json:"digits"`
	SpecialChars int    `json:"special_chars"`
	Letters      int    `json:"letters"`
	UpperCase    int    `json:"upper_case"`
	LowerCase    int    `json:"lower_case"`
	Spaces       int    `json:"spaces"`
	TotalChars   int    `json:"total_hars"`
}

// User struct represents a user for sign-up
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//struct represent a user for sign-in
type SignIn struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

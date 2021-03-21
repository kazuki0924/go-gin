package entity

// Person struct
type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=125"`
	Email     string `json:"email" validate:"required,email"`
}

// Video struct
type Video struct {
	// Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Title       string `json:"title" binding:"min=2,max=100"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

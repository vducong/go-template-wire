package user

type User struct {
	ID       string  `json:"id"`
	FullName *string `json:"full_name"`
}

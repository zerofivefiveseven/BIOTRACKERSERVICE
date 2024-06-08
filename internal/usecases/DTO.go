package usecases

type UserDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

type CredentialsDTO struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AddPlotDTO struct {
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type PlotDTO struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

package domain

type Responsible struct {
	Id_reponsible int `json:"id_reponsible,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Id_user int `json:"id_user"`
}
package model

type UserModel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var UserData = []UserModel{
	{ID: "1", Name: "mas satu"},
	{ID: "2", Name: "mas dua"},
}

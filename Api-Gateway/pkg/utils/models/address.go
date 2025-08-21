package models

type Address struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	District string `json:"district"`
	ZipCode  string `json:"zip_code"`
	Country  string `json:"country"`
}

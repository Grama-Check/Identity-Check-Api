package models

type UserData struct {
	NIC string `json:"nic"`
}

type PersonData struct {
	NIC     string `json:"nic"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

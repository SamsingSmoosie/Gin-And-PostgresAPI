package model

type Friend struct {
	ID   string `json:"id"`
	Name name   `json:"name"`
}

type name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type address struct {
	HouseNumber int    `json:"houseNumber"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     int    `json:"zipCode"`
}

type Person struct {
	ID         string   `json:"id"`
	Index      int      `json:"index"`
	GUID       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       name     `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    address  `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Friends    []Friend `json:"friends"`
}

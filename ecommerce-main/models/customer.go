package models

type Customer struct {
	CustomerId              string            `json:"customerid" bson:"customerid"`
	Firstname               string            `json:"firstname" bson:"firstname"`
	Lastname                string            `json:"lastname" bson:"lastname"`
	HashesAndSaltedPassword string            `json:"hashesandsaltedpassword" bson:"hashedandsaltedpassword"`
	Email                   string            `json:"email" bson:"email"`
	Address                 []Address         `json:"address" bson:"address"`
	ShippingAddress         []ShippingAddress `json:"shippingaddress" bson:"shippingaddress"`
}

type Address struct {
	Country string `json:"country" bson:"country"`
	Street1 string `json:"street1" bson:"street1"`
	Street2 string `json:"street2" bson:"street2"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Zip     string `json:"zip" bson:"zip"`
}

type ShippingAddress struct {
	Street1 string `json:"street1" bson:"street1"`
	Street2 string `json:"street2" bson:"street2"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
}

type CustomerDBResponse struct {
	Customer_id string `json:"customerid" bson:"customerid"`
}

type Token struct {
	CustomerId string `json:"customerid" bson:"customerid"`
	Email      string `json:"email" bson:"email"`
	Token      string `json:"token" bson:"token"`
}

type UpdatePassword struct {
	Email       string `json:"email" bson:"email"`
	OldPassword string `json:"oldpassword" bson:"oldpassword"`
	NewPassword string `json:"newpassword" bson:"newpassword"`
}
type UpdateEmail struct {
	CustomerId string `json:"customerid" bson:"customerid"`
	OldEmail   string `json:"oldemail" bson:"oldemail"`
	NewEmail   string `json:"newemail" bson:"newemail"`
}

type UpdateRequest struct {
	CustomerId string `json:"customerid"`
	Field      string `json:"field"`
	OldValue      string `json:"oldvalue"`
	NewValue    string  `json:"newvalue"`
}

type DeleteRequest struct {
	CustomerId string `json:"customerid"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User model schema
type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductsID []primitive.ObjectID `json:"_productsid,omitempty" bson:"_productsid,required"`
	Firstname string  `json:"firstname,omitempty" bson:"firstname,required"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,required"`
	Username string `json:"username,omitempty" bson:"username,required"`
	Email string `json:"email,omitempty" bson:"email,required"`
}

//Product model schema
type Product struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MerchantID primitive.ObjectID `json:"_merchantid,omitempty" bson:"_merchantid,omitempty"`
	Title string `json:"title,omitempty" bson:"title,required"`
	Location string `json:"location,omitempty" bson:"location,required"`
	PricePerDay string `json:"priceperday,omitempty" bson:"priceperday,required"`
	Pic string `json:"url,omitempty" bson:"url,omitempty"`
}
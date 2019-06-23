package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"
	"models"
)

// Storage type stores instance of database 
type Storage struct {
	db *mongo.Database
	ctx context.Context
}

// NewStorage returns a reference to a storage struct 
func NewStorage(ctx context.Context, databaseName string) *Storage {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if (err != nil){
		fmt.Printf("Error: %v while connecting to mongo", err)
	}

	err = client.Connect(ctx)

	if (err != nil){
		fmt.Printf("Error: %v while connecting mongo client to background context", err)
	}

	return &Storage{client.Database(databaseName), ctx}
}

// GetCollection returns a reference to a mongo Collection 
func (s *Storage) GetCollection(str string) *mongo.Collection{
	return s.db.Collection(str) 
}

//NewProduct returns the id of the new product created
func (s *Storage) NewProduct(p *models.Product, userID string)(string, error){
	collection := s.GetCollection("Products")
	res, err := collection.InsertOne(s.ctx, p)

	if err != nil {
		fmt.Printf("Error: %v while creating a product", err)
	}
	s.AddProductToUser(res.InsertedID.(primitive.ObjectID).Hex(), userID)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

//GetProductsList displays a list of all the products in the database
func (s *Storage) GetProductsList()([] *models.Product){
	var prs []*models.Product

	collection := s.GetCollection("Products")
	cursor, err := collection.Find(s.ctx, bson.D{{}})
	defer cursor.Close(s.ctx)

	if err != nil {
		fmt.Printf("Error: %v while finding all the products", err)
	}

	for cursor.Next(s.ctx) {
		var p models.Product
		err := cursor.Decode(&p)

		if err != nil {
			fmt.Printf("Error: %v while reading a product", err)
		}

		prs = append(prs, &p)

	}

	return prs
}

//GetAllProducts displays all the products being sold by the user
func (s *Storage) GetAllProducts(userID string)([] *models.Product){
	u := s.GetUserByID(userID)
	productsArray := u.ProductsID
	var prs []*models.Product

	collection := s.GetCollection("Products")
	cursor, err := collection.Find(s.ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: productsArray}}}})
	defer cursor.Close(s.ctx)

	if err != nil {
		fmt.Printf("Error: %v while finding all the products", err)
	}

	for cursor.Next(s.ctx) {
		var p models.Product
		err := cursor.Decode(&p)

		if err != nil {
			fmt.Printf("Error: %v while reading a product", err)
		}

		prs = append(prs, &p)

	}

	return prs

}

//AddProductToUser adds product to user model
func (s *Storage) AddProductToUser(prodID string, userID string){
	u := s.GetUserByID(userID)
	objID, _ := primitive.ObjectIDFromHex(prodID)
	u.ProductsID = append(u.ProductsID, objID)
	// fmt.Printf("Products Array: %v", u.ProductsID)
	fmt.Printf("User: %v", u)
	s.UpdateUserByID(userID, &u)
}

// NewUser returns the id of the user created
func (s *Storage) NewUser(u *models.User)(string, error) {
	collection := s.GetCollection("Users")
	res, err := collection.InsertOne(s.ctx, u)

	if err != nil {
		fmt.Printf("Error: %v while creating a user", err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

//GetAllUsers returns all the users stored in the database
func (s *Storage) GetAllUsers()([] *models.User){
	var users []*models.User
	collection := s.GetCollection("Users")
	cursor, err := collection.Find(s.ctx, bson.D{{}})
	defer cursor.Close(s.ctx)

	if err != nil {
		fmt.Printf("Error: %v while finding all the users", err)
	}

	for cursor.Next(s.ctx) {
		var u models.User
		err := cursor.Decode(&u)

		if err != nil {
			fmt.Printf("Error: %v while reading a user", err)
		}

		users = append(users, &u)

	}

	return users
}

// GetUserByID returns user by ID
func (s *Storage) GetUserByID(ID string)(models.User){
	var u models.User
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Users")
	filter := bson.D{{Key: "_id", Value: objID}}
	err := collection.FindOne(s.ctx, filter).Decode(&u)
	fmt.Printf("User First Name: %v and products: %v", u.Firstname, u.ProductsID)
	if err != nil {
		fmt.Printf("Error: %v while finding a user", err)
	}

	return u 
}

//GetProductByID gets product based on ID parameter passed into it 
func (s *Storage) GetProductByID(ID string)(models.Product){
	var p models.Product
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Products")
	filter := bson.D{{Key: "_id", Value: objID}}
	err := collection.FindOne(s.ctx, filter).Decode(&p)
	if err != nil {
		fmt.Printf("Error: %v while finding a product", err)
	}

	return p
}

//DeleteProduct deletes product by id and from user model
func (s *Storage) DeleteProduct(userID string, prodID string)(* mongo.DeleteResult, *mongo.UpdateResult){
	deleteResult := s.DeleteProductByID(prodID)
	u := s.GetUserByID(userID)
	u.ProductsID = RemoveElementFromArray(u.ProductsID, prodID)
	// fmt.Printf("Products Array: %v", u.ProductsID)
	fmt.Printf("User: %v", u)
	updateResult := s.UpdateUserByID(userID, &u)
	return deleteResult,updateResult
}

//RemoveElementFromArray removes element from array
func RemoveElementFromArray(a []primitive.ObjectID, el string)[]primitive.ObjectID{
	objID, _ := primitive.ObjectIDFromHex(el)
    var ind int
	for i,id := range a {
		if (id == objID) {
			ind = i
		}
	}

	a = append(a[:ind], a[ind + 1:]...)
	return a
}

//DeleteProductByID deletes product based on ID parameter passed into it 
func (s *Storage) DeleteProductByID(ID string)(* mongo.DeleteResult){
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Products")
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := collection.DeleteOne(s.ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting a user", err)
	}
	return result

}

// DeleteUserByID deletes user based on ID parameter passed into it 
func (s *Storage) DeleteUserByID(ID string)(* mongo.DeleteResult){
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Users")
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := collection.DeleteOne(s.ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting a user", err)
	}
	return result
}

// DeleteAllUsers returns the count of users that have been succesfully deleted
func (s *Storage) DeleteAllUsers()(* mongo.DeleteResult){
	collection := s.GetCollection("Users")
	filter := bson.D{{}}
	result, err := collection.DeleteMany(s.ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting all users", err)
	}

	return result 
}

//DeleteAllProducts retuns the count of products that have been succesfully deleted
func (s *Storage) DeleteAllProducts()(* mongo.DeleteResult){
	collection := s.GetCollection("Products")
	filter := bson.D{{}}
	result, err := collection.DeleteMany(s.ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting all users", err)
	}

	return result

}

//UpdateProductByID updates the product by ID
func (s *Storage) UpdateProductByID(ID string, p *models.Product) *mongo.UpdateResult{
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Products")
	filter := bson.D{{Key: "_id", Value: objID}}
	updateInterface := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: p.Title}, 
		{Key: "location", Value: p.Location},
		{Key: "priceperday", Value: p.PricePerDay}, 
	}}}
	result, err := collection.UpdateOne(s.ctx, filter, updateInterface)
	if err != nil {
		fmt.Printf("Error: %v while updating a user", err)
	}
	return result
}

// UpdateUserByID returns the count of users that have been successfully updated
func (s *Storage) UpdateUserByID(ID string,u *models.User) *mongo.UpdateResult{
	objID,_ := primitive.ObjectIDFromHex(ID)
	collection := s.GetCollection("Users")
	filter := bson.D{{Key: "_id",Value: objID}}
	updateInterface := bson.D{{Key: "$set", Value: bson.D{
										  {Key: "firstname", Value: u.Firstname}, 
										  {Key: "lastname", Value: u.Lastname}, 
										  {Key: "username", Value: u.Username}, 
										  {Key: "email", Value: u.Email}, 
										  {Key: "_productsid", Value: u.ProductsID},
							}}}
	result, err := collection.UpdateOne(s.ctx, filter, updateInterface) 
	fmt.Printf("Updating array : %v", u.ProductsID)					
	if err != nil {
		fmt.Printf("Error: %v while updating a user", err)
	}

	return result
}
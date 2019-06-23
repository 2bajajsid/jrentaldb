package service 

import (
	"models"
	"storage"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

//Service defines service structure
type Service struct {
	 strg *storage.Storage
}

//NewService creates a new service instance 
func NewService(ctx context.Context, databaseName string) *Service{
	return &Service{strg: storage.NewStorage(ctx, databaseName)}
}

//CreateNewUser creates new user
func (s *Service) CreateNewUser(u *models.User)(string, error) {
	id, err := s.strg.NewUser(u)
	return id, err
}

//CreateNewProduct creates new product
func (s *Service) CreateNewProduct(p *models.Product, userID string)(string, error) {
	id, err := s.strg.NewProduct(p, userID)
	return id, err
}

//GetProductsByUserID retrieves the list of products for every user
func (s *Service) GetProductsByUserID(ID string)([] *models.Product){
	return s.strg.GetAllProducts(ID)
}

//GetProductsList displays a list of all the products in the database
func (s *Service) GetProductsList()([] *models.Product) {
	return s.strg.GetProductsList()
}

//GetUsers returns all users
func (s *Service) GetUsers()([] *models.User){
	return s.strg.GetAllUsers()
}

//GetUser returns user through ID
func (s *Service) GetUser(ID string) models.User {
	return s.strg.GetUserByID(ID)
}

//GetProduct returns product through ID
func (s *Service) GetProduct(ID string) models.Product{
	return s.strg.GetProductByID(ID)
}

//DeleteUser delete user through id
func (s *Service) DeleteUser(ID string) (* mongo.DeleteResult){
	return s.strg.DeleteUserByID(ID)
}

//DeleteAll deletes all users
func (s *Service) DeleteAll() (* mongo.DeleteResult){
	return s.strg.DeleteAllUsers()
}

//DeleteAllProducts deletes all products
func (s *Service) DeleteAllProducts() (* mongo.DeleteResult){
	return s.strg.DeleteAllProducts()
}

//DeleteProduct deletes product based on ID parameter passed into it 
func (s *Service) DeleteProduct(userID string, prodID string) (* mongo.DeleteResult, *mongo.UpdateResult){
	return s.strg.DeleteProduct(userID, prodID)
}

//UpdateUser updates a user
func (s *Service) UpdateUser(ID string, u *models.User) (*mongo.UpdateResult){
	return s.strg.UpdateUserByID(ID, u)
}

//UpdateProduct updates a user
func (s *Service) UpdateProduct(ID string, u *models.Product) (*mongo.UpdateResult){
	return s.strg.UpdateProductByID(ID, u)
}
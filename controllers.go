package controllers

import(
	"service"
	"models"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Controller define controller structure
type Controller struct {
	service *service.Service
}

//NewController initializes a new controller 
func NewController(serv *service.Service) *Controller {
	return &Controller{service: serv}
}

//AddUserHandler is a handler for adding users
func (c *Controller) AddUserHandler(cxt *gin.Context){
	Firstname := cxt.PostForm("firstname")
	Lastname := cxt.PostForm("lastname")
	Username := cxt.PostForm("username")
	Email := cxt.PostForm("email")
	newUser := models.User {
		Firstname: Firstname,
		Lastname: Lastname,
		Username: Username, 
		Email: Email,
	}
	id,_ := c.service.CreateNewUser(&newUser)
	fmt.Printf("New User ID: %s", id)
}

//GetProductsHandler is a handler for retrieving all products
func (c *Controller) GetProductsHandler(cxt *gin.Context)(){
	ID := cxt.Param("id")
	cxt.JSON(http.StatusOK, gin.H{
		"products": c.service.GetProductsByUserID(ID), 
	})
}

//AddProductHandler is a handler for adding users
func (c *Controller) AddProductHandler(cxt *gin.Context){
	Title := cxt.PostForm("title")
	Location := cxt.PostForm("location")
	PricePerDay := cxt.PostForm("priceperday")
	Pic := cxt.PostForm("pic")
	ID := cxt.Param("id")
	MerchantID,_ := primitive.ObjectIDFromHex(ID)
	newProd := models.Product {
		Title: Title,
		Location: Location,
		PricePerDay: PricePerDay, 
		Pic: Pic,
		MerchantID: MerchantID,
	}
	id,_ := c.service.CreateNewProduct(&newProd, ID)
	fmt.Printf("ID of new Product: %v", id)
}

//GetProductsList displays a list of all the products in the database
func (c *Controller) GetProductsList(cxt *gin.Context){
	cxt.JSON(http.StatusOK, gin.H{
		"products": c.service.GetProductsList(), 
	})
}

//GetProductByIDHandler is a handler for getting a product by id
func (c *Controller) GetProductByIDHandler(cxt *gin.Context){
	ID := cxt.Param("id")
	product := c.service.GetProduct(ID)
	cxt.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

//GetUsersHandler is a handler for getting all users
func (c *Controller) GetUsersHandler(cxt *gin.Context){
	users := c.service.GetUsers()
	cxt.JSON(http.StatusOK, gin.H{
		"users": users, 
	})
}

//GetUserByIDHandler is a handler for getting a user by id
func (c *Controller) GetUserByIDHandler(cxt *gin.Context){
	ID := cxt.Param("id")
	user := c.service.GetUser(ID)
	cxt.JSON(http.StatusOK, gin.H{
		"user": user, 
	})
}

//DeleteUsersHandler is a handler for deleting users 
func (c *Controller) DeleteUsersHandler(cxt *gin.Context){
	result := c.service.DeleteAll()
	cxt.JSON(http.StatusOK, gin.H{
		"result": result, 
	})
}

//DeleteUserHandler is a handler for deleting user by id
func (c *Controller) DeleteUserHandler(cxt *gin.Context){
	ID := cxt.Param("id")
	result := c.service.DeleteUser(ID)
	cxt.JSON(http.StatusOK, gin.H{
		"result": result, 
	})
}

//DeleteProductHandler deletes product 
func (c *Controller) DeleteProductHandler(cxt *gin.Context){
	uID := cxt.Param("id")
	pID := cxt.Param("prodid")
	result1, result2 := c.service.DeleteProduct(uID, pID)
	cxt.JSON(http.StatusOK, gin.H{
		"delete result": result1, 
		"update result": result2,
	})
}

//DeleteProducts deletes all products
func (c *Controller) DeleteProducts(cxt *gin.Context) {
	result := c.service.DeleteAllProducts()
	cxt.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//UpdateUserByIDHandler is a handler for updating user by id
func (c *Controller) UpdateUserByIDHandler(cxt *gin.Context){
	ID := cxt.Param("id")
	Firstname := cxt.PostForm("firstname")
	Lastname := cxt.PostForm("lastname")
	Username := cxt.PostForm("username")
	Email := cxt.PostForm("email")
	newUser := models.User {
		Firstname: Firstname,
		Lastname: Lastname,
		Username: Username, 
		Email: Email,
	}
	result := c.service.UpdateUser(ID, &newUser)
	cxt.JSON(http.StatusOK, gin.H{
		"result": result, 
	})
}

//UpdateProductByIDHandler is a handler for updating product by id
func (c *Controller) UpdateProductByIDHandler(cxt *gin.Context){
	ID := cxt.Param("prodid")
	Title := cxt.PostForm("title")
	Location := cxt.PostForm("location")
	PricePerDay := cxt.PostForm("priceperday")
	newProd := models.Product {
		Title: Title,
		Location: Location,
		PricePerDay: PricePerDay,
	}
	result := c.service.UpdateProduct(ID, &newProd)
	cxt.JSON(http.StatusOK, gin.H{
		"result": result, 
	})
}
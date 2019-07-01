package main 

import( 
	"fmt"
	 "github.com/gin-gonic/gin"
	 "github.com/gin-contrib/cors"
 	"go.mongodb.org/mongo-driver/mongo"
 	"go.mongodb.org/mongo-driver/mongo/options"
 	"go.mongodb.org/mongo-driver/bson/primitive"
 	"context"
 	"go.mongodb.org/mongo-driver/bson"
	 "models"
	 "router"
	 "service"
	 "controllers"
	//  "http"
)
// import "time"


func main(){

	//gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()
	serv := service.NewService(ctx, "JRentalsDB")
	contr := controllers.NewController(serv)
	eng := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	eng.Use(cors.New(config))

	router.NewRouters(contr, eng)
	// ctx := context.Background()
	// JRentalsDB,_ := configDB(ctx)

	// r := gin.Default()
	// r.GET("Users", func (c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"users": getUsers(ctx, JRentalsDB), 
	// 	})
	// }) 

	// r.GET("/User/:id", func(c *gin.Context){
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"user": getUser(ctx, JRentalsDB, id), 
	// 	})
	// })

	// r.POST("/User", func(c *gin.Context){
	// 	Firstname := c.PostForm("firstname")
	// 	Lastname := c.PostForm("lastname")
	// 	Username := c.PostForm("username")
	// 	Email := c.PostForm("email")
	// 	newUser := models.User {
	// 		Firstname: Firstname,
	// 		Lastname: Lastname,
	// 		Username: Username, 
	// 		Email: Email,
	// 	}
	// 	id,_ := createUser(ctx, JRentalsDB, newUser)
	// 	fmt.Printf("New User ID: %s", id)
	// })

	// r.PUT("/User/:id", func(c *gin.Context){
	// 	id := c.Param("id")
	// 	Firstname := c.PostForm("firstname")
	// 	Lastname := c.PostForm("lastname")
	// 	Username := c.PostForm("username")
	// 	Email := c.PostForm("email")
	// 	newUser := models.User {
	// 		Firstname: Firstname,
	// 		Lastname: Lastname,
	// 		Username: Username, 
	// 		Email: Email,
	// 	}
	// 	result := updateUser(ctx, JRentalsDB, id, &newUser)
	// 	c.JSON(200, gin.H{
	// 		"result": result, 
	// 	})
	// })

	// r.DELETE("/User/:id", func(c *gin.Context){
	// 	id := c.Param("id")
	// 	result := deleteUser(ctx, JRentalsDB, id)
	// 	c.JSON(200, gin.H{
	// 		"result": result,
	// 	})
	// })

	// r.DELETE("/Users", func(c *gin.Context){
	// 	result := deleteAllUsers(ctx, JRentalsDB)
	// 	c.JSON(200, gin.H{
	// 		"result": result, 
	// 	})
	// })

	// r.Run() 
}

func configDB(ctx context.Context)(*mongo.Database, error){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if (err != nil){
		fmt.Printf("Error: %v while connecting to mongo", err)
	}

	err = client.Connect(ctx)

	if (err != nil){
		fmt.Printf("Error: %v while connecting mongo client to background context", err)
	}

	JRentalsDB := client.Database("JRental")
	return JRentalsDB, nil
}

func getCollection(str string, db *mongo.Database)(*mongo.Collection){
	return db.Collection(str)
}

func createUser(ctx context.Context, db *mongo.Database, u models.User)(string, error){
	collection := getCollection("Users", db)
	res, err := collection.InsertOne(ctx, u)

	if err != nil {
		fmt.Printf("Error: %v while creating a user", err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func getUsers(ctx context.Context, db *mongo.Database)([]*models.User){
	var users []*models.User
	collection := getCollection("Users", db)
	cursor, err := collection.Find(ctx, bson.D{{}})
	defer cursor.Close(ctx)

	if err != nil {
		fmt.Printf("Error: %v while finding all the users", err)
	}

	for cursor.Next(ctx) {
		var u models.User
		err := cursor.Decode(&u)

		if err != nil {
			fmt.Printf("Error: %v while reading a user", err)
		}

		users = append(users, &u)

	}

	return users
}

func getUser(ctx context.Context, db *mongo.Database, id string)(models.User){
	var u models.User
	objID,_ := primitive.ObjectIDFromHex(id)
	collection := getCollection("Users", db)
	filter := bson.D{{Key: "_id", Value: objID}}
	err := collection.FindOne(ctx, filter).Decode(&u)

	if err != nil {
		fmt.Printf("Error: %v while finding a user", err)
	}

	return u 
}

func deleteUser(ctx context.Context, db *mongo.Database, id string)(* mongo.DeleteResult){
	objID,_ := primitive.ObjectIDFromHex(id)
	collection := getCollection("Users", db)
	filter := bson.D{{"_id", objID}}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting a user", err)
	}
	return result
}

func deleteAllUsers(ctx context.Context, db *mongo.Database)(* mongo.DeleteResult){
	collection := getCollection("Users", db)
	filter := bson.D{{}}
	result, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		fmt.Printf("Error: %v while deleting all users", err)
	}

	return result 
}

func updateUser(ctx context.Context, db *mongo.Database, id string, u *models.User)(* mongo.UpdateResult){
	objID,_ := primitive.ObjectIDFromHex(id)
	collection := getCollection("Users", db)
	filter := bson.D{{"_id", objID}}
	updateInterface := bson.D{{"$set", bson.D{
										  {"firstname", u.Firstname}, 
										  {"lastname", u.Lastname}, 
										  {"username", u.Username}, 
										  {"email", u.Email}, 
							}}}
	result, err := collection.UpdateOne(ctx, filter, updateInterface) 

	if err != nil {
		fmt.Printf("Error: %v while updating a user", err)
	}

	return result
}
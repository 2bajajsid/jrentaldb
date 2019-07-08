package router

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

// AppRoute defines application's route structure
type AppRoute struct {
	Group            string
	Routes           []Route
	Middlewares []gin.HandlerFunc
}

// Route defines a single route, e.g. an HTTP method, 
// the pattern, the function that will execute when the route is called
type Route struct {
	Method           string
	Pattern          string
	RouteHandler gin.HandlerFunc
}

//NewRouters groups the routes with their corresponding handlers
func NewRouters(c *controllers.Controller, e *gin.Engine){
	ar := GetRoutes(c)

	groupRoute := e.Group(ar.Group)
	groupRoute.GET("/users", ar.Routes[0].RouteHandler)
	groupRoute.GET("/user/:id", ar.Routes[1].RouteHandler)
	groupRoute.DELETE("/users", ar.Routes[2].RouteHandler)
	groupRoute.POST("/users", ar.Routes[3].RouteHandler)
	groupRoute.PUT("/user/:id", ar.Routes[4].RouteHandler)
	groupRoute.POST("/user/:id/products", ar.Routes[5].RouteHandler)
	groupRoute.GET(ar.Routes[6].Pattern, ar.Routes[6].RouteHandler)
	groupRoute.GET(ar.Routes[7].Pattern, ar.Routes[7].RouteHandler)
	groupRoute.DELETE(ar.Routes[8].Pattern, ar.Routes[8].RouteHandler)
	groupRoute.GET(ar.Routes[9].Pattern, ar.Routes[9].RouteHandler)
	groupRoute.DELETE(ar.Routes[10].Pattern, ar.Routes[10].RouteHandler)
	groupRoute.PUT(ar.Routes[11].Pattern, ar.Routes[11].RouteHandler)
	groupRoute.DELETE(ar.Routes[12].Pattern, ar.Routes[12].RouteHandler)

	// e.Run(":27017")
	e.Run()
}

//GetRoutes initializes all routers
func GetRoutes(c *controllers.Controller) AppRoute {
	return AppRoute{
		Group: "/jrental/v1",
		Middlewares: []gin.HandlerFunc{},
		Routes: []Route{
			Route{"GET", "/users", c.GetUsersHandler},
			Route{"GET", "/user/:id", c.GetUserByIDHandler},
			Route{"DELETE", "/users", c.DeleteUsersHandler},
			Route{"POST", "/users", c.AddUserHandler},
			Route{"PUT", "/user/:id", c.UpdateUserByIDHandler},
			Route{"POST", "/user/:id/products", c.AddProductHandler},
			Route{"GET", "/user/:id/products", c.GetProductsHandler},
			Route{"GET", "/products", c.GetProductsList},
			Route{"DELETE", "/user/:id/product/:prodid", c.DeleteProductHandler},
			Route{"GET", "/product/:id", c.GetProductByIDHandler},
			Route{"DELETE", "/user/:id", c.DeleteUserHandler},
			Route{"PUT", "/user/:id/product/:prodid", c.UpdateProductByIDHandler},
			Route{"DELETE", "/products", c.DeleteProducts},
		},
	}
}

// var getUsersRoute = Route{"GET", "/users", controllers.GetUsersHandler}
// var getUserRoute = Route{"GET", "/user/:id", controllers.GetUserByIDHandler}
// var deleteUsersRoute = Route{"DELETE", "/users", controllers.DeleteUsersHandler}
// var addUserRoute = Route{"POST", "/users", controllers.AddUserHandler}
// var updateUserRoute = Route{"PUT", "/users/:id", controllers.UpdateUserByIDHandler}

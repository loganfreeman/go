import (
  "github.com/iris-contrib/middleware/logger"
  "github.com/iris-contrib/middleware/cors"
  "github.com/iris-contrib/middleware/basicauth"
)
// Root level middleware
iris.Use(logger.New())
iris.Use(cors.Default())

// Group level middleware
authConfig := basicauth.Config{
    Users:      map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
    Realm:      "Authorization Required", // if you don't set it it's "Authorization Required"
    ContextKey: "mycustomkey",            // if you don't set it it's "user"
    Expires:    time.Duration(30) * time.Minute,
}

authentication := basicauth.New(authConfig)

g := iris.Party("/admin")
g.Use(authentication)

// Route level middleware
logme := func(ctx *iris.Context)  {
        println("request to /products")
        ctx.Next()
}
iris.Get("/products", logme, func(ctx *iris.Context) {
     ctx.Text(iris.StatusOK, "/products")
})

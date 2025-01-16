package routes

import (
	_ "crossx/docs"
	privateHandlers "crossx/handlers/private"
	publicHandlers "crossx/handlers/public"
	middlewares "crossx/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

type Routes struct {
	router *gin.Engine
}

func New() *Routes {
	return &Routes{
		router: gin.New(),
	}
}

func (r *Routes) SetupRouter() *gin.Engine {
	// Middlewares de base
	r.router.Use(gin.Recovery())
	r.router.Use(gin.Logger())
	r.router.Use(middlewares.RateLimiter(10)) // 10 requests/second

	//	Security headers
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"localhost:8080"},
		SSLRedirect:           false,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; style-src 'self' 'unsafe-inline'; script-src 'self' 'unsafe-inline'",
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	})

	r.router.Use(func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	})

	// Swagger route
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Routes setup
	r.setupPublicRoutes()
	r.setupProtectedRoutes()

	return r.router
}

func (r *Routes) setupPublicRoutes() {
	public := r.router.Group("/api/v1/public")
	{
		public.GET("/health", publicHandlers.HealthHandler)
	}
}

func (r *Routes) setupProtectedRoutes() {
	protected := r.router.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/test", privateHandlers.TestHandler)
	}
}

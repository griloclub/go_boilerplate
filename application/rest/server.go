package rest

import (
	"fmt"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_boilerplate/application/rest/routes/accountroute"
	"github.com/diegoclair/go_boilerplate/application/rest/routes/authroute"
	"github.com/diegoclair/go_boilerplate/application/rest/routes/pingroute"
	"github.com/diegoclair/go_boilerplate/application/rest/routes/transferroute"
	servermiddleware "github.com/diegoclair/go_boilerplate/application/rest/serverMiddleware"
	"github.com/diegoclair/go_boilerplate/domain/service"
	"github.com/diegoclair/go_boilerplate/infra/auth"
	"github.com/diegoclair/go_boilerplate/infra/config"
	"github.com/diegoclair/go_boilerplate/infra/logger"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// IRouter interface for routers
type IRouter interface {
	RegisterRoutes(appGroup, privateGroup *echo.Group)
}

type Server struct {
	routers []IRouter
	Srv     *echo.Echo
	cfg     *config.Config
}

func StartRestServer(cfg *config.Config, services *service.Services, log logger.Logger, authToken auth.AuthToken, mapper mapper.Mapper) {
	server := newRestServer(services, authToken, cfg, mapper)
	port := cfg.App.Port
	if port == "" {
		port = "5000"
	}

	log.Info(fmt.Sprintf("About to start the application on port: %s...", port))

	if err := server.Start(port); err != nil {
		panic(err)
	}
}

func newRestServer(services *service.Services, authToken auth.AuthToken, cfg *config.Config, mapper mapper.Mapper) *Server {

	srv := echo.New()
	srv.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	pingController := pingroute.NewController()
	accountController := accountroute.NewController(services.AccountService, mapper)
	authController := authroute.NewController(services.AuthService, mapper, authToken)
	transferController := transferroute.NewController(services.TransferService, mapper)

	pingRoute := pingroute.NewRouter(pingController, pingroute.RouteName)
	accountRoute := accountroute.NewRouter(accountController, accountroute.RouteName)
	authRoute := authroute.NewRouter(authController, authroute.RouteName)
	transferRoute := transferroute.NewRouter(transferController, transferroute.RouteName)

	server := &Server{Srv: srv, cfg: cfg}
	server.addRouters(accountRoute)
	server.addRouters(authRoute)
	server.addRouters(pingRoute)
	server.addRouters(transferRoute)
	server.registerAppRouters(authToken)

	server.setupPrometheus()

	return server
}

func (r *Server) addRouters(router IRouter) {
	r.routers = append(r.routers, router)
}

func (r *Server) registerAppRouters(authToken auth.AuthToken) {

	appGroup := r.Srv.Group("/")
	privateGroup := appGroup.Group("",
		servermiddleware.AuthMiddlewarePrivateRoute(authToken),
	)

	for _, appRouter := range r.routers {
		appRouter.RegisterRoutes(appGroup, privateGroup)
	}

}

func (r *Server) setupPrometheus() {

	p := prometheus.NewPrometheus(r.cfg.App.Name, nil)
	p.Use(r.Srv)
}

func (r *Server) Start(port string) error {
	return r.Srv.Start(fmt.Sprintf(":%s", port))
}

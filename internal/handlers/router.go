package handlers

import (
	"BIOTRACKERSERVICE/internal/config"
	"net/http"
)

func (c *Controller) NewServer(cfg config.HTTPServer) http.Server {
	// Router layer
	router := http.NewServeMux()

	router.HandleFunc("POST /login", c.LoginHandler)
	router.HandleFunc("POST /registration", c.RegisterUserHandler)

	router.Handle("POST /user_plots_info", c.UserAuthMiddleware(http.HandlerFunc(c.UserPlotsInfoHandler)))
	router.Handle("POST /user_plots", c.UserAuthMiddleware(http.HandlerFunc(c.UserPlotsByIdsHandler)))

	return http.Server{
		Addr:    cfg.Address,
		Handler: Logging(router),
	}
}

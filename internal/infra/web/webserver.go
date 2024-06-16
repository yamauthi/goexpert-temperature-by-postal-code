package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/configs"
)

type WebServer struct {
	Router chi.Router
	Routes []Route
	Conf   *configs.Conf
}

func NewWebServer(
	conf *configs.Conf,
	routes []Route,
) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		Routes: routes,
		Conf:   conf,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)

	for _, route := range s.Routes {
		s.Router.Method(route.Method, route.Path, route.Handler)
	}
	fmt.Println("Starting webserver on port ", s.Conf.WebServerUrl)
	err := http.ListenAndServe(s.Conf.WebServerUrl, s.Router)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"

	"github.com/eaciit/knot/knot.v1"
)

type Hello struct {
}

// http://servername/hello/morning
func (h *Hello) Morning(r *knot.WebContext) interface{} {
	return "Good morning"
}

// Serve static files and Swagger JSON using http.Handler
func swaggerHandler() http.Handler {
	return http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui")))
}

func swaggerJSONHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
}

func main() {
	ks := new(knot.Server)
	ks.Address = "localhost:8080"
	ks.Register(new(Hello), "")
	// Set up routes to serve Swagger UI and JSON using http package
	// ks.Route("/", http.FileServer(http.Dir("index.html")))
	http.Handle("/swagger/", swaggerHandler())
	http.Handle("/swagger.json", swaggerJSONHandler())
	// http.Handle("/swagger/*any", httpSwagger.WrapHandler)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	// ks.Route("/swagger/*any", SwaggerHandler)
	ks.Route("/hello", HelloHandler)
	ks.Listen()
}

// HelloHandler godoc
// @Summary Return a hello message
// @Description Get a hello message
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /api/v1/hello [get]
func HelloHandler(r *knot.WebContext) interface{} {
	return map[string]string{"message": "Hello, World!"}
}

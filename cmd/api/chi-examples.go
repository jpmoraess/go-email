package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Product struct {
	ID   int
	Name string
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my-handler"))
}

func main() {
	r := chi.NewRouter()

	// r.Use(myMiddleware)

	// r.Handle("/handler", myHandler{})

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		param2 := r.URL.Query().Get("age")
		w.Write([]byte(param + " : " + param2))
	})

	r.Get("/{productId}", func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "productId")
		w.Write([]byte(productId))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		render.DecodeJSON(r.Body, &product)
		product.ID = 10
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("request:", r.Method, "url:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

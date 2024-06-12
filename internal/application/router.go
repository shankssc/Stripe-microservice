package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shankssc/Stripe-microservice/internal/handler"
)

func setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/health", handler.HealthCheckHandler)
	// Add more routes here

	router.Route("/payments", loadPaymentRoutes)

	return router
}

func loadPaymentRoutes(router chi.Router) {
	paymentHandler := &handler.Payment{}

	router.Post("/", paymentHandler.Create)
}

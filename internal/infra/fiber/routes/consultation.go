package routes

import (
  "github.com/gofiber/fiber/v2"
  consultation_controllers "github.com/joaovds/first-go-api/internal/controllers/consultation"
)

func handleConsultationRoutes(router fiber.Router) {
  consultationRouter := router.Group("/consultations")

  consultationRouter.Get("/", consultation_controllers.GetAllConsultations)
  consultationRouter.Get("/:consultationId", consultation_controllers.GetConsultationById)
}


package routes

import (
  "github.com/gofiber/fiber/v2"
  doctor_controllers "github.com/joaovds/first-go-api/internal/controllers/doctor"
)

func handleDoctorRoutes(router fiber.Router) {
  doctorRouter := router.Group("/doctors")

  doctorRouter.Get("/", doctor_controllers.GetAllDoctors)
}


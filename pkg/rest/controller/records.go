package controller

import (
	"interview_test/pkg/helpers"
	"interview_test/pkg/model/records"

	"github.com/gofiber/fiber/v2"
)

// CreateRecord create a new record
func CreateRecord(c *fiber.Ctx) error {
	t := &records.Record{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := records.CreateRecord(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: "record successfully created",
	})
}

// GetRecord retrieves record by ID
func GetRecord(c *fiber.Ctx) error {
	id, err := helpers.StringToInt64(c.Params("id"))
	if err != nil {
		return err
	}

	data, err := records.GetRecord(id)

	if data == nil {
		return c.Status(fiber.StatusNotFound).JSON(Response{
			Status:  "error",
			Message: "no record found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: "record successfuly retrieved",
		Data:    data,
	})
}

// UpdateRecord update record by ID
func UpdateRecord(c *fiber.Ctx) error {
	id, err := helpers.StringToInt64(c.Params("id"))
	if err != nil {
		return err
	}

	t := &records.Record{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := records.UpdateRecord(id, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: "record successfully updated",
	})
}

// DeleteRecord delete record with ID
func DeleteRecord(c *fiber.Ctx) error {
	id, err := helpers.StringToInt64(c.Params("id"))
	if err != nil {
		return err
	}

	if err := records.DeleteRecord(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: "record succesfully deleted",
	})
}

// ReadyzHandler basic health check for availability of the service
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "success",
		Message: "app is healthy",
	})
}

package handlers

import (
	"rest-api-educative/models"
	"rest-api-educative/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()

	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "all items data",
		Data:    items,
	})
}

func GetItemById(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	item, err := services.GetItemById(itemID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item found",
		Data:    item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var createdItem models.Item = services.CreateItem(*itemInput)

	return c.Status(fiber.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var itemID string = c.Params("id")

	updatedItem, err := services.UpdateItem(itemID, *itemInput)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item updated",
		Data:    updatedItem,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	var result bool = services.DeleteItem(itemID)

	if result {
		return c.Status(fiber.StatusOK).JSON(models.Response[any]{
			Success: true,
			Message: "item deleted",
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "item failed to delete",
	})
}

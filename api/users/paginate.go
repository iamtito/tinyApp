package users

import (
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 5
	offset := (page - 1) * limit
	// var total int64
	log.Println(page)

	// db.Offset(offset).Limit(limit).Find(&products)
	// db.Model(&products).Count(&total)
	products := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}

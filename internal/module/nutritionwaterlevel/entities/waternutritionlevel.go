package entities

import (
	"time"
)

type NutritionWaterLevel struct {
	ID        uint      `db:"id"`
	PlantID   uint      `db:"plant_id"`
	Value     float32   `db:"value"`
	CreatedAt time.Time `db:"created_at"`
}

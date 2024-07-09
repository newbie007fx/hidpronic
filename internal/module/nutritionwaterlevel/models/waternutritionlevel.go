package models

import "hidroponic/internal/platform/validation"

type NutritionWaterLevel struct {
	PlantID   uint    `json:"plant_id"`
	Value     float32 `json:"value"`
	CreatedAt int64   `json:"created_at"`
}

type CreateNutritionWaterLevel struct {
	Value   float32 `validate:"required" json:"value"`
	PlantID uint    `validate:"required" json:"plant_id"`
}

func (c *CreateNutritionWaterLevel) Validate() error {
	return validation.Validate(c)
}

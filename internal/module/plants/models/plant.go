package models

import (
	"hidroponic/internal/errors"
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/types"
	"hidroponic/internal/platform/validation"
	"time"
)

type Plant struct {
	ID                  uint              `json:"id"`
	Name                string            `json:"name"`
	Varieties           string            `json:"varieties"`
	PlantType           types.PlantType   `json:"plant_type"`
	GenerativeAge       int               `json:"generative_age"`
	HarvestAge          int               `json:"harvest_age"`
	NutritionMin        float32           `json:"nutrition_min"`
	NutritionMax        float32           `json:"nutrition_max"`
	NutritionAdjustment float32           `json:"nutrition_adjustment"`
	NutritionTargets    []NutritionTarget `json:"nutrition_targets"`
	PHLevel             float32           `json:"ph_level"`
	Temperature         float32           `json:"temperature"`
	PlantAge            int               `json:"plant_age"`
	CurrentAge          int               `json:"current_age,omitempty"`
	CurrentGrowth       types.Growth      `json:"current_growth"`
	Status              types.Status      `json:"status"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	ActivedAt           *time.Time        `json:"actived_at"`
}

type BasicPlant struct {
	ID         uint            `json:"id"`
	Name       string          `json:"name"`
	Varieties  string          `json:"varieties"`
	PlantType  types.PlantType `json:"plant_type"`
	HarvestAge int             `json:"harvest_age"`
	Status     types.Status    `json:"status"`
	CreatedAt  time.Time       `json:"created_at"`
	ActivedAt  *time.Time      `json:"actived_at"`
}

type NutritionTarget struct {
	PlantAge      int     `json:"plant_age"`
	TargetPPM     float32 `json:"target_ppm"`
	AdditionalPPM float32 `json:"additional_ppm"`
}

type CreatePlant struct {
	Name                string          `validate:"required" json:"name"`
	Varieties           string          `validate:"required" json:"varieties"`
	PlantType           types.PlantType `validate:"required" json:"plant_type"`
	GenerativeAge       int             `json:"generative_age"`
	HarvestAge          int             `validate:"required" json:"harvest_age"`
	NutritionMin        float32         `validate:"required" json:"nutrition_min"`
	NutritionMax        float32         `validate:"required" json:"nutrition_max"`
	NutritionAdjustment float32         `json:"nutrition_adjustment"`
	PHLevel             float32         `validate:"required" json:"ph_level"`
	Temperature         float32         `validate:"required" json:"temperature"`
	PlantAge            int             `validate:"required" json:"plant_age"`
}

func (c *CreatePlant) Validate() error {
	if err := validation.Validate(c); err != nil {
		return err
	}

	plantTypeMap := constants.PlantTypeMap()
	if _, ok := plantTypeMap[c.PlantType]; !ok {
		return errors.ErrorInvalidRequestBody.New("invalid plant_type value")
	}

	return nil
}

type UpdatePlant struct {
	ID                  uint              `validate:"required" json:"id"`
	Name                string            `validate:"required" json:"name"`
	Varieties           string            `validate:"required" json:"varieties"`
	PlantType           types.PlantType   `validate:"required" json:"plant_type"`
	GenerativeAge       int               `json:"generative_age"`
	HarvestAge          int               `validate:"required" json:"harvest_age"`
	NutritionMin        float32           `validate:"required" json:"nutrition_min"`
	NutritionMax        float32           `validate:"required" json:"nutrition_max"`
	NutritionAdjustment float32           `json:"nutrition_adjustment"`
	NutritionTargets    []NutritionTarget `validate:"required" json:"nutrition_targets"`
	PHLevel             float32           `validate:"required" json:"ph_level"`
	Temperature         float32           `validate:"required" json:"temperature"`
	PlantAge            int               `validate:"required" json:"plant_age"`
}

func (u *UpdatePlant) Validate() error {
	return validation.Validate(u)
}

type UpdatePlantStatus struct {
	ID     uint         `validate:"required" json:"id"`
	Status types.Status `validate:"required" json:"status"`
}

func (u *UpdatePlantStatus) Validate() error {
	return validation.Validate(u)
}

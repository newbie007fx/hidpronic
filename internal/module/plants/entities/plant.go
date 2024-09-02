package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"hidroponic/internal/module/plants/constants"
	"hidroponic/internal/module/plants/types"
	"time"
)

type Plant struct {
	ID                  uint               `db:"id"`
	Name                string             `db:"name"`
	Description         sql.NullString     `db:"description"`
	Varieties           string             `db:"varieties"`
	PlantType           types.PlantType    `db:"plant_type"`
	GenerativeAge       int                `db:"generative_age"`
	HarvestAge          int                `db:"harvest_age"`
	NutritionMin        float32            `db:"nutrition_min"`
	NutritionMax        float32            `db:"nutrition_max"`
	NutritionAdjustment float32            `db:"nutrition_adjustment"`
	NutritionTargets    NutritionTargetMap `db:"nutrition_targets"`
	PHLevel             float32            `db:"ph_level"`
	Temperature         float32            `db:"temperature"`
	PlantAge            int                `db:"plant_age"`
	CurrentGrowth       types.Growth       `db:"current_growth"`
	Status              types.Status       `db:"status"`
	Yields              int                `db:"yields"`
	CreatedAt           time.Time          `db:"created_at"`
	UpdatedAt           time.Time          `db:"updated_at"`
	ActivatedAt         sql.NullTime       `db:"activated_at"`
	HarvestedAt         sql.NullTime       `db:"harvested_at"`
}

func (p Plant) ValidateStatus(targetStatus types.Status) error {
	if targetStatus == constants.StatusDeactivated || targetStatus == constants.StatusHarvested {
		if p.Status != constants.StatusActivated {
			return errors.New("status transition not llaowed")
		}

		return nil
	}
	if targetStatus == constants.StatusActivated {
		if p.Status != constants.StatusCreated && p.Status != constants.StatusDeactivated {
			return errors.New("status transition not llaowed")
		}
		return nil
	}

	return errors.New("status transition not llaowed")
}

type NutritionTargetMap map[int]NutritionTarget

func (a NutritionTargetMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *NutritionTargetMap) Scan(value any) error {
	if realValue, ok := value.([]byte); ok {
		return json.Unmarshal(realValue, a)
	}
	return nil
}

type NutritionTarget struct {
	TargetPPM     float32 `json:"target_ppm"`
	AdditionalPPM float32 `json:"additional_ppm"`
}

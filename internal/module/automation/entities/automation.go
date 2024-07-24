package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"hidroponic/internal/module/automation/models"
	"hidroponic/internal/module/automation/types"
	plantEntities "hidroponic/internal/module/plants/entities"
	"time"
)

type Automation struct {
	ID          uint                `db:"id"`
	PlantID     uint                `db:"plant_id"`
	BofereData  SensorData          `db:"before_data"`
	AfterData   SensorData          `db:"after_data"`
	Accuration  float32             `db:"accuration"`
	TargetPPM   float32             `db:"target_ppm"`
	Duration    int                 `db:"duration"`
	Status      types.Status        `db:"status"`
	TriggeredAt time.Time           `db:"triggered_at"`
	FinishedAt  sql.NullTime        `db:"finished_at"`
	Plant       plantEntities.Plant `db:"plants"`
}

type SensorData struct {
	models.SensorData
}

func (a SensorData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *SensorData) Scan(value any) error {
	if realValue, ok := value.([]byte); ok {
		return json.Unmarshal(realValue, a)
	}
	return nil
}

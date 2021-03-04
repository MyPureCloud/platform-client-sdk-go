package platformclientv2
import (
	"encoding/json"
)

// Schedulingsettingsrequest - Scheduling Settings
type Schedulingsettingsrequest struct { 
	// MaxOccupancyPercentForDeferredWork - Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`


	// DefaultShrinkagePercent - Default shrinkage percent for scheduling
	DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`


	// ShrinkageOverrides - Shrinkage overrides for scheduling
	ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`


	// PlanningPeriod - Planning period settings for scheduling
	PlanningPeriod *Valuewrapperplanningperiodsettings `json:"planningPeriod,omitempty"`


	// StartDayOfWeekend - Start day of weekend for scheduling
	StartDayOfWeekend *string `json:"startDayOfWeekend,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulingsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

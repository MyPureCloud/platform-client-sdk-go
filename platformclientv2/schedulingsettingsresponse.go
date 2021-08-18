package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingsettingsresponse - Scheduling Settings
type Schedulingsettingsresponse struct { 
	// MaxOccupancyPercentForDeferredWork - Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`


	// DefaultShrinkagePercent - Default shrinkage percent for scheduling
	DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`


	// ShrinkageOverrides - Shrinkage overrides for scheduling
	ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`


	// PlanningPeriod - Planning period settings for scheduling
	PlanningPeriod *Planningperiodsettings `json:"planningPeriod,omitempty"`


	// StartDayOfWeekend - Start day of weekend for scheduling
	StartDayOfWeekend *string `json:"startDayOfWeekend,omitempty"`

}

func (u *Schedulingsettingsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingsettingsresponse

	

	return json.Marshal(&struct { 
		MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`
		
		DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`
		
		ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`
		
		PlanningPeriod *Planningperiodsettings `json:"planningPeriod,omitempty"`
		
		StartDayOfWeekend *string `json:"startDayOfWeekend,omitempty"`
		*Alias
	}{ 
		MaxOccupancyPercentForDeferredWork: u.MaxOccupancyPercentForDeferredWork,
		
		DefaultShrinkagePercent: u.DefaultShrinkagePercent,
		
		ShrinkageOverrides: u.ShrinkageOverrides,
		
		PlanningPeriod: u.PlanningPeriod,
		
		StartDayOfWeekend: u.StartDayOfWeekend,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulingsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

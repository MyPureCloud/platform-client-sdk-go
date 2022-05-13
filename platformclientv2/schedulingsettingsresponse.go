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

func (o *Schedulingsettingsresponse) MarshalJSON() ([]byte, error) {
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
		MaxOccupancyPercentForDeferredWork: o.MaxOccupancyPercentForDeferredWork,
		
		DefaultShrinkagePercent: o.DefaultShrinkagePercent,
		
		ShrinkageOverrides: o.ShrinkageOverrides,
		
		PlanningPeriod: o.PlanningPeriod,
		
		StartDayOfWeekend: o.StartDayOfWeekend,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulingsettingsresponse) UnmarshalJSON(b []byte) error {
	var SchedulingsettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingsettingsresponseMap)
	if err != nil {
		return err
	}
	
	if MaxOccupancyPercentForDeferredWork, ok := SchedulingsettingsresponseMap["maxOccupancyPercentForDeferredWork"].(float64); ok {
		MaxOccupancyPercentForDeferredWorkInt := int(MaxOccupancyPercentForDeferredWork)
		o.MaxOccupancyPercentForDeferredWork = &MaxOccupancyPercentForDeferredWorkInt
	}
	
	if DefaultShrinkagePercent, ok := SchedulingsettingsresponseMap["defaultShrinkagePercent"].(float64); ok {
		o.DefaultShrinkagePercent = &DefaultShrinkagePercent
	}
    
	if ShrinkageOverrides, ok := SchedulingsettingsresponseMap["shrinkageOverrides"].(map[string]interface{}); ok {
		ShrinkageOverridesString, _ := json.Marshal(ShrinkageOverrides)
		json.Unmarshal(ShrinkageOverridesString, &o.ShrinkageOverrides)
	}
	
	if PlanningPeriod, ok := SchedulingsettingsresponseMap["planningPeriod"].(map[string]interface{}); ok {
		PlanningPeriodString, _ := json.Marshal(PlanningPeriod)
		json.Unmarshal(PlanningPeriodString, &o.PlanningPeriod)
	}
	
	if StartDayOfWeekend, ok := SchedulingsettingsresponseMap["startDayOfWeekend"].(string); ok {
		o.StartDayOfWeekend = &StartDayOfWeekend
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

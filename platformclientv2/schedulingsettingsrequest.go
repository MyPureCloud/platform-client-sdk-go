package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Schedulingsettingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingsettingsrequest
	
	return json.Marshal(&struct { 
		MaxOccupancyPercentForDeferredWork *int `json:"maxOccupancyPercentForDeferredWork,omitempty"`
		
		DefaultShrinkagePercent *float64 `json:"defaultShrinkagePercent,omitempty"`
		
		ShrinkageOverrides *Shrinkageoverrides `json:"shrinkageOverrides,omitempty"`
		
		PlanningPeriod *Valuewrapperplanningperiodsettings `json:"planningPeriod,omitempty"`
		
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

func (o *Schedulingsettingsrequest) UnmarshalJSON(b []byte) error {
	var SchedulingsettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingsettingsrequestMap)
	if err != nil {
		return err
	}
	
	if MaxOccupancyPercentForDeferredWork, ok := SchedulingsettingsrequestMap["maxOccupancyPercentForDeferredWork"].(float64); ok {
		MaxOccupancyPercentForDeferredWorkInt := int(MaxOccupancyPercentForDeferredWork)
		o.MaxOccupancyPercentForDeferredWork = &MaxOccupancyPercentForDeferredWorkInt
	}
	
	if DefaultShrinkagePercent, ok := SchedulingsettingsrequestMap["defaultShrinkagePercent"].(float64); ok {
		o.DefaultShrinkagePercent = &DefaultShrinkagePercent
	}
    
	if ShrinkageOverrides, ok := SchedulingsettingsrequestMap["shrinkageOverrides"].(map[string]interface{}); ok {
		ShrinkageOverridesString, _ := json.Marshal(ShrinkageOverrides)
		json.Unmarshal(ShrinkageOverridesString, &o.ShrinkageOverrides)
	}
	
	if PlanningPeriod, ok := SchedulingsettingsrequestMap["planningPeriod"].(map[string]interface{}); ok {
		PlanningPeriodString, _ := json.Marshal(PlanningPeriod)
		json.Unmarshal(PlanningPeriodString, &o.PlanningPeriod)
	}
	
	if StartDayOfWeekend, ok := SchedulingsettingsrequestMap["startDayOfWeekend"].(string); ok {
		o.StartDayOfWeekend = &StartDayOfWeekend
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

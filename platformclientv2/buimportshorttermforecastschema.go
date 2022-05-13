package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buimportshorttermforecastschema
type Buimportshorttermforecastschema struct { 
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`


	// WeekCount - The number of weeks covered by the forecast
	WeekCount *int `json:"weekCount,omitempty"`


	// PlanningGroups - The short term planning group data
	PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`


	// LongTermPlanningGroups - The long term planning group data
	LongTermPlanningGroups *[]Longtermforecastplanninggroupdata `json:"longTermPlanningGroups,omitempty"`


	// CanUseForScheduling - Whether this forecast can be used for scheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

}

func (o *Buimportshorttermforecastschema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buimportshorttermforecastschema
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		PlanningGroups *[]Forecastplanninggroupdata `json:"planningGroups,omitempty"`
		
		LongTermPlanningGroups *[]Longtermforecastplanninggroupdata `json:"longTermPlanningGroups,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		WeekCount: o.WeekCount,
		
		PlanningGroups: o.PlanningGroups,
		
		LongTermPlanningGroups: o.LongTermPlanningGroups,
		
		CanUseForScheduling: o.CanUseForScheduling,
		Alias:    (*Alias)(o),
	})
}

func (o *Buimportshorttermforecastschema) UnmarshalJSON(b []byte) error {
	var BuimportshorttermforecastschemaMap map[string]interface{}
	err := json.Unmarshal(b, &BuimportshorttermforecastschemaMap)
	if err != nil {
		return err
	}
	
	if Description, ok := BuimportshorttermforecastschemaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if WeekCount, ok := BuimportshorttermforecastschemaMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if PlanningGroups, ok := BuimportshorttermforecastschemaMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if LongTermPlanningGroups, ok := BuimportshorttermforecastschemaMap["longTermPlanningGroups"].([]interface{}); ok {
		LongTermPlanningGroupsString, _ := json.Marshal(LongTermPlanningGroups)
		json.Unmarshal(LongTermPlanningGroupsString, &o.LongTermPlanningGroups)
	}
	
	if CanUseForScheduling, ok := BuimportshorttermforecastschemaMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buimportshorttermforecastschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

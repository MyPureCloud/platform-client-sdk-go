package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplanshift
type Createworkplanshift struct { 
	// Name - Name of the shift
	Name *string `json:"name,omitempty"`


	// Days - Days of the week applicable for this shift
	Days *Setwrapperdayofweek `json:"days,omitempty"`


	// FlexibleStartTime - Whether the start time of the shift is flexible
	FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`


	// ExactStartTimeMinutesFromMidnight - Exact start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == false
	ExactStartTimeMinutesFromMidnight *int `json:"exactStartTimeMinutesFromMidnight,omitempty"`


	// EarliestStartTimeMinutesFromMidnight - Earliest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
	EarliestStartTimeMinutesFromMidnight *int `json:"earliestStartTimeMinutesFromMidnight,omitempty"`


	// LatestStartTimeMinutesFromMidnight - Latest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
	LatestStartTimeMinutesFromMidnight *int `json:"latestStartTimeMinutesFromMidnight,omitempty"`


	// ConstrainStopTime - Whether the latest stop time constraint for the shift is enabled.  Deprecated, use constrainLatestStopTime instead
	ConstrainStopTime *bool `json:"constrainStopTime,omitempty"`


	// ConstrainLatestStopTime - Whether the latest stop time constraint for the shift is enabled
	ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`


	// LatestStopTimeMinutesFromMidnight - Latest stop time of the shift defined as offset minutes from midnight. Used if constrainStopTime == true
	LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`


	// ConstrainEarliestStopTime - Whether the earliest stop time constraint for the shift is enabled
	ConstrainEarliestStopTime *bool `json:"constrainEarliestStopTime,omitempty"`


	// EarliestStopTimeMinutesFromMidnight - This is the earliest time a shift can end
	EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`


	// StartIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the shift. Used if flexibleStartTime == true
	StartIncrementMinutes *int `json:"startIncrementMinutes,omitempty"`


	// FlexiblePaidTime - Whether the paid time setting for the shift is flexible
	FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`


	// ExactPaidTimeMinutes - Exact paid time in minutes configured for the shift. Used if flexiblePaidTime == false
	ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`


	// MinimumPaidTimeMinutes - Minimum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`


	// MaximumPaidTimeMinutes - Maximum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`


	// ConstrainContiguousWorkTime - Whether the contiguous time constraint for the shift is enabled
	ConstrainContiguousWorkTime *bool `json:"constrainContiguousWorkTime,omitempty"`


	// MinimumContiguousWorkTimeMinutes - Minimum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
	MinimumContiguousWorkTimeMinutes *int `json:"minimumContiguousWorkTimeMinutes,omitempty"`


	// MaximumContiguousWorkTimeMinutes - Maximum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
	MaximumContiguousWorkTimeMinutes *int `json:"maximumContiguousWorkTimeMinutes,omitempty"`


	// ConstrainDayOff - Whether day off rule is enabled
	ConstrainDayOff *bool `json:"constrainDayOff,omitempty"`


	// DayOffRule - The day off rule for agents to have next day off or previous day off. used if constrainDayOff = true
	DayOffRule *string `json:"dayOffRule,omitempty"`


	// Activities - Activities configured for this shift
	Activities *[]Createworkplanactivity `json:"activities,omitempty"`

}

func (o *Createworkplanshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createworkplanshift
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Days *Setwrapperdayofweek `json:"days,omitempty"`
		
		FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`
		
		ExactStartTimeMinutesFromMidnight *int `json:"exactStartTimeMinutesFromMidnight,omitempty"`
		
		EarliestStartTimeMinutesFromMidnight *int `json:"earliestStartTimeMinutesFromMidnight,omitempty"`
		
		LatestStartTimeMinutesFromMidnight *int `json:"latestStartTimeMinutesFromMidnight,omitempty"`
		
		ConstrainStopTime *bool `json:"constrainStopTime,omitempty"`
		
		ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`
		
		LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`
		
		ConstrainEarliestStopTime *bool `json:"constrainEarliestStopTime,omitempty"`
		
		EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`
		
		StartIncrementMinutes *int `json:"startIncrementMinutes,omitempty"`
		
		FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`
		
		ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`
		
		MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`
		
		MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`
		
		ConstrainContiguousWorkTime *bool `json:"constrainContiguousWorkTime,omitempty"`
		
		MinimumContiguousWorkTimeMinutes *int `json:"minimumContiguousWorkTimeMinutes,omitempty"`
		
		MaximumContiguousWorkTimeMinutes *int `json:"maximumContiguousWorkTimeMinutes,omitempty"`
		
		ConstrainDayOff *bool `json:"constrainDayOff,omitempty"`
		
		DayOffRule *string `json:"dayOffRule,omitempty"`
		
		Activities *[]Createworkplanactivity `json:"activities,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Days: o.Days,
		
		FlexibleStartTime: o.FlexibleStartTime,
		
		ExactStartTimeMinutesFromMidnight: o.ExactStartTimeMinutesFromMidnight,
		
		EarliestStartTimeMinutesFromMidnight: o.EarliestStartTimeMinutesFromMidnight,
		
		LatestStartTimeMinutesFromMidnight: o.LatestStartTimeMinutesFromMidnight,
		
		ConstrainStopTime: o.ConstrainStopTime,
		
		ConstrainLatestStopTime: o.ConstrainLatestStopTime,
		
		LatestStopTimeMinutesFromMidnight: o.LatestStopTimeMinutesFromMidnight,
		
		ConstrainEarliestStopTime: o.ConstrainEarliestStopTime,
		
		EarliestStopTimeMinutesFromMidnight: o.EarliestStopTimeMinutesFromMidnight,
		
		StartIncrementMinutes: o.StartIncrementMinutes,
		
		FlexiblePaidTime: o.FlexiblePaidTime,
		
		ExactPaidTimeMinutes: o.ExactPaidTimeMinutes,
		
		MinimumPaidTimeMinutes: o.MinimumPaidTimeMinutes,
		
		MaximumPaidTimeMinutes: o.MaximumPaidTimeMinutes,
		
		ConstrainContiguousWorkTime: o.ConstrainContiguousWorkTime,
		
		MinimumContiguousWorkTimeMinutes: o.MinimumContiguousWorkTimeMinutes,
		
		MaximumContiguousWorkTimeMinutes: o.MaximumContiguousWorkTimeMinutes,
		
		ConstrainDayOff: o.ConstrainDayOff,
		
		DayOffRule: o.DayOffRule,
		
		Activities: o.Activities,
		Alias:    (*Alias)(o),
	})
}

func (o *Createworkplanshift) UnmarshalJSON(b []byte) error {
	var CreateworkplanshiftMap map[string]interface{}
	err := json.Unmarshal(b, &CreateworkplanshiftMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateworkplanshiftMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Days, ok := CreateworkplanshiftMap["days"].(map[string]interface{}); ok {
		DaysString, _ := json.Marshal(Days)
		json.Unmarshal(DaysString, &o.Days)
	}
	
	if FlexibleStartTime, ok := CreateworkplanshiftMap["flexibleStartTime"].(bool); ok {
		o.FlexibleStartTime = &FlexibleStartTime
	}
    
	if ExactStartTimeMinutesFromMidnight, ok := CreateworkplanshiftMap["exactStartTimeMinutesFromMidnight"].(float64); ok {
		ExactStartTimeMinutesFromMidnightInt := int(ExactStartTimeMinutesFromMidnight)
		o.ExactStartTimeMinutesFromMidnight = &ExactStartTimeMinutesFromMidnightInt
	}
	
	if EarliestStartTimeMinutesFromMidnight, ok := CreateworkplanshiftMap["earliestStartTimeMinutesFromMidnight"].(float64); ok {
		EarliestStartTimeMinutesFromMidnightInt := int(EarliestStartTimeMinutesFromMidnight)
		o.EarliestStartTimeMinutesFromMidnight = &EarliestStartTimeMinutesFromMidnightInt
	}
	
	if LatestStartTimeMinutesFromMidnight, ok := CreateworkplanshiftMap["latestStartTimeMinutesFromMidnight"].(float64); ok {
		LatestStartTimeMinutesFromMidnightInt := int(LatestStartTimeMinutesFromMidnight)
		o.LatestStartTimeMinutesFromMidnight = &LatestStartTimeMinutesFromMidnightInt
	}
	
	if ConstrainStopTime, ok := CreateworkplanshiftMap["constrainStopTime"].(bool); ok {
		o.ConstrainStopTime = &ConstrainStopTime
	}
    
	if ConstrainLatestStopTime, ok := CreateworkplanshiftMap["constrainLatestStopTime"].(bool); ok {
		o.ConstrainLatestStopTime = &ConstrainLatestStopTime
	}
    
	if LatestStopTimeMinutesFromMidnight, ok := CreateworkplanshiftMap["latestStopTimeMinutesFromMidnight"].(float64); ok {
		LatestStopTimeMinutesFromMidnightInt := int(LatestStopTimeMinutesFromMidnight)
		o.LatestStopTimeMinutesFromMidnight = &LatestStopTimeMinutesFromMidnightInt
	}
	
	if ConstrainEarliestStopTime, ok := CreateworkplanshiftMap["constrainEarliestStopTime"].(bool); ok {
		o.ConstrainEarliestStopTime = &ConstrainEarliestStopTime
	}
    
	if EarliestStopTimeMinutesFromMidnight, ok := CreateworkplanshiftMap["earliestStopTimeMinutesFromMidnight"].(float64); ok {
		EarliestStopTimeMinutesFromMidnightInt := int(EarliestStopTimeMinutesFromMidnight)
		o.EarliestStopTimeMinutesFromMidnight = &EarliestStopTimeMinutesFromMidnightInt
	}
	
	if StartIncrementMinutes, ok := CreateworkplanshiftMap["startIncrementMinutes"].(float64); ok {
		StartIncrementMinutesInt := int(StartIncrementMinutes)
		o.StartIncrementMinutes = &StartIncrementMinutesInt
	}
	
	if FlexiblePaidTime, ok := CreateworkplanshiftMap["flexiblePaidTime"].(bool); ok {
		o.FlexiblePaidTime = &FlexiblePaidTime
	}
    
	if ExactPaidTimeMinutes, ok := CreateworkplanshiftMap["exactPaidTimeMinutes"].(float64); ok {
		ExactPaidTimeMinutesInt := int(ExactPaidTimeMinutes)
		o.ExactPaidTimeMinutes = &ExactPaidTimeMinutesInt
	}
	
	if MinimumPaidTimeMinutes, ok := CreateworkplanshiftMap["minimumPaidTimeMinutes"].(float64); ok {
		MinimumPaidTimeMinutesInt := int(MinimumPaidTimeMinutes)
		o.MinimumPaidTimeMinutes = &MinimumPaidTimeMinutesInt
	}
	
	if MaximumPaidTimeMinutes, ok := CreateworkplanshiftMap["maximumPaidTimeMinutes"].(float64); ok {
		MaximumPaidTimeMinutesInt := int(MaximumPaidTimeMinutes)
		o.MaximumPaidTimeMinutes = &MaximumPaidTimeMinutesInt
	}
	
	if ConstrainContiguousWorkTime, ok := CreateworkplanshiftMap["constrainContiguousWorkTime"].(bool); ok {
		o.ConstrainContiguousWorkTime = &ConstrainContiguousWorkTime
	}
    
	if MinimumContiguousWorkTimeMinutes, ok := CreateworkplanshiftMap["minimumContiguousWorkTimeMinutes"].(float64); ok {
		MinimumContiguousWorkTimeMinutesInt := int(MinimumContiguousWorkTimeMinutes)
		o.MinimumContiguousWorkTimeMinutes = &MinimumContiguousWorkTimeMinutesInt
	}
	
	if MaximumContiguousWorkTimeMinutes, ok := CreateworkplanshiftMap["maximumContiguousWorkTimeMinutes"].(float64); ok {
		MaximumContiguousWorkTimeMinutesInt := int(MaximumContiguousWorkTimeMinutes)
		o.MaximumContiguousWorkTimeMinutes = &MaximumContiguousWorkTimeMinutesInt
	}
	
	if ConstrainDayOff, ok := CreateworkplanshiftMap["constrainDayOff"].(bool); ok {
		o.ConstrainDayOff = &ConstrainDayOff
	}
    
	if DayOffRule, ok := CreateworkplanshiftMap["dayOffRule"].(string); ok {
		o.DayOffRule = &DayOffRule
	}
    
	if Activities, ok := CreateworkplanshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createworkplanshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

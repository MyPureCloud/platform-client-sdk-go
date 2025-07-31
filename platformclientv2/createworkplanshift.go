package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplanshift
type Createworkplanshift struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// PlanningPeriodConstraints - Planning period constraints
	PlanningPeriodConstraints *Planningperiodshiftconstraints `json:"planningPeriodConstraints,omitempty"`

	// Activities - Activities configured for this shift
	Activities *[]Createworkplanactivity `json:"activities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createworkplanshift) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Createworkplanshift) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		PlanningPeriodConstraints *Planningperiodshiftconstraints `json:"planningPeriodConstraints,omitempty"`
		
		Activities *[]Createworkplanactivity `json:"activities,omitempty"`
		Alias
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
		
		PlanningPeriodConstraints: o.PlanningPeriodConstraints,
		
		Activities: o.Activities,
		Alias:    (Alias)(o),
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
    
	if PlanningPeriodConstraints, ok := CreateworkplanshiftMap["planningPeriodConstraints"].(map[string]interface{}); ok {
		PlanningPeriodConstraintsString, _ := json.Marshal(PlanningPeriodConstraints)
		json.Unmarshal(PlanningPeriodConstraintsString, &o.PlanningPeriodConstraints)
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

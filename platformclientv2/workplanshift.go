package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanshift
type Workplanshift struct { 
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

	// Activities - Activities configured for this shift
	Activities *[]Workplanactivity `json:"activities,omitempty"`

	// Id - ID of the shift. This is required only for the case of updating an existing shift
	Id *string `json:"id,omitempty"`

	// Delete - If marked true for updating an existing shift, the shift will be permanently deleted
	Delete *bool `json:"delete,omitempty"`

	// ValidationId - ID of shift in the context of work plan validation
	ValidationId *string `json:"validationId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanshift) SetField(field string, fieldValue interface{}) {
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

func (o Workplanshift) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanshift
	
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
		
		Activities *[]Workplanactivity `json:"activities,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		ValidationId *string `json:"validationId,omitempty"`
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
		
		Activities: o.Activities,
		
		Id: o.Id,
		
		Delete: o.Delete,
		
		ValidationId: o.ValidationId,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanshift) UnmarshalJSON(b []byte) error {
	var WorkplanshiftMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanshiftMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkplanshiftMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Days, ok := WorkplanshiftMap["days"].(map[string]interface{}); ok {
		DaysString, _ := json.Marshal(Days)
		json.Unmarshal(DaysString, &o.Days)
	}
	
	if FlexibleStartTime, ok := WorkplanshiftMap["flexibleStartTime"].(bool); ok {
		o.FlexibleStartTime = &FlexibleStartTime
	}
    
	if ExactStartTimeMinutesFromMidnight, ok := WorkplanshiftMap["exactStartTimeMinutesFromMidnight"].(float64); ok {
		ExactStartTimeMinutesFromMidnightInt := int(ExactStartTimeMinutesFromMidnight)
		o.ExactStartTimeMinutesFromMidnight = &ExactStartTimeMinutesFromMidnightInt
	}
	
	if EarliestStartTimeMinutesFromMidnight, ok := WorkplanshiftMap["earliestStartTimeMinutesFromMidnight"].(float64); ok {
		EarliestStartTimeMinutesFromMidnightInt := int(EarliestStartTimeMinutesFromMidnight)
		o.EarliestStartTimeMinutesFromMidnight = &EarliestStartTimeMinutesFromMidnightInt
	}
	
	if LatestStartTimeMinutesFromMidnight, ok := WorkplanshiftMap["latestStartTimeMinutesFromMidnight"].(float64); ok {
		LatestStartTimeMinutesFromMidnightInt := int(LatestStartTimeMinutesFromMidnight)
		o.LatestStartTimeMinutesFromMidnight = &LatestStartTimeMinutesFromMidnightInt
	}
	
	if ConstrainStopTime, ok := WorkplanshiftMap["constrainStopTime"].(bool); ok {
		o.ConstrainStopTime = &ConstrainStopTime
	}
    
	if ConstrainLatestStopTime, ok := WorkplanshiftMap["constrainLatestStopTime"].(bool); ok {
		o.ConstrainLatestStopTime = &ConstrainLatestStopTime
	}
    
	if LatestStopTimeMinutesFromMidnight, ok := WorkplanshiftMap["latestStopTimeMinutesFromMidnight"].(float64); ok {
		LatestStopTimeMinutesFromMidnightInt := int(LatestStopTimeMinutesFromMidnight)
		o.LatestStopTimeMinutesFromMidnight = &LatestStopTimeMinutesFromMidnightInt
	}
	
	if ConstrainEarliestStopTime, ok := WorkplanshiftMap["constrainEarliestStopTime"].(bool); ok {
		o.ConstrainEarliestStopTime = &ConstrainEarliestStopTime
	}
    
	if EarliestStopTimeMinutesFromMidnight, ok := WorkplanshiftMap["earliestStopTimeMinutesFromMidnight"].(float64); ok {
		EarliestStopTimeMinutesFromMidnightInt := int(EarliestStopTimeMinutesFromMidnight)
		o.EarliestStopTimeMinutesFromMidnight = &EarliestStopTimeMinutesFromMidnightInt
	}
	
	if StartIncrementMinutes, ok := WorkplanshiftMap["startIncrementMinutes"].(float64); ok {
		StartIncrementMinutesInt := int(StartIncrementMinutes)
		o.StartIncrementMinutes = &StartIncrementMinutesInt
	}
	
	if FlexiblePaidTime, ok := WorkplanshiftMap["flexiblePaidTime"].(bool); ok {
		o.FlexiblePaidTime = &FlexiblePaidTime
	}
    
	if ExactPaidTimeMinutes, ok := WorkplanshiftMap["exactPaidTimeMinutes"].(float64); ok {
		ExactPaidTimeMinutesInt := int(ExactPaidTimeMinutes)
		o.ExactPaidTimeMinutes = &ExactPaidTimeMinutesInt
	}
	
	if MinimumPaidTimeMinutes, ok := WorkplanshiftMap["minimumPaidTimeMinutes"].(float64); ok {
		MinimumPaidTimeMinutesInt := int(MinimumPaidTimeMinutes)
		o.MinimumPaidTimeMinutes = &MinimumPaidTimeMinutesInt
	}
	
	if MaximumPaidTimeMinutes, ok := WorkplanshiftMap["maximumPaidTimeMinutes"].(float64); ok {
		MaximumPaidTimeMinutesInt := int(MaximumPaidTimeMinutes)
		o.MaximumPaidTimeMinutes = &MaximumPaidTimeMinutesInt
	}
	
	if ConstrainContiguousWorkTime, ok := WorkplanshiftMap["constrainContiguousWorkTime"].(bool); ok {
		o.ConstrainContiguousWorkTime = &ConstrainContiguousWorkTime
	}
    
	if MinimumContiguousWorkTimeMinutes, ok := WorkplanshiftMap["minimumContiguousWorkTimeMinutes"].(float64); ok {
		MinimumContiguousWorkTimeMinutesInt := int(MinimumContiguousWorkTimeMinutes)
		o.MinimumContiguousWorkTimeMinutes = &MinimumContiguousWorkTimeMinutesInt
	}
	
	if MaximumContiguousWorkTimeMinutes, ok := WorkplanshiftMap["maximumContiguousWorkTimeMinutes"].(float64); ok {
		MaximumContiguousWorkTimeMinutesInt := int(MaximumContiguousWorkTimeMinutes)
		o.MaximumContiguousWorkTimeMinutes = &MaximumContiguousWorkTimeMinutesInt
	}
	
	if ConstrainDayOff, ok := WorkplanshiftMap["constrainDayOff"].(bool); ok {
		o.ConstrainDayOff = &ConstrainDayOff
	}
    
	if DayOffRule, ok := WorkplanshiftMap["dayOffRule"].(string); ok {
		o.DayOffRule = &DayOffRule
	}
    
	if Activities, ok := WorkplanshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	
	if Id, ok := WorkplanshiftMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Delete, ok := WorkplanshiftMap["delete"].(bool); ok {
		o.Delete = &Delete
	}
    
	if ValidationId, ok := WorkplanshiftMap["validationId"].(string); ok {
		o.ValidationId = &ValidationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

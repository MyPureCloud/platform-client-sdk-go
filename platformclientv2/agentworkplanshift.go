package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentworkplanshift
type Agentworkplanshift struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// EarliestStopTimeMinutesFromMidnight - This is the earliest time a shift can end
	EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`

	// ConstrainLatestStopTime - Whether the latest stop time constraint for the shift is enabled
	ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`

	// LatestStopTimeMinutesFromMidnight - Latest stop time of the shift defined as offset minutes from midnight. Used if constrainStopTime == true
	LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`

	// FlexiblePaidTime - Whether the paid time setting for the shift is flexible
	FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`

	// ExactPaidTimeMinutes - Exact paid time in minutes configured for the shift. Used if flexiblePaidTime == false
	ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`

	// MinimumPaidTimeMinutes - Minimum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`

	// MaximumPaidTimeMinutes - Maximum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
	MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`

	// Activities - Activities configured for this shift
	Activities *[]Agentworkplanactivity `json:"activities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentworkplanshift) SetField(field string, fieldValue interface{}) {
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

func (o Agentworkplanshift) MarshalJSON() ([]byte, error) {
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
	type Alias Agentworkplanshift
	
	return json.Marshal(&struct { 
		Days *Setwrapperdayofweek `json:"days,omitempty"`
		
		FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`
		
		ExactStartTimeMinutesFromMidnight *int `json:"exactStartTimeMinutesFromMidnight,omitempty"`
		
		EarliestStartTimeMinutesFromMidnight *int `json:"earliestStartTimeMinutesFromMidnight,omitempty"`
		
		LatestStartTimeMinutesFromMidnight *int `json:"latestStartTimeMinutesFromMidnight,omitempty"`
		
		EarliestStopTimeMinutesFromMidnight *int `json:"earliestStopTimeMinutesFromMidnight,omitempty"`
		
		ConstrainLatestStopTime *bool `json:"constrainLatestStopTime,omitempty"`
		
		LatestStopTimeMinutesFromMidnight *int `json:"latestStopTimeMinutesFromMidnight,omitempty"`
		
		FlexiblePaidTime *bool `json:"flexiblePaidTime,omitempty"`
		
		ExactPaidTimeMinutes *int `json:"exactPaidTimeMinutes,omitempty"`
		
		MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`
		
		MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`
		
		Activities *[]Agentworkplanactivity `json:"activities,omitempty"`
		Alias
	}{ 
		Days: o.Days,
		
		FlexibleStartTime: o.FlexibleStartTime,
		
		ExactStartTimeMinutesFromMidnight: o.ExactStartTimeMinutesFromMidnight,
		
		EarliestStartTimeMinutesFromMidnight: o.EarliestStartTimeMinutesFromMidnight,
		
		LatestStartTimeMinutesFromMidnight: o.LatestStartTimeMinutesFromMidnight,
		
		EarliestStopTimeMinutesFromMidnight: o.EarliestStopTimeMinutesFromMidnight,
		
		ConstrainLatestStopTime: o.ConstrainLatestStopTime,
		
		LatestStopTimeMinutesFromMidnight: o.LatestStopTimeMinutesFromMidnight,
		
		FlexiblePaidTime: o.FlexiblePaidTime,
		
		ExactPaidTimeMinutes: o.ExactPaidTimeMinutes,
		
		MinimumPaidTimeMinutes: o.MinimumPaidTimeMinutes,
		
		MaximumPaidTimeMinutes: o.MaximumPaidTimeMinutes,
		
		Activities: o.Activities,
		Alias:    (Alias)(o),
	})
}

func (o *Agentworkplanshift) UnmarshalJSON(b []byte) error {
	var AgentworkplanshiftMap map[string]interface{}
	err := json.Unmarshal(b, &AgentworkplanshiftMap)
	if err != nil {
		return err
	}
	
	if Days, ok := AgentworkplanshiftMap["days"].(map[string]interface{}); ok {
		DaysString, _ := json.Marshal(Days)
		json.Unmarshal(DaysString, &o.Days)
	}
	
	if FlexibleStartTime, ok := AgentworkplanshiftMap["flexibleStartTime"].(bool); ok {
		o.FlexibleStartTime = &FlexibleStartTime
	}
    
	if ExactStartTimeMinutesFromMidnight, ok := AgentworkplanshiftMap["exactStartTimeMinutesFromMidnight"].(float64); ok {
		ExactStartTimeMinutesFromMidnightInt := int(ExactStartTimeMinutesFromMidnight)
		o.ExactStartTimeMinutesFromMidnight = &ExactStartTimeMinutesFromMidnightInt
	}
	
	if EarliestStartTimeMinutesFromMidnight, ok := AgentworkplanshiftMap["earliestStartTimeMinutesFromMidnight"].(float64); ok {
		EarliestStartTimeMinutesFromMidnightInt := int(EarliestStartTimeMinutesFromMidnight)
		o.EarliestStartTimeMinutesFromMidnight = &EarliestStartTimeMinutesFromMidnightInt
	}
	
	if LatestStartTimeMinutesFromMidnight, ok := AgentworkplanshiftMap["latestStartTimeMinutesFromMidnight"].(float64); ok {
		LatestStartTimeMinutesFromMidnightInt := int(LatestStartTimeMinutesFromMidnight)
		o.LatestStartTimeMinutesFromMidnight = &LatestStartTimeMinutesFromMidnightInt
	}
	
	if EarliestStopTimeMinutesFromMidnight, ok := AgentworkplanshiftMap["earliestStopTimeMinutesFromMidnight"].(float64); ok {
		EarliestStopTimeMinutesFromMidnightInt := int(EarliestStopTimeMinutesFromMidnight)
		o.EarliestStopTimeMinutesFromMidnight = &EarliestStopTimeMinutesFromMidnightInt
	}
	
	if ConstrainLatestStopTime, ok := AgentworkplanshiftMap["constrainLatestStopTime"].(bool); ok {
		o.ConstrainLatestStopTime = &ConstrainLatestStopTime
	}
    
	if LatestStopTimeMinutesFromMidnight, ok := AgentworkplanshiftMap["latestStopTimeMinutesFromMidnight"].(float64); ok {
		LatestStopTimeMinutesFromMidnightInt := int(LatestStopTimeMinutesFromMidnight)
		o.LatestStopTimeMinutesFromMidnight = &LatestStopTimeMinutesFromMidnightInt
	}
	
	if FlexiblePaidTime, ok := AgentworkplanshiftMap["flexiblePaidTime"].(bool); ok {
		o.FlexiblePaidTime = &FlexiblePaidTime
	}
    
	if ExactPaidTimeMinutes, ok := AgentworkplanshiftMap["exactPaidTimeMinutes"].(float64); ok {
		ExactPaidTimeMinutesInt := int(ExactPaidTimeMinutes)
		o.ExactPaidTimeMinutes = &ExactPaidTimeMinutesInt
	}
	
	if MinimumPaidTimeMinutes, ok := AgentworkplanshiftMap["minimumPaidTimeMinutes"].(float64); ok {
		MinimumPaidTimeMinutesInt := int(MinimumPaidTimeMinutes)
		o.MinimumPaidTimeMinutes = &MinimumPaidTimeMinutesInt
	}
	
	if MaximumPaidTimeMinutes, ok := AgentworkplanshiftMap["maximumPaidTimeMinutes"].(float64); ok {
		MaximumPaidTimeMinutesInt := int(MaximumPaidTimeMinutes)
		o.MaximumPaidTimeMinutes = &MaximumPaidTimeMinutesInt
	}
	
	if Activities, ok := AgentworkplanshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentworkplanshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

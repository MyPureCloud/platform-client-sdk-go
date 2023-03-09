package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dailypossibleshift
type Dailypossibleshift struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DayOfWeek - Day of the shift
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EarliestShiftStartMinutesFromMidnight - Minutes of the earliest shift start from midnight. Note that midnight is 12:00 am in the time zone specified in the timeZone field (in the top level of the response)
	EarliestShiftStartMinutesFromMidnight *int `json:"earliestShiftStartMinutesFromMidnight,omitempty"`

	// Required - Whether this is a required shift
	Required *bool `json:"required,omitempty"`

	// MinimumPaidTimeMinutes - Minimum paid time in minutes of this daily shift
	MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`

	// MaximumPaidTimeMinutes - Maximum paid time in minutes of this daily shift
	MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`

	// IntervalScheduleProbabilities - The percentage of being scheduled in each interval between the earliest shift start and latest shift end. Range of the values: [0, 100].
	IntervalScheduleProbabilities *[]int `json:"intervalScheduleProbabilities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dailypossibleshift) SetField(field string, fieldValue interface{}) {
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

func (o Dailypossibleshift) MarshalJSON() ([]byte, error) {
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
	type Alias Dailypossibleshift
	
	return json.Marshal(&struct { 
		DayOfWeek *string `json:"dayOfWeek,omitempty"`
		
		EarliestShiftStartMinutesFromMidnight *int `json:"earliestShiftStartMinutesFromMidnight,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		
		MinimumPaidTimeMinutes *int `json:"minimumPaidTimeMinutes,omitempty"`
		
		MaximumPaidTimeMinutes *int `json:"maximumPaidTimeMinutes,omitempty"`
		
		IntervalScheduleProbabilities *[]int `json:"intervalScheduleProbabilities,omitempty"`
		Alias
	}{ 
		DayOfWeek: o.DayOfWeek,
		
		EarliestShiftStartMinutesFromMidnight: o.EarliestShiftStartMinutesFromMidnight,
		
		Required: o.Required,
		
		MinimumPaidTimeMinutes: o.MinimumPaidTimeMinutes,
		
		MaximumPaidTimeMinutes: o.MaximumPaidTimeMinutes,
		
		IntervalScheduleProbabilities: o.IntervalScheduleProbabilities,
		Alias:    (Alias)(o),
	})
}

func (o *Dailypossibleshift) UnmarshalJSON(b []byte) error {
	var DailypossibleshiftMap map[string]interface{}
	err := json.Unmarshal(b, &DailypossibleshiftMap)
	if err != nil {
		return err
	}
	
	if DayOfWeek, ok := DailypossibleshiftMap["dayOfWeek"].(string); ok {
		o.DayOfWeek = &DayOfWeek
	}
    
	if EarliestShiftStartMinutesFromMidnight, ok := DailypossibleshiftMap["earliestShiftStartMinutesFromMidnight"].(float64); ok {
		EarliestShiftStartMinutesFromMidnightInt := int(EarliestShiftStartMinutesFromMidnight)
		o.EarliestShiftStartMinutesFromMidnight = &EarliestShiftStartMinutesFromMidnightInt
	}
	
	if Required, ok := DailypossibleshiftMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if MinimumPaidTimeMinutes, ok := DailypossibleshiftMap["minimumPaidTimeMinutes"].(float64); ok {
		MinimumPaidTimeMinutesInt := int(MinimumPaidTimeMinutes)
		o.MinimumPaidTimeMinutes = &MinimumPaidTimeMinutesInt
	}
	
	if MaximumPaidTimeMinutes, ok := DailypossibleshiftMap["maximumPaidTimeMinutes"].(float64); ok {
		MaximumPaidTimeMinutesInt := int(MaximumPaidTimeMinutes)
		o.MaximumPaidTimeMinutes = &MaximumPaidTimeMinutesInt
	}
	
	if IntervalScheduleProbabilities, ok := DailypossibleshiftMap["intervalScheduleProbabilities"].([]interface{}); ok {
		IntervalScheduleProbabilitiesString, _ := json.Marshal(IntervalScheduleProbabilities)
		json.Unmarshal(IntervalScheduleProbabilitiesString, &o.IntervalScheduleProbabilities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dailypossibleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

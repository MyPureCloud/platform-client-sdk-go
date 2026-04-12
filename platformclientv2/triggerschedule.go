package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Triggerschedule - Schedule configuration for a scheduled trigger
type Triggerschedule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Minutes - Minutes on which the trigger should fire. Coma separated list of up to 2 values 0-59
	Minutes *string `json:"minutes,omitempty"`

	// Hours - Hours on which the trigger should fire. 0-23 or '*' for every hour.
	Hours *string `json:"hours,omitempty"`

	// DaysOfMonth - Days of month on which the trigger should fire. 1-31 or '*' for every or '?' for any
	DaysOfMonth *string `json:"daysOfMonth,omitempty"`

	// Months - Months on which the trigger should fire. 1-12 or '*' for every
	Months *string `json:"months,omitempty"`

	// DaysOfWeek - Days of week on which the trigger should fire. 1-7 or '*' for every or '?' for any
	DaysOfWeek *string `json:"daysOfWeek,omitempty"`

	// Timezone - Timezone for the trigger schedule
	Timezone *string `json:"timezone,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Triggerschedule) SetField(field string, fieldValue interface{}) {
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

func (o Triggerschedule) MarshalJSON() ([]byte, error) {
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
	type Alias Triggerschedule
	
	return json.Marshal(&struct { 
		Minutes *string `json:"minutes,omitempty"`
		
		Hours *string `json:"hours,omitempty"`
		
		DaysOfMonth *string `json:"daysOfMonth,omitempty"`
		
		Months *string `json:"months,omitempty"`
		
		DaysOfWeek *string `json:"daysOfWeek,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		Alias
	}{ 
		Minutes: o.Minutes,
		
		Hours: o.Hours,
		
		DaysOfMonth: o.DaysOfMonth,
		
		Months: o.Months,
		
		DaysOfWeek: o.DaysOfWeek,
		
		Timezone: o.Timezone,
		Alias:    (Alias)(o),
	})
}

func (o *Triggerschedule) UnmarshalJSON(b []byte) error {
	var TriggerscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &TriggerscheduleMap)
	if err != nil {
		return err
	}
	
	if Minutes, ok := TriggerscheduleMap["minutes"].(string); ok {
		o.Minutes = &Minutes
	}
    
	if Hours, ok := TriggerscheduleMap["hours"].(string); ok {
		o.Hours = &Hours
	}
    
	if DaysOfMonth, ok := TriggerscheduleMap["daysOfMonth"].(string); ok {
		o.DaysOfMonth = &DaysOfMonth
	}
    
	if Months, ok := TriggerscheduleMap["months"].(string); ok {
		o.Months = &Months
	}
    
	if DaysOfWeek, ok := TriggerscheduleMap["daysOfWeek"].(string); ok {
		o.DaysOfWeek = &DaysOfWeek
	}
    
	if Timezone, ok := TriggerscheduleMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Triggerschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

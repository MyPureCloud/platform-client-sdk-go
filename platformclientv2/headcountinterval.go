package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Headcountinterval
type Headcountinterval struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - The start date-time for this headcount interval in ISO-8601 format, must be within the 8 day schedule
	Interval *time.Time `json:"interval,omitempty"`

	// Value - Headcount value for this interval
	Value *float64 `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Headcountinterval) SetField(field string, fieldValue interface{}) {
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

func (o Headcountinterval) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Interval", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Headcountinterval
	
	Interval := new(string)
	if o.Interval != nil {
		
		*Interval = timeutil.Strftime(o.Interval, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Interval = nil
	}
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		Alias
	}{ 
		Interval: Interval,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Headcountinterval) UnmarshalJSON(b []byte) error {
	var HeadcountintervalMap map[string]interface{}
	err := json.Unmarshal(b, &HeadcountintervalMap)
	if err != nil {
		return err
	}
	
	if intervalString, ok := HeadcountintervalMap["interval"].(string); ok {
		Interval, _ := time.Parse("2006-01-02T15:04:05.999999Z", intervalString)
		o.Interval = &Interval
	}
	
	if Value, ok := HeadcountintervalMap["value"].(float64); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Headcountinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

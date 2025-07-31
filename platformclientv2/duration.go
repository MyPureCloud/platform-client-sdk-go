package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Duration
type Duration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Seconds
	Seconds *int `json:"seconds,omitempty"`

	// Zero
	Zero *bool `json:"zero,omitempty"`

	// Nano
	Nano *int `json:"nano,omitempty"`

	// Negative
	Negative *bool `json:"negative,omitempty"`

	// Units
	Units *[]Temporalunit `json:"units,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Duration) SetField(field string, fieldValue interface{}) {
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

func (o Duration) MarshalJSON() ([]byte, error) {
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
	type Alias Duration
	
	return json.Marshal(&struct { 
		Seconds *int `json:"seconds,omitempty"`
		
		Zero *bool `json:"zero,omitempty"`
		
		Nano *int `json:"nano,omitempty"`
		
		Negative *bool `json:"negative,omitempty"`
		
		Units *[]Temporalunit `json:"units,omitempty"`
		Alias
	}{ 
		Seconds: o.Seconds,
		
		Zero: o.Zero,
		
		Nano: o.Nano,
		
		Negative: o.Negative,
		
		Units: o.Units,
		Alias:    (Alias)(o),
	})
}

func (o *Duration) UnmarshalJSON(b []byte) error {
	var DurationMap map[string]interface{}
	err := json.Unmarshal(b, &DurationMap)
	if err != nil {
		return err
	}
	
	if Seconds, ok := DurationMap["seconds"].(float64); ok {
		SecondsInt := int(Seconds)
		o.Seconds = &SecondsInt
	}
	
	if Zero, ok := DurationMap["zero"].(bool); ok {
		o.Zero = &Zero
	}
    
	if Nano, ok := DurationMap["nano"].(float64); ok {
		NanoInt := int(Nano)
		o.Nano = &NanoInt
	}
	
	if Negative, ok := DurationMap["negative"].(bool); ok {
		o.Negative = &Negative
	}
    
	if Units, ok := DurationMap["units"].([]interface{}); ok {
		UnitsString, _ := json.Marshal(Units)
		json.Unmarshal(UnitsString, &o.Units)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Duration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

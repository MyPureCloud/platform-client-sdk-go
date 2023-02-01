package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictorworkloadbalancing
type Predictorworkloadbalancing struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Flag to activate and deactivate workload balancing.
	Enabled *bool `json:"enabled,omitempty"`

	// MinimumOccupancy - Desired minimum occupancy threshold of agents. Must be between 0 and 100.
	MinimumOccupancy *int `json:"minimumOccupancy,omitempty"`

	// MaximumOccupancy - Desired maximum occupancy threshold of agents. Must be between 0 and 100.
	MaximumOccupancy *int `json:"maximumOccupancy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictorworkloadbalancing) SetField(field string, fieldValue interface{}) {
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

func (o Predictorworkloadbalancing) MarshalJSON() ([]byte, error) {
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
	type Alias Predictorworkloadbalancing
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		MinimumOccupancy *int `json:"minimumOccupancy,omitempty"`
		
		MaximumOccupancy *int `json:"maximumOccupancy,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		MinimumOccupancy: o.MinimumOccupancy,
		
		MaximumOccupancy: o.MaximumOccupancy,
		Alias:    (Alias)(o),
	})
}

func (o *Predictorworkloadbalancing) UnmarshalJSON(b []byte) error {
	var PredictorworkloadbalancingMap map[string]interface{}
	err := json.Unmarshal(b, &PredictorworkloadbalancingMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := PredictorworkloadbalancingMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MinimumOccupancy, ok := PredictorworkloadbalancingMap["minimumOccupancy"].(float64); ok {
		MinimumOccupancyInt := int(MinimumOccupancy)
		o.MinimumOccupancy = &MinimumOccupancyInt
	}
	
	if MaximumOccupancy, ok := PredictorworkloadbalancingMap["maximumOccupancy"].(float64); ok {
		MaximumOccupancyInt := int(MaximumOccupancy)
		o.MaximumOccupancy = &MaximumOccupancyInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictorworkloadbalancing) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

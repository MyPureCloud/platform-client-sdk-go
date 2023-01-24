package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Useraggregationview
type Useraggregationview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Target - Target metric name
	Target *string `json:"target,omitempty"`

	// Name - A unique name for this view. Must be distinct from other views and built-in metric names.
	Name *string `json:"name,omitempty"`

	// Function - Type of view you wish to create
	Function *string `json:"function,omitempty"`

	// VarRange - Range of numbers for slicing up data
	VarRange *Aggregationrange `json:"range,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Useraggregationview) SetField(field string, fieldValue interface{}) {
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

func (o Useraggregationview) MarshalJSON() ([]byte, error) {
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
	type Alias Useraggregationview
	
	return json.Marshal(&struct { 
		Target *string `json:"target,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Function *string `json:"function,omitempty"`
		
		VarRange *Aggregationrange `json:"range,omitempty"`
		Alias
	}{ 
		Target: o.Target,
		
		Name: o.Name,
		
		Function: o.Function,
		
		VarRange: o.VarRange,
		Alias:    (Alias)(o),
	})
}

func (o *Useraggregationview) UnmarshalJSON(b []byte) error {
	var UseraggregationviewMap map[string]interface{}
	err := json.Unmarshal(b, &UseraggregationviewMap)
	if err != nil {
		return err
	}
	
	if Target, ok := UseraggregationviewMap["target"].(string); ok {
		o.Target = &Target
	}
    
	if Name, ok := UseraggregationviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Function, ok := UseraggregationviewMap["function"].(string); ok {
		o.Function = &Function
	}
    
	if VarRange, ok := UseraggregationviewMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Useraggregationview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentationv2searchaggregation
type Documentationv2searchaggregation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Field - The field used for aggregation
	Field *string `json:"field,omitempty"`

	// Name - The name of the aggregation. The response aggregation uses this name.
	Name *string `json:"name,omitempty"`

	// VarType - The type of aggregation to perform
	VarType *string `json:"type,omitempty"`

	// Value - A value to use for aggregation
	Value *string `json:"value,omitempty"`

	// Size - The number aggregations results to return out of the entire result set
	Size *int `json:"size,omitempty"`

	// Order - The order in which aggregation results are sorted
	Order *[]string `json:"order,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentationv2searchaggregation) SetField(field string, fieldValue interface{}) {
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

func (o Documentationv2searchaggregation) MarshalJSON() ([]byte, error) {
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
	type Alias Documentationv2searchaggregation
	
	return json.Marshal(&struct { 
		Field *string `json:"field,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		Order *[]string `json:"order,omitempty"`
		Alias
	}{ 
		Field: o.Field,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Value: o.Value,
		
		Size: o.Size,
		
		Order: o.Order,
		Alias:    (Alias)(o),
	})
}

func (o *Documentationv2searchaggregation) UnmarshalJSON(b []byte) error {
	var Documentationv2searchaggregationMap map[string]interface{}
	err := json.Unmarshal(b, &Documentationv2searchaggregationMap)
	if err != nil {
		return err
	}
	
	if Field, ok := Documentationv2searchaggregationMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if Name, ok := Documentationv2searchaggregationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := Documentationv2searchaggregationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := Documentationv2searchaggregationMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Size, ok := Documentationv2searchaggregationMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if Order, ok := Documentationv2searchaggregationMap["order"].([]interface{}); ok {
		OrderString, _ := json.Marshal(Order)
		json.Unmarshal(OrderString, &o.Order)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentationv2searchaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Segmentdetailquerypredicate
type Segmentdetailquerypredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`

	// Dimension - Left hand side for dimension predicates
	Dimension *string `json:"dimension,omitempty"`

	// PropertyType - Left hand side for property predicates
	PropertyType *string `json:"propertyType,omitempty"`

	// Property - Left hand side for property predicates
	Property *string `json:"property,omitempty"`

	// Metric - Left hand side for metric predicates
	Metric *string `json:"metric,omitempty"`

	// Operator - Optional operator, default is matches
	Operator *string `json:"operator,omitempty"`

	// Value - Right hand side for dimension, metric, or property predicates
	Value *string `json:"value,omitempty"`

	// VarRange - Right hand side for dimension, metric, or property predicates
	VarRange *Numericrange `json:"range,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Segmentdetailquerypredicate) SetField(field string, fieldValue interface{}) {
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

func (o Segmentdetailquerypredicate) MarshalJSON() ([]byte, error) {
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
	type Alias Segmentdetailquerypredicate
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Numericrange `json:"range,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		PropertyType: o.PropertyType,
		
		Property: o.Property,
		
		Metric: o.Metric,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		Alias:    (Alias)(o),
	})
}

func (o *Segmentdetailquerypredicate) UnmarshalJSON(b []byte) error {
	var SegmentdetailquerypredicateMap map[string]interface{}
	err := json.Unmarshal(b, &SegmentdetailquerypredicateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SegmentdetailquerypredicateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := SegmentdetailquerypredicateMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if PropertyType, ok := SegmentdetailquerypredicateMap["propertyType"].(string); ok {
		o.PropertyType = &PropertyType
	}
    
	if Property, ok := SegmentdetailquerypredicateMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if Metric, ok := SegmentdetailquerypredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Operator, ok := SegmentdetailquerypredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := SegmentdetailquerypredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := SegmentdetailquerypredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Segmentdetailquerypredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

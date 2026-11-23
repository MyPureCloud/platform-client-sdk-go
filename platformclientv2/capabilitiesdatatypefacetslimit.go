package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesdatatypefacetslimit
type Capabilitiesdatatypefacetslimit struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StringValue
	StringValue *string `json:"stringValue,omitempty"`

	// BooleanValue
	BooleanValue *bool `json:"booleanValue,omitempty"`

	// IntegerValue
	IntegerValue *int `json:"integerValue,omitempty"`

	// DecimalValue
	DecimalValue *float64 `json:"decimalValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesdatatypefacetslimit) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesdatatypefacetslimit) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesdatatypefacetslimit
	
	return json.Marshal(&struct { 
		StringValue *string `json:"stringValue,omitempty"`
		
		BooleanValue *bool `json:"booleanValue,omitempty"`
		
		IntegerValue *int `json:"integerValue,omitempty"`
		
		DecimalValue *float64 `json:"decimalValue,omitempty"`
		Alias
	}{ 
		StringValue: o.StringValue,
		
		BooleanValue: o.BooleanValue,
		
		IntegerValue: o.IntegerValue,
		
		DecimalValue: o.DecimalValue,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesdatatypefacetslimit) UnmarshalJSON(b []byte) error {
	var CapabilitiesdatatypefacetslimitMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesdatatypefacetslimitMap)
	if err != nil {
		return err
	}
	
	if StringValue, ok := CapabilitiesdatatypefacetslimitMap["stringValue"].(string); ok {
		o.StringValue = &StringValue
	}
    
	if BooleanValue, ok := CapabilitiesdatatypefacetslimitMap["booleanValue"].(bool); ok {
		o.BooleanValue = &BooleanValue
	}
    
	if IntegerValue, ok := CapabilitiesdatatypefacetslimitMap["integerValue"].(float64); ok {
		IntegerValueInt := int(IntegerValue)
		o.IntegerValue = &IntegerValueInt
	}
	
	if DecimalValue, ok := CapabilitiesdatatypefacetslimitMap["decimalValue"].(float64); ok {
		o.DecimalValue = &DecimalValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypefacetslimit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

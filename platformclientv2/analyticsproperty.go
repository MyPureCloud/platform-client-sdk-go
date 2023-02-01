package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsproperty
type Analyticsproperty struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Property - User-defined rather than intrinsic system-observed values. These are tagged onto segments by other components within PureCloud or by API users directly.  This is the name of the user-defined property.
	Property *string `json:"property,omitempty"`

	// PropertyType - Indicates what the data type is (e.g. integer vs string) and therefore how to evaluate what would constitute a match
	PropertyType *string `json:"propertyType,omitempty"`

	// Value - What property value to match against
	Value *string `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsproperty) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsproperty) MarshalJSON() ([]byte, error) {
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
	type Alias Analyticsproperty
	
	return json.Marshal(&struct { 
		Property *string `json:"property,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		Value *string `json:"value,omitempty"`
		Alias
	}{ 
		Property: o.Property,
		
		PropertyType: o.PropertyType,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsproperty) UnmarshalJSON(b []byte) error {
	var AnalyticspropertyMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticspropertyMap)
	if err != nil {
		return err
	}
	
	if Property, ok := AnalyticspropertyMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if PropertyType, ok := AnalyticspropertyMap["propertyType"].(string); ok {
		o.PropertyType = &PropertyType
	}
    
	if Value, ok := AnalyticspropertyMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsproperty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

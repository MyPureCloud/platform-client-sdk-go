package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontablecolumndefaultrowvalue - Must provide a valid value for exactly one of the fields in this class.
type Decisiontablecolumndefaultrowvalue struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Value - A default string value for this column, will be cast to appropriate type according to the relevant contract schema property.
	Value *string `json:"value,omitempty"`

	// Values - A default list of values for this column, items will be cast to appropriate type according to the relevant contract schema property
	Values *[]string `json:"values,omitempty"`

	// Special - A default special value enum for this column.
	Special *string `json:"special,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontablecolumndefaultrowvalue) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontablecolumndefaultrowvalue) MarshalJSON() ([]byte, error) {
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
	type Alias Decisiontablecolumndefaultrowvalue
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		Special *string `json:"special,omitempty"`
		Alias
	}{ 
		Value: o.Value,
		
		Values: o.Values,
		
		Special: o.Special,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontablecolumndefaultrowvalue) UnmarshalJSON(b []byte) error {
	var DecisiontablecolumndefaultrowvalueMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontablecolumndefaultrowvalueMap)
	if err != nil {
		return err
	}
	
	if Value, ok := DecisiontablecolumndefaultrowvalueMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Values, ok := DecisiontablecolumndefaultrowvalueMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Special, ok := DecisiontablecolumndefaultrowvalueMap["special"].(string); ok {
		o.Special = &Special
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontablecolumndefaultrowvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

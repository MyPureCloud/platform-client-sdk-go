package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemsusereventsnotificationdelta
type Workitemsusereventsnotificationdelta struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Op
	Op *string `json:"op,omitempty"`

	// Field
	Field *string `json:"field,omitempty"`

	// OldValue
	OldValue *string `json:"oldValue,omitempty"`

	// NewValue
	NewValue *string `json:"newValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemsusereventsnotificationdelta) SetField(field string, fieldValue interface{}) {
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

func (o Workitemsusereventsnotificationdelta) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemsusereventsnotificationdelta
	
	return json.Marshal(&struct { 
		Op *string `json:"op,omitempty"`
		
		Field *string `json:"field,omitempty"`
		
		OldValue *string `json:"oldValue,omitempty"`
		
		NewValue *string `json:"newValue,omitempty"`
		Alias
	}{ 
		Op: o.Op,
		
		Field: o.Field,
		
		OldValue: o.OldValue,
		
		NewValue: o.NewValue,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemsusereventsnotificationdelta) UnmarshalJSON(b []byte) error {
	var WorkitemsusereventsnotificationdeltaMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemsusereventsnotificationdeltaMap)
	if err != nil {
		return err
	}
	
	if Op, ok := WorkitemsusereventsnotificationdeltaMap["op"].(string); ok {
		o.Op = &Op
	}
    
	if Field, ok := WorkitemsusereventsnotificationdeltaMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if OldValue, ok := WorkitemsusereventsnotificationdeltaMap["oldValue"].(string); ok {
		o.OldValue = &OldValue
	}
    
	if NewValue, ok := WorkitemsusereventsnotificationdeltaMap["newValue"].(string); ok {
		o.NewValue = &NewValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemsusereventsnotificationdelta) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

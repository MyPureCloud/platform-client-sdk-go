package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Numericrange
type Numericrange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Gt - Greater than
	Gt *float32 `json:"gt,omitempty"`

	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`

	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`

	// Lte - Less than or equal to
	Lte *float32 `json:"lte,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Numericrange) SetField(field string, fieldValue interface{}) {
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

func (o Numericrange) MarshalJSON() ([]byte, error) {
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
	type Alias Numericrange
	
	return json.Marshal(&struct { 
		Gt *float32 `json:"gt,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		
		Lte *float32 `json:"lte,omitempty"`
		Alias
	}{ 
		Gt: o.Gt,
		
		Gte: o.Gte,
		
		Lt: o.Lt,
		
		Lte: o.Lte,
		Alias:    (Alias)(o),
	})
}

func (o *Numericrange) UnmarshalJSON(b []byte) error {
	var NumericrangeMap map[string]interface{}
	err := json.Unmarshal(b, &NumericrangeMap)
	if err != nil {
		return err
	}
	
	if Gt, ok := NumericrangeMap["gt"].(float64); ok {
		GtFloat32 := float32(Gt)
		o.Gt = &GtFloat32
	}
    
	if Gte, ok := NumericrangeMap["gte"].(float64); ok {
		GteFloat32 := float32(Gte)
		o.Gte = &GteFloat32
	}
    
	if Lt, ok := NumericrangeMap["lt"].(float64); ok {
		LtFloat32 := float32(Lt)
		o.Lt = &LtFloat32
	}
    
	if Lte, ok := NumericrangeMap["lte"].(float64); ok {
		LteFloat32 := float32(Lte)
		o.Lte = &LteFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Numericrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

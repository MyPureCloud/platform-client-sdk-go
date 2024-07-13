package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequencescheduleconfigchangealteration
type Dialersequencescheduleconfigchangealteration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType
	VarType *string `json:"type,omitempty"`

	// Start - the end date of an alteration range as an ISO-8601 string
	Start *string `json:"start,omitempty"`

	// End - the end date of an alteration range as an ISO-8601 string
	End *string `json:"end,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialersequencescheduleconfigchangealteration) SetField(field string, fieldValue interface{}) {
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

func (o Dialersequencescheduleconfigchangealteration) MarshalJSON() ([]byte, error) {
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
	type Alias Dialersequencescheduleconfigchangealteration
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Start: o.Start,
		
		End: o.End,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialersequencescheduleconfigchangealteration) UnmarshalJSON(b []byte) error {
	var DialersequencescheduleconfigchangealterationMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequencescheduleconfigchangealterationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DialersequencescheduleconfigchangealterationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Start, ok := DialersequencescheduleconfigchangealterationMap["start"].(string); ok {
		o.Start = &Start
	}
    
	if End, ok := DialersequencescheduleconfigchangealterationMap["end"].(string); ok {
		o.End = &End
	}
    
	if AdditionalProperties, ok := DialersequencescheduleconfigchangealterationMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangealteration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

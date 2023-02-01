package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Validationlimits
type Validationlimits struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MinLength
	MinLength *Minlength `json:"minLength,omitempty"`

	// MaxLength
	MaxLength *Maxlength `json:"maxLength,omitempty"`

	// MinItems
	MinItems *Minlength `json:"minItems,omitempty"`

	// MaxItems
	MaxItems *Maxlength `json:"maxItems,omitempty"`

	// Minimum
	Minimum *Minlength `json:"minimum,omitempty"`

	// Maximum
	Maximum *Maxlength `json:"maximum,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Validationlimits) SetField(field string, fieldValue interface{}) {
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

func (o Validationlimits) MarshalJSON() ([]byte, error) {
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
	type Alias Validationlimits
	
	return json.Marshal(&struct { 
		MinLength *Minlength `json:"minLength,omitempty"`
		
		MaxLength *Maxlength `json:"maxLength,omitempty"`
		
		MinItems *Minlength `json:"minItems,omitempty"`
		
		MaxItems *Maxlength `json:"maxItems,omitempty"`
		
		Minimum *Minlength `json:"minimum,omitempty"`
		
		Maximum *Maxlength `json:"maximum,omitempty"`
		Alias
	}{ 
		MinLength: o.MinLength,
		
		MaxLength: o.MaxLength,
		
		MinItems: o.MinItems,
		
		MaxItems: o.MaxItems,
		
		Minimum: o.Minimum,
		
		Maximum: o.Maximum,
		Alias:    (Alias)(o),
	})
}

func (o *Validationlimits) UnmarshalJSON(b []byte) error {
	var ValidationlimitsMap map[string]interface{}
	err := json.Unmarshal(b, &ValidationlimitsMap)
	if err != nil {
		return err
	}
	
	if MinLength, ok := ValidationlimitsMap["minLength"].(map[string]interface{}); ok {
		MinLengthString, _ := json.Marshal(MinLength)
		json.Unmarshal(MinLengthString, &o.MinLength)
	}
	
	if MaxLength, ok := ValidationlimitsMap["maxLength"].(map[string]interface{}); ok {
		MaxLengthString, _ := json.Marshal(MaxLength)
		json.Unmarshal(MaxLengthString, &o.MaxLength)
	}
	
	if MinItems, ok := ValidationlimitsMap["minItems"].(map[string]interface{}); ok {
		MinItemsString, _ := json.Marshal(MinItems)
		json.Unmarshal(MinItemsString, &o.MinItems)
	}
	
	if MaxItems, ok := ValidationlimitsMap["maxItems"].(map[string]interface{}); ok {
		MaxItemsString, _ := json.Marshal(MaxItems)
		json.Unmarshal(MaxItemsString, &o.MaxItems)
	}
	
	if Minimum, ok := ValidationlimitsMap["minimum"].(map[string]interface{}); ok {
		MinimumString, _ := json.Marshal(Minimum)
		json.Unmarshal(MinimumString, &o.Minimum)
	}
	
	if Maximum, ok := ValidationlimitsMap["maximum"].(map[string]interface{}); ok {
		MaximumString, _ := json.Marshal(Maximum)
		json.Unmarshal(MaximumString, &o.Maximum)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validationlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

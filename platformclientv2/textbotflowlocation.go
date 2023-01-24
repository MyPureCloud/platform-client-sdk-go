package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlocation - Describes a flow location.
type Textbotflowlocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActionName - The name of the action that was active when the event of interest happened.
	ActionName *string `json:"actionName,omitempty"`

	// ActionNumber - The number of the action that was active when the event of interest happened.
	ActionNumber *int `json:"actionNumber,omitempty"`

	// SequenceName - The name of the state or task which was active when the event of interest happened.
	SequenceName *string `json:"sequenceName,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotflowlocation) SetField(field string, fieldValue interface{}) {
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

func (o Textbotflowlocation) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotflowlocation
	
	return json.Marshal(&struct { 
		ActionName *string `json:"actionName,omitempty"`
		
		ActionNumber *int `json:"actionNumber,omitempty"`
		
		SequenceName *string `json:"sequenceName,omitempty"`
		Alias
	}{ 
		ActionName: o.ActionName,
		
		ActionNumber: o.ActionNumber,
		
		SequenceName: o.SequenceName,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotflowlocation) UnmarshalJSON(b []byte) error {
	var TextbotflowlocationMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowlocationMap)
	if err != nil {
		return err
	}
	
	if ActionName, ok := TextbotflowlocationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
    
	if ActionNumber, ok := TextbotflowlocationMap["actionNumber"].(float64); ok {
		ActionNumberInt := int(ActionNumber)
		o.ActionNumber = &ActionNumberInt
	}
	
	if SequenceName, ok := TextbotflowlocationMap["sequenceName"].(string); ok {
		o.SequenceName = &SequenceName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowlocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

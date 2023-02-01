package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnaction
type Reportingturnaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActionId - The ID of the action in the bot flow.
	ActionId *string `json:"actionId,omitempty"`

	// ActionName - The name of the action in the bot flow.
	ActionName *string `json:"actionName,omitempty"`

	// ActionNumber - The number of the action in the bot flow.
	ActionNumber *int `json:"actionNumber,omitempty"`

	// ActionType
	ActionType *string `json:"actionType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingturnaction) SetField(field string, fieldValue interface{}) {
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

func (o Reportingturnaction) MarshalJSON() ([]byte, error) {
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
	type Alias Reportingturnaction
	
	return json.Marshal(&struct { 
		ActionId *string `json:"actionId,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionNumber *int `json:"actionNumber,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		Alias
	}{ 
		ActionId: o.ActionId,
		
		ActionName: o.ActionName,
		
		ActionNumber: o.ActionNumber,
		
		ActionType: o.ActionType,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingturnaction) UnmarshalJSON(b []byte) error {
	var ReportingturnactionMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnactionMap)
	if err != nil {
		return err
	}
	
	if ActionId, ok := ReportingturnactionMap["actionId"].(string); ok {
		o.ActionId = &ActionId
	}
    
	if ActionName, ok := ReportingturnactionMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
    
	if ActionNumber, ok := ReportingturnactionMap["actionNumber"].(float64); ok {
		ActionNumberInt := int(ActionNumber)
		o.ActionNumber = &ActionNumberInt
	}
	
	if ActionType, ok := ReportingturnactionMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

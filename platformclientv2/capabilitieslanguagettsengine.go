package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitieslanguagettsengine
type Capabilitieslanguagettsengine struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FlowTypes
	FlowTypes *map[string]Capabilitiesavailability `json:"flowTypes,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// IsDefault
	IsDefault *bool `json:"isDefault,omitempty"`

	// Voices
	Voices *map[string]Capabilitieslanguagettsenginevoice `json:"voices,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitieslanguagettsengine) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitieslanguagettsengine) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitieslanguagettsengine
	
	return json.Marshal(&struct { 
		FlowTypes *map[string]Capabilitiesavailability `json:"flowTypes,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		Voices *map[string]Capabilitieslanguagettsenginevoice `json:"voices,omitempty"`
		Alias
	}{ 
		FlowTypes: o.FlowTypes,
		
		Id: o.Id,
		
		IsDefault: o.IsDefault,
		
		Voices: o.Voices,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitieslanguagettsengine) UnmarshalJSON(b []byte) error {
	var CapabilitieslanguagettsengineMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitieslanguagettsengineMap)
	if err != nil {
		return err
	}
	
	if FlowTypes, ok := CapabilitieslanguagettsengineMap["flowTypes"].(map[string]interface{}); ok {
		FlowTypesString, _ := json.Marshal(FlowTypes)
		json.Unmarshal(FlowTypesString, &o.FlowTypes)
	}
	
	if Id, ok := CapabilitieslanguagettsengineMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IsDefault, ok := CapabilitieslanguagettsengineMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if Voices, ok := CapabilitieslanguagettsengineMap["voices"].(map[string]interface{}); ok {
		VoicesString, _ := json.Marshal(Voices)
		json.Unmarshal(VoicesString, &o.Voices)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagettsengine) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

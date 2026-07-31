package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createverifierresponse
type Createverifierresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The unique identifier of the verifier.
	Id *string `json:"id,omitempty"`

	// Name - The name of the verifier.
	Name *string `json:"name,omitempty"`

	// VarType - The type of verifier.
	VarType *string `json:"type,omitempty"`

	// Enabled - Indicates whether this verifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// KeyUri - The key URI for TOTP authenticator app registration.
	KeyUri *string `json:"keyUri,omitempty"`

	// VarDefault - Indicates whether this is the default verifier.
	VarDefault *bool `json:"default,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createverifierresponse) SetField(field string, fieldValue interface{}) {
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

func (o Createverifierresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Createverifierresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		KeyUri *string `json:"keyUri,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Enabled: o.Enabled,
		
		KeyUri: o.KeyUri,
		
		VarDefault: o.VarDefault,
		Alias:    (Alias)(o),
	})
}

func (o *Createverifierresponse) UnmarshalJSON(b []byte) error {
	var CreateverifierresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreateverifierresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateverifierresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CreateverifierresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := CreateverifierresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Enabled, ok := CreateverifierresponseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if KeyUri, ok := CreateverifierresponseMap["keyUri"].(string); ok {
		o.KeyUri = &KeyUri
	}
    
	if VarDefault, ok := CreateverifierresponseMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createverifierresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

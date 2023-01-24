package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationencryptionconfiguration
type Conversationencryptionconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Url - keyConfigurationType is always KmsSymmetric, and should be the arn to the key alias for the master key
	Url *string `json:"url,omitempty"`

	// KeyConfigurationType - Type should be 'KmsSymmetric' when create or update Key configurations, 'None' to disable configuration.
	KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`

	// LastError - The error message related to the configuration
	LastError *Errorbody `json:"lastError,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationencryptionconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Conversationencryptionconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationencryptionconfiguration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`
		
		LastError *Errorbody `json:"lastError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Url: o.Url,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		LastError: o.LastError,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationencryptionconfiguration) UnmarshalJSON(b []byte) error {
	var ConversationencryptionconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationencryptionconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationencryptionconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Url, ok := ConversationencryptionconfigurationMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if KeyConfigurationType, ok := ConversationencryptionconfigurationMap["keyConfigurationType"].(string); ok {
		o.KeyConfigurationType = &KeyConfigurationType
	}
    
	if LastError, ok := ConversationencryptionconfigurationMap["lastError"].(map[string]interface{}); ok {
		LastErrorString, _ := json.Marshal(LastError)
		json.Unmarshal(LastErrorString, &o.LastError)
	}
	
	if SelfUri, ok := ConversationencryptionconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationencryptionconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

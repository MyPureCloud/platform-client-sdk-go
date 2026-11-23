package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesttsvoice
type Capabilitiesttsvoice struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DisplayName - The display name of the TTS voice
	DisplayName *string `json:"displayName,omitempty"`

	// Gender - The gender of the TTS voice
	Gender *string `json:"gender,omitempty"`

	// Language - The language supported by the TTS voice, as a lowercase language-region tag, e.g. en-us
	Language *string `json:"language,omitempty"`

	// Name - The name of the TTS voice
	Name *string `json:"name,omitempty"`

	// Provider - The provider of the TTS voice. Values are defined by the upstream Speech Integration Service and driven by third-party vendor integrations, so new values may appear without a corresponding API change; treat as an opaque string
	Provider *string `json:"provider,omitempty"`

	// SupportedModels - The models supported by the TTS voice. Model identifiers are vendor-specific and are defined by the upstream Speech Integration Service, so new values may appear without a corresponding API change; treat as opaque strings
	SupportedModels *[]string `json:"supportedModels,omitempty"`

	// VoiceType - The type of the TTS voice. Values are defined by the upstream Speech Integration Service, so new values may appear without a corresponding API change; treat as an opaque string
	VoiceType *string `json:"voiceType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesttsvoice) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesttsvoice) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesttsvoice
	
	return json.Marshal(&struct { 
		DisplayName *string `json:"displayName,omitempty"`
		
		Gender *string `json:"gender,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SupportedModels *[]string `json:"supportedModels,omitempty"`
		
		VoiceType *string `json:"voiceType,omitempty"`
		Alias
	}{ 
		DisplayName: o.DisplayName,
		
		Gender: o.Gender,
		
		Language: o.Language,
		
		Name: o.Name,
		
		Provider: o.Provider,
		
		SupportedModels: o.SupportedModels,
		
		VoiceType: o.VoiceType,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesttsvoice) UnmarshalJSON(b []byte) error {
	var CapabilitiesttsvoiceMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesttsvoiceMap)
	if err != nil {
		return err
	}
	
	if DisplayName, ok := CapabilitiesttsvoiceMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Gender, ok := CapabilitiesttsvoiceMap["gender"].(string); ok {
		o.Gender = &Gender
	}
    
	if Language, ok := CapabilitiesttsvoiceMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Name, ok := CapabilitiesttsvoiceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Provider, ok := CapabilitiesttsvoiceMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if SupportedModels, ok := CapabilitiesttsvoiceMap["supportedModels"].([]interface{}); ok {
		SupportedModelsString, _ := json.Marshal(SupportedModels)
		json.Unmarshal(SupportedModelsString, &o.SupportedModels)
	}
	
	if VoiceType, ok := CapabilitiesttsvoiceMap["voiceType"].(string); ok {
		o.VoiceType = &VoiceType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesttsvoice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

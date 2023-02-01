package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ttsengineentity
type Ttsengineentity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Languages - The set of languages the TTS engine supports
	Languages *[]string `json:"languages,omitempty"`

	// OutputFormats - The set of output formats the TTS engine can produce
	OutputFormats *[]string `json:"outputFormats,omitempty"`

	// Voices - The set of voices the TTS engine supports
	Voices *[]Ttsvoiceentity `json:"voices,omitempty"`

	// IsDefault - The TTS engine is the global default engine
	IsDefault *bool `json:"isDefault,omitempty"`

	// IsSecure - The TTS engine can be used in a secure call flow
	IsSecure *bool `json:"isSecure,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ttsengineentity) SetField(field string, fieldValue interface{}) {
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

func (o Ttsengineentity) MarshalJSON() ([]byte, error) {
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
	type Alias Ttsengineentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		OutputFormats *[]string `json:"outputFormats,omitempty"`
		
		Voices *[]Ttsvoiceentity `json:"voices,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		IsSecure *bool `json:"isSecure,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Languages: o.Languages,
		
		OutputFormats: o.OutputFormats,
		
		Voices: o.Voices,
		
		IsDefault: o.IsDefault,
		
		IsSecure: o.IsSecure,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Ttsengineentity) UnmarshalJSON(b []byte) error {
	var TtsengineentityMap map[string]interface{}
	err := json.Unmarshal(b, &TtsengineentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TtsengineentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TtsengineentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Languages, ok := TtsengineentityMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if OutputFormats, ok := TtsengineentityMap["outputFormats"].([]interface{}); ok {
		OutputFormatsString, _ := json.Marshal(OutputFormats)
		json.Unmarshal(OutputFormatsString, &o.OutputFormats)
	}
	
	if Voices, ok := TtsengineentityMap["voices"].([]interface{}); ok {
		VoicesString, _ := json.Marshal(Voices)
		json.Unmarshal(VoicesString, &o.Voices)
	}
	
	if IsDefault, ok := TtsengineentityMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if IsSecure, ok := TtsengineentityMap["isSecure"].(bool); ok {
		o.IsSecure = &IsSecure
	}
    
	if SelfUri, ok := TtsengineentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ttsengineentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

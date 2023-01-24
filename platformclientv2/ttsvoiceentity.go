package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ttsvoiceentity
type Ttsvoiceentity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Gender - The gender of the TTS voice
	Gender *string `json:"gender,omitempty"`

	// Language - The language supported by the TTS voice
	Language *string `json:"language,omitempty"`

	// Engine - Ths TTS engine this voice belongs to
	Engine *Ttsengineentity `json:"engine,omitempty"`

	// IsDefault - The voice is the default voice for its language
	IsDefault *bool `json:"isDefault,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ttsvoiceentity) SetField(field string, fieldValue interface{}) {
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

func (o Ttsvoiceentity) MarshalJSON() ([]byte, error) {
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
	type Alias Ttsvoiceentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Gender *string `json:"gender,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Engine *Ttsengineentity `json:"engine,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Gender: o.Gender,
		
		Language: o.Language,
		
		Engine: o.Engine,
		
		IsDefault: o.IsDefault,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Ttsvoiceentity) UnmarshalJSON(b []byte) error {
	var TtsvoiceentityMap map[string]interface{}
	err := json.Unmarshal(b, &TtsvoiceentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TtsvoiceentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TtsvoiceentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Gender, ok := TtsvoiceentityMap["gender"].(string); ok {
		o.Gender = &Gender
	}
    
	if Language, ok := TtsvoiceentityMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Engine, ok := TtsvoiceentityMap["engine"].(map[string]interface{}); ok {
		EngineString, _ := json.Marshal(Engine)
		json.Unmarshal(EngineString, &o.Engine)
	}
	
	if IsDefault, ok := TtsvoiceentityMap["isDefault"].(bool); ok {
		o.IsDefault = &IsDefault
	}
    
	if SelfUri, ok := TtsvoiceentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ttsvoiceentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

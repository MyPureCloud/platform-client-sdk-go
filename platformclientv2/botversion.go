package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Botversion - A version description for a Bot.
type Botversion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - The name of the version.
	Version *string `json:"version,omitempty"`

	// SupportedLanguages - The supported languages for this bot. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code)
	SupportedLanguages *[]string `json:"supportedLanguages,omitempty"`

	// Intents - A list of potential intents this bot will return, limit of 50
	Intents *[]Botconnectorintent `json:"intents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Botversion) SetField(field string, fieldValue interface{}) {
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

func (o Botversion) MarshalJSON() ([]byte, error) {
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
	type Alias Botversion
	
	return json.Marshal(&struct { 
		Version *string `json:"version,omitempty"`
		
		SupportedLanguages *[]string `json:"supportedLanguages,omitempty"`
		
		Intents *[]Botconnectorintent `json:"intents,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		SupportedLanguages: o.SupportedLanguages,
		
		Intents: o.Intents,
		Alias:    (Alias)(o),
	})
}

func (o *Botversion) UnmarshalJSON(b []byte) error {
	var BotversionMap map[string]interface{}
	err := json.Unmarshal(b, &BotversionMap)
	if err != nil {
		return err
	}
	
	if Version, ok := BotversionMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if SupportedLanguages, ok := BotversionMap["supportedLanguages"].([]interface{}); ok {
		SupportedLanguagesString, _ := json.Marshal(SupportedLanguages)
		json.Unmarshal(SupportedLanguagesString, &o.SupportedLanguages)
	}
	
	if Intents, ok := BotversionMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

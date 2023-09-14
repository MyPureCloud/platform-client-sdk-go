package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowhealth
type Flowhealth struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FlowVersionInfo - Info about given flow version.
	FlowVersionInfo *Flowhealthversioninfo `json:"flowVersionInfo,omitempty"`

	// LanguageInfo - Each language's status about its health computation.
	LanguageInfo *map[string]Localeinfo `json:"languageInfo,omitempty"`

	// Intents - Health metrics information for the intents.
	Intents *[]Flowhealthintentinfo `json:"intents,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowhealth) SetField(field string, fieldValue interface{}) {
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

func (o Flowhealth) MarshalJSON() ([]byte, error) {
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
	type Alias Flowhealth
	
	return json.Marshal(&struct { 
		FlowVersionInfo *Flowhealthversioninfo `json:"flowVersionInfo,omitempty"`
		
		LanguageInfo *map[string]Localeinfo `json:"languageInfo,omitempty"`
		
		Intents *[]Flowhealthintentinfo `json:"intents,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		FlowVersionInfo: o.FlowVersionInfo,
		
		LanguageInfo: o.LanguageInfo,
		
		Intents: o.Intents,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flowhealth) UnmarshalJSON(b []byte) error {
	var FlowhealthMap map[string]interface{}
	err := json.Unmarshal(b, &FlowhealthMap)
	if err != nil {
		return err
	}
	
	if FlowVersionInfo, ok := FlowhealthMap["flowVersionInfo"].(map[string]interface{}); ok {
		FlowVersionInfoString, _ := json.Marshal(FlowVersionInfo)
		json.Unmarshal(FlowVersionInfoString, &o.FlowVersionInfo)
	}
	
	if LanguageInfo, ok := FlowhealthMap["languageInfo"].(map[string]interface{}); ok {
		LanguageInfoString, _ := json.Marshal(LanguageInfo)
		json.Unmarshal(LanguageInfoString, &o.LanguageInfo)
	}
	
	if Intents, ok := FlowhealthMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if SelfUri, ok := FlowhealthMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowhealth) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

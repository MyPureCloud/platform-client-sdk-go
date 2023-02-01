package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseconfig - Defines response components of the Action Request.
type Responseconfig struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TranslationMap - Map 'attribute name' and 'JSON path' pairs used to extract data from REST response.
	TranslationMap *map[string]string `json:"translationMap,omitempty"`

	// TranslationMapDefaults - Map 'attribute name' and 'default value' pairs used as fallback values if JSON path extraction fails for specified key.
	TranslationMapDefaults *map[string]string `json:"translationMapDefaults,omitempty"`

	// SuccessTemplate - Velocity template to build response to return from Action.
	SuccessTemplate *string `json:"successTemplate,omitempty"`

	// SuccessTemplateUri - URI to retrieve success template.
	SuccessTemplateUri *string `json:"successTemplateUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Responseconfig) SetField(field string, fieldValue interface{}) {
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

func (o Responseconfig) MarshalJSON() ([]byte, error) {
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
	type Alias Responseconfig
	
	return json.Marshal(&struct { 
		TranslationMap *map[string]string `json:"translationMap,omitempty"`
		
		TranslationMapDefaults *map[string]string `json:"translationMapDefaults,omitempty"`
		
		SuccessTemplate *string `json:"successTemplate,omitempty"`
		
		SuccessTemplateUri *string `json:"successTemplateUri,omitempty"`
		Alias
	}{ 
		TranslationMap: o.TranslationMap,
		
		TranslationMapDefaults: o.TranslationMapDefaults,
		
		SuccessTemplate: o.SuccessTemplate,
		
		SuccessTemplateUri: o.SuccessTemplateUri,
		Alias:    (Alias)(o),
	})
}

func (o *Responseconfig) UnmarshalJSON(b []byte) error {
	var ResponseconfigMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseconfigMap)
	if err != nil {
		return err
	}
	
	if TranslationMap, ok := ResponseconfigMap["translationMap"].(map[string]interface{}); ok {
		TranslationMapString, _ := json.Marshal(TranslationMap)
		json.Unmarshal(TranslationMapString, &o.TranslationMap)
	}
	
	if TranslationMapDefaults, ok := ResponseconfigMap["translationMapDefaults"].(map[string]interface{}); ok {
		TranslationMapDefaultsString, _ := json.Marshal(TranslationMapDefaults)
		json.Unmarshal(TranslationMapDefaultsString, &o.TranslationMapDefaults)
	}
	
	if SuccessTemplate, ok := ResponseconfigMap["successTemplate"].(string); ok {
		o.SuccessTemplate = &SuccessTemplate
	}
    
	if SuccessTemplateUri, ok := ResponseconfigMap["successTemplateUri"].(string); ok {
		o.SuccessTemplateUri = &SuccessTemplateUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responseconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

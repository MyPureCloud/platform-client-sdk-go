package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether or not support center is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// KnowledgeBase - The knowledge base for support center
	KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`

	// CustomMessages - Customizable display texts for support center
	CustomMessages *[]Supportcentercustommessage `json:"customMessages,omitempty"`

	// RouterType - Router type for support center
	RouterType *string `json:"routerType,omitempty"`

	// Screens - Available screens for the support center with its modules
	Screens *[]Supportcenterscreen `json:"screens,omitempty"`

	// EnabledCategories - Enabled article categories for support center
	EnabledCategories *[]Addressableentityref `json:"enabledCategories,omitempty"`

	// StyleSetting - Style attributes for support center
	StyleSetting *Supportcenterstylesetting `json:"styleSetting,omitempty"`

	// Feedback - Customer feedback settings
	Feedback *Supportcenterfeedbacksettings `json:"feedback,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Supportcentersettings) SetField(field string, fieldValue interface{}) {
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

func (o Supportcentersettings) MarshalJSON() ([]byte, error) {
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
	type Alias Supportcentersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`
		
		CustomMessages *[]Supportcentercustommessage `json:"customMessages,omitempty"`
		
		RouterType *string `json:"routerType,omitempty"`
		
		Screens *[]Supportcenterscreen `json:"screens,omitempty"`
		
		EnabledCategories *[]Addressableentityref `json:"enabledCategories,omitempty"`
		
		StyleSetting *Supportcenterstylesetting `json:"styleSetting,omitempty"`
		
		Feedback *Supportcenterfeedbacksettings `json:"feedback,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		KnowledgeBase: o.KnowledgeBase,
		
		CustomMessages: o.CustomMessages,
		
		RouterType: o.RouterType,
		
		Screens: o.Screens,
		
		EnabledCategories: o.EnabledCategories,
		
		StyleSetting: o.StyleSetting,
		
		Feedback: o.Feedback,
		Alias:    (Alias)(o),
	})
}

func (o *Supportcentersettings) UnmarshalJSON(b []byte) error {
	var SupportcentersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SupportcentersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if KnowledgeBase, ok := SupportcentersettingsMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if CustomMessages, ok := SupportcentersettingsMap["customMessages"].([]interface{}); ok {
		CustomMessagesString, _ := json.Marshal(CustomMessages)
		json.Unmarshal(CustomMessagesString, &o.CustomMessages)
	}
	
	if RouterType, ok := SupportcentersettingsMap["routerType"].(string); ok {
		o.RouterType = &RouterType
	}
    
	if Screens, ok := SupportcentersettingsMap["screens"].([]interface{}); ok {
		ScreensString, _ := json.Marshal(Screens)
		json.Unmarshal(ScreensString, &o.Screens)
	}
	
	if EnabledCategories, ok := SupportcentersettingsMap["enabledCategories"].([]interface{}); ok {
		EnabledCategoriesString, _ := json.Marshal(EnabledCategories)
		json.Unmarshal(EnabledCategoriesString, &o.EnabledCategories)
	}
	
	if StyleSetting, ok := SupportcentersettingsMap["styleSetting"].(map[string]interface{}); ok {
		StyleSettingString, _ := json.Marshal(StyleSetting)
		json.Unmarshal(StyleSettingString, &o.StyleSetting)
	}
	
	if Feedback, ok := SupportcentersettingsMap["feedback"].(map[string]interface{}); ok {
		FeedbackString, _ := json.Marshal(Feedback)
		json.Unmarshal(FeedbackString, &o.Feedback)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

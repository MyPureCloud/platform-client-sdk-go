package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionproperties
type Actionproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WebchatPrompt - Prompt message shown to user, used for webchat type action.
	WebchatPrompt *string `json:"webchatPrompt,omitempty"`

	// WebchatTitleText - Title shown to the user, used for webchat type action.
	WebchatTitleText *string `json:"webchatTitleText,omitempty"`

	// WebchatAcceptText - Accept button text shown to user, used for webchat type action.
	WebchatAcceptText *string `json:"webchatAcceptText,omitempty"`

	// WebchatDeclineText - Decline button text shown to user, used for webchat type action.
	WebchatDeclineText *string `json:"webchatDeclineText,omitempty"`

	// WebchatSurvey - Survey provided to the user, used for webchat type action.
	WebchatSurvey *Actionsurvey `json:"webchatSurvey,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionproperties) SetField(field string, fieldValue interface{}) {
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

func (o Actionproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Actionproperties
	
	return json.Marshal(&struct { 
		WebchatPrompt *string `json:"webchatPrompt,omitempty"`
		
		WebchatTitleText *string `json:"webchatTitleText,omitempty"`
		
		WebchatAcceptText *string `json:"webchatAcceptText,omitempty"`
		
		WebchatDeclineText *string `json:"webchatDeclineText,omitempty"`
		
		WebchatSurvey *Actionsurvey `json:"webchatSurvey,omitempty"`
		Alias
	}{ 
		WebchatPrompt: o.WebchatPrompt,
		
		WebchatTitleText: o.WebchatTitleText,
		
		WebchatAcceptText: o.WebchatAcceptText,
		
		WebchatDeclineText: o.WebchatDeclineText,
		
		WebchatSurvey: o.WebchatSurvey,
		Alias:    (Alias)(o),
	})
}

func (o *Actionproperties) UnmarshalJSON(b []byte) error {
	var ActionpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ActionpropertiesMap)
	if err != nil {
		return err
	}
	
	if WebchatPrompt, ok := ActionpropertiesMap["webchatPrompt"].(string); ok {
		o.WebchatPrompt = &WebchatPrompt
	}
    
	if WebchatTitleText, ok := ActionpropertiesMap["webchatTitleText"].(string); ok {
		o.WebchatTitleText = &WebchatTitleText
	}
    
	if WebchatAcceptText, ok := ActionpropertiesMap["webchatAcceptText"].(string); ok {
		o.WebchatAcceptText = &WebchatAcceptText
	}
    
	if WebchatDeclineText, ok := ActionpropertiesMap["webchatDeclineText"].(string); ok {
		o.WebchatDeclineText = &WebchatDeclineText
	}
    
	if WebchatSurvey, ok := ActionpropertiesMap["webchatSurvey"].(map[string]interface{}); ok {
		WebchatSurveyString, _ := json.Marshal(WebchatSurvey)
		json.Unmarshal(WebchatSurveyString, &o.WebchatSurvey)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

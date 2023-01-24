package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nlufeedbackrequest
type Nlufeedbackrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Text - The feedback text.
	Text *string `json:"text,omitempty"`

	// Intents - Detected intent of the utterance
	Intents *[]Intentfeedback `json:"intents,omitempty"`

	// VersionId - The domain version ID of the feedback.
	VersionId *string `json:"versionId,omitempty"`

	// Language - The language of the version to which feedback is linked, e.g. en-us, de-de
	Language *string `json:"language,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nlufeedbackrequest) SetField(field string, fieldValue interface{}) {
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

func (o Nlufeedbackrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Nlufeedbackrequest
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Intents *[]Intentfeedback `json:"intents,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		Alias
	}{ 
		Text: o.Text,
		
		Intents: o.Intents,
		
		VersionId: o.VersionId,
		
		Language: o.Language,
		Alias:    (Alias)(o),
	})
}

func (o *Nlufeedbackrequest) UnmarshalJSON(b []byte) error {
	var NlufeedbackrequestMap map[string]interface{}
	err := json.Unmarshal(b, &NlufeedbackrequestMap)
	if err != nil {
		return err
	}
	
	if Text, ok := NlufeedbackrequestMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Intents, ok := NlufeedbackrequestMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if VersionId, ok := NlufeedbackrequestMap["versionId"].(string); ok {
		o.VersionId = &VersionId
	}
    
	if Language, ok := NlufeedbackrequestMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nlufeedbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

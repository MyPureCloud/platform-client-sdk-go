package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Promptasset
type Promptasset struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// PromptId - Associated prompt ID
	PromptId *string `json:"promptId,omitempty"`

	// Language - Prompt resource language
	Language *string `json:"language,omitempty"`

	// MediaUri - URI of the resource audio
	MediaUri *string `json:"mediaUri,omitempty"`

	// TtsString - Text to speech of the resource
	TtsString *string `json:"ttsString,omitempty"`

	// Text - Text of the resource
	Text *string `json:"text,omitempty"`

	// UploadStatus - Audio upload status
	UploadStatus *string `json:"uploadStatus,omitempty"`

	// UploadUri - Upload URI for the resource audio
	UploadUri *string `json:"uploadUri,omitempty"`

	// LanguageDefault - Whether or not this resource locale is the default for the language
	LanguageDefault *bool `json:"languageDefault,omitempty"`

	// Tags
	Tags *map[string][]string `json:"tags,omitempty"`

	// DurationSeconds
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Promptasset) SetField(field string, fieldValue interface{}) {
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

func (o Promptasset) MarshalJSON() ([]byte, error) {
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
	type Alias Promptasset
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PromptId *string `json:"promptId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		TtsString *string `json:"ttsString,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		UploadUri *string `json:"uploadUri,omitempty"`
		
		LanguageDefault *bool `json:"languageDefault,omitempty"`
		
		Tags *map[string][]string `json:"tags,omitempty"`
		
		DurationSeconds *float64 `json:"durationSeconds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		PromptId: o.PromptId,
		
		Language: o.Language,
		
		MediaUri: o.MediaUri,
		
		TtsString: o.TtsString,
		
		Text: o.Text,
		
		UploadStatus: o.UploadStatus,
		
		UploadUri: o.UploadUri,
		
		LanguageDefault: o.LanguageDefault,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Promptasset) UnmarshalJSON(b []byte) error {
	var PromptassetMap map[string]interface{}
	err := json.Unmarshal(b, &PromptassetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PromptassetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PromptassetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PromptId, ok := PromptassetMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
    
	if Language, ok := PromptassetMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if MediaUri, ok := PromptassetMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if TtsString, ok := PromptassetMap["ttsString"].(string); ok {
		o.TtsString = &TtsString
	}
    
	if Text, ok := PromptassetMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if UploadStatus, ok := PromptassetMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    
	if UploadUri, ok := PromptassetMap["uploadUri"].(string); ok {
		o.UploadUri = &UploadUri
	}
    
	if LanguageDefault, ok := PromptassetMap["languageDefault"].(bool); ok {
		o.LanguageDefault = &LanguageDefault
	}
    
	if Tags, ok := PromptassetMap["tags"].(map[string]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := PromptassetMap["durationSeconds"].(float64); ok {
		o.DurationSeconds = &DurationSeconds
	}
    
	if SelfUri, ok := PromptassetMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Promptasset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

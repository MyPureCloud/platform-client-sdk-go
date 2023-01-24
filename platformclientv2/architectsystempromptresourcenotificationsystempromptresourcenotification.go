package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectsystempromptresourcenotificationsystempromptresourcenotification
type Architectsystempromptresourcenotificationsystempromptresourcenotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PromptId - Id of the prompt that this notification is for.
	PromptId *string `json:"promptId,omitempty"`

	// Id - Id of the prompt resource that this notification is for.
	Id *string `json:"id,omitempty"`

	// Language - Language resource that this notification is for.
	Language *string `json:"language,omitempty"`

	// MediaUri - Uri to the file for this system prompt resource.
	MediaUri *string `json:"mediaUri,omitempty"`

	// UploadStatus - Current upload status of the prompt resource (created, uploaded, transcoded, transcodeFailed).
	UploadStatus *string `json:"uploadStatus,omitempty"`

	// DurationSeconds - Duration (in seconds) for the transcoded audio file.
	DurationSeconds *float32 `json:"durationSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) SetField(field string, fieldValue interface{}) {
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

func (o Architectsystempromptresourcenotificationsystempromptresourcenotification) MarshalJSON() ([]byte, error) {
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
	type Alias Architectsystempromptresourcenotificationsystempromptresourcenotification
	
	return json.Marshal(&struct { 
		PromptId *string `json:"promptId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		DurationSeconds *float32 `json:"durationSeconds,omitempty"`
		Alias
	}{ 
		PromptId: o.PromptId,
		
		Id: o.Id,
		
		Language: o.Language,
		
		MediaUri: o.MediaUri,
		
		UploadStatus: o.UploadStatus,
		
		DurationSeconds: o.DurationSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) UnmarshalJSON(b []byte) error {
	var ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap)
	if err != nil {
		return err
	}
	
	if PromptId, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
    
	if Id, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if MediaUri, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if UploadStatus, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    
	if DurationSeconds, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["durationSeconds"].(float64); ok {
		DurationSecondsFloat32 := float32(DurationSeconds)
		o.DurationSeconds = &DurationSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

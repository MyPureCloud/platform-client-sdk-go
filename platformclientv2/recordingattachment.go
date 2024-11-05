package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingattachment
type Recordingattachment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// MediaType - The type of attachment this instance represents.
	MediaType *string `json:"mediaType,omitempty"`

	// Url - URL of the attachment.
	Url *string `json:"url,omitempty"`

	// Mime - Attachment mime type.
	Mime *string `json:"mime,omitempty"`

	// Text - Text associated with attachment such as an image caption.
	Text *string `json:"text,omitempty"`

	// FileName - Suggested file name for attachment.
	FileName *string `json:"fileName,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingattachment) SetField(field string, fieldValue interface{}) {
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

func (o Recordingattachment) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingattachment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Mime *string `json:"mime,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		Url: o.Url,
		
		Mime: o.Mime,
		
		Text: o.Text,
		
		FileName: o.FileName,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingattachment) UnmarshalJSON(b []byte) error {
	var RecordingattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingattachmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MediaType, ok := RecordingattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Url, ok := RecordingattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Mime, ok := RecordingattachmentMap["mime"].(string); ok {
		o.Mime = &Mime
	}
    
	if Text, ok := RecordingattachmentMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if FileName, ok := RecordingattachmentMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

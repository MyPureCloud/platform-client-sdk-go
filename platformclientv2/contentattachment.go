package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentattachment - Attachment object.
type Contentattachment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Provider specific ID for attachment. For example, a LINE sticker ID.
	Id *string `json:"id,omitempty"`

	// MediaType - The type of attachment this instance represents.
	MediaType *string `json:"mediaType,omitempty"`

	// Url - URL of the attachment.
	Url *string `json:"url,omitempty"`

	// Mime - Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
	Mime *string `json:"mime,omitempty"`

	// Text - Text associated with attachment such as an image caption.
	Text *string `json:"text,omitempty"`

	// Sha256 - Secure hash of the attachment content.
	Sha256 *string `json:"sha256,omitempty"`

	// Filename - Suggested file name for attachment.
	Filename *string `json:"filename,omitempty"`

	// ContentSizeBytes - Size in bytes of the attachment content.
	ContentSizeBytes *int `json:"contentSizeBytes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contentattachment) SetField(field string, fieldValue interface{}) {
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

func (o Contentattachment) MarshalJSON() ([]byte, error) {
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
	type Alias Contentattachment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Mime *string `json:"mime,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Sha256 *string `json:"sha256,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		ContentSizeBytes *int `json:"contentSizeBytes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		Url: o.Url,
		
		Mime: o.Mime,
		
		Text: o.Text,
		
		Sha256: o.Sha256,
		
		Filename: o.Filename,
		
		ContentSizeBytes: o.ContentSizeBytes,
		Alias:    (Alias)(o),
	})
}

func (o *Contentattachment) UnmarshalJSON(b []byte) error {
	var ContentattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &ContentattachmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MediaType, ok := ContentattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Url, ok := ContentattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Mime, ok := ContentattachmentMap["mime"].(string); ok {
		o.Mime = &Mime
	}
    
	if Text, ok := ContentattachmentMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Sha256, ok := ContentattachmentMap["sha256"].(string); ok {
		o.Sha256 = &Sha256
	}
    
	if Filename, ok := ContentattachmentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if ContentSizeBytes, ok := ContentattachmentMap["contentSizeBytes"].(float64); ok {
		ContentSizeBytesInt := int(ContentSizeBytes)
		o.ContentSizeBytes = &ContentSizeBytesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

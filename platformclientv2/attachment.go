package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Attachment
type Attachment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AttachmentId - The unique identifier for the attachment.
	AttachmentId *string `json:"attachmentId,omitempty"`

	// Name - The name of the attachment.
	Name *string `json:"name,omitempty"`

	// ContentUri - The content uri of the attachment. If set, this is commonly a public api download location.
	ContentUri *string `json:"contentUri,omitempty"`

	// ContentType - The type of file the attachment is.
	ContentType *string `json:"contentType,omitempty"`

	// ContentLength - The length of the attachment file.
	ContentLength *int `json:"contentLength,omitempty"`

	// InlineImage - Whether or not the attachment was attached inline.,
	InlineImage *bool `json:"inlineImage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Attachment) SetField(field string, fieldValue interface{}) {
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

func (o Attachment) MarshalJSON() ([]byte, error) {
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
	type Alias Attachment
	
	return json.Marshal(&struct { 
		AttachmentId *string `json:"attachmentId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		InlineImage *bool `json:"inlineImage,omitempty"`
		Alias
	}{ 
		AttachmentId: o.AttachmentId,
		
		Name: o.Name,
		
		ContentUri: o.ContentUri,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		InlineImage: o.InlineImage,
		Alias:    (Alias)(o),
	})
}

func (o *Attachment) UnmarshalJSON(b []byte) error {
	var AttachmentMap map[string]interface{}
	err := json.Unmarshal(b, &AttachmentMap)
	if err != nil {
		return err
	}
	
	if AttachmentId, ok := AttachmentMap["attachmentId"].(string); ok {
		o.AttachmentId = &AttachmentId
	}
    
	if Name, ok := AttachmentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContentUri, ok := AttachmentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if ContentType, ok := AttachmentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := AttachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if InlineImage, ok := AttachmentMap["inlineImage"].(bool); ok {
		o.InlineImage = &InlineImage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Attachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

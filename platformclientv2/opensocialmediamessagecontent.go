package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Opensocialmediamessagecontent - Message content element.
type Opensocialmediamessagecontent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`

	// Attachment - Attachment content.
	Attachment *Conversationcontentattachment `json:"attachment,omitempty"`

	// Text - A content type containing text.
	Text *Conversationcontenttext `json:"text,omitempty"`

	// Reaction - A set of reactions to a message.
	Reaction *Conversationcontentreaction `json:"reaction,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Opensocialmediamessagecontent) SetField(field string, fieldValue interface{}) {
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

func (o Opensocialmediamessagecontent) MarshalJSON() ([]byte, error) {
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
	type Alias Opensocialmediamessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Attachment *Conversationcontentattachment `json:"attachment,omitempty"`
		
		Text *Conversationcontenttext `json:"text,omitempty"`
		
		Reaction *Conversationcontentreaction `json:"reaction,omitempty"`
		Alias
	}{ 
		ContentType: o.ContentType,
		
		Attachment: o.Attachment,
		
		Text: o.Text,
		
		Reaction: o.Reaction,
		Alias:    (Alias)(o),
	})
}

func (o *Opensocialmediamessagecontent) UnmarshalJSON(b []byte) error {
	var OpensocialmediamessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &OpensocialmediamessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := OpensocialmediamessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Attachment, ok := OpensocialmediamessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if Text, ok := OpensocialmediamessagecontentMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if Reaction, ok := OpensocialmediamessagecontentMap["reaction"].(map[string]interface{}); ok {
		ReactionString, _ := json.Marshal(Reaction)
		json.Unmarshal(ReactionString, &o.Reaction)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Opensocialmediamessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationmessagecontent
type V2conversationmessagetypingeventforusertopicconversationmessagecontent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContentType
	ContentType *string `json:"contentType,omitempty"`

	// Location
	Location *V2conversationmessagetypingeventforusertopicconversationcontentlocation `json:"location,omitempty"`

	// Story
	Story *V2conversationmessagetypingeventforusertopicconversationcontentstory `json:"story,omitempty"`

	// Attachment
	Attachment *V2conversationmessagetypingeventforusertopicconversationcontentattachment `json:"attachment,omitempty"`

	// QuickReply
	QuickReply *V2conversationmessagetypingeventforusertopicconversationcontentquickreply `json:"quickReply,omitempty"`

	// Template
	Template *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate `json:"template,omitempty"`

	// ButtonResponse
	ButtonResponse *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`

	// Generic
	Generic *V2conversationmessagetypingeventforusertopicconversationcontentgeneric `json:"generic,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforusertopicconversationmessagecontent) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforusertopicconversationmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *V2conversationmessagetypingeventforusertopicconversationcontentlocation `json:"location,omitempty"`
		
		Story *V2conversationmessagetypingeventforusertopicconversationcontentstory `json:"story,omitempty"`
		
		Attachment *V2conversationmessagetypingeventforusertopicconversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *V2conversationmessagetypingeventforusertopicconversationcontentquickreply `json:"quickReply,omitempty"`
		
		Template *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate `json:"template,omitempty"`
		
		ButtonResponse *V2conversationmessagetypingeventforusertopicconversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *V2conversationmessagetypingeventforusertopicconversationcontentgeneric `json:"generic,omitempty"`
		Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Story: o.Story,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		Template: o.Template,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Location, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Story, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Attachment, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if Template, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if ButtonResponse, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := V2conversationmessagetypingeventforusertopicconversationmessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

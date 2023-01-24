package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmessagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Conversationmessagecontent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`

	// Location - Location content.
	Location *Conversationcontentlocation `json:"location,omitempty"`

	// Attachment - Attachment content.
	Attachment *Conversationcontentattachment `json:"attachment,omitempty"`

	// QuickReply - Quick reply content.
	QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`

	// ButtonResponse - Button response content.
	ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`

	// Template - Template notification content.
	Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`

	// Story - Ephemeral story content.
	Story *Conversationcontentstory `json:"story,omitempty"`

	// Card - Card content
	Card *Conversationcontentcard `json:"card,omitempty"`

	// Carousel - Carousel content
	Carousel *Conversationcontentcarousel `json:"carousel,omitempty"`

	// Text - Text content.
	Text *Conversationcontenttext `json:"text,omitempty"`

	// QuickReplyV2 - Quick reply V2 content.
	QuickReplyV2 *Conversationcontentquickreplyv2 `json:"quickReplyV2,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationmessagecontent) SetField(field string, fieldValue interface{}) {
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

func (o Conversationmessagecontent) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *Conversationcontentlocation `json:"location,omitempty"`
		
		Attachment *Conversationcontentattachment `json:"attachment,omitempty"`
		
		QuickReply *Conversationcontentquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Conversationcontentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Template *Conversationcontentnotificationtemplate `json:"template,omitempty"`
		
		Story *Conversationcontentstory `json:"story,omitempty"`
		
		Card *Conversationcontentcard `json:"card,omitempty"`
		
		Carousel *Conversationcontentcarousel `json:"carousel,omitempty"`
		
		Text *Conversationcontenttext `json:"text,omitempty"`
		
		QuickReplyV2 *Conversationcontentquickreplyv2 `json:"quickReplyV2,omitempty"`
		Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Template: o.Template,
		
		Story: o.Story,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		
		Text: o.Text,
		
		QuickReplyV2: o.QuickReplyV2,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationmessagecontent) UnmarshalJSON(b []byte) error {
	var ConversationmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := ConversationmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Location, ok := ConversationmessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Attachment, ok := ConversationmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := ConversationmessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := ConversationmessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Template, ok := ConversationmessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if Story, ok := ConversationmessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Card, ok := ConversationmessagecontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := ConversationmessagecontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	
	if Text, ok := ConversationmessagecontentMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if QuickReplyV2, ok := ConversationmessagecontentMap["quickReplyV2"].(map[string]interface{}); ok {
		QuickReplyV2String, _ := json.Marshal(QuickReplyV2)
		json.Unmarshal(QuickReplyV2String, &o.QuickReplyV2)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

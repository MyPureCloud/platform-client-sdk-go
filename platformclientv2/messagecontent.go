package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Messagecontent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContentType - Type of this content element.
	ContentType *string `json:"contentType,omitempty"`

	// Location - Location content.
	Location *Contentlocation `json:"location,omitempty"`

	// Attachment - Attachment content.
	Attachment *Contentattachment `json:"attachment,omitempty"`

	// QuickReply - Quick reply content.
	QuickReply *Contentquickreply `json:"quickReply,omitempty"`

	// ButtonResponse - Button response content.
	ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`

	// Generic - Generic content (Deprecated).
	Generic *Contentgeneric `json:"generic,omitempty"`

	// List - List content (Deprecated).
	List *Contentlist `json:"list,omitempty"`

	// Template - Template notification content.
	Template *Contentnotificationtemplate `json:"template,omitempty"`

	// Reactions - A set of reactions to a message.
	Reactions *[]Contentreaction `json:"reactions,omitempty"`

	// Mention - Mention content.
	Mention *Messagingrecipient `json:"mention,omitempty"`

	// Postback - Structured message postback (Deprecated).
	Postback *Contentpostback `json:"postback,omitempty"`

	// Story - Ephemeral story content.
	Story *Contentstory `json:"story,omitempty"`

	// Card - Card content
	Card *Contentcard `json:"card,omitempty"`

	// Carousel - Carousel content
	Carousel *Contentcarousel `json:"carousel,omitempty"`

	// Text - Text content.
	Text *Contenttext `json:"text,omitempty"`

	// QuickReplyV2 - Quick reply V2 content.
	QuickReplyV2 *Contentquickreplyv2 `json:"quickReplyV2,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagecontent) SetField(field string, fieldValue interface{}) {
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

func (o Messagecontent) MarshalJSON() ([]byte, error) {
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
	type Alias Messagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Location *Contentlocation `json:"location,omitempty"`
		
		Attachment *Contentattachment `json:"attachment,omitempty"`
		
		QuickReply *Contentquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Contentbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Contentgeneric `json:"generic,omitempty"`
		
		List *Contentlist `json:"list,omitempty"`
		
		Template *Contentnotificationtemplate `json:"template,omitempty"`
		
		Reactions *[]Contentreaction `json:"reactions,omitempty"`
		
		Mention *Messagingrecipient `json:"mention,omitempty"`
		
		Postback *Contentpostback `json:"postback,omitempty"`
		
		Story *Contentstory `json:"story,omitempty"`
		
		Card *Contentcard `json:"card,omitempty"`
		
		Carousel *Contentcarousel `json:"carousel,omitempty"`
		
		Text *Contenttext `json:"text,omitempty"`
		
		QuickReplyV2 *Contentquickreplyv2 `json:"quickReplyV2,omitempty"`
		Alias
	}{ 
		ContentType: o.ContentType,
		
		Location: o.Location,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		
		List: o.List,
		
		Template: o.Template,
		
		Reactions: o.Reactions,
		
		Mention: o.Mention,
		
		Postback: o.Postback,
		
		Story: o.Story,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		
		Text: o.Text,
		
		QuickReplyV2: o.QuickReplyV2,
		Alias:    (Alias)(o),
	})
}

func (o *Messagecontent) UnmarshalJSON(b []byte) error {
	var MessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &MessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := MessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Location, ok := MessagecontentMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Attachment, ok := MessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := MessagecontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := MessagecontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := MessagecontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	
	if List, ok := MessagecontentMap["list"].(map[string]interface{}); ok {
		ListString, _ := json.Marshal(List)
		json.Unmarshal(ListString, &o.List)
	}
	
	if Template, ok := MessagecontentMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if Reactions, ok := MessagecontentMap["reactions"].([]interface{}); ok {
		ReactionsString, _ := json.Marshal(Reactions)
		json.Unmarshal(ReactionsString, &o.Reactions)
	}
	
	if Mention, ok := MessagecontentMap["mention"].(map[string]interface{}); ok {
		MentionString, _ := json.Marshal(Mention)
		json.Unmarshal(MentionString, &o.Mention)
	}
	
	if Postback, ok := MessagecontentMap["postback"].(map[string]interface{}); ok {
		PostbackString, _ := json.Marshal(Postback)
		json.Unmarshal(PostbackString, &o.Postback)
	}
	
	if Story, ok := MessagecontentMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	
	if Card, ok := MessagecontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := MessagecontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	
	if Text, ok := MessagecontentMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if QuickReplyV2, ok := MessagecontentMap["quickReplyV2"].(map[string]interface{}); ok {
		QuickReplyV2String, _ := json.Marshal(QuickReplyV2)
		json.Unmarshal(QuickReplyV2String, &o.QuickReplyV2)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

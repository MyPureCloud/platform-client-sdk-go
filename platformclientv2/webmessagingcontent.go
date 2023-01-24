package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingcontent - Message content element.
type Webmessagingcontent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`

	// Attachment - Attachment content.
	Attachment *Webmessagingattachment `json:"attachment,omitempty"`

	// QuickReply - Quick reply content.
	QuickReply *Webmessagingquickreply `json:"quickReply,omitempty"`

	// ButtonResponse - Button response content.
	ButtonResponse *Webmessagingbuttonresponse `json:"buttonResponse,omitempty"`

	// Generic - Generic content (Deprecated).
	Generic *Webmessaginggeneric `json:"generic,omitempty"`

	// Card - Card content
	Card *Contentcard `json:"card,omitempty"`

	// Carousel - Carousel content
	Carousel *Contentcarousel `json:"carousel,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webmessagingcontent) SetField(field string, fieldValue interface{}) {
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

func (o Webmessagingcontent) MarshalJSON() ([]byte, error) {
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
	type Alias Webmessagingcontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Attachment *Webmessagingattachment `json:"attachment,omitempty"`
		
		QuickReply *Webmessagingquickreply `json:"quickReply,omitempty"`
		
		ButtonResponse *Webmessagingbuttonresponse `json:"buttonResponse,omitempty"`
		
		Generic *Webmessaginggeneric `json:"generic,omitempty"`
		
		Card *Contentcard `json:"card,omitempty"`
		
		Carousel *Contentcarousel `json:"carousel,omitempty"`
		Alias
	}{ 
		ContentType: o.ContentType,
		
		Attachment: o.Attachment,
		
		QuickReply: o.QuickReply,
		
		ButtonResponse: o.ButtonResponse,
		
		Generic: o.Generic,
		
		Card: o.Card,
		
		Carousel: o.Carousel,
		Alias:    (Alias)(o),
	})
}

func (o *Webmessagingcontent) UnmarshalJSON(b []byte) error {
	var WebmessagingcontentMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingcontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := WebmessagingcontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Attachment, ok := WebmessagingcontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	
	if QuickReply, ok := WebmessagingcontentMap["quickReply"].(map[string]interface{}); ok {
		QuickReplyString, _ := json.Marshal(QuickReply)
		json.Unmarshal(QuickReplyString, &o.QuickReply)
	}
	
	if ButtonResponse, ok := WebmessagingcontentMap["buttonResponse"].(map[string]interface{}); ok {
		ButtonResponseString, _ := json.Marshal(ButtonResponse)
		json.Unmarshal(ButtonResponseString, &o.ButtonResponse)
	}
	
	if Generic, ok := WebmessagingcontentMap["generic"].(map[string]interface{}); ok {
		GenericString, _ := json.Marshal(Generic)
		json.Unmarshal(GenericString, &o.Generic)
	}
	
	if Card, ok := WebmessagingcontentMap["card"].(map[string]interface{}); ok {
		CardString, _ := json.Marshal(Card)
		json.Unmarshal(CardString, &o.Card)
	}
	
	if Carousel, ok := WebmessagingcontentMap["carousel"].(map[string]interface{}); ok {
		CarouselString, _ := json.Marshal(Carousel)
		json.Unmarshal(CarouselString, &o.Carousel)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingcontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentquickreply - Quick reply object.
type Conversationcontentquickreply struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`

	// Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
	Payload *string `json:"payload,omitempty"`

	// Image - URL of an image associated with the quick reply.
	Image *string `json:"image,omitempty"`

	// Action - Specifies the type of action that is triggered upon clicking the quick reply.
	Action *string `json:"action,omitempty"`

	// SummaryText - Summary of what the quick reply relates to.
	SummaryText *string `json:"summaryText,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentquickreply) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentquickreply) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentquickreply
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Payload *string `json:"payload,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		SummaryText *string `json:"summaryText,omitempty"`
		Alias
	}{ 
		Text: o.Text,
		
		Payload: o.Payload,
		
		Image: o.Image,
		
		Action: o.Action,
		
		SummaryText: o.SummaryText,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentquickreply) UnmarshalJSON(b []byte) error {
	var ConversationcontentquickreplyMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentquickreplyMap)
	if err != nil {
		return err
	}
	
	if Text, ok := ConversationcontentquickreplyMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Payload, ok := ConversationcontentquickreplyMap["payload"].(string); ok {
		o.Payload = &Payload
	}
    
	if Image, ok := ConversationcontentquickreplyMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Action, ok := ConversationcontentquickreplyMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if SummaryText, ok := ConversationcontentquickreplyMap["summaryText"].(string); ok {
		o.SummaryText = &SummaryText
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

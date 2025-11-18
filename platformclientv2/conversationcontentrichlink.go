package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentrichlink - A Rich Link attachment
type Conversationcontentrichlink struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Header - Header for the Rich Link.
	Header *Conversationcontentrichlinkheader `json:"header,omitempty"`

	// Footer - Footer text for the Rich Link.
	Footer *string `json:"footer,omitempty"`

	// Text - Text for the body of the Rich Link.
	Text *string `json:"text,omitempty"`

	// UrlLabel - Label for the URL of the Rich Link.
	UrlLabel *string `json:"urlLabel,omitempty"`

	// Url - Url for the Rich Link.
	Url *string `json:"url,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentrichlink) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentrichlink) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentrichlink
	
	return json.Marshal(&struct { 
		Header *Conversationcontentrichlinkheader `json:"header,omitempty"`
		
		Footer *string `json:"footer,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		UrlLabel *string `json:"urlLabel,omitempty"`
		
		Url *string `json:"url,omitempty"`
		Alias
	}{ 
		Header: o.Header,
		
		Footer: o.Footer,
		
		Text: o.Text,
		
		UrlLabel: o.UrlLabel,
		
		Url: o.Url,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentrichlink) UnmarshalJSON(b []byte) error {
	var ConversationcontentrichlinkMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentrichlinkMap)
	if err != nil {
		return err
	}
	
	if Header, ok := ConversationcontentrichlinkMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Footer, ok := ConversationcontentrichlinkMap["footer"].(string); ok {
		o.Footer = &Footer
	}
    
	if Text, ok := ConversationcontentrichlinkMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if UrlLabel, ok := ConversationcontentrichlinkMap["urlLabel"].(string); ok {
		o.UrlLabel = &UrlLabel
	}
    
	if Url, ok := ConversationcontentrichlinkMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentrichlink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

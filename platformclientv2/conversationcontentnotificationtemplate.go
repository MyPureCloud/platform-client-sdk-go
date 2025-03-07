package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentnotificationtemplate - Template notification object.
type Conversationcontentnotificationtemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The identifier of the message template in 'your-namespace@your-template-id/name' format. For External vendor (e.g WhatsApp), 'your-namespace@your-template-name'. For GenesysCloud canned response message template use 'cannedresponse' as your-namespace and use response ID as your-template-id (e.g. response ID=1234 then 'cannedresponse@1234')
	Id *string `json:"id,omitempty"`

	// Language - Template language.
	Language *string `json:"language,omitempty"`

	// Header - The template header.
	Header *Conversationnotificationtemplateheader `json:"header,omitempty"`

	// Body - The template body.
	Body *Conversationnotificationtemplatebody `json:"body,omitempty"`

	// Buttons - Template buttons
	Buttons *[]Conversationnotificationtemplatebutton `json:"buttons,omitempty"`

	// Footer - The template footer.
	Footer *Conversationnotificationtemplatefooter `json:"footer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentnotificationtemplate) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *Conversationnotificationtemplateheader `json:"header,omitempty"`
		
		Body *Conversationnotificationtemplatebody `json:"body,omitempty"`
		
		Buttons *[]Conversationnotificationtemplatebutton `json:"buttons,omitempty"`
		
		Footer *Conversationnotificationtemplatefooter `json:"footer,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Language: o.Language,
		
		Header: o.Header,
		
		Body: o.Body,
		
		Buttons: o.Buttons,
		
		Footer: o.Footer,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var ConversationcontentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcontentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := ConversationcontentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Header, ok := ConversationcontentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := ConversationcontentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Buttons, ok := ConversationcontentnotificationtemplateMap["buttons"].([]interface{}); ok {
		ButtonsString, _ := json.Marshal(Buttons)
		json.Unmarshal(ButtonsString, &o.Buttons)
	}
	
	if Footer, ok := ConversationcontentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

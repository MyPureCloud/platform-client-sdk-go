package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate
type V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`

	// Header
	Header *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateheader `json:"header,omitempty"`

	// Body
	Body *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody `json:"body,omitempty"`

	// Footer
	Footer *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter `json:"footer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Header *V2conversationmessagetypingeventforusertopicconversationnotificationtemplateheader `json:"header,omitempty"`
		
		Body *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatebody `json:"body,omitempty"`
		
		Footer *V2conversationmessagetypingeventforusertopicconversationnotificationtemplatefooter `json:"footer,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Language: o.Language,
		
		Header: o.Header,
		
		Body: o.Body,
		
		Footer: o.Footer,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Header, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["header"].(map[string]interface{}); ok {
		HeaderString, _ := json.Marshal(Header)
		json.Unmarshal(HeaderString, &o.Header)
	}
	
	if Body, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Footer, ok := V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplateMap["footer"].(map[string]interface{}); ok {
		FooterString, _ := json.Marshal(Footer)
		json.Unmarshal(FooterString, &o.Footer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationcontentnotificationtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

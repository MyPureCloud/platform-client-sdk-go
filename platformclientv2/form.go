package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Form - Form configuration for response management
type Form struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FormDescription - Description of the form
	FormDescription *string `json:"formDescription,omitempty"`

	// ReceivedMessage - Message displayed when response is received
	ReceivedMessage *Formmessage `json:"receivedMessage,omitempty"`

	// ReplyMessage - Message displayed as reply
	ReplyMessage *Formmessage `json:"replyMessage,omitempty"`

	// Introduction - Introduction section of the form
	Introduction *Formintroduction `json:"introduction,omitempty"`

	// FormPages - Pages of the form
	FormPages *[]Formpage `json:"formPages,omitempty"`

	// ShowSummary - Whether to show a summary after form completion
	ShowSummary *bool `json:"showSummary,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Form) SetField(field string, fieldValue interface{}) {
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

func (o Form) MarshalJSON() ([]byte, error) {
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
	type Alias Form
	
	return json.Marshal(&struct { 
		FormDescription *string `json:"formDescription,omitempty"`
		
		ReceivedMessage *Formmessage `json:"receivedMessage,omitempty"`
		
		ReplyMessage *Formmessage `json:"replyMessage,omitempty"`
		
		Introduction *Formintroduction `json:"introduction,omitempty"`
		
		FormPages *[]Formpage `json:"formPages,omitempty"`
		
		ShowSummary *bool `json:"showSummary,omitempty"`
		Alias
	}{ 
		FormDescription: o.FormDescription,
		
		ReceivedMessage: o.ReceivedMessage,
		
		ReplyMessage: o.ReplyMessage,
		
		Introduction: o.Introduction,
		
		FormPages: o.FormPages,
		
		ShowSummary: o.ShowSummary,
		Alias:    (Alias)(o),
	})
}

func (o *Form) UnmarshalJSON(b []byte) error {
	var FormMap map[string]interface{}
	err := json.Unmarshal(b, &FormMap)
	if err != nil {
		return err
	}
	
	if FormDescription, ok := FormMap["formDescription"].(string); ok {
		o.FormDescription = &FormDescription
	}
    
	if ReceivedMessage, ok := FormMap["receivedMessage"].(map[string]interface{}); ok {
		ReceivedMessageString, _ := json.Marshal(ReceivedMessage)
		json.Unmarshal(ReceivedMessageString, &o.ReceivedMessage)
	}
	
	if ReplyMessage, ok := FormMap["replyMessage"].(map[string]interface{}); ok {
		ReplyMessageString, _ := json.Marshal(ReplyMessage)
		json.Unmarshal(ReplyMessageString, &o.ReplyMessage)
	}
	
	if Introduction, ok := FormMap["introduction"].(map[string]interface{}); ok {
		IntroductionString, _ := json.Marshal(Introduction)
		json.Unmarshal(IntroductionString, &o.Introduction)
	}
	
	if FormPages, ok := FormMap["formPages"].([]interface{}); ok {
		FormPagesString, _ := json.Marshal(FormPages)
		json.Unmarshal(FormPagesString, &o.FormPages)
	}
	
	if ShowSummary, ok := FormMap["showSummary"].(bool); ok {
		o.ShowSummary = &ShowSummary
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Form) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

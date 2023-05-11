package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historyheaderstranslation
type Historyheaderstranslation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// From - A translation for the word \"from\", for the expected language
	From *string `json:"from,omitempty"`

	// To - A translation for the word \"to\", for the expected language
	To *string `json:"to,omitempty"`

	// Cc - A translation for the word \"cc\", for the expected language
	Cc *string `json:"cc,omitempty"`

	// Subject - A translation for the word \"subject\", for the expected language
	Subject *string `json:"subject,omitempty"`

	// ReplyPrefix - A translation for the subject prefix \"Reply\", for the expected language
	ReplyPrefix *string `json:"replyPrefix,omitempty"`

	// ForwardPrefix - A translation for the subject prefix \"Forward\", for the expected language
	ForwardPrefix *string `json:"forwardPrefix,omitempty"`

	// Sent - A translation for the word \"sent\", for the expected language
	Sent *string `json:"sent,omitempty"`

	// Language - The code of the expected language
	Language *string `json:"language,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historyheaderstranslation) SetField(field string, fieldValue interface{}) {
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

func (o Historyheaderstranslation) MarshalJSON() ([]byte, error) {
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
	type Alias Historyheaderstranslation
	
	return json.Marshal(&struct { 
		From *string `json:"from,omitempty"`
		
		To *string `json:"to,omitempty"`
		
		Cc *string `json:"cc,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		ReplyPrefix *string `json:"replyPrefix,omitempty"`
		
		ForwardPrefix *string `json:"forwardPrefix,omitempty"`
		
		Sent *string `json:"sent,omitempty"`
		
		Language *string `json:"language,omitempty"`
		Alias
	}{ 
		From: o.From,
		
		To: o.To,
		
		Cc: o.Cc,
		
		Subject: o.Subject,
		
		ReplyPrefix: o.ReplyPrefix,
		
		ForwardPrefix: o.ForwardPrefix,
		
		Sent: o.Sent,
		
		Language: o.Language,
		Alias:    (Alias)(o),
	})
}

func (o *Historyheaderstranslation) UnmarshalJSON(b []byte) error {
	var HistoryheaderstranslationMap map[string]interface{}
	err := json.Unmarshal(b, &HistoryheaderstranslationMap)
	if err != nil {
		return err
	}
	
	if From, ok := HistoryheaderstranslationMap["from"].(string); ok {
		o.From = &From
	}
    
	if To, ok := HistoryheaderstranslationMap["to"].(string); ok {
		o.To = &To
	}
    
	if Cc, ok := HistoryheaderstranslationMap["cc"].(string); ok {
		o.Cc = &Cc
	}
    
	if Subject, ok := HistoryheaderstranslationMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if ReplyPrefix, ok := HistoryheaderstranslationMap["replyPrefix"].(string); ok {
		o.ReplyPrefix = &ReplyPrefix
	}
    
	if ForwardPrefix, ok := HistoryheaderstranslationMap["forwardPrefix"].(string); ok {
		o.ForwardPrefix = &ForwardPrefix
	}
    
	if Sent, ok := HistoryheaderstranslationMap["sent"].(string); ok {
		o.Sent = &Sent
	}
    
	if Language, ok := HistoryheaderstranslationMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Historyheaderstranslation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

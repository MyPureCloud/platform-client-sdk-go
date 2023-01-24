package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatconversationresponse
type Createwebchatconversationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Chat Conversation identifier
	Id *string `json:"id,omitempty"`

	// Jwt - The JWT that you can use to identify subsequent calls on this conversation
	Jwt *string `json:"jwt,omitempty"`

	// EventStreamUri - The URI which provides the conversation event stream.
	EventStreamUri *string `json:"eventStreamUri,omitempty"`

	// Member - Chat Member
	Member *Webchatmemberinfo `json:"member,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createwebchatconversationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Createwebchatconversationresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Createwebchatconversationresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Jwt *string `json:"jwt,omitempty"`
		
		EventStreamUri *string `json:"eventStreamUri,omitempty"`
		
		Member *Webchatmemberinfo `json:"member,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Jwt: o.Jwt,
		
		EventStreamUri: o.EventStreamUri,
		
		Member: o.Member,
		Alias:    (Alias)(o),
	})
}

func (o *Createwebchatconversationresponse) UnmarshalJSON(b []byte) error {
	var CreatewebchatconversationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatconversationresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreatewebchatconversationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Jwt, ok := CreatewebchatconversationresponseMap["jwt"].(string); ok {
		o.Jwt = &Jwt
	}
    
	if EventStreamUri, ok := CreatewebchatconversationresponseMap["eventStreamUri"].(string); ok {
		o.EventStreamUri = &EventStreamUri
	}
    
	if Member, ok := CreatewebchatconversationresponseMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatconversationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatmessage
type Webchatmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Conversation - The identifier of the conversation
	Conversation *Webchatconversation `json:"conversation,omitempty"`

	// Sender - The member who sent the message
	Sender *Webchatmemberinfo `json:"sender,omitempty"`

	// Body - The message body.
	Body *string `json:"body,omitempty"`

	// BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
	BodyType *string `json:"bodyType,omitempty"`

	// Timestamp - The timestamp of the message, in ISO-8601 format
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webchatmessage) SetField(field string, fieldValue interface{}) {
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

func (o Webchatmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
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
	type Alias Webchatmessage
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Webchatconversation `json:"conversation,omitempty"`
		
		Sender *Webchatmemberinfo `json:"sender,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		BodyType *string `json:"bodyType,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		Sender: o.Sender,
		
		Body: o.Body,
		
		BodyType: o.BodyType,
		
		Timestamp: Timestamp,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Webchatmessage) UnmarshalJSON(b []byte) error {
	var WebchatmessageMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebchatmessageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := WebchatmessageMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Sender, ok := WebchatmessageMap["sender"].(map[string]interface{}); ok {
		SenderString, _ := json.Marshal(Sender)
		json.Unmarshal(SenderString, &o.Sender)
	}
	
	if Body, ok := WebchatmessageMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if BodyType, ok := WebchatmessageMap["bodyType"].(string); ok {
		o.BodyType = &BodyType
	}
    
	if timestampString, ok := WebchatmessageMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if SelfUri, ok := WebchatmessageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

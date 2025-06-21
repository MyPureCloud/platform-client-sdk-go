package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Internalmessagedata
type Internalmessagedata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Conversation - The conversation of this message.
	Conversation *Addressableentityref `json:"conversation,omitempty"`

	// CommunicationId - The id of the communication of this message.
	CommunicationId *string `json:"communicationId,omitempty"`

	// Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Sender - The sender of the text message.
	Sender *Userreference `json:"sender,omitempty"`

	// Recipient - The recipient of the text message.
	Recipient *Userreference `json:"recipient,omitempty"`

	// NormalizedMessage - The message into normalized format
	NormalizedMessage *Conversationnormalizedmessage `json:"normalizedMessage,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Internalmessagedata) SetField(field string, fieldValue interface{}) {
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

func (o Internalmessagedata) MarshalJSON() ([]byte, error) {
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
	type Alias Internalmessagedata
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Sender *Userreference `json:"sender,omitempty"`
		
		Recipient *Userreference `json:"recipient,omitempty"`
		
		NormalizedMessage *Conversationnormalizedmessage `json:"normalizedMessage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		CommunicationId: o.CommunicationId,
		
		Timestamp: Timestamp,
		
		Sender: o.Sender,
		
		Recipient: o.Recipient,
		
		NormalizedMessage: o.NormalizedMessage,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Internalmessagedata) UnmarshalJSON(b []byte) error {
	var InternalmessagedataMap map[string]interface{}
	err := json.Unmarshal(b, &InternalmessagedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InternalmessagedataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := InternalmessagedataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := InternalmessagedataMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if CommunicationId, ok := InternalmessagedataMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if timestampString, ok := InternalmessagedataMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if Sender, ok := InternalmessagedataMap["sender"].(map[string]interface{}); ok {
		SenderString, _ := json.Marshal(Sender)
		json.Unmarshal(SenderString, &o.Sender)
	}
	
	if Recipient, ok := InternalmessagedataMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if NormalizedMessage, ok := InternalmessagedataMap["normalizedMessage"].(map[string]interface{}); ok {
		NormalizedMessageString, _ := json.Marshal(NormalizedMessage)
		json.Unmarshal(NormalizedMessageString, &o.NormalizedMessage)
	}
	
	if SelfUri, ok := InternalmessagedataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Internalmessagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

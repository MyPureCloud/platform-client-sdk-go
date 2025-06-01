package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Socialmediamessagedata
type Socialmediamessagedata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ProviderMessageId - The unique identifier of the message from provider
	ProviderMessageId *string `json:"providerMessageId,omitempty"`

	// Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// FromAddress - The sender of the text message.
	FromAddress *string `json:"fromAddress,omitempty"`

	// ToAddress - The recipient of the text message.
	ToAddress *string `json:"toAddress,omitempty"`

	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`

	// MessengerType - Type of text messenger.
	MessengerType *string `json:"messengerType,omitempty"`

	// Status - The status of the message.
	Status *string `json:"status,omitempty"`

	// NormalizedMessage - The message into normalized format
	NormalizedMessage *Conversationnormalizedmessage `json:"normalizedMessage,omitempty"`

	// NormalizedReceipts - The delivery event associated with this message in normalized format, if the message direction was outbound
	NormalizedReceipts *[]Conversationnormalizedmessage `json:"normalizedReceipts,omitempty"`

	// CreatedBy - User who sent this message.
	CreatedBy *User `json:"createdBy,omitempty"`

	// ConversationId - The id of the conversation of this message.
	ConversationId *string `json:"conversationId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Socialmediamessagedata) SetField(field string, fieldValue interface{}) {
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

func (o Socialmediamessagedata) MarshalJSON() ([]byte, error) {
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
	type Alias Socialmediamessagedata
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ProviderMessageId *string `json:"providerMessageId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		NormalizedMessage *Conversationnormalizedmessage `json:"normalizedMessage,omitempty"`
		
		NormalizedReceipts *[]Conversationnormalizedmessage `json:"normalizedReceipts,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ProviderMessageId: o.ProviderMessageId,
		
		Timestamp: Timestamp,
		
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		Direction: o.Direction,
		
		MessengerType: o.MessengerType,
		
		Status: o.Status,
		
		NormalizedMessage: o.NormalizedMessage,
		
		NormalizedReceipts: o.NormalizedReceipts,
		
		CreatedBy: o.CreatedBy,
		
		ConversationId: o.ConversationId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Socialmediamessagedata) UnmarshalJSON(b []byte) error {
	var SocialmediamessagedataMap map[string]interface{}
	err := json.Unmarshal(b, &SocialmediamessagedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SocialmediamessagedataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SocialmediamessagedataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ProviderMessageId, ok := SocialmediamessagedataMap["providerMessageId"].(string); ok {
		o.ProviderMessageId = &ProviderMessageId
	}
    
	if timestampString, ok := SocialmediamessagedataMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if FromAddress, ok := SocialmediamessagedataMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if ToAddress, ok := SocialmediamessagedataMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if Direction, ok := SocialmediamessagedataMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if MessengerType, ok := SocialmediamessagedataMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
    
	if Status, ok := SocialmediamessagedataMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if NormalizedMessage, ok := SocialmediamessagedataMap["normalizedMessage"].(map[string]interface{}); ok {
		NormalizedMessageString, _ := json.Marshal(NormalizedMessage)
		json.Unmarshal(NormalizedMessageString, &o.NormalizedMessage)
	}
	
	if NormalizedReceipts, ok := SocialmediamessagedataMap["normalizedReceipts"].([]interface{}); ok {
		NormalizedReceiptsString, _ := json.Marshal(NormalizedReceipts)
		json.Unmarshal(NormalizedReceiptsString, &o.NormalizedReceipts)
	}
	
	if CreatedBy, ok := SocialmediamessagedataMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ConversationId, ok := SocialmediamessagedataMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SelfUri, ok := SocialmediamessagedataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Socialmediamessagedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

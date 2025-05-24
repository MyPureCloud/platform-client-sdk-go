package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentpush - A Push object
type Conversationcontentpush struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DeviceType - The device type used to send the push notification
	DeviceType *string `json:"deviceType,omitempty"`

	// DeviceTokenId - Unique Id of the device token
	DeviceTokenId *string `json:"deviceTokenId,omitempty"`

	// DeviceToken - device token from the notification provider
	DeviceToken *string `json:"deviceToken,omitempty"`

	// FailedMessages - MessageIds failed to be sent which trigger the push event
	FailedMessages *[]Conversationpushfailedmessagereferences `json:"failedMessages,omitempty"`

	// NotificationMessage - Title and body localized according to deployment
	NotificationMessage *Conversationpushnotificationmessagelabel `json:"notificationMessage,omitempty"`

	// PushProviderIntegration - Push provider integrations details configured on the deployment
	PushProviderIntegration *Conversationpushproviderintegration `json:"pushProviderIntegration,omitempty"`

	// Expiration - The time to live of the pushed message
	Expiration *int `json:"expiration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentpush) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentpush) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentpush
	
	return json.Marshal(&struct { 
		DeviceType *string `json:"deviceType,omitempty"`
		
		DeviceTokenId *string `json:"deviceTokenId,omitempty"`
		
		DeviceToken *string `json:"deviceToken,omitempty"`
		
		FailedMessages *[]Conversationpushfailedmessagereferences `json:"failedMessages,omitempty"`
		
		NotificationMessage *Conversationpushnotificationmessagelabel `json:"notificationMessage,omitempty"`
		
		PushProviderIntegration *Conversationpushproviderintegration `json:"pushProviderIntegration,omitempty"`
		
		Expiration *int `json:"expiration,omitempty"`
		Alias
	}{ 
		DeviceType: o.DeviceType,
		
		DeviceTokenId: o.DeviceTokenId,
		
		DeviceToken: o.DeviceToken,
		
		FailedMessages: o.FailedMessages,
		
		NotificationMessage: o.NotificationMessage,
		
		PushProviderIntegration: o.PushProviderIntegration,
		
		Expiration: o.Expiration,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentpush) UnmarshalJSON(b []byte) error {
	var ConversationcontentpushMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentpushMap)
	if err != nil {
		return err
	}
	
	if DeviceType, ok := ConversationcontentpushMap["deviceType"].(string); ok {
		o.DeviceType = &DeviceType
	}
    
	if DeviceTokenId, ok := ConversationcontentpushMap["deviceTokenId"].(string); ok {
		o.DeviceTokenId = &DeviceTokenId
	}
    
	if DeviceToken, ok := ConversationcontentpushMap["deviceToken"].(string); ok {
		o.DeviceToken = &DeviceToken
	}
    
	if FailedMessages, ok := ConversationcontentpushMap["failedMessages"].([]interface{}); ok {
		FailedMessagesString, _ := json.Marshal(FailedMessages)
		json.Unmarshal(FailedMessagesString, &o.FailedMessages)
	}
	
	if NotificationMessage, ok := ConversationcontentpushMap["notificationMessage"].(map[string]interface{}); ok {
		NotificationMessageString, _ := json.Marshal(NotificationMessage)
		json.Unmarshal(NotificationMessageString, &o.NotificationMessage)
	}
	
	if PushProviderIntegration, ok := ConversationcontentpushMap["pushProviderIntegration"].(map[string]interface{}); ok {
		PushProviderIntegrationString, _ := json.Marshal(PushProviderIntegration)
		json.Unmarshal(PushProviderIntegrationString, &o.PushProviderIntegration)
	}
	
	if Expiration, ok := ConversationcontentpushMap["expiration"].(float64); ok {
		ExpirationInt := int(Expiration)
		o.Expiration = &ExpirationInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentpush) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

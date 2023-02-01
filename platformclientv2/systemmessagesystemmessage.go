package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Systemmessagesystemmessage
type Systemmessagesystemmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ChannelId
	ChannelId *string `json:"channelId,omitempty"`

	// SystemTopicType
	SystemTopicType *string `json:"systemTopicType,omitempty"`

	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`

	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// OauthClientId
	OauthClientId *string `json:"oauthClientId,omitempty"`

	// Reason
	Reason *string `json:"reason,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// Data
	Data *map[string]interface{} `json:"data,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Systemmessagesystemmessage) SetField(field string, fieldValue interface{}) {
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

func (o Systemmessagesystemmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Systemmessagesystemmessage
	
	return json.Marshal(&struct { 
		ChannelId *string `json:"channelId,omitempty"`
		
		SystemTopicType *string `json:"systemTopicType,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		OauthClientId *string `json:"oauthClientId,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Data *map[string]interface{} `json:"data,omitempty"`
		Alias
	}{ 
		ChannelId: o.ChannelId,
		
		SystemTopicType: o.SystemTopicType,
		
		CorrelationId: o.CorrelationId,
		
		OrganizationId: o.OrganizationId,
		
		UserId: o.UserId,
		
		OauthClientId: o.OauthClientId,
		
		Reason: o.Reason,
		
		Message: o.Message,
		
		Data: o.Data,
		Alias:    (Alias)(o),
	})
}

func (o *Systemmessagesystemmessage) UnmarshalJSON(b []byte) error {
	var SystemmessagesystemmessageMap map[string]interface{}
	err := json.Unmarshal(b, &SystemmessagesystemmessageMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := SystemmessagesystemmessageMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
    
	if SystemTopicType, ok := SystemmessagesystemmessageMap["systemTopicType"].(string); ok {
		o.SystemTopicType = &SystemTopicType
	}
    
	if CorrelationId, ok := SystemmessagesystemmessageMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if OrganizationId, ok := SystemmessagesystemmessageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if UserId, ok := SystemmessagesystemmessageMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if OauthClientId, ok := SystemmessagesystemmessageMap["oauthClientId"].(string); ok {
		o.OauthClientId = &OauthClientId
	}
    
	if Reason, ok := SystemmessagesystemmessageMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    
	if Message, ok := SystemmessagesystemmessageMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if Data, ok := SystemmessagesystemmessageMap["data"].(map[string]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Systemmessagesystemmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

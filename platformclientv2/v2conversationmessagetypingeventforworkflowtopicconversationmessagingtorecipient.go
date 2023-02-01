package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient
type V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Nickname
	Nickname *string `json:"nickname,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// IdType
	IdType *string `json:"idType,omitempty"`

	// Image
	Image *string `json:"image,omitempty"`

	// FirstName
	FirstName *string `json:"firstName,omitempty"`

	// LastName
	LastName *string `json:"lastName,omitempty"`

	// Email
	Email *string `json:"email,omitempty"`

	// AdditionalIds
	AdditionalIds *[]V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient
	
	return json.Marshal(&struct { 
		Nickname *string `json:"nickname,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		AdditionalIds *[]V2conversationmessagetypingeventforworkflowtopicconversationrecipientadditionalidentifier `json:"additionalIds,omitempty"`
		Alias
	}{ 
		Nickname: o.Nickname,
		
		Id: o.Id,
		
		IdType: o.IdType,
		
		Image: o.Image,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		
		AdditionalIds: o.AdditionalIds,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap)
	if err != nil {
		return err
	}
	
	if Nickname, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
    
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IdType, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
    
	if Image, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if FirstName, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Email, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if AdditionalIds, ok := V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationmessagingtorecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventuserconversationsummary
type Userconversationseventuserconversationsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId
	UserId *string `json:"userId,omitempty"`

	// Call
	Call *Userconversationseventmediasummary `json:"call,omitempty"`

	// Callback
	Callback *Userconversationseventmediasummary `json:"callback,omitempty"`

	// Email
	Email *Userconversationseventmediasummary `json:"email,omitempty"`

	// Message
	Message *Userconversationseventmediasummary `json:"message,omitempty"`

	// Chat
	Chat *Userconversationseventmediasummary `json:"chat,omitempty"`

	// SocialExpression
	SocialExpression *Userconversationseventmediasummary `json:"socialExpression,omitempty"`

	// Video
	Video *Userconversationseventmediasummary `json:"video,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userconversationseventuserconversationsummary) SetField(field string, fieldValue interface{}) {
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

func (o Userconversationseventuserconversationsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Userconversationseventuserconversationsummary
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		Call *Userconversationseventmediasummary `json:"call,omitempty"`
		
		Callback *Userconversationseventmediasummary `json:"callback,omitempty"`
		
		Email *Userconversationseventmediasummary `json:"email,omitempty"`
		
		Message *Userconversationseventmediasummary `json:"message,omitempty"`
		
		Chat *Userconversationseventmediasummary `json:"chat,omitempty"`
		
		SocialExpression *Userconversationseventmediasummary `json:"socialExpression,omitempty"`
		
		Video *Userconversationseventmediasummary `json:"video,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		Call: o.Call,
		
		Callback: o.Callback,
		
		Email: o.Email,
		
		Message: o.Message,
		
		Chat: o.Chat,
		
		SocialExpression: o.SocialExpression,
		
		Video: o.Video,
		Alias:    (Alias)(o),
	})
}

func (o *Userconversationseventuserconversationsummary) UnmarshalJSON(b []byte) error {
	var UserconversationseventuserconversationsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &UserconversationseventuserconversationsummaryMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := UserconversationseventuserconversationsummaryMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Call, ok := UserconversationseventuserconversationsummaryMap["call"].(map[string]interface{}); ok {
		CallString, _ := json.Marshal(Call)
		json.Unmarshal(CallString, &o.Call)
	}
	
	if Callback, ok := UserconversationseventuserconversationsummaryMap["callback"].(map[string]interface{}); ok {
		CallbackString, _ := json.Marshal(Callback)
		json.Unmarshal(CallbackString, &o.Callback)
	}
	
	if Email, ok := UserconversationseventuserconversationsummaryMap["email"].(map[string]interface{}); ok {
		EmailString, _ := json.Marshal(Email)
		json.Unmarshal(EmailString, &o.Email)
	}
	
	if Message, ok := UserconversationseventuserconversationsummaryMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if Chat, ok := UserconversationseventuserconversationsummaryMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if SocialExpression, ok := UserconversationseventuserconversationsummaryMap["socialExpression"].(map[string]interface{}); ok {
		SocialExpressionString, _ := json.Marshal(SocialExpression)
		json.Unmarshal(SocialExpressionString, &o.SocialExpression)
	}
	
	if Video, ok := UserconversationseventuserconversationsummaryMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userconversationseventuserconversationsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

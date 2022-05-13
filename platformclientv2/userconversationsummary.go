package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationsummary
type Userconversationsummary struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`


	// Call
	Call *Mediasummary `json:"call,omitempty"`


	// Callback
	Callback *Mediasummary `json:"callback,omitempty"`


	// Email
	Email *Mediasummary `json:"email,omitempty"`


	// Message
	Message *Mediasummary `json:"message,omitempty"`


	// Chat
	Chat *Mediasummary `json:"chat,omitempty"`


	// SocialExpression
	SocialExpression *Mediasummary `json:"socialExpression,omitempty"`


	// Video
	Video *Mediasummary `json:"video,omitempty"`

}

func (o *Userconversationsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userconversationsummary
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		Call *Mediasummary `json:"call,omitempty"`
		
		Callback *Mediasummary `json:"callback,omitempty"`
		
		Email *Mediasummary `json:"email,omitempty"`
		
		Message *Mediasummary `json:"message,omitempty"`
		
		Chat *Mediasummary `json:"chat,omitempty"`
		
		SocialExpression *Mediasummary `json:"socialExpression,omitempty"`
		
		Video *Mediasummary `json:"video,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		Call: o.Call,
		
		Callback: o.Callback,
		
		Email: o.Email,
		
		Message: o.Message,
		
		Chat: o.Chat,
		
		SocialExpression: o.SocialExpression,
		
		Video: o.Video,
		Alias:    (*Alias)(o),
	})
}

func (o *Userconversationsummary) UnmarshalJSON(b []byte) error {
	var UserconversationsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &UserconversationsummaryMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := UserconversationsummaryMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Call, ok := UserconversationsummaryMap["call"].(map[string]interface{}); ok {
		CallString, _ := json.Marshal(Call)
		json.Unmarshal(CallString, &o.Call)
	}
	
	if Callback, ok := UserconversationsummaryMap["callback"].(map[string]interface{}); ok {
		CallbackString, _ := json.Marshal(Callback)
		json.Unmarshal(CallbackString, &o.Callback)
	}
	
	if Email, ok := UserconversationsummaryMap["email"].(map[string]interface{}); ok {
		EmailString, _ := json.Marshal(Email)
		json.Unmarshal(EmailString, &o.Email)
	}
	
	if Message, ok := UserconversationsummaryMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if Chat, ok := UserconversationsummaryMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if SocialExpression, ok := UserconversationsummaryMap["socialExpression"].(map[string]interface{}); ok {
		SocialExpressionString, _ := json.Marshal(SocialExpression)
		json.Unmarshal(SocialExpressionString, &o.SocialExpression)
	}
	
	if Video, ok := UserconversationsummaryMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userconversationsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

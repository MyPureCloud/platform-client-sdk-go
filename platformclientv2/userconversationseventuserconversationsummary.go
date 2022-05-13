package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventuserconversationsummary
type Userconversationseventuserconversationsummary struct { 
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

func (o *Userconversationseventuserconversationsummary) MarshalJSON() ([]byte, error) {
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

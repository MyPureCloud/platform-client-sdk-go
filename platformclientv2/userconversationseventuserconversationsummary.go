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

func (u *Userconversationseventuserconversationsummary) MarshalJSON() ([]byte, error) {
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
		UserId: u.UserId,
		
		Call: u.Call,
		
		Callback: u.Callback,
		
		Email: u.Email,
		
		Message: u.Message,
		
		Chat: u.Chat,
		
		SocialExpression: u.SocialExpression,
		
		Video: u.Video,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userconversationseventuserconversationsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

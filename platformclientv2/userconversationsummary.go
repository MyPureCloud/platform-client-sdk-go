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

func (u *Userconversationsummary) MarshalJSON() ([]byte, error) {
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
func (o *Userconversationsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Userconversationseventuserconversationsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

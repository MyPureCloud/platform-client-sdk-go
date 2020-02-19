package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Userconversationsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Chat
type Chat struct { 
	// JabberId
	JabberId *string `json:"jabberId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chat) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

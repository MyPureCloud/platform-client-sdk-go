package platformclientv2
import (
	"encoding/json"
)

// Chatbadgetopicbadgeentity
type Chatbadgetopicbadgeentity struct { 
	// JabberId
	JabberId *string `json:"jabberId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatbadgetopicbadgeentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

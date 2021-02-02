package platformclientv2
import (
	"encoding/json"
)

// Analyticsresolution
type Analyticsresolution struct { 
	// QueueId - The ID of the last queue on which the conversation was handled.
	QueueId *string `json:"queueId,omitempty"`


	// UserId - The ID of the last user who handled the conversation.
	UserId *string `json:"userId,omitempty"`


	// GetnNextContactAvoided - The number of interactions for which next contact was avoided.
	GetnNextContactAvoided *int `json:"getnNextContactAvoided,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsresolution) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

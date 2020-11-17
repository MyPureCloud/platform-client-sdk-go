package platformclientv2
import (
	"encoding/json"
)

// Contentreaction - User reaction to public message
type Contentreaction struct { 
	// ReactionType - Type of reaction
	ReactionType *string `json:"reactionType,omitempty"`


	// Count - Number of users that reacted this way to this public message
	Count *int32 `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentreaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

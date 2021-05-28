package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentreaction - User reaction to public message.
type Contentreaction struct { 
	// ReactionType - Type of reaction.
	ReactionType *string `json:"reactionType,omitempty"`


	// Count - Number of users that reacted this way to the message.
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentreaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

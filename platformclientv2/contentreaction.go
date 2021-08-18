package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Contentreaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentreaction

	

	return json.Marshal(&struct { 
		ReactionType *string `json:"reactionType,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		ReactionType: u.ReactionType,
		
		Count: u.Count,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentreaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

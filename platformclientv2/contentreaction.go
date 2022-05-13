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

func (o *Contentreaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentreaction
	
	return json.Marshal(&struct { 
		ReactionType *string `json:"reactionType,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		ReactionType: o.ReactionType,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentreaction) UnmarshalJSON(b []byte) error {
	var ContentreactionMap map[string]interface{}
	err := json.Unmarshal(b, &ContentreactionMap)
	if err != nil {
		return err
	}
	
	if ReactionType, ok := ContentreactionMap["reactionType"].(string); ok {
		o.ReactionType = &ReactionType
	}
    
	if Count, ok := ContentreactionMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentreaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

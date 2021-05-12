package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Reaction
type Reaction struct { 
	// Data - Parameter for this reaction. For transfer_flow, this would be the outbound flow id.
	Data *string `json:"data,omitempty"`


	// Name - Name of the parameter for this reaction. For transfer_flow, this would be the outbound flow name.
	Name *string `json:"name,omitempty"`


	// ReactionType - The reaction to take for a given call analysis result.
	ReactionType *string `json:"reactionType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

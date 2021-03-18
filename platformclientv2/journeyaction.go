package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyaction
type Journeyaction struct { 
	// Id - The ID of an action from the Journey System (an action is spawned from an actionMap)
	Id *string `json:"id,omitempty"`


	// ActionMap - Details about the action map from the Journey System which triggered this action
	ActionMap *Journeyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

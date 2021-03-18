package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerresponsesetconfigchangereaction
type Dialerresponsesetconfigchangereaction struct { 
	// Data
	Data *string `json:"data,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ReactionType
	ReactionType *string `json:"reactionType,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerresponsesetconfigchangereaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Webchatroutingtarget
type Webchatroutingtarget struct { 
	// TargetType - The target type of the routing target, such as 'QUEUE'.
	TargetType *string `json:"targetType,omitempty"`


	// TargetAddress - The target of the route, in the format appropriate given the 'targetType'.
	TargetAddress *string `json:"targetAddress,omitempty"`


	// Skills - The list of skill names to use for routing.
	Skills *[]string `json:"skills,omitempty"`


	// Language - The language name to use for routing.
	Language *string `json:"language,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int64 `json:"priority,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatroutingtarget) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

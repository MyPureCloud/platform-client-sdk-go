package platformclientv2
import (
	"encoding/json"
)

// Schedulegenerationmessage
type Schedulegenerationmessage struct { 
	// VarType - The type of the message
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments describing the message
	Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

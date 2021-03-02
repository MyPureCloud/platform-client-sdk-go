package platformclientv2
import (
	"encoding/json"
)

// Schedulermessageargument
type Schedulermessageargument struct { 
	// VarType - The type of this message parameter
	VarType *string `json:"type,omitempty"`


	// Value - The value of this message parameter
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulermessageargument) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

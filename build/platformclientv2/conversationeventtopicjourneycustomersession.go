package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicjourneycustomersession
type Conversationeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

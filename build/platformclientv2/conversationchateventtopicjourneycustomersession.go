package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicjourneycustomersession
type Conversationchateventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

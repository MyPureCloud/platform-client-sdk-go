package platformclientv2
import (
	"encoding/json"
)

// Conversationvideoeventtopicjourneycustomersession
type Conversationvideoeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

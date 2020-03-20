package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicjourneycustomersession
type Queueconversationeventtopicjourneycustomersession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicjourneycustomersession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

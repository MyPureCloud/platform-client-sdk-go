package platformclientv2
import (
	"encoding/json"
)

// Faxtopicworkspacedata
type Faxtopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxtopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

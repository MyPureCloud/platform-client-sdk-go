package platformclientv2
import (
	"encoding/json"
)

// Adhocrecordingtopicworkspacedata
type Adhocrecordingtopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

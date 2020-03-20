package platformclientv2
import (
	"encoding/json"
)

// Stateventqueuetopicstatsnotification
type Stateventqueuetopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventqueuetopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Stateventusertopicstatsnotification
type Stateventusertopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventusertopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventusertopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

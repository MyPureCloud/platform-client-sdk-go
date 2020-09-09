package platformclientv2
import (
	"encoding/json"
)

// Stateventflowoutcometopicstatsnotification
type Stateventflowoutcometopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventflowoutcometopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventflowoutcometopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

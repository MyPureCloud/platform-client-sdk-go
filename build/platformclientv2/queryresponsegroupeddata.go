package platformclientv2
import (
	"encoding/json"
)

// Queryresponsegroupeddata
type Queryresponsegroupeddata struct { 
	// Group - The group values for this data
	Group *map[string]string `json:"group,omitempty"`


	// Data - The metrics in this group
	Data *[]Queryresponsedata `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryresponsegroupeddata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

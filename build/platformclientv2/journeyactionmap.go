package platformclientv2
import (
	"encoding/json"
)

// Journeyactionmap
type Journeyactionmap struct { 
	// Id - The ID of the actionMap in the Journey System which triggered this action
	Id *string `json:"id,omitempty"`


	// Version - The version number of the actionMap in the Journey System at the time this action was triggered
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeyactionmap) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

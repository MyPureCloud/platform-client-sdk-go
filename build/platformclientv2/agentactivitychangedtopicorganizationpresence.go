package platformclientv2
import (
	"encoding/json"
)

// Agentactivitychangedtopicorganizationpresence
type Agentactivitychangedtopicorganizationpresence struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicorganizationpresence) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

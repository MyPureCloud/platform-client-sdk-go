package platformclientv2
import (
	"encoding/json"
)

// Wemcoachingappointmenttopiccoachingappointmentexternallink
type Wemcoachingappointmenttopiccoachingappointmentexternallink struct { 
	// ExternalLink
	ExternalLink *string `json:"externalLink,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

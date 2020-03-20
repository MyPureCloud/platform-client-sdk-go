package platformclientv2
import (
	"encoding/json"
)

// Wemcoachingappointmenttopiccoachingappointmentdocument
type Wemcoachingappointmenttopiccoachingappointmentdocument struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentdocument) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

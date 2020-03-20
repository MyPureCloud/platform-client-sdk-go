package platformclientv2
import (
	"encoding/json"
)

// Wemcoachingappointmenttopiccoachingappointmentconversation
type Wemcoachingappointmenttopiccoachingappointmentconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

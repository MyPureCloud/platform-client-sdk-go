package platformclientv2
import (
	"encoding/json"
)

// Consulttransferresponse
type Consulttransferresponse struct { 
	// DestinationParticipantId - Participant ID to whom the call is being transferred.
	DestinationParticipantId *string `json:"destinationParticipantId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Consulttransferresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

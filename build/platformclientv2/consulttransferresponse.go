package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransferresponse
type Consulttransferresponse struct { 
	// DestinationParticipantId - Participant ID to whom the call is being transferred.
	DestinationParticipantId *string `json:"destinationParticipantId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Consulttransferresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransferresponse
type Consulttransferresponse struct { 
	// DestinationParticipantId - Participant ID to whom the call is being transferred.
	DestinationParticipantId *string `json:"destinationParticipantId,omitempty"`

}

func (u *Consulttransferresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransferresponse

	

	return json.Marshal(&struct { 
		DestinationParticipantId *string `json:"destinationParticipantId,omitempty"`
		*Alias
	}{ 
		DestinationParticipantId: u.DestinationParticipantId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Consulttransferresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

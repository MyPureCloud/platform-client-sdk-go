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

func (o *Consulttransferresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransferresponse
	
	return json.Marshal(&struct { 
		DestinationParticipantId *string `json:"destinationParticipantId,omitempty"`
		*Alias
	}{ 
		DestinationParticipantId: o.DestinationParticipantId,
		Alias:    (*Alias)(o),
	})
}

func (o *Consulttransferresponse) UnmarshalJSON(b []byte) error {
	var ConsulttransferresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ConsulttransferresponseMap)
	if err != nil {
		return err
	}
	
	if DestinationParticipantId, ok := ConsulttransferresponseMap["destinationParticipantId"].(string); ok {
		o.DestinationParticipantId = &DestinationParticipantId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Consulttransferresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

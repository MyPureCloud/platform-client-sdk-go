package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Maxparticipants
type Maxparticipants struct { 
	// MaxParticipants - The maximum number of participants that are allowed on a conversation.
	MaxParticipants *int `json:"maxParticipants,omitempty"`

}

func (o *Maxparticipants) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Maxparticipants
	
	return json.Marshal(&struct { 
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		*Alias
	}{ 
		MaxParticipants: o.MaxParticipants,
		Alias:    (*Alias)(o),
	})
}

func (o *Maxparticipants) UnmarshalJSON(b []byte) error {
	var MaxparticipantsMap map[string]interface{}
	err := json.Unmarshal(b, &MaxparticipantsMap)
	if err != nil {
		return err
	}
	
	if MaxParticipants, ok := MaxparticipantsMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Maxparticipants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

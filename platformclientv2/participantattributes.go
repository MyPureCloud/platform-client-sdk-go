package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantattributes
type Participantattributes struct { 
	// Attributes - The map of attribute keys to values.
	Attributes *map[string]string `json:"attributes,omitempty"`

}

func (o *Participantattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participantattributes
	
	return json.Marshal(&struct { 
		Attributes *map[string]string `json:"attributes,omitempty"`
		*Alias
	}{ 
		Attributes: o.Attributes,
		Alias:    (*Alias)(o),
	})
}

func (o *Participantattributes) UnmarshalJSON(b []byte) error {
	var ParticipantattributesMap map[string]interface{}
	err := json.Unmarshal(b, &ParticipantattributesMap)
	if err != nil {
		return err
	}
	
	if Attributes, ok := ParticipantattributesMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Participantattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

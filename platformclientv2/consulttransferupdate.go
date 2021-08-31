package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransferupdate
type Consulttransferupdate struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking.
	SpeakTo *string `json:"speakTo,omitempty"`

}

func (o *Consulttransferupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransferupdate
	
	return json.Marshal(&struct { 
		SpeakTo *string `json:"speakTo,omitempty"`
		*Alias
	}{ 
		SpeakTo: o.SpeakTo,
		Alias:    (*Alias)(o),
	})
}

func (o *Consulttransferupdate) UnmarshalJSON(b []byte) error {
	var ConsulttransferupdateMap map[string]interface{}
	err := json.Unmarshal(b, &ConsulttransferupdateMap)
	if err != nil {
		return err
	}
	
	if SpeakTo, ok := ConsulttransferupdateMap["speakTo"].(string); ok {
		o.SpeakTo = &SpeakTo
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Consulttransferupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

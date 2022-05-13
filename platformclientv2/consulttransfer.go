package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Consulttransfer
type Consulttransfer struct { 
	// SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
	SpeakTo *string `json:"speakTo,omitempty"`


	// Destination - Destination phone number and name.
	Destination *Destination `json:"destination,omitempty"`

}

func (o *Consulttransfer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Consulttransfer
	
	return json.Marshal(&struct { 
		SpeakTo *string `json:"speakTo,omitempty"`
		
		Destination *Destination `json:"destination,omitempty"`
		*Alias
	}{ 
		SpeakTo: o.SpeakTo,
		
		Destination: o.Destination,
		Alias:    (*Alias)(o),
	})
}

func (o *Consulttransfer) UnmarshalJSON(b []byte) error {
	var ConsulttransferMap map[string]interface{}
	err := json.Unmarshal(b, &ConsulttransferMap)
	if err != nil {
		return err
	}
	
	if SpeakTo, ok := ConsulttransferMap["speakTo"].(string); ok {
		o.SpeakTo = &SpeakTo
	}
    
	if Destination, ok := ConsulttransferMap["destination"].(map[string]interface{}); ok {
		DestinationString, _ := json.Marshal(Destination)
		json.Unmarshal(DestinationString, &o.Destination)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Consulttransfer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

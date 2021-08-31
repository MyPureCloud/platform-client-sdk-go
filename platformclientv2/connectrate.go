package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Connectrate
type Connectrate struct { 
	// Attempts - Number of call attempts made
	Attempts *int `json:"attempts,omitempty"`


	// Connects - Number of calls with a live voice detected
	Connects *int `json:"connects,omitempty"`


	// ConnectRatio - Ratio of connects to attempts
	ConnectRatio *float64 `json:"connectRatio,omitempty"`

}

func (o *Connectrate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Connectrate
	
	return json.Marshal(&struct { 
		Attempts *int `json:"attempts,omitempty"`
		
		Connects *int `json:"connects,omitempty"`
		
		ConnectRatio *float64 `json:"connectRatio,omitempty"`
		*Alias
	}{ 
		Attempts: o.Attempts,
		
		Connects: o.Connects,
		
		ConnectRatio: o.ConnectRatio,
		Alias:    (*Alias)(o),
	})
}

func (o *Connectrate) UnmarshalJSON(b []byte) error {
	var ConnectrateMap map[string]interface{}
	err := json.Unmarshal(b, &ConnectrateMap)
	if err != nil {
		return err
	}
	
	if Attempts, ok := ConnectrateMap["attempts"].(float64); ok {
		AttemptsInt := int(Attempts)
		o.Attempts = &AttemptsInt
	}
	
	if Connects, ok := ConnectrateMap["connects"].(float64); ok {
		ConnectsInt := int(Connects)
		o.Connects = &ConnectsInt
	}
	
	if ConnectRatio, ok := ConnectrateMap["connectRatio"].(float64); ok {
		o.ConnectRatio = &ConnectRatio
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Connectrate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

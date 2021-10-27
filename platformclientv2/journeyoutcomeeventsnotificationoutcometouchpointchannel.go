package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcometouchpointchannel
type Journeyoutcomeeventsnotificationoutcometouchpointchannel struct { 
	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationoutcometouchpointchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationoutcometouchpointchannel
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcometouchpointchannel) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcometouchpointchannelMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcometouchpointchannelMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneyoutcomeeventsnotificationoutcometouchpointchannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcometouchpointchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

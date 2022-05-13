package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Freeseatingconfiguration
type Freeseatingconfiguration struct { 
	// FreeSeatingState - The FreeSeatingState for FreeSeatingConfiguration. Can be ON, OFF, or PARTIAL. ON meaning disassociate the user after the ttl expires, OFF meaning never disassociate the user, and PARTIAL meaning only disassociate when a user explicitly clicks log out.
	FreeSeatingState *string `json:"freeSeatingState,omitempty"`


	// TtlMinutes - The amount of time in minutes until an offline user is disassociated from their station
	TtlMinutes *int `json:"ttlMinutes,omitempty"`

}

func (o *Freeseatingconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Freeseatingconfiguration
	
	return json.Marshal(&struct { 
		FreeSeatingState *string `json:"freeSeatingState,omitempty"`
		
		TtlMinutes *int `json:"ttlMinutes,omitempty"`
		*Alias
	}{ 
		FreeSeatingState: o.FreeSeatingState,
		
		TtlMinutes: o.TtlMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Freeseatingconfiguration) UnmarshalJSON(b []byte) error {
	var FreeseatingconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &FreeseatingconfigurationMap)
	if err != nil {
		return err
	}
	
	if FreeSeatingState, ok := FreeseatingconfigurationMap["freeSeatingState"].(string); ok {
		o.FreeSeatingState = &FreeSeatingState
	}
    
	if TtlMinutes, ok := FreeseatingconfigurationMap["ttlMinutes"].(float64); ok {
		TtlMinutesInt := int(TtlMinutes)
		o.TtlMinutes = &TtlMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Freeseatingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

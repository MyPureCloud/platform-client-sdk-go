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

func (u *Freeseatingconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Freeseatingconfiguration

	

	return json.Marshal(&struct { 
		FreeSeatingState *string `json:"freeSeatingState,omitempty"`
		
		TtlMinutes *int `json:"ttlMinutes,omitempty"`
		*Alias
	}{ 
		FreeSeatingState: u.FreeSeatingState,
		
		TtlMinutes: u.TtlMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Freeseatingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

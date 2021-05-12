package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Freeseatingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

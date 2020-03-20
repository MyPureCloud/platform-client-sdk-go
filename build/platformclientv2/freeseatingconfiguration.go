package platformclientv2
import (
	"encoding/json"
)

// Freeseatingconfiguration
type Freeseatingconfiguration struct { 
	// FreeSeatingState - The FreeSeatingState for FreeSeatingConfiguration. Can be ON, OFF, or PARTIAL. ON meaning disassociate the user after the ttl expires, OFF meaning never disassociate the user, and PARTIAL meaning only disassociate when a user explicitly clicks log out.
	FreeSeatingState *string `json:"freeSeatingState,omitempty"`


	// TtlMinutes - The amount of time in minutes until an offline user is disassociated from their station
	TtlMinutes *int32 `json:"ttlMinutes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Freeseatingconfiguration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

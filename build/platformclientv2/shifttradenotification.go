package platformclientv2
import (
	"time"
	"encoding/json"
)

// Shifttradenotification
type Shifttradenotification struct { 
	// WeekDate - The start date of the schedule with which this trade is associated
	WeekDate *string `json:"weekDate,omitempty"`


	// TradeId - The ID of the shift trade
	TradeId *string `json:"tradeId,omitempty"`


	// OneSided - Whether this is a one sided shift trade
	OneSided *bool `json:"oneSided,omitempty"`


	// NewState - The new state of the shift trade, null if there was no change
	NewState *string `json:"newState,omitempty"`


	// InitiatingUser - The user who initiated the shift trade
	InitiatingUser *Userreference `json:"initiatingUser,omitempty"`


	// InitiatingShiftDate - The start date and time of the initiating shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	InitiatingShiftDate *time.Time `json:"initiatingShiftDate,omitempty"`


	// ReceivingUser - The user on the receiving side of this shift trade (null if not matched)
	ReceivingUser *Userreference `json:"receivingUser,omitempty"`


	// ReceivingShiftDate - The start date and time of the receiving shift (null if not matched or if one-sided. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReceivingShiftDate *time.Time `json:"receivingShiftDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

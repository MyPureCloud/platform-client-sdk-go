package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmusernotificationtopicshifttradenotification
type Wfmusernotificationtopicshifttradenotification struct { 
	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// TradeId
	TradeId *string `json:"tradeId,omitempty"`


	// OneSided
	OneSided *bool `json:"oneSided,omitempty"`


	// NewState
	NewState *string `json:"newState,omitempty"`


	// InitiatingUser
	InitiatingUser *Wfmusernotificationtopicuserreference `json:"initiatingUser,omitempty"`


	// InitiatingShiftDate
	InitiatingShiftDate *time.Time `json:"initiatingShiftDate,omitempty"`


	// ReceivingUser
	ReceivingUser *Wfmusernotificationtopicuserreference `json:"receivingUser,omitempty"`


	// ReceivingShiftDate
	ReceivingShiftDate *time.Time `json:"receivingShiftDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicshifttradenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

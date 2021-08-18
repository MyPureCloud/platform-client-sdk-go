package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// InitiatingShiftDate - The start date and time of the initiating shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InitiatingShiftDate *time.Time `json:"initiatingShiftDate,omitempty"`


	// ReceivingUser - The user on the receiving side of this shift trade (null if not matched)
	ReceivingUser *Userreference `json:"receivingUser,omitempty"`


	// ReceivingShiftDate - The start date and time of the receiving shift (null if not matched or if one-sided. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReceivingShiftDate *time.Time `json:"receivingShiftDate,omitempty"`

}

func (u *Shifttradenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradenotification

	
	InitiatingShiftDate := new(string)
	if u.InitiatingShiftDate != nil {
		
		*InitiatingShiftDate = timeutil.Strftime(u.InitiatingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftDate = nil
	}
	
	ReceivingShiftDate := new(string)
	if u.ReceivingShiftDate != nil {
		
		*ReceivingShiftDate = timeutil.Strftime(u.ReceivingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReceivingShiftDate = nil
	}
	

	return json.Marshal(&struct { 
		WeekDate *string `json:"weekDate,omitempty"`
		
		TradeId *string `json:"tradeId,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		NewState *string `json:"newState,omitempty"`
		
		InitiatingUser *Userreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftDate *string `json:"initiatingShiftDate,omitempty"`
		
		ReceivingUser *Userreference `json:"receivingUser,omitempty"`
		
		ReceivingShiftDate *string `json:"receivingShiftDate,omitempty"`
		*Alias
	}{ 
		WeekDate: u.WeekDate,
		
		TradeId: u.TradeId,
		
		OneSided: u.OneSided,
		
		NewState: u.NewState,
		
		InitiatingUser: u.InitiatingUser,
		
		InitiatingShiftDate: InitiatingShiftDate,
		
		ReceivingUser: u.ReceivingUser,
		
		ReceivingShiftDate: ReceivingShiftDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

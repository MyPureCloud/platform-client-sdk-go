package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Wfmusernotificationtopicshifttradenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotificationtopicshifttradenotification

	
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
		
		InitiatingUser *Wfmusernotificationtopicuserreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftDate *string `json:"initiatingShiftDate,omitempty"`
		
		ReceivingUser *Wfmusernotificationtopicuserreference `json:"receivingUser,omitempty"`
		
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
func (o *Wfmusernotificationtopicshifttradenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Shifttradenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradenotification
	
	InitiatingShiftDate := new(string)
	if o.InitiatingShiftDate != nil {
		
		*InitiatingShiftDate = timeutil.Strftime(o.InitiatingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftDate = nil
	}
	
	ReceivingShiftDate := new(string)
	if o.ReceivingShiftDate != nil {
		
		*ReceivingShiftDate = timeutil.Strftime(o.ReceivingShiftDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		WeekDate: o.WeekDate,
		
		TradeId: o.TradeId,
		
		OneSided: o.OneSided,
		
		NewState: o.NewState,
		
		InitiatingUser: o.InitiatingUser,
		
		InitiatingShiftDate: InitiatingShiftDate,
		
		ReceivingUser: o.ReceivingUser,
		
		ReceivingShiftDate: ReceivingShiftDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradenotification) UnmarshalJSON(b []byte) error {
	var ShifttradenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradenotificationMap)
	if err != nil {
		return err
	}
	
	if WeekDate, ok := ShifttradenotificationMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if TradeId, ok := ShifttradenotificationMap["tradeId"].(string); ok {
		o.TradeId = &TradeId
	}
	
	if OneSided, ok := ShifttradenotificationMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
	
	if NewState, ok := ShifttradenotificationMap["newState"].(string); ok {
		o.NewState = &NewState
	}
	
	if InitiatingUser, ok := ShifttradenotificationMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if initiatingShiftDateString, ok := ShifttradenotificationMap["initiatingShiftDate"].(string); ok {
		InitiatingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftDateString)
		o.InitiatingShiftDate = &InitiatingShiftDate
	}
	
	if ReceivingUser, ok := ShifttradenotificationMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if receivingShiftDateString, ok := ShifttradenotificationMap["receivingShiftDate"].(string); ok {
		ReceivingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftDateString)
		o.ReceivingShiftDate = &ReceivingShiftDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

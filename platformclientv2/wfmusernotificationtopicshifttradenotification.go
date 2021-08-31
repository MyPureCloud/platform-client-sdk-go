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

func (o *Wfmusernotificationtopicshifttradenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotificationtopicshifttradenotification
	
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
		
		InitiatingUser *Wfmusernotificationtopicuserreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftDate *string `json:"initiatingShiftDate,omitempty"`
		
		ReceivingUser *Wfmusernotificationtopicuserreference `json:"receivingUser,omitempty"`
		
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

func (o *Wfmusernotificationtopicshifttradenotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationtopicshifttradenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationtopicshifttradenotificationMap)
	if err != nil {
		return err
	}
	
	if WeekDate, ok := WfmusernotificationtopicshifttradenotificationMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if TradeId, ok := WfmusernotificationtopicshifttradenotificationMap["tradeId"].(string); ok {
		o.TradeId = &TradeId
	}
	
	if OneSided, ok := WfmusernotificationtopicshifttradenotificationMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
	
	if NewState, ok := WfmusernotificationtopicshifttradenotificationMap["newState"].(string); ok {
		o.NewState = &NewState
	}
	
	if InitiatingUser, ok := WfmusernotificationtopicshifttradenotificationMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if initiatingShiftDateString, ok := WfmusernotificationtopicshifttradenotificationMap["initiatingShiftDate"].(string); ok {
		InitiatingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftDateString)
		o.InitiatingShiftDate = &InitiatingShiftDate
	}
	
	if ReceivingUser, ok := WfmusernotificationtopicshifttradenotificationMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if receivingShiftDateString, ok := WfmusernotificationtopicshifttradenotificationMap["receivingShiftDate"].(string); ok {
		ReceivingShiftDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftDateString)
		o.ReceivingShiftDate = &ReceivingShiftDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicshifttradenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

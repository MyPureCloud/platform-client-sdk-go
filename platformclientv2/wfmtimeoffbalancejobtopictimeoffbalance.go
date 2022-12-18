package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffbalancejobtopictimeoffbalance
type Wfmtimeoffbalancejobtopictimeoffbalance struct { 
	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// HrisTimeOffTypeId
	HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`


	// HrisTimeOffTypeSecondaryId
	HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// BalanceMinutesPerDay
	BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`

}

func (o *Wfmtimeoffbalancejobtopictimeoffbalance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmtimeoffbalancejobtopictimeoffbalance
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`
		
		HrisTimeOffTypeSecondaryId *string `json:"hrisTimeOffTypeSecondaryId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		HrisTimeOffTypeId: o.HrisTimeOffTypeId,
		
		HrisTimeOffTypeSecondaryId: o.HrisTimeOffTypeSecondaryId,
		
		StartDate: StartDate,
		
		BalanceMinutesPerDay: o.BalanceMinutesPerDay,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmtimeoffbalancejobtopictimeoffbalance) UnmarshalJSON(b []byte) error {
	var WfmtimeoffbalancejobtopictimeoffbalanceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffbalancejobtopictimeoffbalanceMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if HrisTimeOffTypeId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["hrisTimeOffTypeId"].(string); ok {
		o.HrisTimeOffTypeId = &HrisTimeOffTypeId
	}
    
	if HrisTimeOffTypeSecondaryId, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["hrisTimeOffTypeSecondaryId"].(string); ok {
		o.HrisTimeOffTypeSecondaryId = &HrisTimeOffTypeSecondaryId
	}
    
	if startDateString, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if BalanceMinutesPerDay, ok := WfmtimeoffbalancejobtopictimeoffbalanceMap["balanceMinutesPerDay"].([]interface{}); ok {
		BalanceMinutesPerDayString, _ := json.Marshal(BalanceMinutesPerDay)
		json.Unmarshal(BalanceMinutesPerDayString, &o.BalanceMinutesPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffbalancejobtopictimeoffbalance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

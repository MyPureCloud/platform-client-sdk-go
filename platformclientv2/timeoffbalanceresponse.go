package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffbalanceresponse
type Timeoffbalanceresponse struct { 
	// ActivityCodeId - The ID for activity code associated with time off balance
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// HrisTimeOffTypeId - The ID of the time off type configured in HRIS integration
	HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`


	// StartDate - The Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`


	// BalanceMinutesPerDay - The list of available time off balance values in minutes for each day
	BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`

}

func (o *Timeoffbalanceresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffbalanceresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		HrisTimeOffTypeId *string `json:"hrisTimeOffTypeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		BalanceMinutesPerDay *[]int `json:"balanceMinutesPerDay,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		HrisTimeOffTypeId: o.HrisTimeOffTypeId,
		
		StartDate: StartDate,
		
		BalanceMinutesPerDay: o.BalanceMinutesPerDay,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffbalanceresponse) UnmarshalJSON(b []byte) error {
	var TimeoffbalanceresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffbalanceresponseMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := TimeoffbalanceresponseMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if HrisTimeOffTypeId, ok := TimeoffbalanceresponseMap["hrisTimeOffTypeId"].(string); ok {
		o.HrisTimeOffTypeId = &HrisTimeOffTypeId
	}
    
	if startDateString, ok := TimeoffbalanceresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if BalanceMinutesPerDay, ok := TimeoffbalanceresponseMap["balanceMinutesPerDay"].([]interface{}); ok {
		BalanceMinutesPerDayString, _ := json.Marshal(BalanceMinutesPerDay)
		json.Unmarshal(BalanceMinutesPerDayString, &o.BalanceMinutesPerDay)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffbalanceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

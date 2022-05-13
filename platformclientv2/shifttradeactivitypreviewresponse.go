package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradeactivitypreviewresponse
type Shifttradeactivitypreviewresponse struct { 
	// StartDate - The start date and time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length in minutes of this activity
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// ActivityCodeId - The ID of the activity code for this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// CountsAsPaidTime - Whether this activity counts as paid time
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

}

func (o *Shifttradeactivitypreviewresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradeactivitypreviewresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		ActivityCodeId: o.ActivityCodeId,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradeactivitypreviewresponse) UnmarshalJSON(b []byte) error {
	var ShifttradeactivitypreviewresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradeactivitypreviewresponseMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := ShifttradeactivitypreviewresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := ShifttradeactivitypreviewresponseMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if ActivityCodeId, ok := ShifttradeactivitypreviewresponseMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if CountsAsPaidTime, ok := ShifttradeactivitypreviewresponseMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradeactivitypreviewresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchreviewuserresponse
type Shifttradematchreviewuserresponse struct { 
	// WeeklyMinimumPaidMinutes - The minimum weekly paid minutes for this user per the work plan tied to the agent schedule
	WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`


	// WeeklyMaximumPaidMinutes - The maximum weekly paid minutes for this user per the work plan tied to the agent schedule
	WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`


	// PreTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user prior to the shift trade
	PreTradeSchedulePaidMinutes *int `json:"preTradeSchedulePaidMinutes,omitempty"`


	// PostTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user if the shift trade is approved
	PostTradeSchedulePaidMinutes *int `json:"postTradeSchedulePaidMinutes,omitempty"`


	// PostTradeNewShift - Preview of what the shift will look like for the opposite side of this trade after the match is approved
	PostTradeNewShift *Shifttradepreviewresponse `json:"postTradeNewShift,omitempty"`

}

func (o *Shifttradematchreviewuserresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradematchreviewuserresponse
	
	return json.Marshal(&struct { 
		WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`
		
		WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`
		
		PreTradeSchedulePaidMinutes *int `json:"preTradeSchedulePaidMinutes,omitempty"`
		
		PostTradeSchedulePaidMinutes *int `json:"postTradeSchedulePaidMinutes,omitempty"`
		
		PostTradeNewShift *Shifttradepreviewresponse `json:"postTradeNewShift,omitempty"`
		*Alias
	}{ 
		WeeklyMinimumPaidMinutes: o.WeeklyMinimumPaidMinutes,
		
		WeeklyMaximumPaidMinutes: o.WeeklyMaximumPaidMinutes,
		
		PreTradeSchedulePaidMinutes: o.PreTradeSchedulePaidMinutes,
		
		PostTradeSchedulePaidMinutes: o.PostTradeSchedulePaidMinutes,
		
		PostTradeNewShift: o.PostTradeNewShift,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradematchreviewuserresponse) UnmarshalJSON(b []byte) error {
	var ShifttradematchreviewuserresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchreviewuserresponseMap)
	if err != nil {
		return err
	}
	
	if WeeklyMinimumPaidMinutes, ok := ShifttradematchreviewuserresponseMap["weeklyMinimumPaidMinutes"].(float64); ok {
		WeeklyMinimumPaidMinutesInt := int(WeeklyMinimumPaidMinutes)
		o.WeeklyMinimumPaidMinutes = &WeeklyMinimumPaidMinutesInt
	}
	
	if WeeklyMaximumPaidMinutes, ok := ShifttradematchreviewuserresponseMap["weeklyMaximumPaidMinutes"].(float64); ok {
		WeeklyMaximumPaidMinutesInt := int(WeeklyMaximumPaidMinutes)
		o.WeeklyMaximumPaidMinutes = &WeeklyMaximumPaidMinutesInt
	}
	
	if PreTradeSchedulePaidMinutes, ok := ShifttradematchreviewuserresponseMap["preTradeSchedulePaidMinutes"].(float64); ok {
		PreTradeSchedulePaidMinutesInt := int(PreTradeSchedulePaidMinutes)
		o.PreTradeSchedulePaidMinutes = &PreTradeSchedulePaidMinutesInt
	}
	
	if PostTradeSchedulePaidMinutes, ok := ShifttradematchreviewuserresponseMap["postTradeSchedulePaidMinutes"].(float64); ok {
		PostTradeSchedulePaidMinutesInt := int(PostTradeSchedulePaidMinutes)
		o.PostTradeSchedulePaidMinutes = &PostTradeSchedulePaidMinutesInt
	}
	
	if PostTradeNewShift, ok := ShifttradematchreviewuserresponseMap["postTradeNewShift"].(map[string]interface{}); ok {
		PostTradeNewShiftString, _ := json.Marshal(PostTradeNewShift)
		json.Unmarshal(PostTradeNewShiftString, &o.PostTradeNewShift)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewuserresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

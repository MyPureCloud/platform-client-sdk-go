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

func (u *Shifttradematchreviewuserresponse) MarshalJSON() ([]byte, error) {
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
		WeeklyMinimumPaidMinutes: u.WeeklyMinimumPaidMinutes,
		
		WeeklyMaximumPaidMinutes: u.WeeklyMaximumPaidMinutes,
		
		PreTradeSchedulePaidMinutes: u.PreTradeSchedulePaidMinutes,
		
		PostTradeSchedulePaidMinutes: u.PostTradeSchedulePaidMinutes,
		
		PostTradeNewShift: u.PostTradeNewShift,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewuserresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

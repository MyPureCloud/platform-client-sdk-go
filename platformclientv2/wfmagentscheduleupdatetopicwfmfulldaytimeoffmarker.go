package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker
type Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker struct { 
	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// ManagementUnitDate
	ManagementUnitDate *string `json:"managementUnitDate,omitempty"`


	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// IsPaid
	IsPaid *bool `json:"isPaid,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Paid
	Paid *bool `json:"paid,omitempty"`

}

func (u *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker

	

	return json.Marshal(&struct { 
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ManagementUnitDate *string `json:"managementUnitDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		*Alias
	}{ 
		TimeOffRequestId: u.TimeOffRequestId,
		
		ManagementUnitDate: u.ManagementUnitDate,
		
		ActivityCodeId: u.ActivityCodeId,
		
		IsPaid: u.IsPaid,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Description: u.Description,
		
		Paid: u.Paid,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

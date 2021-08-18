package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bufulldaytimeoffmarker
type Bufulldaytimeoffmarker struct { 
	// BusinessUnitDate - The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`


	// LengthMinutes - The length of the time off marker in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of the time off marker
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with the time off marker
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid - Whether the time off marker is paid
	Paid *bool `json:"paid,omitempty"`


	// TimeOffRequestId - The ID of the time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

}

func (u *Bufulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bufulldaytimeoffmarker

	
	BusinessUnitDate := new(string)
	if u.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(u.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	

	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		*Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		LengthMinutes: u.LengthMinutes,
		
		Description: u.Description,
		
		ActivityCodeId: u.ActivityCodeId,
		
		Paid: u.Paid,
		
		TimeOffRequestId: u.TimeOffRequestId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bufulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

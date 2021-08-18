package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userschedulefulldaytimeoffmarker - Marker to indicate an approved full day time off request
type Userschedulefulldaytimeoffmarker struct { 
	// ManagementUnitDate - The date associated with the time off request that this marker corresponds to.  Date only, in ISO-8601 format.
	ManagementUnitDate *string `json:"managementUnitDate,omitempty"`


	// ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// IsPaid - Whether this is paid time off
	IsPaid *bool `json:"isPaid,omitempty"`


	// LengthInMinutes - The length in minutes of this time off marker
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description - The description associated with the time off request that this marker corresponds to
	Description *string `json:"description,omitempty"`


	// Delete - If marked true for updating an existing full day time off marker, it will be deleted
	Delete *bool `json:"delete,omitempty"`

}

func (u *Userschedulefulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedulefulldaytimeoffmarker

	

	return json.Marshal(&struct { 
		ManagementUnitDate *string `json:"managementUnitDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		*Alias
	}{ 
		ManagementUnitDate: u.ManagementUnitDate,
		
		ActivityCodeId: u.ActivityCodeId,
		
		IsPaid: u.IsPaid,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Description: u.Description,
		
		Delete: u.Delete,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userschedulefulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

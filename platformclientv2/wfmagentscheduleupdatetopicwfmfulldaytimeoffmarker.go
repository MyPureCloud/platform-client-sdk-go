package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

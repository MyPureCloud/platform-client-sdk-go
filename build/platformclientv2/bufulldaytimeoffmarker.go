package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bufulldaytimeoffmarker
type Bufulldaytimeoffmarker struct { 
	// BusinessUnitDate - The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`


	// LengthMinutes - The length of the time off marker in minutes
	LengthMinutes *int32 `json:"lengthMinutes,omitempty"`


	// Description - The description of the time off marker
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with the time off marker
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid - Whether the time off marker is paid
	Paid *bool `json:"paid,omitempty"`


	// TimeOffRequestId - The ID of the time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bufulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

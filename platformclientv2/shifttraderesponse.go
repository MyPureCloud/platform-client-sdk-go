package platformclientv2
import (
	"time"
	"encoding/json"
)

// Shifttraderesponse
type Shifttraderesponse struct { 
	// Id - The ID of this shift trade
	Id *string `json:"id,omitempty"`


	// WeekDate - The start week date of the associated schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// Schedule - The ID of the associated schedule
	Schedule *Weekschedulereference `json:"schedule,omitempty"`


	// State - The state of this shift trade
	State *string `json:"state,omitempty"`


	// InitiatingUser - The user who initiated this trade
	InitiatingUser *Userreference `json:"initiatingUser,omitempty"`


	// InitiatingShiftId - The ID of the shift offered for trade by the initiating user
	InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`


	// InitiatingShiftStart - The start date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	InitiatingShiftStart *time.Time `json:"initiatingShiftStart,omitempty"`


	// InitiatingShiftEnd - The end date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	InitiatingShiftEnd *time.Time `json:"initiatingShiftEnd,omitempty"`


	// ReceivingUser - The user matching the trade, or if the state is not Matched, the user to whom the trade request was sent
	ReceivingUser *Userreference `json:"receivingUser,omitempty"`


	// ReceivingShiftId - The ID of the shift being exchanged for the initiating shift, null if the receiving user is picking up a shift
	ReceivingShiftId *string `json:"receivingShiftId,omitempty"`


	// ReceivingShiftStart - The start date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReceivingShiftStart *time.Time `json:"receivingShiftStart,omitempty"`


	// ReceivingShiftEnd - The end date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReceivingShiftEnd *time.Time `json:"receivingShiftEnd,omitempty"`


	// Expiration - When this shift trade offer will expire if not matched or approved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Expiration *time.Time `json:"expiration,omitempty"`


	// OneSided - Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
	OneSided *bool `json:"oneSided,omitempty"`


	// AcceptableIntervals
	AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`


	// ReviewedBy - The user who reviewed this shift trade
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`


	// ReviewedDate - The timestamp when this shift trade was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// Metadata - Version data for this trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

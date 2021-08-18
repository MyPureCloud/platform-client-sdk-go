package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Addshifttraderequest
type Addshifttraderequest struct { 
	// ScheduleId - The ID of the schedule to which the initiating and receiving shifts belong
	ScheduleId *string `json:"scheduleId,omitempty"`


	// InitiatingShiftId - The ID of the shift that the initiating user wants to give up
	InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`


	// ReceivingUserId - The ID of the user to whom to send the request (for use in direct trade requests)
	ReceivingUserId *string `json:"receivingUserId,omitempty"`


	// Expiration - When this shift trade request should expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`


	// AcceptableIntervals
	AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`

}

func (u *Addshifttraderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Addshifttraderequest

	
	Expiration := new(string)
	if u.Expiration != nil {
		
		*Expiration = timeutil.Strftime(u.Expiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expiration = nil
	}
	

	return json.Marshal(&struct { 
		ScheduleId *string `json:"scheduleId,omitempty"`
		
		InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`
		
		ReceivingUserId *string `json:"receivingUserId,omitempty"`
		
		Expiration *string `json:"expiration,omitempty"`
		
		AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`
		*Alias
	}{ 
		ScheduleId: u.ScheduleId,
		
		InitiatingShiftId: u.InitiatingShiftId,
		
		ReceivingUserId: u.ReceivingUserId,
		
		Expiration: Expiration,
		
		AcceptableIntervals: u.AcceptableIntervals,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Addshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

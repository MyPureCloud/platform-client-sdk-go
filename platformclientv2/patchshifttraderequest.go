package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchshifttraderequest
type Patchshifttraderequest struct { 
	// ReceivingUserId - Update the ID of the receiving user to direct the request at a specific user, or set the wrapped id to null to open up a trade to be matched by any user.
	ReceivingUserId *Valuewrapperstring `json:"receivingUserId,omitempty"`


	// Expiration - Update the expiration time for this shift trade.
	Expiration *Valuewrapperdate `json:"expiration,omitempty"`


	// AcceptableIntervals - Update the acceptable intervals the initiating user is willing to accept in trade. Setting the enclosed list to empty will make this a one sided trade request
	AcceptableIntervals *Listwrapperinterval `json:"acceptableIntervals,omitempty"`


	// Metadata - Version metadata
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (u *Patchshifttraderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchshifttraderequest

	

	return json.Marshal(&struct { 
		ReceivingUserId *Valuewrapperstring `json:"receivingUserId,omitempty"`
		
		Expiration *Valuewrapperdate `json:"expiration,omitempty"`
		
		AcceptableIntervals *Listwrapperinterval `json:"acceptableIntervals,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		ReceivingUserId: u.ReceivingUserId,
		
		Expiration: u.Expiration,
		
		AcceptableIntervals: u.AcceptableIntervals,
		
		Metadata: u.Metadata,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

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

func (o *Patchshifttraderequest) MarshalJSON() ([]byte, error) {
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
		ReceivingUserId: o.ReceivingUserId,
		
		Expiration: o.Expiration,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchshifttraderequest) UnmarshalJSON(b []byte) error {
	var PatchshifttraderequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchshifttraderequestMap)
	if err != nil {
		return err
	}
	
	if ReceivingUserId, ok := PatchshifttraderequestMap["receivingUserId"].(map[string]interface{}); ok {
		ReceivingUserIdString, _ := json.Marshal(ReceivingUserId)
		json.Unmarshal(ReceivingUserIdString, &o.ReceivingUserId)
	}
	
	if Expiration, ok := PatchshifttraderequestMap["expiration"].(map[string]interface{}); ok {
		ExpirationString, _ := json.Marshal(Expiration)
		json.Unmarshal(ExpirationString, &o.Expiration)
	}
	
	if AcceptableIntervals, ok := PatchshifttraderequestMap["acceptableIntervals"].(map[string]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if Metadata, ok := PatchshifttraderequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

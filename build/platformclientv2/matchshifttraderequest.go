package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Matchshifttraderequest
type Matchshifttraderequest struct { 
	// ReceivingScheduleId - The ID of the schedule with which the shift trade is associated
	ReceivingScheduleId *string `json:"receivingScheduleId,omitempty"`


	// ReceivingShiftId - The ID of the shift the receiving user is giving up in trade, if applicable
	ReceivingShiftId *string `json:"receivingShiftId,omitempty"`


	// Metadata - Version metadata for the shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Matchshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

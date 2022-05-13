package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Matchshifttraderequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Matchshifttraderequest
	
	return json.Marshal(&struct { 
		ReceivingScheduleId *string `json:"receivingScheduleId,omitempty"`
		
		ReceivingShiftId *string `json:"receivingShiftId,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		ReceivingScheduleId: o.ReceivingScheduleId,
		
		ReceivingShiftId: o.ReceivingShiftId,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Matchshifttraderequest) UnmarshalJSON(b []byte) error {
	var MatchshifttraderequestMap map[string]interface{}
	err := json.Unmarshal(b, &MatchshifttraderequestMap)
	if err != nil {
		return err
	}
	
	if ReceivingScheduleId, ok := MatchshifttraderequestMap["receivingScheduleId"].(string); ok {
		o.ReceivingScheduleId = &ReceivingScheduleId
	}
    
	if ReceivingShiftId, ok := MatchshifttraderequestMap["receivingShiftId"].(string); ok {
		o.ReceivingShiftId = &ReceivingShiftId
	}
    
	if Metadata, ok := MatchshifttraderequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Matchshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapestimaterequest
type Actionmapestimaterequest struct { 
	// SegmentIds - List of Segment IDs.
	SegmentIds *[]string `json:"segmentIds,omitempty"`


	// OutcomeCriteria - Outcome Criteria containing outcomeId and probability thresholds.
	OutcomeCriteria *Actionmapestimateoutcomecriteria `json:"outcomeCriteria,omitempty"`

}

func (o *Actionmapestimaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapestimaterequest
	
	return json.Marshal(&struct { 
		SegmentIds *[]string `json:"segmentIds,omitempty"`
		
		OutcomeCriteria *Actionmapestimateoutcomecriteria `json:"outcomeCriteria,omitempty"`
		*Alias
	}{ 
		SegmentIds: o.SegmentIds,
		
		OutcomeCriteria: o.OutcomeCriteria,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionmapestimaterequest) UnmarshalJSON(b []byte) error {
	var ActionmapestimaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapestimaterequestMap)
	if err != nil {
		return err
	}
	
	if SegmentIds, ok := ActionmapestimaterequestMap["segmentIds"].([]interface{}); ok {
		SegmentIdsString, _ := json.Marshal(SegmentIds)
		json.Unmarshal(SegmentIdsString, &o.SegmentIds)
	}
	
	if OutcomeCriteria, ok := ActionmapestimaterequestMap["outcomeCriteria"].(map[string]interface{}); ok {
		OutcomeCriteriaString, _ := json.Marshal(OutcomeCriteria)
		json.Unmarshal(OutcomeCriteriaString, &o.OutcomeCriteria)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapestimaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

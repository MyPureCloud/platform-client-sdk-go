package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Segmentestimatecount
type Segmentestimatecount struct { 
	// SegmentId - ID of Segment.
	SegmentId *string `json:"segmentId,omitempty"`


	// Count - Estimate count per segment.
	Count *int `json:"count,omitempty"`

}

func (o *Segmentestimatecount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Segmentestimatecount
	
	return json.Marshal(&struct { 
		SegmentId *string `json:"segmentId,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		SegmentId: o.SegmentId,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Segmentestimatecount) UnmarshalJSON(b []byte) error {
	var SegmentestimatecountMap map[string]interface{}
	err := json.Unmarshal(b, &SegmentestimatecountMap)
	if err != nil {
		return err
	}
	
	if SegmentId, ok := SegmentestimatecountMap["segmentId"].(string); ok {
		o.SegmentId = &SegmentId
	}
    
	if Count, ok := SegmentestimatecountMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Segmentestimatecount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapestimateresult
type Actionmapestimateresult struct { 
	// QualifiedSessionCount - Number of sessions qualified for Action map.
	QualifiedSessionCount *int `json:"qualifiedSessionCount,omitempty"`


	// TotalSessionCount - Total number of sessions.
	TotalSessionCount *int `json:"totalSessionCount,omitempty"`


	// PerSegmentCounts - Number of sessions qualified for Action map per segment.
	PerSegmentCounts *[]Segmentestimatecount `json:"perSegmentCounts,omitempty"`


	// OutcomesScoresCount - Difference made by outcome criteria to number of sessions qualified for Action map.
	OutcomesScoresCount *int `json:"outcomesScoresCount,omitempty"`

}

func (o *Actionmapestimateresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapestimateresult
	
	return json.Marshal(&struct { 
		QualifiedSessionCount *int `json:"qualifiedSessionCount,omitempty"`
		
		TotalSessionCount *int `json:"totalSessionCount,omitempty"`
		
		PerSegmentCounts *[]Segmentestimatecount `json:"perSegmentCounts,omitempty"`
		
		OutcomesScoresCount *int `json:"outcomesScoresCount,omitempty"`
		*Alias
	}{ 
		QualifiedSessionCount: o.QualifiedSessionCount,
		
		TotalSessionCount: o.TotalSessionCount,
		
		PerSegmentCounts: o.PerSegmentCounts,
		
		OutcomesScoresCount: o.OutcomesScoresCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionmapestimateresult) UnmarshalJSON(b []byte) error {
	var ActionmapestimateresultMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapestimateresultMap)
	if err != nil {
		return err
	}
	
	if QualifiedSessionCount, ok := ActionmapestimateresultMap["qualifiedSessionCount"].(float64); ok {
		QualifiedSessionCountInt := int(QualifiedSessionCount)
		o.QualifiedSessionCount = &QualifiedSessionCountInt
	}
	
	if TotalSessionCount, ok := ActionmapestimateresultMap["totalSessionCount"].(float64); ok {
		TotalSessionCountInt := int(TotalSessionCount)
		o.TotalSessionCount = &TotalSessionCountInt
	}
	
	if PerSegmentCounts, ok := ActionmapestimateresultMap["perSegmentCounts"].([]interface{}); ok {
		PerSegmentCountsString, _ := json.Marshal(PerSegmentCounts)
		json.Unmarshal(PerSegmentCountsString, &o.PerSegmentCounts)
	}
	
	if OutcomesScoresCount, ok := ActionmapestimateresultMap["outcomesScoresCount"].(float64); ok {
		OutcomesScoresCountInt := int(OutcomesScoresCount)
		o.OutcomesScoresCount = &OutcomesScoresCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapestimateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

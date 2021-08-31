package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherenceactuals
type Historicaladherenceactuals struct { 
	// ActualActivityCategory - Activity in which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`


	// StartOffsetSeconds - Actual start offset in seconds relative to query start time
	StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`


	// EndOffsetSeconds - Actual end offset in seconds relative to query start time
	EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`

}

func (o *Historicaladherenceactuals) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicaladherenceactuals
	
	return json.Marshal(&struct { 
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`
		
		EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`
		*Alias
	}{ 
		ActualActivityCategory: o.ActualActivityCategory,
		
		StartOffsetSeconds: o.StartOffsetSeconds,
		
		EndOffsetSeconds: o.EndOffsetSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicaladherenceactuals) UnmarshalJSON(b []byte) error {
	var HistoricaladherenceactualsMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricaladherenceactualsMap)
	if err != nil {
		return err
	}
	
	if ActualActivityCategory, ok := HistoricaladherenceactualsMap["actualActivityCategory"].(string); ok {
		o.ActualActivityCategory = &ActualActivityCategory
	}
	
	if StartOffsetSeconds, ok := HistoricaladherenceactualsMap["startOffsetSeconds"].(float64); ok {
		StartOffsetSecondsInt := int(StartOffsetSeconds)
		o.StartOffsetSeconds = &StartOffsetSecondsInt
	}
	
	if EndOffsetSeconds, ok := HistoricaladherenceactualsMap["endOffsetSeconds"].(float64); ok {
		EndOffsetSecondsInt := int(EndOffsetSeconds)
		o.EndOffsetSeconds = &EndOffsetSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicaladherenceactuals) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

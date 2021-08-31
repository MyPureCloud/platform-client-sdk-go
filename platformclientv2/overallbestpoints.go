package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Overallbestpoints
type Overallbestpoints struct { 
	// Division - The requested division
	Division *Division `json:"division,omitempty"`


	// BestPoints - List of gamification best point items
	BestPoints *[]Overallbestpointsitem `json:"bestPoints,omitempty"`

}

func (o *Overallbestpoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Overallbestpoints
	
	return json.Marshal(&struct { 
		Division *Division `json:"division,omitempty"`
		
		BestPoints *[]Overallbestpointsitem `json:"bestPoints,omitempty"`
		*Alias
	}{ 
		Division: o.Division,
		
		BestPoints: o.BestPoints,
		Alias:    (*Alias)(o),
	})
}

func (o *Overallbestpoints) UnmarshalJSON(b []byte) error {
	var OverallbestpointsMap map[string]interface{}
	err := json.Unmarshal(b, &OverallbestpointsMap)
	if err != nil {
		return err
	}
	
	if Division, ok := OverallbestpointsMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if BestPoints, ok := OverallbestpointsMap["bestPoints"].([]interface{}); ok {
		BestPointsString, _ := json.Marshal(BestPoints)
		json.Unmarshal(BestPointsString, &o.BestPoints)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Overallbestpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

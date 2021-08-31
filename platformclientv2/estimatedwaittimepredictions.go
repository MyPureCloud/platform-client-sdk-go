package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Estimatedwaittimepredictions
type Estimatedwaittimepredictions struct { 
	// Results - Returned upon a successful estimated wait time request.
	Results *[]Predictionresults `json:"results,omitempty"`

}

func (o *Estimatedwaittimepredictions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Estimatedwaittimepredictions
	
	return json.Marshal(&struct { 
		Results *[]Predictionresults `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Estimatedwaittimepredictions) UnmarshalJSON(b []byte) error {
	var EstimatedwaittimepredictionsMap map[string]interface{}
	err := json.Unmarshal(b, &EstimatedwaittimepredictionsMap)
	if err != nil {
		return err
	}
	
	if Results, ok := EstimatedwaittimepredictionsMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Estimatedwaittimepredictions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

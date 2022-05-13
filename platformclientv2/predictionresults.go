package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictionresults
type Predictionresults struct { 
	// Intent - Indicates the media type scope of this estimated wait time
	Intent *string `json:"intent,omitempty"`


	// Formula - Indicates the estimated wait time Formula
	Formula *string `json:"formula,omitempty"`


	// EstimatedWaitTimeSeconds - Estimated wait time in seconds
	EstimatedWaitTimeSeconds *int `json:"estimatedWaitTimeSeconds,omitempty"`

}

func (o *Predictionresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictionresults
	
	return json.Marshal(&struct { 
		Intent *string `json:"intent,omitempty"`
		
		Formula *string `json:"formula,omitempty"`
		
		EstimatedWaitTimeSeconds *int `json:"estimatedWaitTimeSeconds,omitempty"`
		*Alias
	}{ 
		Intent: o.Intent,
		
		Formula: o.Formula,
		
		EstimatedWaitTimeSeconds: o.EstimatedWaitTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictionresults) UnmarshalJSON(b []byte) error {
	var PredictionresultsMap map[string]interface{}
	err := json.Unmarshal(b, &PredictionresultsMap)
	if err != nil {
		return err
	}
	
	if Intent, ok := PredictionresultsMap["intent"].(string); ok {
		o.Intent = &Intent
	}
    
	if Formula, ok := PredictionresultsMap["formula"].(string); ok {
		o.Formula = &Formula
	}
    
	if EstimatedWaitTimeSeconds, ok := PredictionresultsMap["estimatedWaitTimeSeconds"].(float64); ok {
		EstimatedWaitTimeSecondsInt := int(EstimatedWaitTimeSeconds)
		o.EstimatedWaitTimeSeconds = &EstimatedWaitTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictionresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

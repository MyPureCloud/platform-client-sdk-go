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

func (u *Predictionresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictionresults

	

	return json.Marshal(&struct { 
		Intent *string `json:"intent,omitempty"`
		
		Formula *string `json:"formula,omitempty"`
		
		EstimatedWaitTimeSeconds *int `json:"estimatedWaitTimeSeconds,omitempty"`
		*Alias
	}{ 
		Intent: u.Intent,
		
		Formula: u.Formula,
		
		EstimatedWaitTimeSeconds: u.EstimatedWaitTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Predictionresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

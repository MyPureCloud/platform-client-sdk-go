package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Estimatedwaittimepredictions
type Estimatedwaittimepredictions struct { 
	// Results - Returned upon a successful estimated wait time request.
	Results *[]Predictionresults `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Estimatedwaittimepredictions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

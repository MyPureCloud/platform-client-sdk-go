package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkrelationshipsresponse
type Bulkrelationshipsresponse struct { 
	// Results
	Results *[]Bulkresponseresultrelationshiprelationship `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkrelationshipsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkorganizationsresponse
type Bulkorganizationsresponse struct { 
	// Results
	Results *[]Bulkresponseresultexternalorganizationexternalorganization `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkorganizationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

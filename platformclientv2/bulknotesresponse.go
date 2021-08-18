package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulknotesresponse
type Bulknotesresponse struct { 
	// Results
	Results *[]Bulkresponseresultnotenote `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

func (u *Bulknotesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulknotesresponse

	

	return json.Marshal(&struct { 
		Results *[]Bulkresponseresultnotenote `json:"results,omitempty"`
		
		ErrorCount *int `json:"errorCount,omitempty"`
		
		ErrorIndexes *[]int `json:"errorIndexes,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		
		ErrorCount: u.ErrorCount,
		
		ErrorIndexes: u.ErrorIndexes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bulknotesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"encoding/json"
)

// Bulkfetchorganizationsresponse
type Bulkfetchorganizationsresponse struct { 
	// Results
	Results *[]Bulkresponseresultexternalorganizationentity `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchorganizationsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

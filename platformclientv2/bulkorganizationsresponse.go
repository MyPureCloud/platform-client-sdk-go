package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}

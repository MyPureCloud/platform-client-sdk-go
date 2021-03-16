package platformclientv2
import (
	"encoding/json"
)

// Bulkdeleteresponse
type Bulkdeleteresponse struct { 
	// Results
	Results *[]Bulkresponseresultvoidentity `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkdeleteresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

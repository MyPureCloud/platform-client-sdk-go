package platformclientv2
import (
	"encoding/json"
)

// Bulkfetchcontactsresponse
type Bulkfetchcontactsresponse struct { 
	// Results
	Results *[]Bulkresponseresultexternalcontactentity `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkfetchcontactsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

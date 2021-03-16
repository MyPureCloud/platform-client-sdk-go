package platformclientv2
import (
	"encoding/json"
)

// Bulkcontactsresponse
type Bulkcontactsresponse struct { 
	// Results
	Results *[]Bulkresponseresultexternalcontactexternalcontact `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkcontactsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

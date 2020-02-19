package platformclientv2
import (
	"encoding/json"
)

// Tagqueryrequest
type Tagqueryrequest struct { 
	// Query
	Query *string `json:"query,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`

}

// String returns a JSON representation of the model
func (o *Tagqueryrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Entitylisting
type Entitylisting struct { 
	// Entities
	Entities *[]Datatableimportjob `json:"entities,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Total
	Total *int64 `json:"total,omitempty"`


	// PageCount
	PageCount *int32 `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

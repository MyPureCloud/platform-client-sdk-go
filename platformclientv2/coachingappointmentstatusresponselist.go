package platformclientv2
import (
	"encoding/json"
)

// Coachingappointmentstatusresponselist
type Coachingappointmentstatusresponselist struct { 
	// Entities
	Entities *[]Coachingappointmentstatusresponse `json:"entities,omitempty"`


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
func (o *Coachingappointmentstatusresponselist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

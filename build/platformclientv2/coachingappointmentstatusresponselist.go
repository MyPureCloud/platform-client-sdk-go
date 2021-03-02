package platformclientv2
import (
	"encoding/json"
)

// Coachingappointmentstatusresponselist
type Coachingappointmentstatusresponselist struct { 
	// Entities
	Entities *[]Coachingappointmentstatusresponse `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponselist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

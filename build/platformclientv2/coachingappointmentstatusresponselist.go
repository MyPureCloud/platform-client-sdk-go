package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

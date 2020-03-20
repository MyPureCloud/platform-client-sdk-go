package platformclientv2
import (
	"time"
	"encoding/json"
)

// Routegrouplist
type Routegrouplist struct { 
	// StartDate - The reference start date for the route group arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// RouteGroups - The route group data for this forecast
	RouteGroups *[]Routegroup `json:"routeGroups,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routegrouplist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

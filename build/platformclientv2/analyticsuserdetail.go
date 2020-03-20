package platformclientv2
import (
	"encoding/json"
)

// Analyticsuserdetail
type Analyticsuserdetail struct { 
	// UserId - The identifier for the user
	UserId *string `json:"userId,omitempty"`


	// PrimaryPresence - The presence records for the user
	PrimaryPresence *[]Analyticsuserpresencerecord `json:"primaryPresence,omitempty"`


	// RoutingStatus - The ACD routing status records for the user
	RoutingStatus *[]Analyticsroutingstatusrecord `json:"routingStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserdetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

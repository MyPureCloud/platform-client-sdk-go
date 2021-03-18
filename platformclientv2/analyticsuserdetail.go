package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

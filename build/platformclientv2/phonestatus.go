package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Phonestatus
type Phonestatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// OperationalStatus - The Operational Status of this phone
	OperationalStatus *string `json:"operationalStatus,omitempty"`


	// EdgesStatus - The status of the primary or secondary Edges assigned to the phone lines.
	EdgesStatus *string `json:"edgesStatus,omitempty"`


	// EventCreationTime - Event Creation Time represents an ISO-8601 string. For example: UTC, UTC+01:00, or Europe/London
	EventCreationTime *string `json:"eventCreationTime,omitempty"`


	// Provision - Provision information for this phone
	Provision *Provisioninfo `json:"provision,omitempty"`


	// LineStatuses - A list of LineStatus information for each of the lines of this phone
	LineStatuses *[]Linestatus `json:"lineStatuses,omitempty"`


	// PhoneAssignmentToEdgeType - The phone status's edge assignment type.
	PhoneAssignmentToEdgeType *string `json:"phoneAssignmentToEdgeType,omitempty"`


	// Edge - The URI of the edge that provided this status information.
	Edge *Domainentityref `json:"edge,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

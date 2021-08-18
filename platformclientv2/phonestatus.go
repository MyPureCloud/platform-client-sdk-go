package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Phonestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonestatus

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		OperationalStatus *string `json:"operationalStatus,omitempty"`
		
		EdgesStatus *string `json:"edgesStatus,omitempty"`
		
		EventCreationTime *string `json:"eventCreationTime,omitempty"`
		
		Provision *Provisioninfo `json:"provision,omitempty"`
		
		LineStatuses *[]Linestatus `json:"lineStatuses,omitempty"`
		
		PhoneAssignmentToEdgeType *string `json:"phoneAssignmentToEdgeType,omitempty"`
		
		Edge *Domainentityref `json:"edge,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		OperationalStatus: u.OperationalStatus,
		
		EdgesStatus: u.EdgesStatus,
		
		EventCreationTime: u.EventCreationTime,
		
		Provision: u.Provision,
		
		LineStatuses: u.LineStatuses,
		
		PhoneAssignmentToEdgeType: u.PhoneAssignmentToEdgeType,
		
		Edge: u.Edge,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

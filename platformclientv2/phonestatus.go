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

func (o *Phonestatus) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		OperationalStatus: o.OperationalStatus,
		
		EdgesStatus: o.EdgesStatus,
		
		EventCreationTime: o.EventCreationTime,
		
		Provision: o.Provision,
		
		LineStatuses: o.LineStatuses,
		
		PhoneAssignmentToEdgeType: o.PhoneAssignmentToEdgeType,
		
		Edge: o.Edge,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonestatus) UnmarshalJSON(b []byte) error {
	var PhonestatusMap map[string]interface{}
	err := json.Unmarshal(b, &PhonestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonestatusMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PhonestatusMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if OperationalStatus, ok := PhonestatusMap["operationalStatus"].(string); ok {
		o.OperationalStatus = &OperationalStatus
	}
	
	if EdgesStatus, ok := PhonestatusMap["edgesStatus"].(string); ok {
		o.EdgesStatus = &EdgesStatus
	}
	
	if EventCreationTime, ok := PhonestatusMap["eventCreationTime"].(string); ok {
		o.EventCreationTime = &EventCreationTime
	}
	
	if Provision, ok := PhonestatusMap["provision"].(map[string]interface{}); ok {
		ProvisionString, _ := json.Marshal(Provision)
		json.Unmarshal(ProvisionString, &o.Provision)
	}
	
	if LineStatuses, ok := PhonestatusMap["lineStatuses"].([]interface{}); ok {
		LineStatusesString, _ := json.Marshal(LineStatuses)
		json.Unmarshal(LineStatusesString, &o.LineStatuses)
	}
	
	if PhoneAssignmentToEdgeType, ok := PhonestatusMap["phoneAssignmentToEdgeType"].(string); ok {
		o.PhoneAssignmentToEdgeType = &PhoneAssignmentToEdgeType
	}
	
	if Edge, ok := PhonestatusMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if SelfUri, ok := PhonestatusMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

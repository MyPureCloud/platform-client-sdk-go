package platformclientv2
import (
	"encoding/json"
)

// Planninggroup - Planning Group
type Planninggroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ServiceGoalTemplate - The ID of the service goal template associated with this planning group
	ServiceGoalTemplate *Servicegoaltemplatereference `json:"serviceGoalTemplate,omitempty"`


	// RoutePaths - Set of route paths associated with the planning group
	RoutePaths *[]Routepathresponse `json:"routePaths,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Planninggroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

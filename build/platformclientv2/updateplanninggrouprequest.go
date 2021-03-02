package platformclientv2
import (
	"encoding/json"
)

// Updateplanninggrouprequest
type Updateplanninggrouprequest struct { 
	// Name - The name of the planning group
	Name *string `json:"name,omitempty"`


	// RoutePaths - Set of route paths to associate with the planning group
	RoutePaths *Setwrapperroutepathrequest `json:"routePaths,omitempty"`


	// ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
	ServiceGoalTemplateId *string `json:"serviceGoalTemplateId,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

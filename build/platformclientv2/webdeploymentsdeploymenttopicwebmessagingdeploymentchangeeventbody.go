package platformclientv2
import (
	"encoding/json"
)

// Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody
type Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Configuration
	Configuration *Webdeploymentsdeploymenttopicwebmessagingconfigchangeeventbody `json:"configuration,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsdeploymenttopicwebmessagingdeploymentchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

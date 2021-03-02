package platformclientv2
import (
	"encoding/json"
)

// Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody
type Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Createqueuemediaassociationrequest - A combination of a single queue and one or more media types to associate with a service goal group
type Createqueuemediaassociationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Queue - The queue to associate with the service goal group
	Queue *Queuereference `json:"queue,omitempty"`


	// MediaTypes - The media types of the given queue to associate with the service goal group
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createqueuemediaassociationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

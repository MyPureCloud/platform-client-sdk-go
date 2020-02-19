package platformclientv2
import (
	"encoding/json"
)

// Queuemediaassociation - A combination of a single queue and one or more media types to associate with a service goal group
type Queuemediaassociation struct { 
	// Id - The reference ID for this QueueMediaAssociation
	Id *string `json:"id,omitempty"`


	// Queue - The queue to associate with the service goal group
	Queue *Queuereference `json:"queue,omitempty"`


	// MediaTypes - The media types of the given queue to associate with the service goal group
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// Delete - If marked true on a PATCH, this QueueMediaAssociation will be permanently deleted
	Delete *bool `json:"delete,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queuemediaassociation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

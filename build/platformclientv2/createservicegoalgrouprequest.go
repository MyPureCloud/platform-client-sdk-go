package platformclientv2
import (
	"encoding/json"
)

// Createservicegoalgrouprequest - Service Goal Group
type Createservicegoalgrouprequest struct { 
	// Name - name
	Name *string `json:"name,omitempty"`


	// Goals - Goals defined for this service goal group
	Goals *Servicegoalgroupgoals `json:"goals,omitempty"`


	// QueueMediaAssociations - List of queues and media types from that queue to associate with this service goal group
	QueueMediaAssociations *[]Createqueuemediaassociationrequest `json:"queueMediaAssociations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createservicegoalgrouprequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

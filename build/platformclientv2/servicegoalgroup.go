package platformclientv2
import (
	"encoding/json"
)

// Servicegoalgroup - Service Goal Group
type Servicegoalgroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Goals - Goals defined for this service goal group
	Goals *Servicegoalgroupgoals `json:"goals,omitempty"`


	// QueueMediaAssociations - List of queues and media types from that queue to associate with this service goal group
	QueueMediaAssociations *[]Queuemediaassociation `json:"queueMediaAssociations,omitempty"`


	// Metadata - Version metadata for the list of service goal groups for the associated management unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicegoalgroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

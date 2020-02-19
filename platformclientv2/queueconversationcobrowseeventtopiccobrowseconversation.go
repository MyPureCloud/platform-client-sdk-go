package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcobrowseeventtopiccobrowseconversation
type Queueconversationcobrowseeventtopiccobrowseconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Queueconversationcobrowseeventtopiccobrowsemediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopiccobrowseconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

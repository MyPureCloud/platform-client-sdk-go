package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcallbackeventtopiccallbackconversation
type Queueconversationcallbackeventtopiccallbackconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants
	Participants *[]Queueconversationcallbackeventtopiccallbackmediaparticipant `json:"participants,omitempty"`


	// OtherMediaUris
	OtherMediaUris *[]string `json:"otherMediaUris,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopiccallbackconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

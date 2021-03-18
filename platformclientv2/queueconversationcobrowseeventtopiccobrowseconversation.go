package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

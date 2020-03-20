package platformclientv2
import (
	"encoding/json"
)

// Messagestickerattachment
type Messagestickerattachment struct { 
	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagestickerattachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

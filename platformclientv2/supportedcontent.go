package platformclientv2
import (
	"encoding/json"
)

// Supportedcontent - Supported content for inbound and outbound messages
type Supportedcontent struct { 
	// MediaTypes - Defines the allowable media that may be accepted for an inbound message or to be sent in an outbound message. The following is an example of allowing all inbound media, and for outbound all images and only mpeg video: {   \"mediaTypes\": {     \"allow\": {       \"inbound\": [{\"type\": \"*/*\"}],       \"outbound\": [{\"type\": \"image/*\"}, {\"type\": \"video/mpeg\"}]     }   } }
	MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Supportedcontent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

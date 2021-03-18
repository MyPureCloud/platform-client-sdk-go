package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatypeaccess - Media type access definitions
type Mediatypeaccess struct { 
	// Inbound - List of media types allowed for inbound messages from customers. If inbound messages from a customer contain media that is not in this list, the media will be dropped from the outbound message.
	Inbound *[]Mediatype `json:"inbound,omitempty"`


	// Outbound - List of media types allowed for outbound messages to customers. If an outbound message is sent that contains media that is not in this list, the message will not be sent.
	Outbound *[]Mediatype `json:"outbound,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediatypeaccess) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Mediatypeaccess) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatypeaccess
	
	return json.Marshal(&struct { 
		Inbound *[]Mediatype `json:"inbound,omitempty"`
		
		Outbound *[]Mediatype `json:"outbound,omitempty"`
		*Alias
	}{ 
		Inbound: o.Inbound,
		
		Outbound: o.Outbound,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediatypeaccess) UnmarshalJSON(b []byte) error {
	var MediatypeaccessMap map[string]interface{}
	err := json.Unmarshal(b, &MediatypeaccessMap)
	if err != nil {
		return err
	}
	
	if Inbound, ok := MediatypeaccessMap["inbound"].([]interface{}); ok {
		InboundString, _ := json.Marshal(Inbound)
		json.Unmarshal(InboundString, &o.Inbound)
	}
	
	if Outbound, ok := MediatypeaccessMap["outbound"].([]interface{}); ok {
		OutboundString, _ := json.Marshal(Outbound)
		json.Unmarshal(OutboundString, &o.Outbound)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediatypeaccess) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

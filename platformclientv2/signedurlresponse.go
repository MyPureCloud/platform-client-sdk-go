package platformclientv2
import (
	"encoding/json"
)

// Signedurlresponse
type Signedurlresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Url - Url of the downloaded pcap file
	Url *string `json:"url,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Signedurlresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

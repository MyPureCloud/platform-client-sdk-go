package platformclientv2
import (
	"encoding/json"
)

// Queueemailaddress
type Queueemailaddress struct { 
	// Domain
	Domain *Domainentityref `json:"domain,omitempty"`


	// Route
	Route *Inboundroute `json:"route,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueemailaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

package platformclientv2
import (
	"encoding/json"
)

// Outboundmessagingmessagingcampaignconfigchangeerrordetail
type Outboundmessagingmessagingcampaignconfigchangeerrordetail struct { 
	// VarError
	VarError *string `json:"error,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeerrordetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

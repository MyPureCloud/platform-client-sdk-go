package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

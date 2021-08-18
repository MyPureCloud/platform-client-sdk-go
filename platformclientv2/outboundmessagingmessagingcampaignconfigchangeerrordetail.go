package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Outboundmessagingmessagingcampaignconfigchangeerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangeerrordetail

	

	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		VarError: u.VarError,
		
		Details: u.Details,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

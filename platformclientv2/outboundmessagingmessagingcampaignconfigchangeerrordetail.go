package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangeerrordetail
type Outboundmessagingmessagingcampaignconfigchangeerrordetail struct { 
	// VarError - The name of the error code.
	VarError *string `json:"error,omitempty"`


	// Details - The additional information regarding the error message.
	Details *string `json:"details,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangeerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangeerrordetail
	
	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		VarError: o.VarError,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangeerrordetail) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangeerrordetailMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangeerrordetailMap)
	if err != nil {
		return err
	}
	
	if VarError, ok := OutboundmessagingmessagingcampaignconfigchangeerrordetailMap["error"].(string); ok {
		o.VarError = &VarError
	}
    
	if Details, ok := OutboundmessagingmessagingcampaignconfigchangeerrordetailMap["details"].(string); ok {
		o.Details = &Details
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

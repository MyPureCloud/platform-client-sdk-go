package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangeresponseref - A reference for a Response
type Outboundmessagingmessagingcampaignconfigchangeresponseref struct { 
	// Id - The unique response id
	Id *string `json:"id,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangeresponseref) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangeresponseref
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangeresponseref) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangeresponserefMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangeresponserefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutboundmessagingmessagingcampaignconfigchangeresponserefMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangeresponseref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

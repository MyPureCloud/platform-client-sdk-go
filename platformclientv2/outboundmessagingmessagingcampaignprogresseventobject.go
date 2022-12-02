package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignprogresseventobject
type Outboundmessagingmessagingcampaignprogresseventobject struct { }

func (o *Outboundmessagingmessagingcampaignprogresseventobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignprogresseventobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignprogresseventobject) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignprogresseventobjectMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignprogresseventobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignprogresseventobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

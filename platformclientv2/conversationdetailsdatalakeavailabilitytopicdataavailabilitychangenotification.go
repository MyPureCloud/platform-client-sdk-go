package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
type Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification struct { }

func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) UnmarshalJSON(b []byte) error {
	var ConversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotificationMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdataavailabilitychangenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

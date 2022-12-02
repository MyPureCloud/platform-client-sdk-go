package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignscheduleconfigchangeobject
type Dialercampaignscheduleconfigchangeobject struct { }

func (o *Dialercampaignscheduleconfigchangeobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignscheduleconfigchangeobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignscheduleconfigchangeobject) UnmarshalJSON(b []byte) error {
	var DialercampaignscheduleconfigchangeobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignscheduleconfigchangeobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangeobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangeobject
type Dialercontactlistfilterconfigchangeobject struct { }

func (o *Dialercontactlistfilterconfigchangeobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangeobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangeobject) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangeobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangeobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangeobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangeobject
type Dialercontactlistconfigchangeobject struct { }

func (o *Dialercontactlistconfigchangeobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangeobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistconfigchangeobject) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangeobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangeobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangeobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerdnclistconfigchangeobject
type Dialerdnclistconfigchangeobject struct { }

func (o *Dialerdnclistconfigchangeobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerdnclistconfigchangeobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dialerdnclistconfigchangeobject) UnmarshalJSON(b []byte) error {
	var DialerdnclistconfigchangeobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DialerdnclistconfigchangeobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangeobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

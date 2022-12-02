package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistdownloadreadyobject
type Dnclistdownloadreadyobject struct { }

func (o *Dnclistdownloadreadyobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistdownloadreadyobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dnclistdownloadreadyobject) UnmarshalJSON(b []byte) error {
	var DnclistdownloadreadyobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistdownloadreadyobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclistdownloadreadyobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

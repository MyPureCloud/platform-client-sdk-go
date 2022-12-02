package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistdownloadreadyobject
type Contactlistdownloadreadyobject struct { }

func (o *Contactlistdownloadreadyobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistdownloadreadyobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Contactlistdownloadreadyobject) UnmarshalJSON(b []byte) error {
	var ContactlistdownloadreadyobjectMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistdownloadreadyobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistdownloadreadyobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistimportstatusobject
type Contactlistimportstatusobject struct { }

func (o *Contactlistimportstatusobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistimportstatusobject
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Contactlistimportstatusobject) UnmarshalJSON(b []byte) error {
	var ContactlistimportstatusobjectMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistimportstatusobjectMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistimportstatusobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

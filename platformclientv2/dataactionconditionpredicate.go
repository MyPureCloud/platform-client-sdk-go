package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataactionconditionpredicate
type Dataactionconditionpredicate struct { }

func (o *Dataactionconditionpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataactionconditionpredicate
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Dataactionconditionpredicate) UnmarshalJSON(b []byte) error {
	var DataactionconditionpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DataactionconditionpredicateMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

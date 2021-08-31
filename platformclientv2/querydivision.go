package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Querydivision
type Querydivision struct { }

func (o *Querydivision) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Querydivision
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Querydivision) UnmarshalJSON(b []byte) error {
	var QuerydivisionMap map[string]interface{}
	err := json.Unmarshal(b, &QuerydivisionMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Querydivision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

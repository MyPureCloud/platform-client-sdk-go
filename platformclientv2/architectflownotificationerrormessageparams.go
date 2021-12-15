package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationerrormessageparams - The error message params, if the action failed
type Architectflownotificationerrormessageparams struct { }

func (o *Architectflownotificationerrormessageparams) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationerrormessageparams
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationerrormessageparams) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationerrormessageparamsMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationerrormessageparamsMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

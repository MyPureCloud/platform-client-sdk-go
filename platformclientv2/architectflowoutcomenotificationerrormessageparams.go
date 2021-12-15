package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationerrormessageparams - The error message params, if the action failed
type Architectflowoutcomenotificationerrormessageparams struct { }

func (o *Architectflowoutcomenotificationerrormessageparams) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationerrormessageparams
	
	return json.Marshal(&struct { *Alias
	}{ Alias:    (*Alias)(o),
	})
}

func (o *Architectflowoutcomenotificationerrormessageparams) UnmarshalJSON(b []byte) error {
	var ArchitectflowoutcomenotificationerrormessageparamsMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowoutcomenotificationerrormessageparamsMap)
	if err != nil {
		return err
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

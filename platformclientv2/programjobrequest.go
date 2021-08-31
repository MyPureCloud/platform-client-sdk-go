package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programjobrequest
type Programjobrequest struct { 
	// ProgramIds - The ids of the programs used for this job
	ProgramIds *[]string `json:"programIds,omitempty"`

}

func (o *Programjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programjobrequest
	
	return json.Marshal(&struct { 
		ProgramIds *[]string `json:"programIds,omitempty"`
		*Alias
	}{ 
		ProgramIds: o.ProgramIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Programjobrequest) UnmarshalJSON(b []byte) error {
	var ProgramjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ProgramjobrequestMap)
	if err != nil {
		return err
	}
	
	if ProgramIds, ok := ProgramjobrequestMap["programIds"].([]interface{}); ok {
		ProgramIdsString, _ := json.Marshal(ProgramIds)
		json.Unmarshal(ProgramIdsString, &o.ProgramIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Programjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

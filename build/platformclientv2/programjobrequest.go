package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Programjobrequest
type Programjobrequest struct { 
	// ProgramIds - The ids of the programs used for this job
	ProgramIds *[]string `json:"programIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Programjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

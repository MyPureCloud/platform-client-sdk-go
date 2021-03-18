package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatusroutingstatus
type Userroutingstatusroutingstatus struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutingstatusroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

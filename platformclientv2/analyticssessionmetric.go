package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssessionmetric
type Analyticssessionmetric struct { 
	// EmitDate - Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EmitDate *time.Time `json:"emitDate,omitempty"`


	// Name - Unique name of this metric
	Name *string `json:"name,omitempty"`


	// Value - The metric value
	Value *int `json:"value,omitempty"`

}

func (u *Analyticssessionmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticssessionmetric

	
	EmitDate := new(string)
	if u.EmitDate != nil {
		
		*EmitDate = timeutil.Strftime(u.EmitDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EmitDate = nil
	}
	

	return json.Marshal(&struct { 
		EmitDate *string `json:"emitDate,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Value *int `json:"value,omitempty"`
		*Alias
	}{ 
		EmitDate: EmitDate,
		
		Name: u.Name,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticssessionmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

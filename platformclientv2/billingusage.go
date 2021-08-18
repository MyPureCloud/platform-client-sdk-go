package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingusage
type Billingusage struct { 
	// Name - Identifies the billable usage.
	Name *string `json:"name,omitempty"`


	// TotalUsage - The total amount of usage, expressed as a decimal number in string format.
	TotalUsage *string `json:"totalUsage,omitempty"`


	// Resources - The resources for which usage was observed (e.g. license users, devices).
	Resources *[]Billingusageresource `json:"resources,omitempty"`

}

func (u *Billingusage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Billingusage

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		TotalUsage *string `json:"totalUsage,omitempty"`
		
		Resources *[]Billingusageresource `json:"resources,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		TotalUsage: u.TotalUsage,
		
		Resources: u.Resources,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Billingusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

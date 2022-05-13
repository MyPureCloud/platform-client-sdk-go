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

func (o *Billingusage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Billingusage
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		TotalUsage *string `json:"totalUsage,omitempty"`
		
		Resources *[]Billingusageresource `json:"resources,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		TotalUsage: o.TotalUsage,
		
		Resources: o.Resources,
		Alias:    (*Alias)(o),
	})
}

func (o *Billingusage) UnmarshalJSON(b []byte) error {
	var BillingusageMap map[string]interface{}
	err := json.Unmarshal(b, &BillingusageMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BillingusageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TotalUsage, ok := BillingusageMap["totalUsage"].(string); ok {
		o.TotalUsage = &TotalUsage
	}
    
	if Resources, ok := BillingusageMap["resources"].([]interface{}); ok {
		ResourcesString, _ := json.Marshal(Resources)
		json.Unmarshal(ResourcesString, &o.Resources)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Billingusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingusageresource
type Billingusageresource struct { 
	// Name - Identifies the resource (e.g. license user, device).
	Name *string `json:"name,omitempty"`


	// Date - The date that the usage was first observed by the billing subsystem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`

}

func (o *Billingusageresource) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Billingusageresource
	
	Date := new(string)
	if o.Date != nil {
		
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Date *string `json:"date,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Date: Date,
		Alias:    (*Alias)(o),
	})
}

func (o *Billingusageresource) UnmarshalJSON(b []byte) error {
	var BillingusageresourceMap map[string]interface{}
	err := json.Unmarshal(b, &BillingusageresourceMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BillingusageresourceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateString, ok := BillingusageresourceMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateString)
		o.Date = &Date
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Billingusageresource) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

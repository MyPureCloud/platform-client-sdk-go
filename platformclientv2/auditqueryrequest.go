package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryrequest
type Auditqueryrequest struct { 
	// Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// ServiceName - Name of the service to query audits for.
	ServiceName *string `json:"serviceName,omitempty"`


	// Filters - Additional filters for the query.
	Filters *[]Auditqueryfilter `json:"filters,omitempty"`


	// Sort - Sort parameter for the query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`

}

func (u *Auditqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryrequest

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Filters *[]Auditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		ServiceName: u.ServiceName,
		
		Filters: u.Filters,
		
		Sort: u.Sort,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

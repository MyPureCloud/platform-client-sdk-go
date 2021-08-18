package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditrealtimequeryrequest
type Auditrealtimequeryrequest struct { 
	// Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// ServiceName - Name of the service to query audits for.
	ServiceName *string `json:"serviceName,omitempty"`


	// Filters - Additional filters for the query.
	Filters *[]Auditqueryfilter `json:"filters,omitempty"`


	// Sort - Sort parameter for the query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`


	// PageNumber - Page number
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize - Page size
	PageSize *int `json:"pageSize,omitempty"`

}

func (u *Auditrealtimequeryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditrealtimequeryrequest

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Filters *[]Auditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		ServiceName: u.ServiceName,
		
		Filters: u.Filters,
		
		Sort: u.Sort,
		
		PageNumber: u.PageNumber,
		
		PageSize: u.PageSize,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditrealtimequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

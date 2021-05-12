package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Auditrealtimequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

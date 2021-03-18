package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryexecutionstatusresponse
type Auditqueryexecutionstatusresponse struct { 
	// Id - Id of the audit query execution request.
	Id *string `json:"id,omitempty"`


	// State - Status of the audit query execution request.
	State *string `json:"state,omitempty"`


	// StartDate - Start date and time of the audit query execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// Interval - Interval for the audit query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// ServiceName - Service name for the audit query.
	ServiceName *string `json:"serviceName,omitempty"`


	// Filters - Filters for the audit query.
	Filters *[]Auditqueryfilter `json:"filters,omitempty"`


	// Sort - Sort parameter for the audit query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

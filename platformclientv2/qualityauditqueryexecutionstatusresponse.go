package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditqueryexecutionstatusresponse
type Qualityauditqueryexecutionstatusresponse struct { 
	// Id - Id of the audit query execution request.
	Id *string `json:"id,omitempty"`


	// State - Status of the audit query execution request.
	State *string `json:"state,omitempty"`


	// DateStart - Start date and time of the audit query execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// Interval - Interval for the audit query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Filters - Filters for the audit query.
	Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`


	// Sort - Sort parameter for the audit query.
	Sort *[]Auditquerysort `json:"sort,omitempty"`

}

func (u *Qualityauditqueryexecutionstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qualityauditqueryexecutionstatusresponse

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		Filters *[]Qualityauditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		State: u.State,
		
		DateStart: DateStart,
		
		Interval: u.Interval,
		
		Filters: u.Filters,
		
		Sort: u.Sort,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Qualityauditqueryexecutionstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

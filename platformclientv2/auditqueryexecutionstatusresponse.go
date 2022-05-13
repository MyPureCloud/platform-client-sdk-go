package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (o *Auditqueryexecutionstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryexecutionstatusresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Filters *[]Auditqueryfilter `json:"filters,omitempty"`
		
		Sort *[]Auditquerysort `json:"sort,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		StartDate: StartDate,
		
		Interval: o.Interval,
		
		ServiceName: o.ServiceName,
		
		Filters: o.Filters,
		
		Sort: o.Sort,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryexecutionstatusresponse) UnmarshalJSON(b []byte) error {
	var AuditqueryexecutionstatusresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryexecutionstatusresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuditqueryexecutionstatusresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := AuditqueryexecutionstatusresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if startDateString, ok := AuditqueryexecutionstatusresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if Interval, ok := AuditqueryexecutionstatusresponseMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if ServiceName, ok := AuditqueryexecutionstatusresponseMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
    
	if Filters, ok := AuditqueryexecutionstatusresponseMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Sort, ok := AuditqueryexecutionstatusresponseMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

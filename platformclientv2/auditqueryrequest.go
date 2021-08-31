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

func (o *Auditqueryrequest) MarshalJSON() ([]byte, error) {
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
		Interval: o.Interval,
		
		ServiceName: o.ServiceName,
		
		Filters: o.Filters,
		
		Sort: o.Sort,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryrequest) UnmarshalJSON(b []byte) error {
	var AuditqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := AuditqueryrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if ServiceName, ok := AuditqueryrequestMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
	
	if Filters, ok := AuditqueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Sort, ok := AuditqueryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditrealtimequeryrequest
type Auditrealtimequeryrequest struct { 
	// Interval - Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ssZ/YYYY-MM-DDThh:mm:ssZ
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

func (o *Auditrealtimequeryrequest) MarshalJSON() ([]byte, error) {
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
		Interval: o.Interval,
		
		ServiceName: o.ServiceName,
		
		Filters: o.Filters,
		
		Sort: o.Sort,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditrealtimequeryrequest) UnmarshalJSON(b []byte) error {
	var AuditrealtimequeryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AuditrealtimequeryrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := AuditrealtimequeryrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if ServiceName, ok := AuditrealtimequeryrequestMap["serviceName"].(string); ok {
		o.ServiceName = &ServiceName
	}
    
	if Filters, ok := AuditrealtimequeryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Sort, ok := AuditrealtimequeryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if PageNumber, ok := AuditrealtimequeryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := AuditrealtimequeryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditrealtimequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

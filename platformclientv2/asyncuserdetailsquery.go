package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncuserdetailsquery
type Asyncuserdetailsquery struct { 
	// Interval - Specifies the date and time range of data being queried. Conversations MUST have started within this time range to potentially be included within the result set. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// UserFilters - Filters that target the users to retrieve data for
	UserFilters *[]Userdetailqueryfilter `json:"userFilters,omitempty"`


	// PresenceFilters - Filters that target system and organization presence-level data
	PresenceFilters *[]Presencedetailqueryfilter `json:"presenceFilters,omitempty"`


	// RoutingStatusFilters - Filters that target agent routing status-level data
	RoutingStatusFilters *[]Routingstatusdetailqueryfilter `json:"routingStatusFilters,omitempty"`


	// Order - Sort the result set in ascending/descending order. Default is ascending
	Order *string `json:"order,omitempty"`


	// Limit - Specify number of results to be returned
	Limit *int `json:"limit,omitempty"`

}

func (o *Asyncuserdetailsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncuserdetailsquery
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		UserFilters *[]Userdetailqueryfilter `json:"userFilters,omitempty"`
		
		PresenceFilters *[]Presencedetailqueryfilter `json:"presenceFilters,omitempty"`
		
		RoutingStatusFilters *[]Routingstatusdetailqueryfilter `json:"routingStatusFilters,omitempty"`
		
		Order *string `json:"order,omitempty"`
		
		Limit *int `json:"limit,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		UserFilters: o.UserFilters,
		
		PresenceFilters: o.PresenceFilters,
		
		RoutingStatusFilters: o.RoutingStatusFilters,
		
		Order: o.Order,
		
		Limit: o.Limit,
		Alias:    (*Alias)(o),
	})
}

func (o *Asyncuserdetailsquery) UnmarshalJSON(b []byte) error {
	var AsyncuserdetailsqueryMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncuserdetailsqueryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := AsyncuserdetailsqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if UserFilters, ok := AsyncuserdetailsqueryMap["userFilters"].([]interface{}); ok {
		UserFiltersString, _ := json.Marshal(UserFilters)
		json.Unmarshal(UserFiltersString, &o.UserFilters)
	}
	
	if PresenceFilters, ok := AsyncuserdetailsqueryMap["presenceFilters"].([]interface{}); ok {
		PresenceFiltersString, _ := json.Marshal(PresenceFilters)
		json.Unmarshal(PresenceFiltersString, &o.PresenceFilters)
	}
	
	if RoutingStatusFilters, ok := AsyncuserdetailsqueryMap["routingStatusFilters"].([]interface{}); ok {
		RoutingStatusFiltersString, _ := json.Marshal(RoutingStatusFilters)
		json.Unmarshal(RoutingStatusFiltersString, &o.RoutingStatusFilters)
	}
	
	if Order, ok := AsyncuserdetailsqueryMap["order"].(string); ok {
		o.Order = &Order
	}
	
	if Limit, ok := AsyncuserdetailsqueryMap["limit"].(float64); ok {
		LimitInt := int(Limit)
		o.Limit = &LimitInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncuserdetailsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

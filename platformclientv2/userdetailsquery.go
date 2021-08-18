package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailsquery
type Userdetailsquery struct { 
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


	// PresenceAggregations - Include faceted search and aggregate roll-ups of presence data in your search results. This does not function as a filter, but rather, summary data about the presence results matching your filters
	PresenceAggregations *[]Analyticsqueryaggregation `json:"presenceAggregations,omitempty"`


	// RoutingStatusAggregations - Include faceted search and aggregate roll-ups of agent routing status data in your search results. This does not function as a filter, but rather, summary data about the agent routing status results matching your filters
	RoutingStatusAggregations *[]Analyticsqueryaggregation `json:"routingStatusAggregations,omitempty"`


	// Paging - Page size and number to control iterating through large result sets. Default page size is 25
	Paging *Pagingspec `json:"paging,omitempty"`

}

func (u *Userdetailsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userdetailsquery

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		UserFilters *[]Userdetailqueryfilter `json:"userFilters,omitempty"`
		
		PresenceFilters *[]Presencedetailqueryfilter `json:"presenceFilters,omitempty"`
		
		RoutingStatusFilters *[]Routingstatusdetailqueryfilter `json:"routingStatusFilters,omitempty"`
		
		Order *string `json:"order,omitempty"`
		
		PresenceAggregations *[]Analyticsqueryaggregation `json:"presenceAggregations,omitempty"`
		
		RoutingStatusAggregations *[]Analyticsqueryaggregation `json:"routingStatusAggregations,omitempty"`
		
		Paging *Pagingspec `json:"paging,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		UserFilters: u.UserFilters,
		
		PresenceFilters: u.PresenceFilters,
		
		RoutingStatusFilters: u.RoutingStatusFilters,
		
		Order: u.Order,
		
		PresenceAggregations: u.PresenceAggregations,
		
		RoutingStatusAggregations: u.RoutingStatusAggregations,
		
		Paging: u.Paging,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userdetailsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

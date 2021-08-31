package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationaggregationquery
type Conversationaggregationquery struct { 
	// Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Granularity - Granularity aggregates metrics into subpartitions within the time interval specified. The default granularity is the same duration as the interval. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Granularity *string `json:"granularity,omitempty"`


	// TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`


	// GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
	GroupBy *[]string `json:"groupBy,omitempty"`


	// Filter - Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
	Filter *Conversationaggregatequeryfilter `json:"filter,omitempty"`


	// Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
	Metrics *[]string `json:"metrics,omitempty"`


	// FlattenMultivaluedDimensions - Flattens any multivalued dimensions used in response groups (e.g. ['a','b','c']->'a,b,c')
	FlattenMultivaluedDimensions *bool `json:"flattenMultivaluedDimensions,omitempty"`


	// Views - Custom derived metric views
	Views *[]Conversationaggregationview `json:"views,omitempty"`


	// AlternateTimeDimension - Dimension to use as the alternative timestamp for data in the aggregate.  Choosing \"eventTime\" uses the actual time of the data event.
	AlternateTimeDimension *string `json:"alternateTimeDimension,omitempty"`

}

func (o *Conversationaggregationquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationaggregationquery
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Filter *Conversationaggregatequeryfilter `json:"filter,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		FlattenMultivaluedDimensions *bool `json:"flattenMultivaluedDimensions,omitempty"`
		
		Views *[]Conversationaggregationview `json:"views,omitempty"`
		
		AlternateTimeDimension *string `json:"alternateTimeDimension,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Granularity: o.Granularity,
		
		TimeZone: o.TimeZone,
		
		GroupBy: o.GroupBy,
		
		Filter: o.Filter,
		
		Metrics: o.Metrics,
		
		FlattenMultivaluedDimensions: o.FlattenMultivaluedDimensions,
		
		Views: o.Views,
		
		AlternateTimeDimension: o.AlternateTimeDimension,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationaggregationquery) UnmarshalJSON(b []byte) error {
	var ConversationaggregationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationaggregationqueryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := ConversationaggregationqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Granularity, ok := ConversationaggregationqueryMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
	
	if TimeZone, ok := ConversationaggregationqueryMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if GroupBy, ok := ConversationaggregationqueryMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Filter, ok := ConversationaggregationqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Metrics, ok := ConversationaggregationqueryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if FlattenMultivaluedDimensions, ok := ConversationaggregationqueryMap["flattenMultivaluedDimensions"].(bool); ok {
		o.FlattenMultivaluedDimensions = &FlattenMultivaluedDimensions
	}
	
	if Views, ok := ConversationaggregationqueryMap["views"].([]interface{}); ok {
		ViewsString, _ := json.Marshal(Views)
		json.Unmarshal(ViewsString, &o.Views)
	}
	
	if AlternateTimeDimension, ok := ConversationaggregationqueryMap["alternateTimeDimension"].(string); ok {
		o.AlternateTimeDimension = &AlternateTimeDimension
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationaggregationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

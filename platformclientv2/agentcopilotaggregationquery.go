package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentcopilotaggregationquery
type Agentcopilotaggregationquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Granularity - Granularity aggregates metrics into subpartitions within the time interval specified. The default granularity is the same duration as the interval. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Granularity *string `json:"granularity,omitempty"`

	// TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`

	// GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
	GroupBy *[]string `json:"groupBy,omitempty"`

	// Filter - Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
	Filter *Agentcopilotaggregatequeryfilter `json:"filter,omitempty"`

	// Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
	Metrics *[]string `json:"metrics,omitempty"`

	// FlattenMultivaluedDimensions - Flattens any multivalued dimensions used in response groups (e.g. ['a','b','c']->'a,b,c')
	FlattenMultivaluedDimensions *bool `json:"flattenMultivaluedDimensions,omitempty"`

	// Views - Custom derived metric views
	Views *[]Agentcopilotaggregationview `json:"views,omitempty"`

	// AlternateTimeDimension - Dimension to use as the alternative timestamp for data in the aggregate.  Choosing \"eventTime\" uses the actual time of the data event.
	AlternateTimeDimension *string `json:"alternateTimeDimension,omitempty"`

	// QueryType - Query type to use. Use groupBy for all matching results, and topN/bottomN for N results ordered by the sortMetric. Default is groupBy.
	QueryType *string `json:"queryType,omitempty"`

	// SortMetric - Required when requesting multiple metrics. Only applicable for topN/bottomN query type.
	SortMetric *Agentcopilotaggregationsort `json:"sortMetric,omitempty"`

	// Limit - How many results you want in an ordered list. Only applicable for topN/bottomN query type.
	Limit *int `json:"limit,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentcopilotaggregationquery) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Agentcopilotaggregationquery) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentcopilotaggregationquery
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Filter *Agentcopilotaggregatequeryfilter `json:"filter,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		FlattenMultivaluedDimensions *bool `json:"flattenMultivaluedDimensions,omitempty"`
		
		Views *[]Agentcopilotaggregationview `json:"views,omitempty"`
		
		AlternateTimeDimension *string `json:"alternateTimeDimension,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SortMetric *Agentcopilotaggregationsort `json:"sortMetric,omitempty"`
		
		Limit *int `json:"limit,omitempty"`
		Alias
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
		
		QueryType: o.QueryType,
		
		SortMetric: o.SortMetric,
		
		Limit: o.Limit,
		Alias:    (Alias)(o),
	})
}

func (o *Agentcopilotaggregationquery) UnmarshalJSON(b []byte) error {
	var AgentcopilotaggregationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &AgentcopilotaggregationqueryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := AgentcopilotaggregationqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Granularity, ok := AgentcopilotaggregationqueryMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if TimeZone, ok := AgentcopilotaggregationqueryMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if GroupBy, ok := AgentcopilotaggregationqueryMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Filter, ok := AgentcopilotaggregationqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Metrics, ok := AgentcopilotaggregationqueryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if FlattenMultivaluedDimensions, ok := AgentcopilotaggregationqueryMap["flattenMultivaluedDimensions"].(bool); ok {
		o.FlattenMultivaluedDimensions = &FlattenMultivaluedDimensions
	}
    
	if Views, ok := AgentcopilotaggregationqueryMap["views"].([]interface{}); ok {
		ViewsString, _ := json.Marshal(Views)
		json.Unmarshal(ViewsString, &o.Views)
	}
	
	if AlternateTimeDimension, ok := AgentcopilotaggregationqueryMap["alternateTimeDimension"].(string); ok {
		o.AlternateTimeDimension = &AlternateTimeDimension
	}
    
	if QueryType, ok := AgentcopilotaggregationqueryMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SortMetric, ok := AgentcopilotaggregationqueryMap["sortMetric"].(map[string]interface{}); ok {
		SortMetricString, _ := json.Marshal(SortMetric)
		json.Unmarshal(SortMetricString, &o.SortMetric)
	}
	
	if Limit, ok := AgentcopilotaggregationqueryMap["limit"].(float64); ok {
		LimitInt := int(Limit)
		o.Limit = &LimitInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentcopilotaggregationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

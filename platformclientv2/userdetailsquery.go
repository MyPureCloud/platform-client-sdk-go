package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userdetailsquery
type Userdetailsquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userdetailsquery) SetField(field string, fieldValue interface{}) {
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

func (o Userdetailsquery) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
		Alias
	}{ 
		Interval: o.Interval,
		
		UserFilters: o.UserFilters,
		
		PresenceFilters: o.PresenceFilters,
		
		RoutingStatusFilters: o.RoutingStatusFilters,
		
		Order: o.Order,
		
		PresenceAggregations: o.PresenceAggregations,
		
		RoutingStatusAggregations: o.RoutingStatusAggregations,
		
		Paging: o.Paging,
		Alias:    (Alias)(o),
	})
}

func (o *Userdetailsquery) UnmarshalJSON(b []byte) error {
	var UserdetailsqueryMap map[string]interface{}
	err := json.Unmarshal(b, &UserdetailsqueryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := UserdetailsqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if UserFilters, ok := UserdetailsqueryMap["userFilters"].([]interface{}); ok {
		UserFiltersString, _ := json.Marshal(UserFilters)
		json.Unmarshal(UserFiltersString, &o.UserFilters)
	}
	
	if PresenceFilters, ok := UserdetailsqueryMap["presenceFilters"].([]interface{}); ok {
		PresenceFiltersString, _ := json.Marshal(PresenceFilters)
		json.Unmarshal(PresenceFiltersString, &o.PresenceFilters)
	}
	
	if RoutingStatusFilters, ok := UserdetailsqueryMap["routingStatusFilters"].([]interface{}); ok {
		RoutingStatusFiltersString, _ := json.Marshal(RoutingStatusFilters)
		json.Unmarshal(RoutingStatusFiltersString, &o.RoutingStatusFilters)
	}
	
	if Order, ok := UserdetailsqueryMap["order"].(string); ok {
		o.Order = &Order
	}
    
	if PresenceAggregations, ok := UserdetailsqueryMap["presenceAggregations"].([]interface{}); ok {
		PresenceAggregationsString, _ := json.Marshal(PresenceAggregations)
		json.Unmarshal(PresenceAggregationsString, &o.PresenceAggregations)
	}
	
	if RoutingStatusAggregations, ok := UserdetailsqueryMap["routingStatusAggregations"].([]interface{}); ok {
		RoutingStatusAggregationsString, _ := json.Marshal(RoutingStatusAggregations)
		json.Unmarshal(RoutingStatusAggregationsString, &o.RoutingStatusAggregations)
	}
	
	if Paging, ok := UserdetailsqueryMap["paging"].(map[string]interface{}); ok {
		PagingString, _ := json.Marshal(Paging)
		json.Unmarshal(PagingString, &o.Paging)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userdetailsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

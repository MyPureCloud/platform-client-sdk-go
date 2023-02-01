package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncconversationquery
type Asyncconversationquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationFilters - Filters that target conversation-level data
	ConversationFilters *[]Conversationdetailqueryfilter `json:"conversationFilters,omitempty"`

	// SegmentFilters - Filters that target individual segments within a conversation
	SegmentFilters *[]Segmentdetailqueryfilter `json:"segmentFilters,omitempty"`

	// EvaluationFilters - Filters that target evaluations
	EvaluationFilters *[]Evaluationdetailqueryfilter `json:"evaluationFilters,omitempty"`

	// SurveyFilters - Filters that target surveys
	SurveyFilters *[]Surveydetailqueryfilter `json:"surveyFilters,omitempty"`

	// ResolutionFilters - Filters that target resolutions
	ResolutionFilters *[]Resolutiondetailqueryfilter `json:"resolutionFilters,omitempty"`

	// Order - Sort the result set in ascending/descending order. Default is ascending
	Order *string `json:"order,omitempty"`

	// OrderBy - Specify which data element within the result set to use for sorting. The options  to use as a basis for sorting the results: conversationStart, segmentStart, and segmentEnd. If not specified, the default is conversationStart
	OrderBy *string `json:"orderBy,omitempty"`

	// Interval - Specifies the date and time range of data being queried. Results will include all conversations that had activity during the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Limit - Specify number of results to be returned
	Limit *int `json:"limit,omitempty"`

	// StartOfDayIntervalMatching - Add a filter to only include conversations that started after the beginning of the interval start date (UTC)
	StartOfDayIntervalMatching *bool `json:"startOfDayIntervalMatching,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Asyncconversationquery) SetField(field string, fieldValue interface{}) {
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

func (o Asyncconversationquery) MarshalJSON() ([]byte, error) {
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
	type Alias Asyncconversationquery
	
	return json.Marshal(&struct { 
		ConversationFilters *[]Conversationdetailqueryfilter `json:"conversationFilters,omitempty"`
		
		SegmentFilters *[]Segmentdetailqueryfilter `json:"segmentFilters,omitempty"`
		
		EvaluationFilters *[]Evaluationdetailqueryfilter `json:"evaluationFilters,omitempty"`
		
		SurveyFilters *[]Surveydetailqueryfilter `json:"surveyFilters,omitempty"`
		
		ResolutionFilters *[]Resolutiondetailqueryfilter `json:"resolutionFilters,omitempty"`
		
		Order *string `json:"order,omitempty"`
		
		OrderBy *string `json:"orderBy,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		Limit *int `json:"limit,omitempty"`
		
		StartOfDayIntervalMatching *bool `json:"startOfDayIntervalMatching,omitempty"`
		Alias
	}{ 
		ConversationFilters: o.ConversationFilters,
		
		SegmentFilters: o.SegmentFilters,
		
		EvaluationFilters: o.EvaluationFilters,
		
		SurveyFilters: o.SurveyFilters,
		
		ResolutionFilters: o.ResolutionFilters,
		
		Order: o.Order,
		
		OrderBy: o.OrderBy,
		
		Interval: o.Interval,
		
		Limit: o.Limit,
		
		StartOfDayIntervalMatching: o.StartOfDayIntervalMatching,
		Alias:    (Alias)(o),
	})
}

func (o *Asyncconversationquery) UnmarshalJSON(b []byte) error {
	var AsyncconversationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncconversationqueryMap)
	if err != nil {
		return err
	}
	
	if ConversationFilters, ok := AsyncconversationqueryMap["conversationFilters"].([]interface{}); ok {
		ConversationFiltersString, _ := json.Marshal(ConversationFilters)
		json.Unmarshal(ConversationFiltersString, &o.ConversationFilters)
	}
	
	if SegmentFilters, ok := AsyncconversationqueryMap["segmentFilters"].([]interface{}); ok {
		SegmentFiltersString, _ := json.Marshal(SegmentFilters)
		json.Unmarshal(SegmentFiltersString, &o.SegmentFilters)
	}
	
	if EvaluationFilters, ok := AsyncconversationqueryMap["evaluationFilters"].([]interface{}); ok {
		EvaluationFiltersString, _ := json.Marshal(EvaluationFilters)
		json.Unmarshal(EvaluationFiltersString, &o.EvaluationFilters)
	}
	
	if SurveyFilters, ok := AsyncconversationqueryMap["surveyFilters"].([]interface{}); ok {
		SurveyFiltersString, _ := json.Marshal(SurveyFilters)
		json.Unmarshal(SurveyFiltersString, &o.SurveyFilters)
	}
	
	if ResolutionFilters, ok := AsyncconversationqueryMap["resolutionFilters"].([]interface{}); ok {
		ResolutionFiltersString, _ := json.Marshal(ResolutionFilters)
		json.Unmarshal(ResolutionFiltersString, &o.ResolutionFilters)
	}
	
	if Order, ok := AsyncconversationqueryMap["order"].(string); ok {
		o.Order = &Order
	}
    
	if OrderBy, ok := AsyncconversationqueryMap["orderBy"].(string); ok {
		o.OrderBy = &OrderBy
	}
    
	if Interval, ok := AsyncconversationqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Limit, ok := AsyncconversationqueryMap["limit"].(float64); ok {
		LimitInt := int(Limit)
		o.Limit = &LimitInt
	}
	
	if StartOfDayIntervalMatching, ok := AsyncconversationqueryMap["startOfDayIntervalMatching"].(bool); ok {
		o.StartOfDayIntervalMatching = &StartOfDayIntervalMatching
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncconversationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

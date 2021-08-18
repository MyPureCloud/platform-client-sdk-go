package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationquery
type Conversationquery struct { 
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


	// Interval - Specifies the date and time range of data being queried. Results will only include conversations that started on a day touched by the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// Aggregations - Include faceted search and aggregate roll-ups describing your search results. This does not function as a filter, but rather, summary data about the data matching your filters
	Aggregations *[]Analyticsqueryaggregation `json:"aggregations,omitempty"`


	// Paging - Page size and number to control iterating through large result sets. Default page size is 25
	Paging *Pagingspec `json:"paging,omitempty"`

}

func (u *Conversationquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationquery

	

	return json.Marshal(&struct { 
		ConversationFilters *[]Conversationdetailqueryfilter `json:"conversationFilters,omitempty"`
		
		SegmentFilters *[]Segmentdetailqueryfilter `json:"segmentFilters,omitempty"`
		
		EvaluationFilters *[]Evaluationdetailqueryfilter `json:"evaluationFilters,omitempty"`
		
		SurveyFilters *[]Surveydetailqueryfilter `json:"surveyFilters,omitempty"`
		
		ResolutionFilters *[]Resolutiondetailqueryfilter `json:"resolutionFilters,omitempty"`
		
		Order *string `json:"order,omitempty"`
		
		OrderBy *string `json:"orderBy,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		Aggregations *[]Analyticsqueryaggregation `json:"aggregations,omitempty"`
		
		Paging *Pagingspec `json:"paging,omitempty"`
		*Alias
	}{ 
		ConversationFilters: u.ConversationFilters,
		
		SegmentFilters: u.SegmentFilters,
		
		EvaluationFilters: u.EvaluationFilters,
		
		SurveyFilters: u.SurveyFilters,
		
		ResolutionFilters: u.ResolutionFilters,
		
		Order: u.Order,
		
		OrderBy: u.OrderBy,
		
		Interval: u.Interval,
		
		Aggregations: u.Aggregations,
		
		Paging: u.Paging,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

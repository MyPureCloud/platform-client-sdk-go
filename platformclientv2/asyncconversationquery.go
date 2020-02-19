package platformclientv2
import (
	"encoding/json"
)

// Asyncconversationquery
type Asyncconversationquery struct { 
	// Interval - Specifies the date and time range of data being queried. Results will include conversations that both started on a day touched by the interval AND either started, ended, or any activity during the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// ConversationFilters - Filters that target conversation-level data
	ConversationFilters *[]Conversationdetailqueryfilter `json:"conversationFilters,omitempty"`


	// SegmentFilters - Filters that target individual segments within a conversation
	SegmentFilters *[]Segmentdetailqueryfilter `json:"segmentFilters,omitempty"`


	// EvaluationFilters - Filters that target evaluations
	EvaluationFilters *[]Evaluationdetailqueryfilter `json:"evaluationFilters,omitempty"`


	// MediaEndpointStatFilters - Filters that target mediaEndpointStats
	MediaEndpointStatFilters *[]Mediaendpointstatdetailqueryfilter `json:"mediaEndpointStatFilters,omitempty"`


	// SurveyFilters - Filters that target surveys
	SurveyFilters *[]Surveydetailqueryfilter `json:"surveyFilters,omitempty"`


	// Order - Sort the result set in ascending/descending order. Default is ascending
	Order *string `json:"order,omitempty"`


	// OrderBy - Specify which data element within the result set to use for sorting. The options  to use as a basis for sorting the results: conversationStart, segmentStart, and segmentEnd. If not specified, the default is conversationStart
	OrderBy *string `json:"orderBy,omitempty"`


	// Limit - Specify number of results to be returned
	Limit *int32 `json:"limit,omitempty"`


	// StartOfDayIntervalMatching - Add a filter to only include conversations that started after the beginning of the interval start date (UTC)
	StartOfDayIntervalMatching *bool `json:"startOfDayIntervalMatching,omitempty"`

}

// String returns a JSON representation of the model
func (o *Asyncconversationquery) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

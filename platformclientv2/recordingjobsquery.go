package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobsquery
type Recordingjobsquery struct { 
	// Action - Operation to perform bulk task. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date.
	Action *string `json:"action,omitempty"`


	// ActionDate - The date when the action will be performed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActionDate *time.Time `json:"actionDate,omitempty"`


	// IntegrationId - IntegrationId to Access AWS S3 bucket for bulk recording exports. This field is required and used only for EXPORT action.
	IntegrationId *string `json:"integrationId,omitempty"`


	// IncludeRecordingsWithSensitiveData - Whether to include recordings with PCI DSS and/or PII data, default value = false 
	IncludeRecordingsWithSensitiveData *bool `json:"includeRecordingsWithSensitiveData,omitempty"`


	// IncludeScreenRecordings - Whether to include Screen recordings for the action, default value = true 
	IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`


	// ClearExport - For DELETE action, setting this to true will clear any pending exports for recordings. This field is not used for EXPORT action. Default value = false
	ClearExport *bool `json:"clearExport,omitempty"`


	// ConversationQuery - Conversation Query. Note: After the recording is created, it might take up to 48 hours for the recording to be included in the submitted job query.  This result depends on the analytics data lake job completion. See also: https://developer.genesys.cloud/analyticsdatamanagement/analytics/jobs/conversation-details-job#data-availability
	ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`

}

func (o *Recordingjobsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingjobsquery
	
	ActionDate := new(string)
	if o.ActionDate != nil {
		
		*ActionDate = timeutil.Strftime(o.ActionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActionDate = nil
	}
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		ActionDate *string `json:"actionDate,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		IncludeRecordingsWithSensitiveData *bool `json:"includeRecordingsWithSensitiveData,omitempty"`
		
		IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`
		
		ClearExport *bool `json:"clearExport,omitempty"`
		
		ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		ActionDate: ActionDate,
		
		IntegrationId: o.IntegrationId,
		
		IncludeRecordingsWithSensitiveData: o.IncludeRecordingsWithSensitiveData,
		
		IncludeScreenRecordings: o.IncludeScreenRecordings,
		
		ClearExport: o.ClearExport,
		
		ConversationQuery: o.ConversationQuery,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingjobsquery) UnmarshalJSON(b []byte) error {
	var RecordingjobsqueryMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingjobsqueryMap)
	if err != nil {
		return err
	}
	
	if Action, ok := RecordingjobsqueryMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if actionDateString, ok := RecordingjobsqueryMap["actionDate"].(string); ok {
		ActionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", actionDateString)
		o.ActionDate = &ActionDate
	}
	
	if IntegrationId, ok := RecordingjobsqueryMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if IncludeRecordingsWithSensitiveData, ok := RecordingjobsqueryMap["includeRecordingsWithSensitiveData"].(bool); ok {
		o.IncludeRecordingsWithSensitiveData = &IncludeRecordingsWithSensitiveData
	}
    
	if IncludeScreenRecordings, ok := RecordingjobsqueryMap["includeScreenRecordings"].(bool); ok {
		o.IncludeScreenRecordings = &IncludeScreenRecordings
	}
    
	if ClearExport, ok := RecordingjobsqueryMap["clearExport"].(bool); ok {
		o.ClearExport = &ClearExport
	}
    
	if ConversationQuery, ok := RecordingjobsqueryMap["conversationQuery"].(map[string]interface{}); ok {
		ConversationQueryString, _ := json.Marshal(ConversationQuery)
		json.Unmarshal(ConversationQueryString, &o.ConversationQuery)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

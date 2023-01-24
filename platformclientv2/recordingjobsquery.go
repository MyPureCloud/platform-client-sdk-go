package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobsquery
type Recordingjobsquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action - Operation to perform bulk task. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date.
	Action *string `json:"action,omitempty"`

	// ActionDate - The date when the action will be performed. If screenRecordingActionDate is also provided, this value is only used for non-screen recordings. Otherwise this value is used for all recordings. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActionDate *time.Time `json:"actionDate,omitempty"`

	// ActionAge - The number of days after each recording's creation date when the action will be performed. If screenRecordingActionAge is also provided, this value is only used for non-screen recordings. Otherwise this value is used for all recordings.
	ActionAge *int `json:"actionAge,omitempty"`

	// ScreenRecordingActionDate - The date when the action will be performed for screen recordings. If this is provided then includeScreenRecordings must be true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ScreenRecordingActionDate *time.Time `json:"screenRecordingActionDate,omitempty"`

	// ScreenRecordingActionAge - The number of days after each screen recording's creation date when the action will be performed. If this is provided then includeScreenRecordings must be true.
	ScreenRecordingActionAge *int `json:"screenRecordingActionAge,omitempty"`

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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingjobsquery) SetField(field string, fieldValue interface{}) {
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

func (o Recordingjobsquery) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ActionDate","ScreenRecordingActionDate", }
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
	type Alias Recordingjobsquery
	
	ActionDate := new(string)
	if o.ActionDate != nil {
		
		*ActionDate = timeutil.Strftime(o.ActionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActionDate = nil
	}
	
	ScreenRecordingActionDate := new(string)
	if o.ScreenRecordingActionDate != nil {
		
		*ScreenRecordingActionDate = timeutil.Strftime(o.ScreenRecordingActionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ScreenRecordingActionDate = nil
	}
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		ActionDate *string `json:"actionDate,omitempty"`
		
		ActionAge *int `json:"actionAge,omitempty"`
		
		ScreenRecordingActionDate *string `json:"screenRecordingActionDate,omitempty"`
		
		ScreenRecordingActionAge *int `json:"screenRecordingActionAge,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		IncludeRecordingsWithSensitiveData *bool `json:"includeRecordingsWithSensitiveData,omitempty"`
		
		IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`
		
		ClearExport *bool `json:"clearExport,omitempty"`
		
		ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`
		Alias
	}{ 
		Action: o.Action,
		
		ActionDate: ActionDate,
		
		ActionAge: o.ActionAge,
		
		ScreenRecordingActionDate: ScreenRecordingActionDate,
		
		ScreenRecordingActionAge: o.ScreenRecordingActionAge,
		
		IntegrationId: o.IntegrationId,
		
		IncludeRecordingsWithSensitiveData: o.IncludeRecordingsWithSensitiveData,
		
		IncludeScreenRecordings: o.IncludeScreenRecordings,
		
		ClearExport: o.ClearExport,
		
		ConversationQuery: o.ConversationQuery,
		Alias:    (Alias)(o),
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
	
	if ActionAge, ok := RecordingjobsqueryMap["actionAge"].(float64); ok {
		ActionAgeInt := int(ActionAge)
		o.ActionAge = &ActionAgeInt
	}
	
	if screenRecordingActionDateString, ok := RecordingjobsqueryMap["screenRecordingActionDate"].(string); ok {
		ScreenRecordingActionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", screenRecordingActionDateString)
		o.ScreenRecordingActionDate = &ScreenRecordingActionDate
	}
	
	if ScreenRecordingActionAge, ok := RecordingjobsqueryMap["screenRecordingActionAge"].(float64); ok {
		ScreenRecordingActionAgeInt := int(ScreenRecordingActionAge)
		o.ScreenRecordingActionAge = &ScreenRecordingActionAgeInt
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

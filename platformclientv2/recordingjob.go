package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjob
type Recordingjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// State - The current state of the job.
	State *string `json:"state,omitempty"`

	// RecordingJobsQuery - Original query of the job.
	RecordingJobsQuery *Recordingjobsquery `json:"recordingJobsQuery,omitempty"`

	// DateCreated - Date when the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// TotalConversations - Total number of conversations affected.
	TotalConversations *int `json:"totalConversations,omitempty"`

	// TotalRecordings - Total number of recordings affected.
	TotalRecordings *int `json:"totalRecordings,omitempty"`

	// TotalSkippedRecordings - Total number of recordings that have been skipped.
	TotalSkippedRecordings *int `json:"totalSkippedRecordings,omitempty"`

	// TotalFailedRecordings - Total number of recordings that the bulk job failed to process.
	TotalFailedRecordings *int `json:"totalFailedRecordings,omitempty"`

	// TotalProcessedRecordings - Total number of recordings have been processed.
	TotalProcessedRecordings *int `json:"totalProcessedRecordings,omitempty"`

	// PercentProgress - Progress in percentage based on the number of recordings
	PercentProgress *int `json:"percentProgress,omitempty"`

	// ErrorMessage - Error occurred during the job execution
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// FailedRecordings - Get IDs of recordings that the bulk job failed for
	FailedRecordings *string `json:"failedRecordings,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingjob) SetField(field string, fieldValue interface{}) {
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

func (o Recordingjob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Recordingjob
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		RecordingJobsQuery *Recordingjobsquery `json:"recordingJobsQuery,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		TotalConversations *int `json:"totalConversations,omitempty"`
		
		TotalRecordings *int `json:"totalRecordings,omitempty"`
		
		TotalSkippedRecordings *int `json:"totalSkippedRecordings,omitempty"`
		
		TotalFailedRecordings *int `json:"totalFailedRecordings,omitempty"`
		
		TotalProcessedRecordings *int `json:"totalProcessedRecordings,omitempty"`
		
		PercentProgress *int `json:"percentProgress,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		FailedRecordings *string `json:"failedRecordings,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		User *Addressableentityref `json:"user,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		RecordingJobsQuery: o.RecordingJobsQuery,
		
		DateCreated: DateCreated,
		
		TotalConversations: o.TotalConversations,
		
		TotalRecordings: o.TotalRecordings,
		
		TotalSkippedRecordings: o.TotalSkippedRecordings,
		
		TotalFailedRecordings: o.TotalFailedRecordings,
		
		TotalProcessedRecordings: o.TotalProcessedRecordings,
		
		PercentProgress: o.PercentProgress,
		
		ErrorMessage: o.ErrorMessage,
		
		FailedRecordings: o.FailedRecordings,
		
		SelfUri: o.SelfUri,
		
		User: o.User,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingjob) UnmarshalJSON(b []byte) error {
	var RecordingjobMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := RecordingjobMap["state"].(string); ok {
		o.State = &State
	}
    
	if RecordingJobsQuery, ok := RecordingjobMap["recordingJobsQuery"].(map[string]interface{}); ok {
		RecordingJobsQueryString, _ := json.Marshal(RecordingJobsQuery)
		json.Unmarshal(RecordingJobsQueryString, &o.RecordingJobsQuery)
	}
	
	if dateCreatedString, ok := RecordingjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if TotalConversations, ok := RecordingjobMap["totalConversations"].(float64); ok {
		TotalConversationsInt := int(TotalConversations)
		o.TotalConversations = &TotalConversationsInt
	}
	
	if TotalRecordings, ok := RecordingjobMap["totalRecordings"].(float64); ok {
		TotalRecordingsInt := int(TotalRecordings)
		o.TotalRecordings = &TotalRecordingsInt
	}
	
	if TotalSkippedRecordings, ok := RecordingjobMap["totalSkippedRecordings"].(float64); ok {
		TotalSkippedRecordingsInt := int(TotalSkippedRecordings)
		o.TotalSkippedRecordings = &TotalSkippedRecordingsInt
	}
	
	if TotalFailedRecordings, ok := RecordingjobMap["totalFailedRecordings"].(float64); ok {
		TotalFailedRecordingsInt := int(TotalFailedRecordings)
		o.TotalFailedRecordings = &TotalFailedRecordingsInt
	}
	
	if TotalProcessedRecordings, ok := RecordingjobMap["totalProcessedRecordings"].(float64); ok {
		TotalProcessedRecordingsInt := int(TotalProcessedRecordings)
		o.TotalProcessedRecordings = &TotalProcessedRecordingsInt
	}
	
	if PercentProgress, ok := RecordingjobMap["percentProgress"].(float64); ok {
		PercentProgressInt := int(PercentProgress)
		o.PercentProgress = &PercentProgressInt
	}
	
	if ErrorMessage, ok := RecordingjobMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if FailedRecordings, ok := RecordingjobMap["failedRecordings"].(string); ok {
		o.FailedRecordings = &FailedRecordings
	}
    
	if SelfUri, ok := RecordingjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if User, ok := RecordingjobMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

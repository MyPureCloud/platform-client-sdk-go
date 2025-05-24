package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage
type V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// SummaryId
	SummaryId *string `json:"summaryId,omitempty"`

	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// SummarySettingsId
	SummarySettingsId *string `json:"summarySettingsId,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Summary
	Summary *V2sessionconversationssummarysettingsprevieweventconversationsummarymessage `json:"summary,omitempty"`

	// Reason
	Reason *V2sessionconversationssummarysettingsprevieweventconversationreasonmessage `json:"reason,omitempty"`

	// Resolution
	Resolution *V2sessionconversationssummarysettingsprevieweventconversationresolutionmessage `json:"resolution,omitempty"`

	// FollowupActions
	FollowupActions *[]V2sessionconversationssummarysettingsprevieweventconversationfollowupaction `json:"followupActions,omitempty"`

	// ExtractedEntities
	ExtractedEntities *[]V2sessionconversationssummarysettingsprevieweventconversationsummaryextractedentity `json:"extractedEntities,omitempty"`

	// ErrorType
	ErrorType *string `json:"errorType,omitempty"`

	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage) SetField(field string, fieldValue interface{}) {
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

func (o V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		CreatedDate *string `json:"createdDate,omitempty"`
		
		SummaryId *string `json:"summaryId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		SummarySettingsId *string `json:"summarySettingsId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Summary *V2sessionconversationssummarysettingsprevieweventconversationsummarymessage `json:"summary,omitempty"`
		
		Reason *V2sessionconversationssummarysettingsprevieweventconversationreasonmessage `json:"reason,omitempty"`
		
		Resolution *V2sessionconversationssummarysettingsprevieweventconversationresolutionmessage `json:"resolution,omitempty"`
		
		FollowupActions *[]V2sessionconversationssummarysettingsprevieweventconversationfollowupaction `json:"followupActions,omitempty"`
		
		ExtractedEntities *[]V2sessionconversationssummarysettingsprevieweventconversationsummaryextractedentity `json:"extractedEntities,omitempty"`
		
		ErrorType *string `json:"errorType,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		Alias
	}{ 
		CreatedDate: CreatedDate,
		
		SummaryId: o.SummaryId,
		
		SessionId: o.SessionId,
		
		UserId: o.UserId,
		
		SummarySettingsId: o.SummarySettingsId,
		
		Language: o.Language,
		
		MediaType: o.MediaType,
		
		Summary: o.Summary,
		
		Reason: o.Reason,
		
		Resolution: o.Resolution,
		
		FollowupActions: o.FollowupActions,
		
		ExtractedEntities: o.ExtractedEntities,
		
		ErrorType: o.ErrorType,
		
		DurationMs: o.DurationMs,
		Alias:    (Alias)(o),
	})
}

func (o *V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage) UnmarshalJSON(b []byte) error {
	var V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap)
	if err != nil {
		return err
	}
	
	if createdDateString, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if SummaryId, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["summaryId"].(string); ok {
		o.SummaryId = &SummaryId
	}
    
	if SessionId, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if UserId, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if SummarySettingsId, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["summarySettingsId"].(string); ok {
		o.SummarySettingsId = &SummarySettingsId
	}
    
	if Language, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if MediaType, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Summary, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["summary"].(map[string]interface{}); ok {
		SummaryString, _ := json.Marshal(Summary)
		json.Unmarshal(SummaryString, &o.Summary)
	}
	
	if Reason, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["reason"].(map[string]interface{}); ok {
		ReasonString, _ := json.Marshal(Reason)
		json.Unmarshal(ReasonString, &o.Reason)
	}
	
	if Resolution, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["resolution"].(map[string]interface{}); ok {
		ResolutionString, _ := json.Marshal(Resolution)
		json.Unmarshal(ResolutionString, &o.Resolution)
	}
	
	if FollowupActions, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["followupActions"].([]interface{}); ok {
		FollowupActionsString, _ := json.Marshal(FollowupActions)
		json.Unmarshal(FollowupActionsString, &o.FollowupActions)
	}
	
	if ExtractedEntities, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["extractedEntities"].([]interface{}); ok {
		ExtractedEntitiesString, _ := json.Marshal(ExtractedEntities)
		json.Unmarshal(ExtractedEntitiesString, &o.ExtractedEntities)
	}
	
	if ErrorType, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["errorType"].(string); ok {
		o.ErrorType = &ErrorType
	}
    
	if DurationMs, ok := V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessageMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2sessionconversationssummarysettingsprevieweventconversationsummarysettingspreviewmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

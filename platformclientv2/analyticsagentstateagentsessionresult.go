package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentstateagentsessionresult
type Analyticsagentstateagentsessionresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId - Conversation Id
	ConversationId *string `json:"conversationId,omitempty"`

	// SessionId - Session Id
	SessionId *string `json:"sessionId,omitempty"`

	// SessionStart - Session start datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SessionStart *time.Time `json:"sessionStart,omitempty"`

	// SegmentStart - Segment start datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentStart *time.Time `json:"segmentStart,omitempty"`

	// SegmentType - Segment type
	SegmentType *string `json:"segmentType,omitempty"`

	// RoutedQueueId - Routed queue Id
	RoutedQueueId *string `json:"routedQueueId,omitempty"`

	// RequestedRoutingSkillIds - List of requested routing skill Id
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`

	// RequestedLanguageId - Requested language Id
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// OriginatingDirection - Originating direction
	OriginatingDirection *string `json:"originatingDirection,omitempty"`

	// MediaType - Media type
	MediaType *string `json:"mediaType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsagentstateagentsessionresult) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsagentstateagentsessionresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "SessionStart","SegmentStart", }
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
	type Alias Analyticsagentstateagentsessionresult
	
	SessionStart := new(string)
	if o.SessionStart != nil {
		
		*SessionStart = timeutil.Strftime(o.SessionStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SessionStart = nil
	}
	
	SegmentStart := new(string)
	if o.SegmentStart != nil {
		
		*SegmentStart = timeutil.Strftime(o.SegmentStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SegmentStart = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		SessionStart *string `json:"sessionStart,omitempty"`
		
		SegmentStart *string `json:"segmentStart,omitempty"`
		
		SegmentType *string `json:"segmentType,omitempty"`
		
		RoutedQueueId *string `json:"routedQueueId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		SessionId: o.SessionId,
		
		SessionStart: SessionStart,
		
		SegmentStart: SegmentStart,
		
		SegmentType: o.SegmentType,
		
		RoutedQueueId: o.RoutedQueueId,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		OriginatingDirection: o.OriginatingDirection,
		
		MediaType: o.MediaType,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsagentstateagentsessionresult) UnmarshalJSON(b []byte) error {
	var AnalyticsagentstateagentsessionresultMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentstateagentsessionresultMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := AnalyticsagentstateagentsessionresultMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SessionId, ok := AnalyticsagentstateagentsessionresultMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if sessionStartString, ok := AnalyticsagentstateagentsessionresultMap["sessionStart"].(string); ok {
		SessionStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", sessionStartString)
		o.SessionStart = &SessionStart
	}
	
	if segmentStartString, ok := AnalyticsagentstateagentsessionresultMap["segmentStart"].(string); ok {
		SegmentStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", segmentStartString)
		o.SegmentStart = &SegmentStart
	}
	
	if SegmentType, ok := AnalyticsagentstateagentsessionresultMap["segmentType"].(string); ok {
		o.SegmentType = &SegmentType
	}
    
	if RoutedQueueId, ok := AnalyticsagentstateagentsessionresultMap["routedQueueId"].(string); ok {
		o.RoutedQueueId = &RoutedQueueId
	}
    
	if RequestedRoutingSkillIds, ok := AnalyticsagentstateagentsessionresultMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedLanguageId, ok := AnalyticsagentstateagentsessionresultMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if OriginatingDirection, ok := AnalyticsagentstateagentsessionresultMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
    
	if MediaType, ok := AnalyticsagentstateagentsessionresultMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentsessionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

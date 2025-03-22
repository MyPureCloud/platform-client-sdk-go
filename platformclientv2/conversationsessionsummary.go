package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsessionsummary
type Conversationsessionsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Text - The text of the summary.
	Text *string `json:"text,omitempty"`

	// Status - The status of the conversation summary.
	Status *string `json:"status,omitempty"`

	// MediaType - The media type of the conversation.
	MediaType *string `json:"mediaType,omitempty"`

	// Language - The language of the conversation.
	Language *string `json:"language,omitempty"`

	// PredictedWrapupCodes - The wrapup codes of the conversation summary.
	PredictedWrapupCodes *[]Conversationsummarywrapupcode `json:"predictedWrapupCodes,omitempty"`

	// EditedSummary - The edited summary of the conversation.
	EditedSummary *Conversationeditedinput `json:"editedSummary,omitempty"`

	// Reason - The reason of the conversation summary.
	Reason *Conversationsummaryreason `json:"reason,omitempty"`

	// Followup - The followup of the conversation summary.
	Followup *Conversationsummaryfollowup `json:"followup,omitempty"`

	// Resolution - The resolution of the conversation summary.
	Resolution *Conversationsummaryresolution `json:"resolution,omitempty"`

	// DateCreated - The created date of the summary. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Id - The id of the summary.
	Id *string `json:"id,omitempty"`

	// Confidence - The AI confidence value.
	Confidence *float32 `json:"confidence,omitempty"`

	// Participants - The list of participants.
	Participants *[]Addressableentityref `json:"participants,omitempty"`

	// Communication - The communication object of the summary.
	Communication *Entity `json:"communication,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationsessionsummary) SetField(field string, fieldValue interface{}) {
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

func (o Conversationsessionsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationsessionsummary
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		PredictedWrapupCodes *[]Conversationsummarywrapupcode `json:"predictedWrapupCodes,omitempty"`
		
		EditedSummary *Conversationeditedinput `json:"editedSummary,omitempty"`
		
		Reason *Conversationsummaryreason `json:"reason,omitempty"`
		
		Followup *Conversationsummaryfollowup `json:"followup,omitempty"`
		
		Resolution *Conversationsummaryresolution `json:"resolution,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		
		Participants *[]Addressableentityref `json:"participants,omitempty"`
		
		Communication *Entity `json:"communication,omitempty"`
		Alias
	}{ 
		Text: o.Text,
		
		Status: o.Status,
		
		MediaType: o.MediaType,
		
		Language: o.Language,
		
		PredictedWrapupCodes: o.PredictedWrapupCodes,
		
		EditedSummary: o.EditedSummary,
		
		Reason: o.Reason,
		
		Followup: o.Followup,
		
		Resolution: o.Resolution,
		
		DateCreated: DateCreated,
		
		Id: o.Id,
		
		Confidence: o.Confidence,
		
		Participants: o.Participants,
		
		Communication: o.Communication,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationsessionsummary) UnmarshalJSON(b []byte) error {
	var ConversationsessionsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsessionsummaryMap)
	if err != nil {
		return err
	}
	
	if Text, ok := ConversationsessionsummaryMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Status, ok := ConversationsessionsummaryMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if MediaType, ok := ConversationsessionsummaryMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Language, ok := ConversationsessionsummaryMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if PredictedWrapupCodes, ok := ConversationsessionsummaryMap["predictedWrapupCodes"].([]interface{}); ok {
		PredictedWrapupCodesString, _ := json.Marshal(PredictedWrapupCodes)
		json.Unmarshal(PredictedWrapupCodesString, &o.PredictedWrapupCodes)
	}
	
	if EditedSummary, ok := ConversationsessionsummaryMap["editedSummary"].(map[string]interface{}); ok {
		EditedSummaryString, _ := json.Marshal(EditedSummary)
		json.Unmarshal(EditedSummaryString, &o.EditedSummary)
	}
	
	if Reason, ok := ConversationsessionsummaryMap["reason"].(map[string]interface{}); ok {
		ReasonString, _ := json.Marshal(Reason)
		json.Unmarshal(ReasonString, &o.Reason)
	}
	
	if Followup, ok := ConversationsessionsummaryMap["followup"].(map[string]interface{}); ok {
		FollowupString, _ := json.Marshal(Followup)
		json.Unmarshal(FollowupString, &o.Followup)
	}
	
	if Resolution, ok := ConversationsessionsummaryMap["resolution"].(map[string]interface{}); ok {
		ResolutionString, _ := json.Marshal(Resolution)
		json.Unmarshal(ResolutionString, &o.Resolution)
	}
	
	if dateCreatedString, ok := ConversationsessionsummaryMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Id, ok := ConversationsessionsummaryMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Confidence, ok := ConversationsessionsummaryMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if Participants, ok := ConversationsessionsummaryMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if Communication, ok := ConversationsessionsummaryMap["communication"].(map[string]interface{}); ok {
		CommunicationString, _ := json.Marshal(Communication)
		json.Unmarshal(CommunicationString, &o.Communication)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsessionsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

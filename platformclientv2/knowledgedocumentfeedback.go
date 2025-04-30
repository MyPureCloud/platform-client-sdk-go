package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentfeedback
type Knowledgedocumentfeedback struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DocumentVariation - The variation of the document on which feedback was given.
	DocumentVariation *Entityreference `json:"documentVariation,omitempty"`

	// Rating - Feedback rating.
	Rating *string `json:"rating,omitempty"`

	// Reason - Feedback reason.
	Reason *string `json:"reason,omitempty"`

	// Comment - Free-text comment of the feedback. Maximum length: 2000 characters.
	Comment *string `json:"comment,omitempty"`

	// Search - The search that surfaced the document on which feedback was given.
	Search *Entityreference `json:"search,omitempty"`

	// SessionId - Knowledge guest session ID.
	SessionId *string `json:"sessionId,omitempty"`

	// DateCreated - The date and time of the feedback. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// QueryType - The type of the query that surfaced the document on which the feedback was given.
	QueryType *string `json:"queryType,omitempty"`

	// SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.Generative: A generated answer in a snippet was shown.
	SurfacingMethod *string `json:"surfacingMethod,omitempty"`

	// State - The state of the feedback.
	State *string `json:"state,omitempty"`

	// Document - The document on which feedback was given.
	Document *Knowledgedocumentversionreference `json:"document,omitempty"`

	// Application - The client application from which feedback was given.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`

	// ConversationContext - Conversation context information if the feedback is given in the context of a conversation.
	ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`

	// UserId - The ID of the user who created the feedback.
	UserId *string `json:"userId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentfeedback) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentfeedback) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentfeedback
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DocumentVariation *Entityreference `json:"documentVariation,omitempty"`
		
		Rating *string `json:"rating,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		
		Comment *string `json:"comment,omitempty"`
		
		Search *Entityreference `json:"search,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SurfacingMethod *string `json:"surfacingMethod,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Document *Knowledgedocumentversionreference `json:"document,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DocumentVariation: o.DocumentVariation,
		
		Rating: o.Rating,
		
		Reason: o.Reason,
		
		Comment: o.Comment,
		
		Search: o.Search,
		
		SessionId: o.SessionId,
		
		DateCreated: DateCreated,
		
		QueryType: o.QueryType,
		
		SurfacingMethod: o.SurfacingMethod,
		
		State: o.State,
		
		Document: o.Document,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		
		UserId: o.UserId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentfeedback) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentfeedbackMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentfeedbackMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentfeedbackMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DocumentVariation, ok := KnowledgedocumentfeedbackMap["documentVariation"].(map[string]interface{}); ok {
		DocumentVariationString, _ := json.Marshal(DocumentVariation)
		json.Unmarshal(DocumentVariationString, &o.DocumentVariation)
	}
	
	if Rating, ok := KnowledgedocumentfeedbackMap["rating"].(string); ok {
		o.Rating = &Rating
	}
    
	if Reason, ok := KnowledgedocumentfeedbackMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    
	if Comment, ok := KnowledgedocumentfeedbackMap["comment"].(string); ok {
		o.Comment = &Comment
	}
    
	if Search, ok := KnowledgedocumentfeedbackMap["search"].(map[string]interface{}); ok {
		SearchString, _ := json.Marshal(Search)
		json.Unmarshal(SearchString, &o.Search)
	}
	
	if SessionId, ok := KnowledgedocumentfeedbackMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if dateCreatedString, ok := KnowledgedocumentfeedbackMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if QueryType, ok := KnowledgedocumentfeedbackMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SurfacingMethod, ok := KnowledgedocumentfeedbackMap["surfacingMethod"].(string); ok {
		o.SurfacingMethod = &SurfacingMethod
	}
    
	if State, ok := KnowledgedocumentfeedbackMap["state"].(string); ok {
		o.State = &State
	}
    
	if Document, ok := KnowledgedocumentfeedbackMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if Application, ok := KnowledgedocumentfeedbackMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgedocumentfeedbackMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if UserId, ok := KnowledgedocumentfeedbackMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if SelfUri, ok := KnowledgedocumentfeedbackMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentfeedback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

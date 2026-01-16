package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assistant
type Assistant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the assistant that will assist the agent.
	Name *string `json:"name,omitempty"`

	// DateCreated - Date when the assistant was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date when the assistant was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - The user who created the assistant.
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// ModifiedBy - The user who last modified the assistant.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// GoogleDialogflowConfig - (Deprecated: use the 'knowledgeSuggestionConfig' for genesys knowledge suggestions) Configuration of Dialogflow used to assist the agent with transcriptions and knowledge suggestions.
	GoogleDialogflowConfig *Googledialogflowconfig `json:"googleDialogflowConfig,omitempty"`

	// TranscriptionConfig - Configuration for speech transcription used to assist the agent.
	TranscriptionConfig *Transcriptionconfig `json:"transcriptionConfig,omitempty"`

	// KnowledgeSuggestionConfig - Configuration that defines how to produce knowledge suggestions.
	KnowledgeSuggestionConfig *Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig,omitempty"`

	// State - State of the assistant.
	State *string `json:"state,omitempty"`

	// Copilot - Agent copilot configuration.
	Copilot *Copilot `json:"copilot,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// VariationParent - The parent assistant if this assistant is a variation of an assistant
	VariationParent *Addressableentityref `json:"variationParent,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assistant) SetField(field string, fieldValue interface{}) {
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

func (o Assistant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Assistant
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		GoogleDialogflowConfig *Googledialogflowconfig `json:"googleDialogflowConfig,omitempty"`
		
		TranscriptionConfig *Transcriptionconfig `json:"transcriptionConfig,omitempty"`
		
		KnowledgeSuggestionConfig *Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Copilot *Copilot `json:"copilot,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VariationParent *Addressableentityref `json:"variationParent,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		GoogleDialogflowConfig: o.GoogleDialogflowConfig,
		
		TranscriptionConfig: o.TranscriptionConfig,
		
		KnowledgeSuggestionConfig: o.KnowledgeSuggestionConfig,
		
		State: o.State,
		
		Copilot: o.Copilot,
		
		SelfUri: o.SelfUri,
		
		VariationParent: o.VariationParent,
		Alias:    (Alias)(o),
	})
}

func (o *Assistant) UnmarshalJSON(b []byte) error {
	var AssistantMap map[string]interface{}
	err := json.Unmarshal(b, &AssistantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssistantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AssistantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := AssistantMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := AssistantMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := AssistantMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := AssistantMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if GoogleDialogflowConfig, ok := AssistantMap["googleDialogflowConfig"].(map[string]interface{}); ok {
		GoogleDialogflowConfigString, _ := json.Marshal(GoogleDialogflowConfig)
		json.Unmarshal(GoogleDialogflowConfigString, &o.GoogleDialogflowConfig)
	}
	
	if TranscriptionConfig, ok := AssistantMap["transcriptionConfig"].(map[string]interface{}); ok {
		TranscriptionConfigString, _ := json.Marshal(TranscriptionConfig)
		json.Unmarshal(TranscriptionConfigString, &o.TranscriptionConfig)
	}
	
	if KnowledgeSuggestionConfig, ok := AssistantMap["knowledgeSuggestionConfig"].(map[string]interface{}); ok {
		KnowledgeSuggestionConfigString, _ := json.Marshal(KnowledgeSuggestionConfig)
		json.Unmarshal(KnowledgeSuggestionConfigString, &o.KnowledgeSuggestionConfig)
	}
	
	if State, ok := AssistantMap["state"].(string); ok {
		o.State = &State
	}
    
	if Copilot, ok := AssistantMap["copilot"].(map[string]interface{}); ok {
		CopilotString, _ := json.Marshal(Copilot)
		json.Unmarshal(CopilotString, &o.Copilot)
	}
	
	if SelfUri, ok := AssistantMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if VariationParent, ok := AssistantMap["variationParent"].(map[string]interface{}); ok {
		VariationParentString, _ := json.Marshal(VariationParent)
		json.Unmarshal(VariationParentString, &o.VariationParent)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assistant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

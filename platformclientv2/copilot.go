package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Copilot
type Copilot struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Copilot is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// LiveOnQueue - Copilot is live on selected queue.
	LiveOnQueue *bool `json:"liveOnQueue,omitempty"`

	// DefaultLanguage - Copilot default language, e.g. [en-US, es-US, es-ES]. Once set, it can not be modified.
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`

	// KnowledgeAnswerConfig - Deprecated: Please use AutoSearchConfig and ManualSearchConfig fields instead.
	KnowledgeAnswerConfig *Knowledgeanswerconfig `json:"knowledgeAnswerConfig,omitempty"`

	// SummaryGenerationConfig - Copilot generated summary configuration.
	SummaryGenerationConfig *Summarygenerationconfig `json:"summaryGenerationConfig,omitempty"`

	// WrapupCodePredictionConfig - Copilot generated wrapup code prediction configuration.
	WrapupCodePredictionConfig *Wrapupcodepredictionconfig `json:"wrapupCodePredictionConfig,omitempty"`

	// AnswerGenerationConfig - Deprecated: Please use AutoSearchConfig and ManualSearchConfig fields instead.
	AnswerGenerationConfig *Answergenerationconfig `json:"answerGenerationConfig,omitempty"`

	// NluEngineType - Language understanding engine type.
	NluEngineType *string `json:"nluEngineType,omitempty"`

	// NluConfig - NLU configuration.
	NluConfig *Nluconfig `json:"nluConfig,omitempty"`

	// RuleEngineConfig - Rule engine configuration.
	RuleEngineConfig *Ruleengineconfig `json:"ruleEngineConfig,omitempty"`

	// AutoSearchConfig - Auto search configuration.
	AutoSearchConfig *Autosearchconfig `json:"autoSearchConfig,omitempty"`

	// ManualSearchConfig - Manual Search configuration.
	ManualSearchConfig *Manualsearchconfig `json:"manualSearchConfig,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Copilot) SetField(field string, fieldValue interface{}) {
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

func (o Copilot) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Copilot
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		LiveOnQueue *bool `json:"liveOnQueue,omitempty"`
		
		DefaultLanguage *string `json:"defaultLanguage,omitempty"`
		
		KnowledgeAnswerConfig *Knowledgeanswerconfig `json:"knowledgeAnswerConfig,omitempty"`
		
		SummaryGenerationConfig *Summarygenerationconfig `json:"summaryGenerationConfig,omitempty"`
		
		WrapupCodePredictionConfig *Wrapupcodepredictionconfig `json:"wrapupCodePredictionConfig,omitempty"`
		
		AnswerGenerationConfig *Answergenerationconfig `json:"answerGenerationConfig,omitempty"`
		
		NluEngineType *string `json:"nluEngineType,omitempty"`
		
		NluConfig *Nluconfig `json:"nluConfig,omitempty"`
		
		RuleEngineConfig *Ruleengineconfig `json:"ruleEngineConfig,omitempty"`
		
		AutoSearchConfig *Autosearchconfig `json:"autoSearchConfig,omitempty"`
		
		ManualSearchConfig *Manualsearchconfig `json:"manualSearchConfig,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		LiveOnQueue: o.LiveOnQueue,
		
		DefaultLanguage: o.DefaultLanguage,
		
		KnowledgeAnswerConfig: o.KnowledgeAnswerConfig,
		
		SummaryGenerationConfig: o.SummaryGenerationConfig,
		
		WrapupCodePredictionConfig: o.WrapupCodePredictionConfig,
		
		AnswerGenerationConfig: o.AnswerGenerationConfig,
		
		NluEngineType: o.NluEngineType,
		
		NluConfig: o.NluConfig,
		
		RuleEngineConfig: o.RuleEngineConfig,
		
		AutoSearchConfig: o.AutoSearchConfig,
		
		ManualSearchConfig: o.ManualSearchConfig,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Copilot) UnmarshalJSON(b []byte) error {
	var CopilotMap map[string]interface{}
	err := json.Unmarshal(b, &CopilotMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := CopilotMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if LiveOnQueue, ok := CopilotMap["liveOnQueue"].(bool); ok {
		o.LiveOnQueue = &LiveOnQueue
	}
    
	if DefaultLanguage, ok := CopilotMap["defaultLanguage"].(string); ok {
		o.DefaultLanguage = &DefaultLanguage
	}
    
	if KnowledgeAnswerConfig, ok := CopilotMap["knowledgeAnswerConfig"].(map[string]interface{}); ok {
		KnowledgeAnswerConfigString, _ := json.Marshal(KnowledgeAnswerConfig)
		json.Unmarshal(KnowledgeAnswerConfigString, &o.KnowledgeAnswerConfig)
	}
	
	if SummaryGenerationConfig, ok := CopilotMap["summaryGenerationConfig"].(map[string]interface{}); ok {
		SummaryGenerationConfigString, _ := json.Marshal(SummaryGenerationConfig)
		json.Unmarshal(SummaryGenerationConfigString, &o.SummaryGenerationConfig)
	}
	
	if WrapupCodePredictionConfig, ok := CopilotMap["wrapupCodePredictionConfig"].(map[string]interface{}); ok {
		WrapupCodePredictionConfigString, _ := json.Marshal(WrapupCodePredictionConfig)
		json.Unmarshal(WrapupCodePredictionConfigString, &o.WrapupCodePredictionConfig)
	}
	
	if AnswerGenerationConfig, ok := CopilotMap["answerGenerationConfig"].(map[string]interface{}); ok {
		AnswerGenerationConfigString, _ := json.Marshal(AnswerGenerationConfig)
		json.Unmarshal(AnswerGenerationConfigString, &o.AnswerGenerationConfig)
	}
	
	if NluEngineType, ok := CopilotMap["nluEngineType"].(string); ok {
		o.NluEngineType = &NluEngineType
	}
    
	if NluConfig, ok := CopilotMap["nluConfig"].(map[string]interface{}); ok {
		NluConfigString, _ := json.Marshal(NluConfig)
		json.Unmarshal(NluConfigString, &o.NluConfig)
	}
	
	if RuleEngineConfig, ok := CopilotMap["ruleEngineConfig"].(map[string]interface{}); ok {
		RuleEngineConfigString, _ := json.Marshal(RuleEngineConfig)
		json.Unmarshal(RuleEngineConfigString, &o.RuleEngineConfig)
	}
	
	if AutoSearchConfig, ok := CopilotMap["autoSearchConfig"].(map[string]interface{}); ok {
		AutoSearchConfigString, _ := json.Marshal(AutoSearchConfig)
		json.Unmarshal(AutoSearchConfigString, &o.AutoSearchConfig)
	}
	
	if ManualSearchConfig, ok := CopilotMap["manualSearchConfig"].(map[string]interface{}); ok {
		ManualSearchConfigString, _ := json.Marshal(ManualSearchConfig)
		json.Unmarshal(ManualSearchConfigString, &o.ManualSearchConfig)
	}
	
	if SelfUri, ok := CopilotMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Copilot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

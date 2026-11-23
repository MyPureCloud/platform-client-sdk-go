package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitieslanguagefeatures
type Capabilitieslanguagefeatures struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CollectInput
	CollectInput *Capabilitiesflowtypesavailability `json:"collectInput,omitempty"`

	// Copilot
	Copilot *Capabilitiescopilot `json:"copilot,omitempty"`

	// EstimatedWaitTime
	EstimatedWaitTime *Capabilitiesflowtypesavailability `json:"estimatedWaitTime,omitempty"`

	// KnowledgeBase
	KnowledgeBase *Capabilitiesknowledgebase `json:"knowledgeBase,omitempty"`

	// PositionInQueue
	PositionInQueue *Capabilitiesflowtypesavailability `json:"positionInQueue,omitempty"`

	// RuntimeDataPlayback
	RuntimeDataPlayback *Capabilitiesflowtypesavailability `json:"runtimeDataPlayback,omitempty"`

	// SpeechRec
	SpeechRec *Capabilitiesflowtypesavailability `json:"speechRec,omitempty"`

	// TextToSpeech
	TextToSpeech *Capabilitieslanguagetts `json:"textToSpeech,omitempty"`

	// Transcription
	Transcription *Capabilitiesflowtypesavailability `json:"transcription,omitempty"`

	// Voicemail
	Voicemail *map[string]Capabilitiesflowtypesavailability `json:"voicemail,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitieslanguagefeatures) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitieslanguagefeatures) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitieslanguagefeatures
	
	return json.Marshal(&struct { 
		CollectInput *Capabilitiesflowtypesavailability `json:"collectInput,omitempty"`
		
		Copilot *Capabilitiescopilot `json:"copilot,omitempty"`
		
		EstimatedWaitTime *Capabilitiesflowtypesavailability `json:"estimatedWaitTime,omitempty"`
		
		KnowledgeBase *Capabilitiesknowledgebase `json:"knowledgeBase,omitempty"`
		
		PositionInQueue *Capabilitiesflowtypesavailability `json:"positionInQueue,omitempty"`
		
		RuntimeDataPlayback *Capabilitiesflowtypesavailability `json:"runtimeDataPlayback,omitempty"`
		
		SpeechRec *Capabilitiesflowtypesavailability `json:"speechRec,omitempty"`
		
		TextToSpeech *Capabilitieslanguagetts `json:"textToSpeech,omitempty"`
		
		Transcription *Capabilitiesflowtypesavailability `json:"transcription,omitempty"`
		
		Voicemail *map[string]Capabilitiesflowtypesavailability `json:"voicemail,omitempty"`
		Alias
	}{ 
		CollectInput: o.CollectInput,
		
		Copilot: o.Copilot,
		
		EstimatedWaitTime: o.EstimatedWaitTime,
		
		KnowledgeBase: o.KnowledgeBase,
		
		PositionInQueue: o.PositionInQueue,
		
		RuntimeDataPlayback: o.RuntimeDataPlayback,
		
		SpeechRec: o.SpeechRec,
		
		TextToSpeech: o.TextToSpeech,
		
		Transcription: o.Transcription,
		
		Voicemail: o.Voicemail,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitieslanguagefeatures) UnmarshalJSON(b []byte) error {
	var CapabilitieslanguagefeaturesMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitieslanguagefeaturesMap)
	if err != nil {
		return err
	}
	
	if CollectInput, ok := CapabilitieslanguagefeaturesMap["collectInput"].(map[string]interface{}); ok {
		CollectInputString, _ := json.Marshal(CollectInput)
		json.Unmarshal(CollectInputString, &o.CollectInput)
	}
	
	if Copilot, ok := CapabilitieslanguagefeaturesMap["copilot"].(map[string]interface{}); ok {
		CopilotString, _ := json.Marshal(Copilot)
		json.Unmarshal(CopilotString, &o.Copilot)
	}
	
	if EstimatedWaitTime, ok := CapabilitieslanguagefeaturesMap["estimatedWaitTime"].(map[string]interface{}); ok {
		EstimatedWaitTimeString, _ := json.Marshal(EstimatedWaitTime)
		json.Unmarshal(EstimatedWaitTimeString, &o.EstimatedWaitTime)
	}
	
	if KnowledgeBase, ok := CapabilitieslanguagefeaturesMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if PositionInQueue, ok := CapabilitieslanguagefeaturesMap["positionInQueue"].(map[string]interface{}); ok {
		PositionInQueueString, _ := json.Marshal(PositionInQueue)
		json.Unmarshal(PositionInQueueString, &o.PositionInQueue)
	}
	
	if RuntimeDataPlayback, ok := CapabilitieslanguagefeaturesMap["runtimeDataPlayback"].(map[string]interface{}); ok {
		RuntimeDataPlaybackString, _ := json.Marshal(RuntimeDataPlayback)
		json.Unmarshal(RuntimeDataPlaybackString, &o.RuntimeDataPlayback)
	}
	
	if SpeechRec, ok := CapabilitieslanguagefeaturesMap["speechRec"].(map[string]interface{}); ok {
		SpeechRecString, _ := json.Marshal(SpeechRec)
		json.Unmarshal(SpeechRecString, &o.SpeechRec)
	}
	
	if TextToSpeech, ok := CapabilitieslanguagefeaturesMap["textToSpeech"].(map[string]interface{}); ok {
		TextToSpeechString, _ := json.Marshal(TextToSpeech)
		json.Unmarshal(TextToSpeechString, &o.TextToSpeech)
	}
	
	if Transcription, ok := CapabilitieslanguagefeaturesMap["transcription"].(map[string]interface{}); ok {
		TranscriptionString, _ := json.Marshal(Transcription)
		json.Unmarshal(TranscriptionString, &o.Transcription)
	}
	
	if Voicemail, ok := CapabilitieslanguagefeaturesMap["voicemail"].(map[string]interface{}); ok {
		VoicemailString, _ := json.Marshal(Voicemail)
		json.Unmarshal(VoicemailString, &o.Voicemail)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagefeatures) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

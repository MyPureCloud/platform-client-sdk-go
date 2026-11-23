package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesfeaturesupport
type Capabilitiesfeaturesupport struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TextToSpeech - Architect features metadata text to speech
	TextToSpeech *Capabilitiesttsconfiguration `json:"textToSpeech,omitempty"`

	// SpeechToText - Architect features metadata speech to text
	SpeechToText *Capabilitiessttconfiguration `json:"speechToText,omitempty"`

	// NaturalLanguageUnderstanding - Architect features metadata natural language understanding
	NaturalLanguageUnderstanding *Capabilitiesnlusupport `json:"naturalLanguageUnderstanding,omitempty"`

	// Grammars - Architect features metadata grammars
	Grammars *Capabilitiesgrammarssupport `json:"grammars,omitempty"`

	// Workitems - Architect features metadata workitems
	Workitems *Capabilitiesworkitemssupport `json:"workitems,omitempty"`

	// Navigator - Architect features metadata navigator
	Navigator *Capabilitiesnavigatorsupport `json:"navigator,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesfeaturesupport) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesfeaturesupport) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesfeaturesupport
	
	return json.Marshal(&struct { 
		TextToSpeech *Capabilitiesttsconfiguration `json:"textToSpeech,omitempty"`
		
		SpeechToText *Capabilitiessttconfiguration `json:"speechToText,omitempty"`
		
		NaturalLanguageUnderstanding *Capabilitiesnlusupport `json:"naturalLanguageUnderstanding,omitempty"`
		
		Grammars *Capabilitiesgrammarssupport `json:"grammars,omitempty"`
		
		Workitems *Capabilitiesworkitemssupport `json:"workitems,omitempty"`
		
		Navigator *Capabilitiesnavigatorsupport `json:"navigator,omitempty"`
		Alias
	}{ 
		TextToSpeech: o.TextToSpeech,
		
		SpeechToText: o.SpeechToText,
		
		NaturalLanguageUnderstanding: o.NaturalLanguageUnderstanding,
		
		Grammars: o.Grammars,
		
		Workitems: o.Workitems,
		
		Navigator: o.Navigator,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesfeaturesupport) UnmarshalJSON(b []byte) error {
	var CapabilitiesfeaturesupportMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesfeaturesupportMap)
	if err != nil {
		return err
	}
	
	if TextToSpeech, ok := CapabilitiesfeaturesupportMap["textToSpeech"].(map[string]interface{}); ok {
		TextToSpeechString, _ := json.Marshal(TextToSpeech)
		json.Unmarshal(TextToSpeechString, &o.TextToSpeech)
	}
	
	if SpeechToText, ok := CapabilitiesfeaturesupportMap["speechToText"].(map[string]interface{}); ok {
		SpeechToTextString, _ := json.Marshal(SpeechToText)
		json.Unmarshal(SpeechToTextString, &o.SpeechToText)
	}
	
	if NaturalLanguageUnderstanding, ok := CapabilitiesfeaturesupportMap["naturalLanguageUnderstanding"].(map[string]interface{}); ok {
		NaturalLanguageUnderstandingString, _ := json.Marshal(NaturalLanguageUnderstanding)
		json.Unmarshal(NaturalLanguageUnderstandingString, &o.NaturalLanguageUnderstanding)
	}
	
	if Grammars, ok := CapabilitiesfeaturesupportMap["grammars"].(map[string]interface{}); ok {
		GrammarsString, _ := json.Marshal(Grammars)
		json.Unmarshal(GrammarsString, &o.Grammars)
	}
	
	if Workitems, ok := CapabilitiesfeaturesupportMap["workitems"].(map[string]interface{}); ok {
		WorkitemsString, _ := json.Marshal(Workitems)
		json.Unmarshal(WorkitemsString, &o.Workitems)
	}
	
	if Navigator, ok := CapabilitiesfeaturesupportMap["navigator"].(map[string]interface{}); ok {
		NavigatorString, _ := json.Marshal(Navigator)
		json.Unmarshal(NavigatorString, &o.Navigator)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesfeaturesupport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

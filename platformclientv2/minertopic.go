package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Minertopic
type Minertopic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Topic name.
	Name *string `json:"name,omitempty"`

	// Miner - The miner to which the topic belongs.
	Miner **Miner `json:"miner,omitempty"`

	// ConversationCount - Number of conversations where a topic has occurred.
	ConversationCount *int `json:"conversationCount,omitempty"`

	// ConversationPercent - Percentage of conversations where a topic has occurred.
	ConversationPercent *float32 `json:"conversationPercent,omitempty"`

	// UtteranceCount - Number of unique utterances where a topic has occurred.
	UtteranceCount *int `json:"utteranceCount,omitempty"`

	// PhraseCount - Number of unique phrases (sub-utterances) where a topic has occurred.
	PhraseCount *int `json:"phraseCount,omitempty"`

	// Phrases - Phrases associated with a topic.
	Phrases *[]Topicphrase `json:"phrases,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Minertopic) SetField(field string, fieldValue interface{}) {
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

func (o Minertopic) MarshalJSON() ([]byte, error) {
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
	type Alias Minertopic
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Miner **Miner `json:"miner,omitempty"`
		
		ConversationCount *int `json:"conversationCount,omitempty"`
		
		ConversationPercent *float32 `json:"conversationPercent,omitempty"`
		
		UtteranceCount *int `json:"utteranceCount,omitempty"`
		
		PhraseCount *int `json:"phraseCount,omitempty"`
		
		Phrases *[]Topicphrase `json:"phrases,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Miner: o.Miner,
		
		ConversationCount: o.ConversationCount,
		
		ConversationPercent: o.ConversationPercent,
		
		UtteranceCount: o.UtteranceCount,
		
		PhraseCount: o.PhraseCount,
		
		Phrases: o.Phrases,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Minertopic) UnmarshalJSON(b []byte) error {
	var MinertopicMap map[string]interface{}
	err := json.Unmarshal(b, &MinertopicMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MinertopicMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MinertopicMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Miner, ok := MinertopicMap["miner"].(map[string]interface{}); ok {
		MinerString, _ := json.Marshal(Miner)
		json.Unmarshal(MinerString, &o.Miner)
	}
	
	if ConversationCount, ok := MinertopicMap["conversationCount"].(float64); ok {
		ConversationCountInt := int(ConversationCount)
		o.ConversationCount = &ConversationCountInt
	}
	
	if ConversationPercent, ok := MinertopicMap["conversationPercent"].(float64); ok {
		ConversationPercentFloat32 := float32(ConversationPercent)
		o.ConversationPercent = &ConversationPercentFloat32
	}
	
	if UtteranceCount, ok := MinertopicMap["utteranceCount"].(float64); ok {
		UtteranceCountInt := int(UtteranceCount)
		o.UtteranceCount = &UtteranceCountInt
	}
	
	if PhraseCount, ok := MinertopicMap["phraseCount"].(float64); ok {
		PhraseCountInt := int(PhraseCount)
		o.PhraseCount = &PhraseCountInt
	}
	
	if Phrases, ok := MinertopicMap["phrases"].([]interface{}); ok {
		PhrasesString, _ := json.Marshal(Phrases)
		json.Unmarshal(PhrasesString, &o.Phrases)
	}
	
	if SelfUri, ok := MinertopicMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Minertopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

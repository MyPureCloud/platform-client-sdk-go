package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Minertopicphrase
type Minertopicphrase struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Phrase name.
	Name *string `json:"name,omitempty"`

	// Topic - Topic associated with a phrase.
	Topic *Minertopic `json:"topic,omitempty"`

	// Utterances - List of utterances related to a phrase.
	Utterances *[]Utterance `json:"utterances,omitempty"`

	// UtteranceCount - Number of utterances belonging to a phrase
	UtteranceCount *int `json:"utteranceCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Minertopicphrase) SetField(field string, fieldValue interface{}) {
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

func (o Minertopicphrase) MarshalJSON() ([]byte, error) {
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
	type Alias Minertopicphrase
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Topic *Minertopic `json:"topic,omitempty"`
		
		Utterances *[]Utterance `json:"utterances,omitempty"`
		
		UtteranceCount *int `json:"utteranceCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Topic: o.Topic,
		
		Utterances: o.Utterances,
		
		UtteranceCount: o.UtteranceCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Minertopicphrase) UnmarshalJSON(b []byte) error {
	var MinertopicphraseMap map[string]interface{}
	err := json.Unmarshal(b, &MinertopicphraseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MinertopicphraseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MinertopicphraseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Topic, ok := MinertopicphraseMap["topic"].(map[string]interface{}); ok {
		TopicString, _ := json.Marshal(Topic)
		json.Unmarshal(TopicString, &o.Topic)
	}
	
	if Utterances, ok := MinertopicphraseMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	
	if UtteranceCount, ok := MinertopicphraseMap["utteranceCount"].(float64); ok {
		UtteranceCountInt := int(UtteranceCount)
		o.UtteranceCount = &UtteranceCountInt
	}
	
	if SelfUri, ok := MinertopicphraseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Minertopicphrase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

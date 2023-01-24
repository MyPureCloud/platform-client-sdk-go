package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredgroup
type Unansweredgroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Label - Knowledge base unanswered group label
	Label *string `json:"label,omitempty"`

	// PhraseGroups - Represents a list of phrase groups inside an unanswered group
	PhraseGroups *[]Unansweredphrasegroup `json:"phraseGroups,omitempty"`

	// SuggestedDocuments - Represents a list of documents that may be linked to an unanswered group
	SuggestedDocuments *[]Unansweredgroupsuggesteddocument `json:"suggestedDocuments,omitempty"`

	// Statistics - Statistics object containing the various hit counts for an unanswered group
	Statistics *Knowledgegroupstatistics `json:"statistics,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Unansweredgroup) SetField(field string, fieldValue interface{}) {
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

func (o Unansweredgroup) MarshalJSON() ([]byte, error) {
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
	type Alias Unansweredgroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		PhraseGroups *[]Unansweredphrasegroup `json:"phraseGroups,omitempty"`
		
		SuggestedDocuments *[]Unansweredgroupsuggesteddocument `json:"suggestedDocuments,omitempty"`
		
		Statistics *Knowledgegroupstatistics `json:"statistics,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Label: o.Label,
		
		PhraseGroups: o.PhraseGroups,
		
		SuggestedDocuments: o.SuggestedDocuments,
		
		Statistics: o.Statistics,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Unansweredgroup) UnmarshalJSON(b []byte) error {
	var UnansweredgroupMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredgroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UnansweredgroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Label, ok := UnansweredgroupMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if PhraseGroups, ok := UnansweredgroupMap["phraseGroups"].([]interface{}); ok {
		PhraseGroupsString, _ := json.Marshal(PhraseGroups)
		json.Unmarshal(PhraseGroupsString, &o.PhraseGroups)
	}
	
	if SuggestedDocuments, ok := UnansweredgroupMap["suggestedDocuments"].([]interface{}); ok {
		SuggestedDocumentsString, _ := json.Marshal(SuggestedDocuments)
		json.Unmarshal(SuggestedDocumentsString, &o.SuggestedDocuments)
	}
	
	if Statistics, ok := UnansweredgroupMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	
	if SelfUri, ok := UnansweredgroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

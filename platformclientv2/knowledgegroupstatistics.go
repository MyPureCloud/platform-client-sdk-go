package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgegroupstatistics
type Knowledgegroupstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UnlinkedPhraseCount - Knowledge Group unique phrase count
	UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`

	// UnlinkedPhraseHitCount - Knowledge Group unlinked phrases hit count
	UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`

	// TotalPhraseHitCount - Total number of phrase hit counts of an unanswered group
	TotalPhraseHitCount *int `json:"totalPhraseHitCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgegroupstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgegroupstatistics) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgegroupstatistics
	
	return json.Marshal(&struct { 
		UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`
		
		UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`
		
		TotalPhraseHitCount *int `json:"totalPhraseHitCount,omitempty"`
		Alias
	}{ 
		UnlinkedPhraseCount: o.UnlinkedPhraseCount,
		
		UnlinkedPhraseHitCount: o.UnlinkedPhraseHitCount,
		
		TotalPhraseHitCount: o.TotalPhraseHitCount,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgegroupstatistics) UnmarshalJSON(b []byte) error {
	var KnowledgegroupstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgegroupstatisticsMap)
	if err != nil {
		return err
	}
	
	if UnlinkedPhraseCount, ok := KnowledgegroupstatisticsMap["unlinkedPhraseCount"].(float64); ok {
		UnlinkedPhraseCountInt := int(UnlinkedPhraseCount)
		o.UnlinkedPhraseCount = &UnlinkedPhraseCountInt
	}
	
	if UnlinkedPhraseHitCount, ok := KnowledgegroupstatisticsMap["unlinkedPhraseHitCount"].(float64); ok {
		UnlinkedPhraseHitCountInt := int(UnlinkedPhraseHitCount)
		o.UnlinkedPhraseHitCount = &UnlinkedPhraseHitCountInt
	}
	
	if TotalPhraseHitCount, ok := KnowledgegroupstatisticsMap["totalPhraseHitCount"].(float64); ok {
		TotalPhraseHitCountInt := int(TotalPhraseHitCount)
		o.TotalPhraseHitCount = &TotalPhraseHitCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgegroupstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

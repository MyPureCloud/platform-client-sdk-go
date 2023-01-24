package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptalternative
type Transcriptiontopictranscriptalternative struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Confidence
	Confidence *float32 `json:"confidence,omitempty"`

	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`

	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`

	// Transcript
	Transcript *string `json:"transcript,omitempty"`

	// Words
	Words *[]Transcriptiontopictranscriptword `json:"words,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transcriptiontopictranscriptalternative) SetField(field string, fieldValue interface{}) {
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

func (o Transcriptiontopictranscriptalternative) MarshalJSON() ([]byte, error) {
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
	type Alias Transcriptiontopictranscriptalternative
	
	return json.Marshal(&struct { 
		Confidence *float32 `json:"confidence,omitempty"`
		
		OffsetMs *int `json:"offsetMs,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		
		Transcript *string `json:"transcript,omitempty"`
		
		Words *[]Transcriptiontopictranscriptword `json:"words,omitempty"`
		Alias
	}{ 
		Confidence: o.Confidence,
		
		OffsetMs: o.OffsetMs,
		
		DurationMs: o.DurationMs,
		
		Transcript: o.Transcript,
		
		Words: o.Words,
		Alias:    (Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptalternative) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptalternativeMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptalternativeMap)
	if err != nil {
		return err
	}
	
	if Confidence, ok := TranscriptiontopictranscriptalternativeMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
    
	if OffsetMs, ok := TranscriptiontopictranscriptalternativeMap["offsetMs"].(float64); ok {
		OffsetMsInt := int(OffsetMs)
		o.OffsetMs = &OffsetMsInt
	}
	
	if DurationMs, ok := TranscriptiontopictranscriptalternativeMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	
	if Transcript, ok := TranscriptiontopictranscriptalternativeMap["transcript"].(string); ok {
		o.Transcript = &Transcript
	}
    
	if Words, ok := TranscriptiontopictranscriptalternativeMap["words"].([]interface{}); ok {
		WordsString, _ := json.Marshal(Words)
		json.Unmarshal(WordsString, &o.Words)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptalternative) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

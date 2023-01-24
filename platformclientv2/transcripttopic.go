package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcripttopic
type Transcripttopic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the object.
	Name *string `json:"name,omitempty"`

	// TopicPhrase - The phrase which detected the topic. 
	TopicPhrase *string `json:"topicPhrase,omitempty"`

	// TranscriptPhrase - The transcript phrase which detected the topic.
	TranscriptPhrase *string `json:"transcriptPhrase,omitempty"`

	// Confidence - The detection confidence of the topic.
	Confidence *int `json:"confidence,omitempty"`

	// StartTimeMilliseconds - The start time of the topic phrase.
	StartTimeMilliseconds *int `json:"startTimeMilliseconds,omitempty"`

	// Duration
	Duration *Topicduration `json:"duration,omitempty"`

	// Offset - Location of the phrase
	Offset *Topicoffset `json:"offset,omitempty"`

	// RecordingLocation - Location of the phrase in the recording in milliseconds
	RecordingLocation *int `json:"recordingLocation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transcripttopic) SetField(field string, fieldValue interface{}) {
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

func (o Transcripttopic) MarshalJSON() ([]byte, error) {
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
	type Alias Transcripttopic
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TopicPhrase *string `json:"topicPhrase,omitempty"`
		
		TranscriptPhrase *string `json:"transcriptPhrase,omitempty"`
		
		Confidence *int `json:"confidence,omitempty"`
		
		StartTimeMilliseconds *int `json:"startTimeMilliseconds,omitempty"`
		
		Duration *Topicduration `json:"duration,omitempty"`
		
		Offset *Topicoffset `json:"offset,omitempty"`
		
		RecordingLocation *int `json:"recordingLocation,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TopicPhrase: o.TopicPhrase,
		
		TranscriptPhrase: o.TranscriptPhrase,
		
		Confidence: o.Confidence,
		
		StartTimeMilliseconds: o.StartTimeMilliseconds,
		
		Duration: o.Duration,
		
		Offset: o.Offset,
		
		RecordingLocation: o.RecordingLocation,
		Alias:    (Alias)(o),
	})
}

func (o *Transcripttopic) UnmarshalJSON(b []byte) error {
	var TranscripttopicMap map[string]interface{}
	err := json.Unmarshal(b, &TranscripttopicMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TranscripttopicMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TranscripttopicMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TopicPhrase, ok := TranscripttopicMap["topicPhrase"].(string); ok {
		o.TopicPhrase = &TopicPhrase
	}
    
	if TranscriptPhrase, ok := TranscripttopicMap["transcriptPhrase"].(string); ok {
		o.TranscriptPhrase = &TranscriptPhrase
	}
    
	if Confidence, ok := TranscripttopicMap["confidence"].(float64); ok {
		ConfidenceInt := int(Confidence)
		o.Confidence = &ConfidenceInt
	}
	
	if StartTimeMilliseconds, ok := TranscripttopicMap["startTimeMilliseconds"].(float64); ok {
		StartTimeMillisecondsInt := int(StartTimeMilliseconds)
		o.StartTimeMilliseconds = &StartTimeMillisecondsInt
	}
	
	if Duration, ok := TranscripttopicMap["duration"].(map[string]interface{}); ok {
		DurationString, _ := json.Marshal(Duration)
		json.Unmarshal(DurationString, &o.Duration)
	}
	
	if Offset, ok := TranscripttopicMap["offset"].(map[string]interface{}); ok {
		OffsetString, _ := json.Marshal(Offset)
		json.Unmarshal(OffsetString, &o.Offset)
	}
	
	if RecordingLocation, ok := TranscripttopicMap["recordingLocation"].(float64); ok {
		RecordingLocationInt := int(RecordingLocation)
		o.RecordingLocation = &RecordingLocationInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcripttopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

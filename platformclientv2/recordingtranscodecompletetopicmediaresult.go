package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingtranscodecompletetopicmediaresult
type Recordingtranscodecompletetopicmediaresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ChannelId
	ChannelId *string `json:"channelId,omitempty"`

	// WaveUri
	WaveUri *string `json:"waveUri,omitempty"`

	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`

	// WaveformData
	WaveformData *[]float32 `json:"waveformData,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingtranscodecompletetopicmediaresult) SetField(field string, fieldValue interface{}) {
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

func (o Recordingtranscodecompletetopicmediaresult) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingtranscodecompletetopicmediaresult
	
	return json.Marshal(&struct { 
		ChannelId *string `json:"channelId,omitempty"`
		
		WaveUri *string `json:"waveUri,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		WaveformData *[]float32 `json:"waveformData,omitempty"`
		Alias
	}{ 
		ChannelId: o.ChannelId,
		
		WaveUri: o.WaveUri,
		
		MediaUri: o.MediaUri,
		
		WaveformData: o.WaveformData,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingtranscodecompletetopicmediaresult) UnmarshalJSON(b []byte) error {
	var RecordingtranscodecompletetopicmediaresultMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingtranscodecompletetopicmediaresultMap)
	if err != nil {
		return err
	}
	
	if ChannelId, ok := RecordingtranscodecompletetopicmediaresultMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
    
	if WaveUri, ok := RecordingtranscodecompletetopicmediaresultMap["waveUri"].(string); ok {
		o.WaveUri = &WaveUri
	}
    
	if MediaUri, ok := RecordingtranscodecompletetopicmediaresultMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if WaveformData, ok := RecordingtranscodecompletetopicmediaresultMap["waveformData"].([]interface{}); ok {
		WaveformDataString, _ := json.Marshal(WaveformData)
		json.Unmarshal(WaveformDataString, &o.WaveformData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingtranscodecompletetopicmediaresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

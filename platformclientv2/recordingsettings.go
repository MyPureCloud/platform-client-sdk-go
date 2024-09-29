package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingsettings
type Recordingsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MaxSimultaneousStreams - Maximum number of simultaneous screen recording streams
	MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`

	// MaxConfigurableScreenRecordingStreams - Upper limit that maxSimultaneousStreams can be configured
	MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`

	// RegionalRecordingStorageEnabled - Store call recordings in the region where they are intended to be recorded, otherwise in the organization's home region
	RegionalRecordingStorageEnabled *bool `json:"regionalRecordingStorageEnabled,omitempty"`

	// RecordingPlaybackUrlTtl - The duration in minutes for which the generated URL for recording playback remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.
	RecordingPlaybackUrlTtl *int `json:"recordingPlaybackUrlTtl,omitempty"`

	// RecordingBatchDownloadUrlTtl - The duration in minutes for which the generated URL for recording batch download remains valid.The default duration is set to 60 minutes, with a minimum allowable duration of 2 minutes and a maximum of 60 minutes.
	RecordingBatchDownloadUrlTtl *int `json:"recordingBatchDownloadUrlTtl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingsettings) SetField(field string, fieldValue interface{}) {
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

func (o Recordingsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingsettings
	
	return json.Marshal(&struct { 
		MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`
		
		MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`
		
		RegionalRecordingStorageEnabled *bool `json:"regionalRecordingStorageEnabled,omitempty"`
		
		RecordingPlaybackUrlTtl *int `json:"recordingPlaybackUrlTtl,omitempty"`
		
		RecordingBatchDownloadUrlTtl *int `json:"recordingBatchDownloadUrlTtl,omitempty"`
		Alias
	}{ 
		MaxSimultaneousStreams: o.MaxSimultaneousStreams,
		
		MaxConfigurableScreenRecordingStreams: o.MaxConfigurableScreenRecordingStreams,
		
		RegionalRecordingStorageEnabled: o.RegionalRecordingStorageEnabled,
		
		RecordingPlaybackUrlTtl: o.RecordingPlaybackUrlTtl,
		
		RecordingBatchDownloadUrlTtl: o.RecordingBatchDownloadUrlTtl,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingsettings) UnmarshalJSON(b []byte) error {
	var RecordingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingsettingsMap)
	if err != nil {
		return err
	}
	
	if MaxSimultaneousStreams, ok := RecordingsettingsMap["maxSimultaneousStreams"].(float64); ok {
		MaxSimultaneousStreamsInt := int(MaxSimultaneousStreams)
		o.MaxSimultaneousStreams = &MaxSimultaneousStreamsInt
	}
	
	if MaxConfigurableScreenRecordingStreams, ok := RecordingsettingsMap["maxConfigurableScreenRecordingStreams"].(float64); ok {
		MaxConfigurableScreenRecordingStreamsInt := int(MaxConfigurableScreenRecordingStreams)
		o.MaxConfigurableScreenRecordingStreams = &MaxConfigurableScreenRecordingStreamsInt
	}
	
	if RegionalRecordingStorageEnabled, ok := RecordingsettingsMap["regionalRecordingStorageEnabled"].(bool); ok {
		o.RegionalRecordingStorageEnabled = &RegionalRecordingStorageEnabled
	}
    
	if RecordingPlaybackUrlTtl, ok := RecordingsettingsMap["recordingPlaybackUrlTtl"].(float64); ok {
		RecordingPlaybackUrlTtlInt := int(RecordingPlaybackUrlTtl)
		o.RecordingPlaybackUrlTtl = &RecordingPlaybackUrlTtlInt
	}
	
	if RecordingBatchDownloadUrlTtl, ok := RecordingsettingsMap["recordingBatchDownloadUrlTtl"].(float64); ok {
		RecordingBatchDownloadUrlTtlInt := int(RecordingBatchDownloadUrlTtl)
		o.RecordingBatchDownloadUrlTtl = &RecordingBatchDownloadUrlTtlInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

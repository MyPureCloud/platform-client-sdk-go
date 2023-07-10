package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediastatistics
type Mediastatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`

	// DateStart - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// CreationMilliseconds - Relative milliseconds to create media session
	CreationMilliseconds *int `json:"creationMilliseconds,omitempty"`

	// PreferredRegion - Preferred media region
	PreferredRegion *string `json:"preferredRegion,omitempty"`

	// EffectiveRegion - Actual media region
	EffectiveRegion *string `json:"effectiveRegion,omitempty"`

	// MediaStatistics - Media statistics for each media endpoint
	MediaStatistics *[]Mediaendpointstatistics `json:"mediaStatistics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediastatistics) SetField(field string, fieldValue interface{}) {
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

func (o Mediastatistics) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart", }
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
	type Alias Mediastatistics
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		CommunicationId *string `json:"communicationId,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		CreationMilliseconds *int `json:"creationMilliseconds,omitempty"`
		
		PreferredRegion *string `json:"preferredRegion,omitempty"`
		
		EffectiveRegion *string `json:"effectiveRegion,omitempty"`
		
		MediaStatistics *[]Mediaendpointstatistics `json:"mediaStatistics,omitempty"`
		Alias
	}{ 
		CommunicationId: o.CommunicationId,
		
		DateStart: DateStart,
		
		CreationMilliseconds: o.CreationMilliseconds,
		
		PreferredRegion: o.PreferredRegion,
		
		EffectiveRegion: o.EffectiveRegion,
		
		MediaStatistics: o.MediaStatistics,
		Alias:    (Alias)(o),
	})
}

func (o *Mediastatistics) UnmarshalJSON(b []byte) error {
	var MediastatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &MediastatisticsMap)
	if err != nil {
		return err
	}
	
	if CommunicationId, ok := MediastatisticsMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if dateStartString, ok := MediastatisticsMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if CreationMilliseconds, ok := MediastatisticsMap["creationMilliseconds"].(float64); ok {
		CreationMillisecondsInt := int(CreationMilliseconds)
		o.CreationMilliseconds = &CreationMillisecondsInt
	}
	
	if PreferredRegion, ok := MediastatisticsMap["preferredRegion"].(string); ok {
		o.PreferredRegion = &PreferredRegion
	}
    
	if EffectiveRegion, ok := MediastatisticsMap["effectiveRegion"].(string); ok {
		o.EffectiveRegion = &EffectiveRegion
	}
    
	if MediaStatistics, ok := MediastatisticsMap["mediaStatistics"].([]interface{}); ok {
		MediaStatisticsString, _ := json.Marshal(MediaStatistics)
		json.Unmarshal(MediaStatisticsString, &o.MediaStatistics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediastatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sessionmetadataresult
type Sessionmetadataresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SessionInfo - Information about the continuous forecast session
	SessionInfo *Sessioninfo `json:"sessionInfo,omitempty"`

	// Snapshots - Captured snapshots
	Snapshots *[]Snapshots `json:"snapshots,omitempty"`

	// DateForecastStart - Start date of the forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateForecastStart *time.Time `json:"dateForecastStart,omitempty"`

	// DateHistoricalStart - Start date of the oldest available historical week. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateHistoricalStart *time.Time `json:"dateHistoricalStart,omitempty"`

	// AggregateOfferedHistoricalAvailability - Total historical availability for offered metric across all planning groups
	AggregateOfferedHistoricalAvailability *Aggregatehistoricalavailability `json:"aggregateOfferedHistoricalAvailability,omitempty"`

	// AggregateAverageHandleTimeHistoricalAvailability - Total historical availability for average handle time metric across all planning groups
	AggregateAverageHandleTimeHistoricalAvailability *Aggregatehistoricalavailability `json:"aggregateAverageHandleTimeHistoricalAvailability,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sessionmetadataresult) SetField(field string, fieldValue interface{}) {
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

func (o Sessionmetadataresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateForecastStart","DateHistoricalStart", }
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
	type Alias Sessionmetadataresult
	
	DateForecastStart := new(string)
	if o.DateForecastStart != nil {
		
		*DateForecastStart = timeutil.Strftime(o.DateForecastStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateForecastStart = nil
	}
	
	DateHistoricalStart := new(string)
	if o.DateHistoricalStart != nil {
		
		*DateHistoricalStart = timeutil.Strftime(o.DateHistoricalStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateHistoricalStart = nil
	}
	
	return json.Marshal(&struct { 
		SessionInfo *Sessioninfo `json:"sessionInfo,omitempty"`
		
		Snapshots *[]Snapshots `json:"snapshots,omitempty"`
		
		DateForecastStart *string `json:"dateForecastStart,omitempty"`
		
		DateHistoricalStart *string `json:"dateHistoricalStart,omitempty"`
		
		AggregateOfferedHistoricalAvailability *Aggregatehistoricalavailability `json:"aggregateOfferedHistoricalAvailability,omitempty"`
		
		AggregateAverageHandleTimeHistoricalAvailability *Aggregatehistoricalavailability `json:"aggregateAverageHandleTimeHistoricalAvailability,omitempty"`
		Alias
	}{ 
		SessionInfo: o.SessionInfo,
		
		Snapshots: o.Snapshots,
		
		DateForecastStart: DateForecastStart,
		
		DateHistoricalStart: DateHistoricalStart,
		
		AggregateOfferedHistoricalAvailability: o.AggregateOfferedHistoricalAvailability,
		
		AggregateAverageHandleTimeHistoricalAvailability: o.AggregateAverageHandleTimeHistoricalAvailability,
		Alias:    (Alias)(o),
	})
}

func (o *Sessionmetadataresult) UnmarshalJSON(b []byte) error {
	var SessionmetadataresultMap map[string]interface{}
	err := json.Unmarshal(b, &SessionmetadataresultMap)
	if err != nil {
		return err
	}
	
	if SessionInfo, ok := SessionmetadataresultMap["sessionInfo"].(map[string]interface{}); ok {
		SessionInfoString, _ := json.Marshal(SessionInfo)
		json.Unmarshal(SessionInfoString, &o.SessionInfo)
	}
	
	if Snapshots, ok := SessionmetadataresultMap["snapshots"].([]interface{}); ok {
		SnapshotsString, _ := json.Marshal(Snapshots)
		json.Unmarshal(SnapshotsString, &o.Snapshots)
	}
	
	if dateForecastStartString, ok := SessionmetadataresultMap["dateForecastStart"].(string); ok {
		DateForecastStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateForecastStartString)
		o.DateForecastStart = &DateForecastStart
	}
	
	if dateHistoricalStartString, ok := SessionmetadataresultMap["dateHistoricalStart"].(string); ok {
		DateHistoricalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateHistoricalStartString)
		o.DateHistoricalStart = &DateHistoricalStart
	}
	
	if AggregateOfferedHistoricalAvailability, ok := SessionmetadataresultMap["aggregateOfferedHistoricalAvailability"].(map[string]interface{}); ok {
		AggregateOfferedHistoricalAvailabilityString, _ := json.Marshal(AggregateOfferedHistoricalAvailability)
		json.Unmarshal(AggregateOfferedHistoricalAvailabilityString, &o.AggregateOfferedHistoricalAvailability)
	}
	
	if AggregateAverageHandleTimeHistoricalAvailability, ok := SessionmetadataresultMap["aggregateAverageHandleTimeHistoricalAvailability"].(map[string]interface{}); ok {
		AggregateAverageHandleTimeHistoricalAvailabilityString, _ := json.Marshal(AggregateAverageHandleTimeHistoricalAvailability)
		json.Unmarshal(AggregateAverageHandleTimeHistoricalAvailabilityString, &o.AggregateAverageHandleTimeHistoricalAvailability)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sessionmetadataresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

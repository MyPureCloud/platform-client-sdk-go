package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitsettingsresponse
type Managementunitsettingsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Adherence - Adherence settings for this management unit
	Adherence *Adherencesettings `json:"adherence,omitempty"`

	// ShortTermForecasting - Short term forecasting settings for this management unit
	ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`

	// TimeOff - Time off request settings for this management unit
	TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`

	// Scheduling - Scheduling settings for this management unit. These settings are only available if you have the permission wfm:managementUnit:view
	Scheduling *Schedulingsettingsresponse `json:"scheduling,omitempty"`

	// ShiftTrading - Shift trade settings for this management unit
	ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`

	// Metadata - Version info metadata for the associated management unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Managementunitsettingsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Managementunitsettingsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Managementunitsettingsresponse
	
	return json.Marshal(&struct { 
		Adherence *Adherencesettings `json:"adherence,omitempty"`
		
		ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`
		
		TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`
		
		Scheduling *Schedulingsettingsresponse `json:"scheduling,omitempty"`
		
		ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Adherence: o.Adherence,
		
		ShortTermForecasting: o.ShortTermForecasting,
		
		TimeOff: o.TimeOff,
		
		Scheduling: o.Scheduling,
		
		ShiftTrading: o.ShiftTrading,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Managementunitsettingsresponse) UnmarshalJSON(b []byte) error {
	var ManagementunitsettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ManagementunitsettingsresponseMap)
	if err != nil {
		return err
	}
	
	if Adherence, ok := ManagementunitsettingsresponseMap["adherence"].(map[string]interface{}); ok {
		AdherenceString, _ := json.Marshal(Adherence)
		json.Unmarshal(AdherenceString, &o.Adherence)
	}
	
	if ShortTermForecasting, ok := ManagementunitsettingsresponseMap["shortTermForecasting"].(map[string]interface{}); ok {
		ShortTermForecastingString, _ := json.Marshal(ShortTermForecasting)
		json.Unmarshal(ShortTermForecastingString, &o.ShortTermForecasting)
	}
	
	if TimeOff, ok := ManagementunitsettingsresponseMap["timeOff"].(map[string]interface{}); ok {
		TimeOffString, _ := json.Marshal(TimeOff)
		json.Unmarshal(TimeOffString, &o.TimeOff)
	}
	
	if Scheduling, ok := ManagementunitsettingsresponseMap["scheduling"].(map[string]interface{}); ok {
		SchedulingString, _ := json.Marshal(Scheduling)
		json.Unmarshal(SchedulingString, &o.Scheduling)
	}
	
	if ShiftTrading, ok := ManagementunitsettingsresponseMap["shiftTrading"].(map[string]interface{}); ok {
		ShiftTradingString, _ := json.Marshal(ShiftTrading)
		json.Unmarshal(ShiftTradingString, &o.ShiftTrading)
	}
	
	if Metadata, ok := ManagementunitsettingsresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Managementunitsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

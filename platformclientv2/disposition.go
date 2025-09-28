package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Disposition
type Disposition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - Name of the disposition. Either a platform predefined value, or the name of the disposition in the disposition table..
	Name *string `json:"name,omitempty"`

	// Analyzer - The final media analyzer result that triggered the disposition result, if any.
	Analyzer *string `json:"analyzer,omitempty"`

	// DispositionParameters - Contains various parameters related to call analysis.
	DispositionParameters *Dispositionparameters `json:"dispositionParameters,omitempty"`

	// DetectedSpeechStart - Absolute time when the speech started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DetectedSpeechStart *time.Time `json:"detectedSpeechStart,omitempty"`

	// DetectedSpeechEnd - Absolute time when the speech ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DetectedSpeechEnd *time.Time `json:"detectedSpeechEnd,omitempty"`

	// AmdTimeout - Answering Machine Detection timeout configuration.
	AmdTimeout *Dispositionamdtimeout `json:"amdTimeout,omitempty"`

	// SilentCallTimeout - Silent Call timeout configuration.
	SilentCallTimeout *Dispositionsilentcalltimeout `json:"silentCallTimeout,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Disposition) SetField(field string, fieldValue interface{}) {
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

func (o Disposition) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DetectedSpeechStart","DetectedSpeechEnd", }
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
	type Alias Disposition
	
	DetectedSpeechStart := new(string)
	if o.DetectedSpeechStart != nil {
		
		*DetectedSpeechStart = timeutil.Strftime(o.DetectedSpeechStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DetectedSpeechStart = nil
	}
	
	DetectedSpeechEnd := new(string)
	if o.DetectedSpeechEnd != nil {
		
		*DetectedSpeechEnd = timeutil.Strftime(o.DetectedSpeechEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DetectedSpeechEnd = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Analyzer *string `json:"analyzer,omitempty"`
		
		DispositionParameters *Dispositionparameters `json:"dispositionParameters,omitempty"`
		
		DetectedSpeechStart *string `json:"detectedSpeechStart,omitempty"`
		
		DetectedSpeechEnd *string `json:"detectedSpeechEnd,omitempty"`
		
		AmdTimeout *Dispositionamdtimeout `json:"amdTimeout,omitempty"`
		
		SilentCallTimeout *Dispositionsilentcalltimeout `json:"silentCallTimeout,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Analyzer: o.Analyzer,
		
		DispositionParameters: o.DispositionParameters,
		
		DetectedSpeechStart: DetectedSpeechStart,
		
		DetectedSpeechEnd: DetectedSpeechEnd,
		
		AmdTimeout: o.AmdTimeout,
		
		SilentCallTimeout: o.SilentCallTimeout,
		Alias:    (Alias)(o),
	})
}

func (o *Disposition) UnmarshalJSON(b []byte) error {
	var DispositionMap map[string]interface{}
	err := json.Unmarshal(b, &DispositionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := DispositionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Analyzer, ok := DispositionMap["analyzer"].(string); ok {
		o.Analyzer = &Analyzer
	}
    
	if DispositionParameters, ok := DispositionMap["dispositionParameters"].(map[string]interface{}); ok {
		DispositionParametersString, _ := json.Marshal(DispositionParameters)
		json.Unmarshal(DispositionParametersString, &o.DispositionParameters)
	}
	
	if detectedSpeechStartString, ok := DispositionMap["detectedSpeechStart"].(string); ok {
		DetectedSpeechStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", detectedSpeechStartString)
		o.DetectedSpeechStart = &DetectedSpeechStart
	}
	
	if detectedSpeechEndString, ok := DispositionMap["detectedSpeechEnd"].(string); ok {
		DetectedSpeechEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", detectedSpeechEndString)
		o.DetectedSpeechEnd = &DetectedSpeechEnd
	}
	
	if AmdTimeout, ok := DispositionMap["amdTimeout"].(map[string]interface{}); ok {
		AmdTimeoutString, _ := json.Marshal(AmdTimeout)
		json.Unmarshal(AmdTimeoutString, &o.AmdTimeout)
	}
	
	if SilentCallTimeout, ok := DispositionMap["silentCallTimeout"].(map[string]interface{}); ok {
		SilentCallTimeoutString, _ := json.Marshal(SilentCallTimeout)
		json.Unmarshal(SilentCallTimeoutString, &o.SilentCallTimeout)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Disposition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

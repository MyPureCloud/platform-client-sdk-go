package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainedgesoftwareupdatedto
type Domainedgesoftwareupdatedto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - Version
	Version *Domainedgesoftwareversiondto `json:"version,omitempty"`

	// MaxDownloadRate
	MaxDownloadRate *int `json:"maxDownloadRate,omitempty"`

	// DownloadStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DownloadStartTime *time.Time `json:"downloadStartTime,omitempty"`

	// ExecuteStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExecuteStartTime *time.Time `json:"executeStartTime,omitempty"`

	// ExecuteStopTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExecuteStopTime *time.Time `json:"executeStopTime,omitempty"`

	// ExecuteOnIdle
	ExecuteOnIdle *bool `json:"executeOnIdle,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`

	// CallDrainingWaitTimeSeconds
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`

	// Current
	Current *bool `json:"current,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainedgesoftwareupdatedto) SetField(field string, fieldValue interface{}) {
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

func (o Domainedgesoftwareupdatedto) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DownloadStartTime","ExecuteStartTime","ExecuteStopTime", }
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
	type Alias Domainedgesoftwareupdatedto
	
	DownloadStartTime := new(string)
	if o.DownloadStartTime != nil {
		
		*DownloadStartTime = timeutil.Strftime(o.DownloadStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DownloadStartTime = nil
	}
	
	ExecuteStartTime := new(string)
	if o.ExecuteStartTime != nil {
		
		*ExecuteStartTime = timeutil.Strftime(o.ExecuteStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExecuteStartTime = nil
	}
	
	ExecuteStopTime := new(string)
	if o.ExecuteStopTime != nil {
		
		*ExecuteStopTime = timeutil.Strftime(o.ExecuteStopTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExecuteStopTime = nil
	}
	
	return json.Marshal(&struct { 
		Version *Domainedgesoftwareversiondto `json:"version,omitempty"`
		
		MaxDownloadRate *int `json:"maxDownloadRate,omitempty"`
		
		DownloadStartTime *string `json:"downloadStartTime,omitempty"`
		
		ExecuteStartTime *string `json:"executeStartTime,omitempty"`
		
		ExecuteStopTime *string `json:"executeStopTime,omitempty"`
		
		ExecuteOnIdle *bool `json:"executeOnIdle,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		MaxDownloadRate: o.MaxDownloadRate,
		
		DownloadStartTime: DownloadStartTime,
		
		ExecuteStartTime: ExecuteStartTime,
		
		ExecuteStopTime: ExecuteStopTime,
		
		ExecuteOnIdle: o.ExecuteOnIdle,
		
		Status: o.Status,
		
		EdgeUri: o.EdgeUri,
		
		CallDrainingWaitTimeSeconds: o.CallDrainingWaitTimeSeconds,
		
		Current: o.Current,
		Alias:    (Alias)(o),
	})
}

func (o *Domainedgesoftwareupdatedto) UnmarshalJSON(b []byte) error {
	var DomainedgesoftwareupdatedtoMap map[string]interface{}
	err := json.Unmarshal(b, &DomainedgesoftwareupdatedtoMap)
	if err != nil {
		return err
	}
	
	if Version, ok := DomainedgesoftwareupdatedtoMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if MaxDownloadRate, ok := DomainedgesoftwareupdatedtoMap["maxDownloadRate"].(float64); ok {
		MaxDownloadRateInt := int(MaxDownloadRate)
		o.MaxDownloadRate = &MaxDownloadRateInt
	}
	
	if downloadStartTimeString, ok := DomainedgesoftwareupdatedtoMap["downloadStartTime"].(string); ok {
		DownloadStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", downloadStartTimeString)
		o.DownloadStartTime = &DownloadStartTime
	}
	
	if executeStartTimeString, ok := DomainedgesoftwareupdatedtoMap["executeStartTime"].(string); ok {
		ExecuteStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStartTimeString)
		o.ExecuteStartTime = &ExecuteStartTime
	}
	
	if executeStopTimeString, ok := DomainedgesoftwareupdatedtoMap["executeStopTime"].(string); ok {
		ExecuteStopTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStopTimeString)
		o.ExecuteStopTime = &ExecuteStopTime
	}
	
	if ExecuteOnIdle, ok := DomainedgesoftwareupdatedtoMap["executeOnIdle"].(bool); ok {
		o.ExecuteOnIdle = &ExecuteOnIdle
	}
    
	if Status, ok := DomainedgesoftwareupdatedtoMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if EdgeUri, ok := DomainedgesoftwareupdatedtoMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
    
	if CallDrainingWaitTimeSeconds, ok := DomainedgesoftwareupdatedtoMap["callDrainingWaitTimeSeconds"].(float64); ok {
		CallDrainingWaitTimeSecondsInt := int(CallDrainingWaitTimeSeconds)
		o.CallDrainingWaitTimeSeconds = &CallDrainingWaitTimeSecondsInt
	}
	
	if Current, ok := DomainedgesoftwareupdatedtoMap["current"].(bool); ok {
		o.Current = &Current
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareupdatedto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

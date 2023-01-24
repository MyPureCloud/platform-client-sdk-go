package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportrunentry
type Reportrunentry struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ReportId
	ReportId *string `json:"reportId,omitempty"`

	// RunTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RunTime *time.Time `json:"runTime,omitempty"`

	// RunStatus
	RunStatus *string `json:"runStatus,omitempty"`

	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// RunDurationMsec
	RunDurationMsec *int `json:"runDurationMsec,omitempty"`

	// ReportUrl
	ReportUrl *string `json:"reportUrl,omitempty"`

	// ReportFormat
	ReportFormat *string `json:"reportFormat,omitempty"`

	// ScheduleUri
	ScheduleUri *string `json:"scheduleUri,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportrunentry) SetField(field string, fieldValue interface{}) {
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

func (o Reportrunentry) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "RunTime", }
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
	type Alias Reportrunentry
	
	RunTime := new(string)
	if o.RunTime != nil {
		
		*RunTime = timeutil.Strftime(o.RunTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RunTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ReportId *string `json:"reportId,omitempty"`
		
		RunTime *string `json:"runTime,omitempty"`
		
		RunStatus *string `json:"runStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		RunDurationMsec *int `json:"runDurationMsec,omitempty"`
		
		ReportUrl *string `json:"reportUrl,omitempty"`
		
		ReportFormat *string `json:"reportFormat,omitempty"`
		
		ScheduleUri *string `json:"scheduleUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ReportId: o.ReportId,
		
		RunTime: RunTime,
		
		RunStatus: o.RunStatus,
		
		ErrorMessage: o.ErrorMessage,
		
		RunDurationMsec: o.RunDurationMsec,
		
		ReportUrl: o.ReportUrl,
		
		ReportFormat: o.ReportFormat,
		
		ScheduleUri: o.ScheduleUri,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Reportrunentry) UnmarshalJSON(b []byte) error {
	var ReportrunentryMap map[string]interface{}
	err := json.Unmarshal(b, &ReportrunentryMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportrunentryMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportrunentryMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ReportId, ok := ReportrunentryMap["reportId"].(string); ok {
		o.ReportId = &ReportId
	}
    
	if runTimeString, ok := ReportrunentryMap["runTime"].(string); ok {
		RunTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", runTimeString)
		o.RunTime = &RunTime
	}
	
	if RunStatus, ok := ReportrunentryMap["runStatus"].(string); ok {
		o.RunStatus = &RunStatus
	}
    
	if ErrorMessage, ok := ReportrunentryMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if RunDurationMsec, ok := ReportrunentryMap["runDurationMsec"].(float64); ok {
		RunDurationMsecInt := int(RunDurationMsec)
		o.RunDurationMsec = &RunDurationMsecInt
	}
	
	if ReportUrl, ok := ReportrunentryMap["reportUrl"].(string); ok {
		o.ReportUrl = &ReportUrl
	}
    
	if ReportFormat, ok := ReportrunentryMap["reportFormat"].(string); ok {
		o.ReportFormat = &ReportFormat
	}
    
	if ScheduleUri, ok := ReportrunentryMap["scheduleUri"].(string); ok {
		o.ScheduleUri = &ScheduleUri
	}
    
	if SelfUri, ok := ReportrunentryMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportrunentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

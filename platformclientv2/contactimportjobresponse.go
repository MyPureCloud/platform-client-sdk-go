package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactimportjobresponse
type Contactimportjobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// StatusDetails - Detailed description for JobStatus.
	StatusDetails *string `json:"statusDetails,omitempty"`

	// ExecutionStep - Detailed description for the Job execution state
	ExecutionStep *string `json:"executionStep,omitempty"`

	// Metadata - Metadata for the job
	Metadata *Contactimportjobmetadata `json:"metadata,omitempty"`

	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Division - Division for the job
	Division *Starrabledivision `json:"division,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// Settings - Settings
	Settings *Addressableentityref `json:"settings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactimportjobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Contactimportjobresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Contactimportjobresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StatusDetails *string `json:"statusDetails,omitempty"`
		
		ExecutionStep *string `json:"executionStep,omitempty"`
		
		Metadata *Contactimportjobmetadata `json:"metadata,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Division *Starrabledivision `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Settings *Addressableentityref `json:"settings,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		StatusDetails: o.StatusDetails,
		
		ExecutionStep: o.ExecutionStep,
		
		Metadata: o.Metadata,
		
		DateCreated: DateCreated,
		
		Division: o.Division,
		
		SelfUri: o.SelfUri,
		
		Settings: o.Settings,
		Alias:    (Alias)(o),
	})
}

func (o *Contactimportjobresponse) UnmarshalJSON(b []byte) error {
	var ContactimportjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ContactimportjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactimportjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := ContactimportjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if StatusDetails, ok := ContactimportjobresponseMap["statusDetails"].(string); ok {
		o.StatusDetails = &StatusDetails
	}
    
	if ExecutionStep, ok := ContactimportjobresponseMap["executionStep"].(string); ok {
		o.ExecutionStep = &ExecutionStep
	}
    
	if Metadata, ok := ContactimportjobresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if dateCreatedString, ok := ContactimportjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Division, ok := ContactimportjobresponseMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if SelfUri, ok := ContactimportjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Settings, ok := ContactimportjobresponseMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactimportjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

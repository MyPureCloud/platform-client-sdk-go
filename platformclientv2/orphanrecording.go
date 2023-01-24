package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Orphanrecording
type Orphanrecording struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// CreatedTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	// RecoveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecoveredTime *time.Time `json:"recoveredTime,omitempty"`

	// ProviderType
	ProviderType *string `json:"providerType,omitempty"`

	// MediaSizeBytes
	MediaSizeBytes *int `json:"mediaSizeBytes,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// FileState
	FileState *string `json:"fileState,omitempty"`

	// ProviderEndpoint
	ProviderEndpoint *Endpoint `json:"providerEndpoint,omitempty"`

	// Recording
	Recording *Recording `json:"recording,omitempty"`

	// OrphanStatus - The status of the orphaned recording's conversation.
	OrphanStatus *string `json:"orphanStatus,omitempty"`

	// SourceOrphaningId - An identifier used during recovery operations by the supplying hybrid platform to track back and determine which interaction this recording is associated with
	SourceOrphaningId *string `json:"sourceOrphaningId,omitempty"`

	// Region
	Region *string `json:"region,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Orphanrecording) SetField(field string, fieldValue interface{}) {
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

func (o Orphanrecording) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedTime","RecoveredTime", }
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
	type Alias Orphanrecording
	
	CreatedTime := new(string)
	if o.CreatedTime != nil {
		
		*CreatedTime = timeutil.Strftime(o.CreatedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedTime = nil
	}
	
	RecoveredTime := new(string)
	if o.RecoveredTime != nil {
		
		*RecoveredTime = timeutil.Strftime(o.RecoveredTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecoveredTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CreatedTime *string `json:"createdTime,omitempty"`
		
		RecoveredTime *string `json:"recoveredTime,omitempty"`
		
		ProviderType *string `json:"providerType,omitempty"`
		
		MediaSizeBytes *int `json:"mediaSizeBytes,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		ProviderEndpoint *Endpoint `json:"providerEndpoint,omitempty"`
		
		Recording *Recording `json:"recording,omitempty"`
		
		OrphanStatus *string `json:"orphanStatus,omitempty"`
		
		SourceOrphaningId *string `json:"sourceOrphaningId,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CreatedTime: CreatedTime,
		
		RecoveredTime: RecoveredTime,
		
		ProviderType: o.ProviderType,
		
		MediaSizeBytes: o.MediaSizeBytes,
		
		MediaType: o.MediaType,
		
		FileState: o.FileState,
		
		ProviderEndpoint: o.ProviderEndpoint,
		
		Recording: o.Recording,
		
		OrphanStatus: o.OrphanStatus,
		
		SourceOrphaningId: o.SourceOrphaningId,
		
		Region: o.Region,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Orphanrecording) UnmarshalJSON(b []byte) error {
	var OrphanrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &OrphanrecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OrphanrecordingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OrphanrecordingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if createdTimeString, ok := OrphanrecordingMap["createdTime"].(string); ok {
		CreatedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdTimeString)
		o.CreatedTime = &CreatedTime
	}
	
	if recoveredTimeString, ok := OrphanrecordingMap["recoveredTime"].(string); ok {
		RecoveredTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", recoveredTimeString)
		o.RecoveredTime = &RecoveredTime
	}
	
	if ProviderType, ok := OrphanrecordingMap["providerType"].(string); ok {
		o.ProviderType = &ProviderType
	}
    
	if MediaSizeBytes, ok := OrphanrecordingMap["mediaSizeBytes"].(float64); ok {
		MediaSizeBytesInt := int(MediaSizeBytes)
		o.MediaSizeBytes = &MediaSizeBytesInt
	}
	
	if MediaType, ok := OrphanrecordingMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if FileState, ok := OrphanrecordingMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
    
	if ProviderEndpoint, ok := OrphanrecordingMap["providerEndpoint"].(map[string]interface{}); ok {
		ProviderEndpointString, _ := json.Marshal(ProviderEndpoint)
		json.Unmarshal(ProviderEndpointString, &o.ProviderEndpoint)
	}
	
	if Recording, ok := OrphanrecordingMap["recording"].(map[string]interface{}); ok {
		RecordingString, _ := json.Marshal(Recording)
		json.Unmarshal(RecordingString, &o.Recording)
	}
	
	if OrphanStatus, ok := OrphanrecordingMap["orphanStatus"].(string); ok {
		o.OrphanStatus = &OrphanStatus
	}
    
	if SourceOrphaningId, ok := OrphanrecordingMap["sourceOrphaningId"].(string); ok {
		o.SourceOrphaningId = &SourceOrphaningId
	}
    
	if Region, ok := OrphanrecordingMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if SelfUri, ok := OrphanrecordingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Orphanrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

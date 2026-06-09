package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3synchronizationuploadurlresponse
type V3synchronizationuploadurlresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FileId - The unique identifier for the upload object.
	FileId *string `json:"fileId,omitempty"`

	// FileName - Name of the uploaded file.
	FileName *string `json:"fileName,omitempty"`

	// Metadata - The metadata of the uploaded file
	Metadata *V3synchronizationuploadmetadata `json:"metadata,omitempty"`

	// Synchronization - The synchronization of the file upload.
	Synchronization *V3synchronizationref `json:"synchronization,omitempty"`

	// Url - Pre-signed URL to PUT the file to.
	Url *string `json:"url,omitempty"`

	// Headers - Required headers when uploading a file through PUT request to the URL.
	Headers *map[string]string `json:"headers,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3synchronizationuploadurlresponse) SetField(field string, fieldValue interface{}) {
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

func (o V3synchronizationuploadurlresponse) MarshalJSON() ([]byte, error) {
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
	type Alias V3synchronizationuploadurlresponse
	
	return json.Marshal(&struct { 
		FileId *string `json:"fileId,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		Metadata *V3synchronizationuploadmetadata `json:"metadata,omitempty"`
		
		Synchronization *V3synchronizationref `json:"synchronization,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		FileId: o.FileId,
		
		FileName: o.FileName,
		
		Metadata: o.Metadata,
		
		Synchronization: o.Synchronization,
		
		Url: o.Url,
		
		Headers: o.Headers,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *V3synchronizationuploadurlresponse) UnmarshalJSON(b []byte) error {
	var V3synchronizationuploadurlresponseMap map[string]interface{}
	err := json.Unmarshal(b, &V3synchronizationuploadurlresponseMap)
	if err != nil {
		return err
	}
	
	if FileId, ok := V3synchronizationuploadurlresponseMap["fileId"].(string); ok {
		o.FileId = &FileId
	}
    
	if FileName, ok := V3synchronizationuploadurlresponseMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if Metadata, ok := V3synchronizationuploadurlresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if Synchronization, ok := V3synchronizationuploadurlresponseMap["synchronization"].(map[string]interface{}); ok {
		SynchronizationString, _ := json.Marshal(Synchronization)
		json.Unmarshal(SynchronizationString, &o.Synchronization)
	}
	
	if Url, ok := V3synchronizationuploadurlresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Headers, ok := V3synchronizationuploadurlresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	
	if SelfUri, ok := V3synchronizationuploadurlresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

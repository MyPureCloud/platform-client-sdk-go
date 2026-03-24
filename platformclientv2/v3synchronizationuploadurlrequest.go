package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3synchronizationuploadurlrequest
type V3synchronizationuploadurlrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	FileName *string `json:"fileName,omitempty"`

	// ContentMd5 - Content MD5 of the file to upload encoded in base64, example: \"f8VicOenD6gaWTW3Lqy+KQ==\". Not the hexadecimal representation as \"7fc56270e7a70fa81a5935b72eacbe29\".
	ContentMd5 *string `json:"contentMd5,omitempty"`

	// ContentType - The content type of the file to upload
	ContentType *string `json:"contentType,omitempty"`

	// ContentLength - The length of the file to upload in bytes
	ContentLength *int `json:"contentLength,omitempty"`

	// Metadata - The metadata of the file to upload
	Metadata *V3synchronizationuploadmetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3synchronizationuploadurlrequest) SetField(field string, fieldValue interface{}) {
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

func (o V3synchronizationuploadurlrequest) MarshalJSON() ([]byte, error) {
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
	type Alias V3synchronizationuploadurlrequest
	
	return json.Marshal(&struct { 
		FileName *string `json:"fileName,omitempty"`
		
		ContentMd5 *string `json:"contentMd5,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Metadata *V3synchronizationuploadmetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		FileName: o.FileName,
		
		ContentMd5: o.ContentMd5,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *V3synchronizationuploadurlrequest) UnmarshalJSON(b []byte) error {
	var V3synchronizationuploadurlrequestMap map[string]interface{}
	err := json.Unmarshal(b, &V3synchronizationuploadurlrequestMap)
	if err != nil {
		return err
	}
	
	if FileName, ok := V3synchronizationuploadurlrequestMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if ContentMd5, ok := V3synchronizationuploadurlrequestMap["contentMd5"].(string); ok {
		o.ContentMd5 = &ContentMd5
	}
    
	if ContentType, ok := V3synchronizationuploadurlrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := V3synchronizationuploadurlrequestMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Metadata, ok := V3synchronizationuploadurlrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

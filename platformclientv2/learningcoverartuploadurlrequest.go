package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningcoverartuploadurlrequest
type Learningcoverartuploadurlrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	FileName *string `json:"fileName,omitempty"`

	// ContentMd5 - Content MD5 of the file to upload
	ContentMd5 *string `json:"contentMd5,omitempty"`

	// SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`

	// ContentType - The content type of the file to upload.
	ContentType *string `json:"contentType,omitempty"`

	// ServerSideEncryption
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningcoverartuploadurlrequest) SetField(field string, fieldValue interface{}) {
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

func (o Learningcoverartuploadurlrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Learningcoverartuploadurlrequest
	
	return json.Marshal(&struct { 
		FileName *string `json:"fileName,omitempty"`
		
		ContentMd5 *string `json:"contentMd5,omitempty"`
		
		SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ServerSideEncryption *string `json:"serverSideEncryption,omitempty"`
		Alias
	}{ 
		FileName: o.FileName,
		
		ContentMd5: o.ContentMd5,
		
		SignedUrlTimeoutSeconds: o.SignedUrlTimeoutSeconds,
		
		ContentType: o.ContentType,
		
		ServerSideEncryption: o.ServerSideEncryption,
		Alias:    (Alias)(o),
	})
}

func (o *Learningcoverartuploadurlrequest) UnmarshalJSON(b []byte) error {
	var LearningcoverartuploadurlrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningcoverartuploadurlrequestMap)
	if err != nil {
		return err
	}
	
	if FileName, ok := LearningcoverartuploadurlrequestMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if ContentMd5, ok := LearningcoverartuploadurlrequestMap["contentMd5"].(string); ok {
		o.ContentMd5 = &ContentMd5
	}
    
	if SignedUrlTimeoutSeconds, ok := LearningcoverartuploadurlrequestMap["signedUrlTimeoutSeconds"].(float64); ok {
		SignedUrlTimeoutSecondsInt := int(SignedUrlTimeoutSeconds)
		o.SignedUrlTimeoutSeconds = &SignedUrlTimeoutSecondsInt
	}
	
	if ContentType, ok := LearningcoverartuploadurlrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ServerSideEncryption, ok := LearningcoverartuploadurlrequestMap["serverSideEncryption"].(string); ok {
		o.ServerSideEncryption = &ServerSideEncryption
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningcoverartuploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

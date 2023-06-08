package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentcontentupload
type Knowledgedocumentcontentupload struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ContentType - Type of Article Content.
	ContentType *string `json:"contentType,omitempty"`

	// FileName - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	FileName *string `json:"fileName,omitempty"`

	// Status - Status of the upload operation
	Status *string `json:"status,omitempty"`

	// UploadKey - Key that identifies the file in the storage including the file name
	UploadKey *string `json:"uploadKey,omitempty"`

	// Url - Presigned URL to PUT the file to
	Url *string `json:"url,omitempty"`

	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`

	// Document - ID of the document for which article content is to be uploaded
	Document *Addressableentityref `json:"document,omitempty"`

	// ErrorMessage - Error message when upload fails
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentcontentupload) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentcontentupload) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentcontentupload
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		
		Document *Addressableentityref `json:"document,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ContentType: o.ContentType,
		
		FileName: o.FileName,
		
		Status: o.Status,
		
		UploadKey: o.UploadKey,
		
		Url: o.Url,
		
		Headers: o.Headers,
		
		Document: o.Document,
		
		ErrorMessage: o.ErrorMessage,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentcontentupload) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentcontentuploadMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentcontentuploadMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentcontentuploadMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ContentType, ok := KnowledgedocumentcontentuploadMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if FileName, ok := KnowledgedocumentcontentuploadMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if Status, ok := KnowledgedocumentcontentuploadMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if UploadKey, ok := KnowledgedocumentcontentuploadMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if Url, ok := KnowledgedocumentcontentuploadMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Headers, ok := KnowledgedocumentcontentuploadMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	
	if Document, ok := KnowledgedocumentcontentuploadMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if ErrorMessage, ok := KnowledgedocumentcontentuploadMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if SelfUri, ok := KnowledgedocumentcontentuploadMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentcontentupload) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

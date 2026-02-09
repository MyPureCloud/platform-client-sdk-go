package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistuploadurlrequest
type Contactlistuploadurlrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`

	// ContentType - The content type of the file to upload. Allows MIME types are text/csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
	ContentType *string `json:"contentType,omitempty"`

	// Id - Id of your contact list to upload to
	Id *string `json:"id,omitempty"`

	// ContactIdName - The column name from your file to use as the contact id.
	ContactIdName *string `json:"contactIdName,omitempty"`

	// ImportTemplateId - Id of your import template to use.
	ImportTemplateId *string `json:"importTemplateId,omitempty"`

	// ListNamePrefix - String that will replace %N in the listNameFormat specified on the import template.
	ListNamePrefix *string `json:"listNamePrefix,omitempty"`

	// ClearSystemData - Whether to clear system data
	ClearSystemData *bool `json:"clearSystemData,omitempty"`

	// DivisionIdForTargetContactLists - Id of the division to be used for the creation of the target contact lists. If not provided, Home division will be used.
	DivisionIdForTargetContactLists *string `json:"divisionIdForTargetContactLists,omitempty"`

	// FileSpecificationTemplateId - File specification template ID
	FileSpecificationTemplateId *string `json:"fileSpecificationTemplateId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactlistuploadurlrequest) SetField(field string, fieldValue interface{}) {
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

func (o Contactlistuploadurlrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Contactlistuploadurlrequest
	
	return json.Marshal(&struct { 
		SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		ContactIdName *string `json:"contactIdName,omitempty"`
		
		ImportTemplateId *string `json:"importTemplateId,omitempty"`
		
		ListNamePrefix *string `json:"listNamePrefix,omitempty"`
		
		ClearSystemData *bool `json:"clearSystemData,omitempty"`
		
		DivisionIdForTargetContactLists *string `json:"divisionIdForTargetContactLists,omitempty"`
		
		FileSpecificationTemplateId *string `json:"fileSpecificationTemplateId,omitempty"`
		Alias
	}{ 
		SignedUrlTimeoutSeconds: o.SignedUrlTimeoutSeconds,
		
		ContentType: o.ContentType,
		
		Id: o.Id,
		
		ContactIdName: o.ContactIdName,
		
		ImportTemplateId: o.ImportTemplateId,
		
		ListNamePrefix: o.ListNamePrefix,
		
		ClearSystemData: o.ClearSystemData,
		
		DivisionIdForTargetContactLists: o.DivisionIdForTargetContactLists,
		
		FileSpecificationTemplateId: o.FileSpecificationTemplateId,
		Alias:    (Alias)(o),
	})
}

func (o *Contactlistuploadurlrequest) UnmarshalJSON(b []byte) error {
	var ContactlistuploadurlrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistuploadurlrequestMap)
	if err != nil {
		return err
	}
	
	if SignedUrlTimeoutSeconds, ok := ContactlistuploadurlrequestMap["signedUrlTimeoutSeconds"].(float64); ok {
		SignedUrlTimeoutSecondsInt := int(SignedUrlTimeoutSeconds)
		o.SignedUrlTimeoutSeconds = &SignedUrlTimeoutSecondsInt
	}
	
	if ContentType, ok := ContactlistuploadurlrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Id, ok := ContactlistuploadurlrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ContactIdName, ok := ContactlistuploadurlrequestMap["contactIdName"].(string); ok {
		o.ContactIdName = &ContactIdName
	}
    
	if ImportTemplateId, ok := ContactlistuploadurlrequestMap["importTemplateId"].(string); ok {
		o.ImportTemplateId = &ImportTemplateId
	}
    
	if ListNamePrefix, ok := ContactlistuploadurlrequestMap["listNamePrefix"].(string); ok {
		o.ListNamePrefix = &ListNamePrefix
	}
    
	if ClearSystemData, ok := ContactlistuploadurlrequestMap["clearSystemData"].(bool); ok {
		o.ClearSystemData = &ClearSystemData
	}
    
	if DivisionIdForTargetContactLists, ok := ContactlistuploadurlrequestMap["divisionIdForTargetContactLists"].(string); ok {
		o.DivisionIdForTargetContactLists = &DivisionIdForTargetContactLists
	}
    
	if FileSpecificationTemplateId, ok := ContactlistuploadurlrequestMap["fileSpecificationTemplateId"].(string); ok {
		o.FileSpecificationTemplateId = &FileSpecificationTemplateId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistuploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

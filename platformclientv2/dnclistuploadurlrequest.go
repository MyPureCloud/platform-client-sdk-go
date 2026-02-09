package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistuploadurlrequest
type Dnclistuploadurlrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`

	// ContentType - The content type of the file to upload. Allows all types are text/csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
	ContentType *string `json:"contentType,omitempty"`

	// Id - Id of your dnc list to upload to
	Id *string `json:"id,omitempty"`

	// PhoneColumns - The column names from your file from which to upload dnc phone numbers.
	PhoneColumns *string `json:"phoneColumns,omitempty"`

	// EmailColumns - The column names from your file from which to upload dnc emails.
	EmailColumns *string `json:"emailColumns,omitempty"`

	// CustomExclusionColumns - The column names from your file from which to upload dnc custom exclusion column entries.
	CustomExclusionColumns *string `json:"customExclusionColumns,omitempty"`

	// ExpirationDateTimeColumn - The column name from your file to use as the dnc expiration date time.
	ExpirationDateTimeColumn *string `json:"expirationDateTimeColumn,omitempty"`

	// WhatsAppColumns - The column names from your file from which to upload dnc whatsapp.
	WhatsAppColumns *string `json:"whatsAppColumns,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dnclistuploadurlrequest) SetField(field string, fieldValue interface{}) {
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

func (o Dnclistuploadurlrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Dnclistuploadurlrequest
	
	return json.Marshal(&struct { 
		SignedUrlTimeoutSeconds *int `json:"signedUrlTimeoutSeconds,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		PhoneColumns *string `json:"phoneColumns,omitempty"`
		
		EmailColumns *string `json:"emailColumns,omitempty"`
		
		CustomExclusionColumns *string `json:"customExclusionColumns,omitempty"`
		
		ExpirationDateTimeColumn *string `json:"expirationDateTimeColumn,omitempty"`
		
		WhatsAppColumns *string `json:"whatsAppColumns,omitempty"`
		Alias
	}{ 
		SignedUrlTimeoutSeconds: o.SignedUrlTimeoutSeconds,
		
		ContentType: o.ContentType,
		
		Id: o.Id,
		
		PhoneColumns: o.PhoneColumns,
		
		EmailColumns: o.EmailColumns,
		
		CustomExclusionColumns: o.CustomExclusionColumns,
		
		ExpirationDateTimeColumn: o.ExpirationDateTimeColumn,
		
		WhatsAppColumns: o.WhatsAppColumns,
		Alias:    (Alias)(o),
	})
}

func (o *Dnclistuploadurlrequest) UnmarshalJSON(b []byte) error {
	var DnclistuploadurlrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistuploadurlrequestMap)
	if err != nil {
		return err
	}
	
	if SignedUrlTimeoutSeconds, ok := DnclistuploadurlrequestMap["signedUrlTimeoutSeconds"].(float64); ok {
		SignedUrlTimeoutSecondsInt := int(SignedUrlTimeoutSeconds)
		o.SignedUrlTimeoutSeconds = &SignedUrlTimeoutSecondsInt
	}
	
	if ContentType, ok := DnclistuploadurlrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Id, ok := DnclistuploadurlrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PhoneColumns, ok := DnclistuploadurlrequestMap["phoneColumns"].(string); ok {
		o.PhoneColumns = &PhoneColumns
	}
    
	if EmailColumns, ok := DnclistuploadurlrequestMap["emailColumns"].(string); ok {
		o.EmailColumns = &EmailColumns
	}
    
	if CustomExclusionColumns, ok := DnclistuploadurlrequestMap["customExclusionColumns"].(string); ok {
		o.CustomExclusionColumns = &CustomExclusionColumns
	}
    
	if ExpirationDateTimeColumn, ok := DnclistuploadurlrequestMap["expirationDateTimeColumn"].(string); ok {
		o.ExpirationDateTimeColumn = &ExpirationDateTimeColumn
	}
    
	if WhatsAppColumns, ok := DnclistuploadurlrequestMap["whatsAppColumns"].(string); ok {
		o.WhatsAppColumns = &WhatsAppColumns
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclistuploadurlrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

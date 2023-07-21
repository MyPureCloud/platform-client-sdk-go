package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobresponse
type Knowledgeimportjobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Id of the import job
	Id *string `json:"id,omitempty"`

	// UploadKey - Upload key
	UploadKey *string `json:"uploadKey,omitempty"`

	// FileType - File type of the document
	FileType *string `json:"fileType,omitempty"`

	// Settings - Additional optional settings
	Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`

	// Status - Status of the import job
	Status *string `json:"status,omitempty"`

	// Report - Report of the import job
	Report *Knowledgeimportjobreport `json:"report,omitempty"`

	// KnowledgeBase - Knowledge base which document import does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`

	// DateCreated - Created date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SkipConfirmationStep - If enabled pre-validation step will be skipped.
	SkipConfirmationStep *bool `json:"skipConfirmationStep,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeimportjobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeimportjobresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Knowledgeimportjobresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		
		Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Report *Knowledgeimportjobreport `json:"report,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SkipConfirmationStep *bool `json:"skipConfirmationStep,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		UploadKey: o.UploadKey,
		
		FileType: o.FileType,
		
		Settings: o.Settings,
		
		Status: o.Status,
		
		Report: o.Report,
		
		KnowledgeBase: o.KnowledgeBase,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SkipConfirmationStep: o.SkipConfirmationStep,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeimportjobresponse) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeimportjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UploadKey, ok := KnowledgeimportjobresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if FileType, ok := KnowledgeimportjobresponseMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    
	if Settings, ok := KnowledgeimportjobresponseMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if Status, ok := KnowledgeimportjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Report, ok := KnowledgeimportjobresponseMap["report"].(map[string]interface{}); ok {
		ReportString, _ := json.Marshal(Report)
		json.Unmarshal(ReportString, &o.Report)
	}
	
	if KnowledgeBase, ok := KnowledgeimportjobresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if dateCreatedString, ok := KnowledgeimportjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeimportjobresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SkipConfirmationStep, ok := KnowledgeimportjobresponseMap["skipConfirmationStep"].(bool); ok {
		o.SkipConfirmationStep = &SkipConfirmationStep
	}
    
	if SelfUri, ok := KnowledgeimportjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

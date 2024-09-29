package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesyncjobresponse
type Knowledgesyncjobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Id of the sync job.
	Id *string `json:"id,omitempty"`

	// UploadKey
	UploadKey *string `json:"uploadKey,omitempty"`

	// Status - The status of the export job.
	Status *string `json:"status,omitempty"`

	// Report - Report of the sync job
	Report *Knowledgesyncjobreport `json:"report,omitempty"`

	// KnowledgeBase - Knowledge base which document export belongs to.
	KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`

	// DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The timestamp of when the export stopped. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - The user who created the operation
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DownloadURL - The URL of the location at which the caller can download the sync file, when available.
	DownloadURL *string `json:"downloadURL,omitempty"`

	// FailedEntitiesURL - The URL of the location at which the caller can download the entities in json format that failed during the sync.
	FailedEntitiesURL *string `json:"failedEntitiesURL,omitempty"`

	// Source - Source of the sync job.
	Source *Knowledgeoperationsource `json:"source,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesyncjobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesyncjobresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesyncjobresponse
	
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
		
		Status *string `json:"status,omitempty"`
		
		Report *Knowledgesyncjobreport `json:"report,omitempty"`
		
		KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DownloadURL *string `json:"downloadURL,omitempty"`
		
		FailedEntitiesURL *string `json:"failedEntitiesURL,omitempty"`
		
		Source *Knowledgeoperationsource `json:"source,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		UploadKey: o.UploadKey,
		
		Status: o.Status,
		
		Report: o.Report,
		
		KnowledgeBase: o.KnowledgeBase,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		DownloadURL: o.DownloadURL,
		
		FailedEntitiesURL: o.FailedEntitiesURL,
		
		Source: o.Source,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesyncjobresponse) UnmarshalJSON(b []byte) error {
	var KnowledgesyncjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesyncjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgesyncjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UploadKey, ok := KnowledgesyncjobresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if Status, ok := KnowledgesyncjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Report, ok := KnowledgesyncjobresponseMap["report"].(map[string]interface{}); ok {
		ReportString, _ := json.Marshal(Report)
		json.Unmarshal(ReportString, &o.Report)
	}
	
	if KnowledgeBase, ok := KnowledgesyncjobresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if dateCreatedString, ok := KnowledgesyncjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgesyncjobresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := KnowledgesyncjobresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if DownloadURL, ok := KnowledgesyncjobresponseMap["downloadURL"].(string); ok {
		o.DownloadURL = &DownloadURL
	}
    
	if FailedEntitiesURL, ok := KnowledgesyncjobresponseMap["failedEntitiesURL"].(string); ok {
		o.FailedEntitiesURL = &FailedEntitiesURL
	}
    
	if Source, ok := KnowledgesyncjobresponseMap["source"].(map[string]interface{}); ok {
		SourceString, _ := json.Marshal(Source)
		json.Unmarshal(SourceString, &o.Source)
	}
	
	if SelfUri, ok := KnowledgesyncjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

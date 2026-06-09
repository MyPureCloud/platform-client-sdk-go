package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontableimportjob - State of a decision table row import job
type Decisiontableimportjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// TableVersion - The table version to be replaced by this import
	TableVersion *int `json:"tableVersion,omitempty"`

	// Status - Current status of the import job
	Status *string `json:"status,omitempty"`

	// UploadUrl - Pre-signed URL to upload the import file (PUT)
	UploadUrl *string `json:"uploadUrl,omitempty"`

	// UploadHeaders - Headers required when uploading file with data to be imported to uploadUrl
	UploadHeaders *map[string]string `json:"uploadHeaders,omitempty"`

	// ImportMode - Whether rows are appended to existing rows or rows are replaced
	ImportMode *string `json:"importMode,omitempty"`

	// FileName - Original file name supplied when the job was created, including the file extension
	FileName *string `json:"fileName,omitempty"`

	// CreatedBy - The user who created the job
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`

	// DateCreated - When the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - When the job was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateCompleted - When processing finished, successfully or not. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

	// DateExpires - When upload credentials expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`

	// RowMetrics - Row-level metrics populated incrementally during import processing
	RowMetrics *Decisiontableimportrowmetrics `json:"rowMetrics,omitempty"`

	// VarError - Present when the import job could not be successfully finished
	VarError *Decisiontableimportjoberror `json:"error,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontableimportjob) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontableimportjob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateCompleted","DateExpires", }
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
	type Alias Decisiontableimportjob
	
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
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TableVersion *int `json:"tableVersion,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		UploadUrl *string `json:"uploadUrl,omitempty"`
		
		UploadHeaders *map[string]string `json:"uploadHeaders,omitempty"`
		
		ImportMode *string `json:"importMode,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		RowMetrics *Decisiontableimportrowmetrics `json:"rowMetrics,omitempty"`
		
		VarError *Decisiontableimportjoberror `json:"error,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		TableVersion: o.TableVersion,
		
		Status: o.Status,
		
		UploadUrl: o.UploadUrl,
		
		UploadHeaders: o.UploadHeaders,
		
		ImportMode: o.ImportMode,
		
		FileName: o.FileName,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateCompleted: DateCompleted,
		
		DateExpires: DateExpires,
		
		RowMetrics: o.RowMetrics,
		
		VarError: o.VarError,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontableimportjob) UnmarshalJSON(b []byte) error {
	var DecisiontableimportjobMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontableimportjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DecisiontableimportjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if TableVersion, ok := DecisiontableimportjobMap["tableVersion"].(float64); ok {
		TableVersionInt := int(TableVersion)
		o.TableVersion = &TableVersionInt
	}
	
	if Status, ok := DecisiontableimportjobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if UploadUrl, ok := DecisiontableimportjobMap["uploadUrl"].(string); ok {
		o.UploadUrl = &UploadUrl
	}
    
	if UploadHeaders, ok := DecisiontableimportjobMap["uploadHeaders"].(map[string]interface{}); ok {
		UploadHeadersString, _ := json.Marshal(UploadHeaders)
		json.Unmarshal(UploadHeadersString, &o.UploadHeaders)
	}
	
	if ImportMode, ok := DecisiontableimportjobMap["importMode"].(string); ok {
		o.ImportMode = &ImportMode
	}
    
	if FileName, ok := DecisiontableimportjobMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if CreatedBy, ok := DecisiontableimportjobMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := DecisiontableimportjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DecisiontableimportjobMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateCompletedString, ok := DecisiontableimportjobMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if dateExpiresString, ok := DecisiontableimportjobMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if RowMetrics, ok := DecisiontableimportjobMap["rowMetrics"].(map[string]interface{}); ok {
		RowMetricsString, _ := json.Marshal(RowMetrics)
		json.Unmarshal(RowMetricsString, &o.RowMetrics)
	}
	
	if VarError, ok := DecisiontableimportjobMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if SelfUri, ok := DecisiontableimportjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontableimportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

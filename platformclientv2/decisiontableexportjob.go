package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisiontableexportjob
type Decisiontableexportjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// TableVersion - The version of the decision table that was exported.
	TableVersion *int `json:"tableVersion,omitempty"`

	// Status - Current status of the export job.
	Status *string `json:"status,omitempty"`

	// CreatedBy - The user who created the export job.
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`

	// DateCreated - Date when this export job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date when this export job was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// FileName - Name of the exported file.
	FileName *string `json:"fileName,omitempty"`

	// Download - Reference to the download resource for obtaining the exported file.
	Download *Addressableentityref `json:"download,omitempty"`

	// DateDownloadExpires - Date when the download link expires. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDownloadExpires *time.Time `json:"dateDownloadExpires,omitempty"`

	// ExportType - The type of export that was performed.
	ExportType *string `json:"exportType,omitempty"`

	// TotalRows - Total number of rows to export (set when row loading begins).
	TotalRows *int `json:"totalRows,omitempty"`

	// RowsExported - The number of rows exported.
	RowsExported *int `json:"rowsExported,omitempty"`

	// Format - The format of the exported file.
	Format *string `json:"format,omitempty"`

	// VarError - Error details if the export job failed.
	VarError *Decisiontableexportjoberror `json:"error,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisiontableexportjob) SetField(field string, fieldValue interface{}) {
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

func (o Decisiontableexportjob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateDownloadExpires", }
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
	type Alias Decisiontableexportjob
	
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
	
	DateDownloadExpires := new(string)
	if o.DateDownloadExpires != nil {
		
		*DateDownloadExpires = timeutil.Strftime(o.DateDownloadExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDownloadExpires = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TableVersion *int `json:"tableVersion,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		Download *Addressableentityref `json:"download,omitempty"`
		
		DateDownloadExpires *string `json:"dateDownloadExpires,omitempty"`
		
		ExportType *string `json:"exportType,omitempty"`
		
		TotalRows *int `json:"totalRows,omitempty"`
		
		RowsExported *int `json:"rowsExported,omitempty"`
		
		Format *string `json:"format,omitempty"`
		
		VarError *Decisiontableexportjoberror `json:"error,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		TableVersion: o.TableVersion,
		
		Status: o.Status,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		FileName: o.FileName,
		
		Download: o.Download,
		
		DateDownloadExpires: DateDownloadExpires,
		
		ExportType: o.ExportType,
		
		TotalRows: o.TotalRows,
		
		RowsExported: o.RowsExported,
		
		Format: o.Format,
		
		VarError: o.VarError,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Decisiontableexportjob) UnmarshalJSON(b []byte) error {
	var DecisiontableexportjobMap map[string]interface{}
	err := json.Unmarshal(b, &DecisiontableexportjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DecisiontableexportjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if TableVersion, ok := DecisiontableexportjobMap["tableVersion"].(float64); ok {
		TableVersionInt := int(TableVersion)
		o.TableVersion = &TableVersionInt
	}
	
	if Status, ok := DecisiontableexportjobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if CreatedBy, ok := DecisiontableexportjobMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := DecisiontableexportjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DecisiontableexportjobMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if FileName, ok := DecisiontableexportjobMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if Download, ok := DecisiontableexportjobMap["download"].(map[string]interface{}); ok {
		DownloadString, _ := json.Marshal(Download)
		json.Unmarshal(DownloadString, &o.Download)
	}
	
	if dateDownloadExpiresString, ok := DecisiontableexportjobMap["dateDownloadExpires"].(string); ok {
		DateDownloadExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDownloadExpiresString)
		o.DateDownloadExpires = &DateDownloadExpires
	}
	
	if ExportType, ok := DecisiontableexportjobMap["exportType"].(string); ok {
		o.ExportType = &ExportType
	}
    
	if TotalRows, ok := DecisiontableexportjobMap["totalRows"].(float64); ok {
		TotalRowsInt := int(TotalRows)
		o.TotalRows = &TotalRowsInt
	}
	
	if RowsExported, ok := DecisiontableexportjobMap["rowsExported"].(float64); ok {
		RowsExportedInt := int(RowsExported)
		o.RowsExported = &RowsExportedInt
	}
	
	if Format, ok := DecisiontableexportjobMap["format"].(string); ok {
		o.Format = &Format
	}
    
	if VarError, ok := DecisiontableexportjobMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if SelfUri, ok := DecisiontableexportjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Decisiontableexportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

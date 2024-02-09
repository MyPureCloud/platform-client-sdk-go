package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeparsejobresponse
type Knowledgeparsejobresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Id of the parse job
	Id *string `json:"id,omitempty"`

	// DownloadURL - The URL of the location at which the caller can download the original html file.
	DownloadURL *string `json:"downloadURL,omitempty"`

	// Hints - Hinted titles for the parser.
	Hints *[]string `json:"hints,omitempty"`

	// Status - Status of the parse job
	Status *string `json:"status,omitempty"`

	// ParseResults - Results of the parse
	ParseResults *[]Knowledgeparserecord `json:"parseResults,omitempty"`

	// ImportResult - Result of the import phase
	ImportResult *Knowledgeparseimportresult `json:"importResult,omitempty"`

	// CreatedBy - The user who created the operation
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - Created date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeparsejobresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeparsejobresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeparsejobresponse
	
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
		
		DownloadURL *string `json:"downloadURL,omitempty"`
		
		Hints *[]string `json:"hints,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ParseResults *[]Knowledgeparserecord `json:"parseResults,omitempty"`
		
		ImportResult *Knowledgeparseimportresult `json:"importResult,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DownloadURL: o.DownloadURL,
		
		Hints: o.Hints,
		
		Status: o.Status,
		
		ParseResults: o.ParseResults,
		
		ImportResult: o.ImportResult,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeparsejobresponse) UnmarshalJSON(b []byte) error {
	var KnowledgeparsejobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeparsejobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeparsejobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DownloadURL, ok := KnowledgeparsejobresponseMap["downloadURL"].(string); ok {
		o.DownloadURL = &DownloadURL
	}
    
	if Hints, ok := KnowledgeparsejobresponseMap["hints"].([]interface{}); ok {
		HintsString, _ := json.Marshal(Hints)
		json.Unmarshal(HintsString, &o.Hints)
	}
	
	if Status, ok := KnowledgeparsejobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ParseResults, ok := KnowledgeparsejobresponseMap["parseResults"].([]interface{}); ok {
		ParseResultsString, _ := json.Marshal(ParseResults)
		json.Unmarshal(ParseResultsString, &o.ParseResults)
	}
	
	if ImportResult, ok := KnowledgeparsejobresponseMap["importResult"].(map[string]interface{}); ok {
		ImportResultString, _ := json.Marshal(ImportResult)
		json.Unmarshal(ImportResultString, &o.ImportResult)
	}
	
	if CreatedBy, ok := KnowledgeparsejobresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := KnowledgeparsejobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeparsejobresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := KnowledgeparsejobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeparsejobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

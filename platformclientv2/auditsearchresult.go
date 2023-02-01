package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditsearchresult
type Auditsearchresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageNumber - Which page was returned.
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - The number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`

	// Total - The total number of results.
	Total *int `json:"total,omitempty"`

	// PageCount - The number of pages of results.
	PageCount *int `json:"pageCount,omitempty"`

	// FacetInfo
	FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`

	// AuditMessages
	AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Auditsearchresult) SetField(field string, fieldValue interface{}) {
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

func (o Auditsearchresult) MarshalJSON() ([]byte, error) {
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
	type Alias Auditsearchresult
	
	return json.Marshal(&struct { 
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`
		
		AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`
		Alias
	}{ 
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		FacetInfo: o.FacetInfo,
		
		AuditMessages: o.AuditMessages,
		Alias:    (Alias)(o),
	})
}

func (o *Auditsearchresult) UnmarshalJSON(b []byte) error {
	var AuditsearchresultMap map[string]interface{}
	err := json.Unmarshal(b, &AuditsearchresultMap)
	if err != nil {
		return err
	}
	
	if PageNumber, ok := AuditsearchresultMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := AuditsearchresultMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Total, ok := AuditsearchresultMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := AuditsearchresultMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if FacetInfo, ok := AuditsearchresultMap["facetInfo"].([]interface{}); ok {
		FacetInfoString, _ := json.Marshal(FacetInfo)
		json.Unmarshal(FacetInfoString, &o.FacetInfo)
	}
	
	if AuditMessages, ok := AuditsearchresultMap["auditMessages"].([]interface{}); ok {
		AuditMessagesString, _ := json.Marshal(AuditMessages)
		json.Unmarshal(AuditMessagesString, &o.AuditMessages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

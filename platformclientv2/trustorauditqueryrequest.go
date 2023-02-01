package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustorauditqueryrequest
type Trustorauditqueryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TrustorOrganizationId - Limit returned audits to this trustor organizationId.
	TrustorOrganizationId *string `json:"trustorOrganizationId,omitempty"`

	// TrusteeUserIds - Limit returned audits to these trustee userIds.
	TrusteeUserIds *[]string `json:"trusteeUserIds,omitempty"`

	// StartDate - Starting date/time for the audit search. ISO-8601 formatted date-time, UTC.
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - Ending date/time for the audit search. ISO-8601 formatted date-time, UTC.
	EndDate *time.Time `json:"endDate,omitempty"`

	// QueryPhrase - Word or phrase to look for in audit bodies.
	QueryPhrase *string `json:"queryPhrase,omitempty"`

	// Facets - Facet information to be returned with the query results.
	Facets *[]Facet `json:"facets,omitempty"`

	// Filters - Additional custom filters to be applied to the query.
	Filters *[]Filter `json:"filters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trustorauditqueryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Trustorauditqueryrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
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
	type Alias Trustorauditqueryrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		TrustorOrganizationId *string `json:"trustorOrganizationId,omitempty"`
		
		TrusteeUserIds *[]string `json:"trusteeUserIds,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		Facets *[]Facet `json:"facets,omitempty"`
		
		Filters *[]Filter `json:"filters,omitempty"`
		Alias
	}{ 
		TrustorOrganizationId: o.TrustorOrganizationId,
		
		TrusteeUserIds: o.TrusteeUserIds,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		QueryPhrase: o.QueryPhrase,
		
		Facets: o.Facets,
		
		Filters: o.Filters,
		Alias:    (Alias)(o),
	})
}

func (o *Trustorauditqueryrequest) UnmarshalJSON(b []byte) error {
	var TrustorauditqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TrustorauditqueryrequestMap)
	if err != nil {
		return err
	}
	
	if TrustorOrganizationId, ok := TrustorauditqueryrequestMap["trustorOrganizationId"].(string); ok {
		o.TrustorOrganizationId = &TrustorOrganizationId
	}
    
	if TrusteeUserIds, ok := TrustorauditqueryrequestMap["trusteeUserIds"].([]interface{}); ok {
		TrusteeUserIdsString, _ := json.Marshal(TrusteeUserIds)
		json.Unmarshal(TrusteeUserIdsString, &o.TrusteeUserIds)
	}
	
	if startDateString, ok := TrustorauditqueryrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := TrustorauditqueryrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if QueryPhrase, ok := TrustorauditqueryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if Facets, ok := TrustorauditqueryrequestMap["facets"].([]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Filters, ok := TrustorauditqueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustorauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

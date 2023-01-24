package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usersearchrequest
type Usersearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`

	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`

	// PageSize - The number of results per page
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - The page of resources you want to retrieve
	PageNumber *int `json:"pageNumber,omitempty"`

	// Sort - Multi-value sort order, list of multiple sort values
	Sort *[]Searchsort `json:"sort,omitempty"`

	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`

	// Query
	Query *[]Usersearchcriteria `json:"query,omitempty"`

	// IntegrationPresenceSource - Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an \"expand\". When using this parameter the maximum number of users that can be returned is 100.
	IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`

	// EnforcePermissions - This property only applies to api/v2/user/search; when set to true add additional search criteria to filter users by: directory:user:view
	EnforcePermissions *bool `json:"enforcePermissions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usersearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Usersearchrequest) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Usersearchrequest
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		Expand *[]string `json:"expand,omitempty"`
		
		Query *[]Usersearchcriteria `json:"query,omitempty"`
		
		IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`
		
		EnforcePermissions *bool `json:"enforcePermissions,omitempty"`
		Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Sort: o.Sort,
		
		Expand: o.Expand,
		
		Query: o.Query,
		
		IntegrationPresenceSource: o.IntegrationPresenceSource,
		
		EnforcePermissions: o.EnforcePermissions,
		Alias:    (Alias)(o),
	})
}

func (o *Usersearchrequest) UnmarshalJSON(b []byte) error {
	var UsersearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UsersearchrequestMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := UsersearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortBy, ok := UsersearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if PageSize, ok := UsersearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := UsersearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Sort, ok := UsersearchrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Expand, ok := UsersearchrequestMap["expand"].([]interface{}); ok {
		ExpandString, _ := json.Marshal(Expand)
		json.Unmarshal(ExpandString, &o.Expand)
	}
	
	if Query, ok := UsersearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	
	if IntegrationPresenceSource, ok := UsersearchrequestMap["integrationPresenceSource"].(string); ok {
		o.IntegrationPresenceSource = &IntegrationPresenceSource
	}
    
	if EnforcePermissions, ok := UsersearchrequestMap["enforcePermissions"].(bool); ok {
		o.EnforcePermissions = &EnforcePermissions
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usersearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

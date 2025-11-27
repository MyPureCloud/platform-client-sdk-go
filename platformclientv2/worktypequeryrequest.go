package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Worktypequeryrequest
type Worktypequeryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageSize - Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an `after` key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200.
	PageSize *int `json:"pageSize,omitempty"`

	// VarSelect - Specify the value 'Count' for this parameter in order to return only the record count.
	VarSelect *string `json:"select,omitempty"`

	// Filters - List of filter objects to be used in the search. Valid filter names are: 'divisionId', 'id', 'name', 'description', 'defaultWorkbinId', 'defaultDurationSeconds', 'defaultExpirationSeconds', 'defaultDueDurationSeconds', 'defaultPriority', 'defaultLanguageId', 'defaultTtlSeconds', 'assignmentEnabled', 'defaultQueueId', 'schemaId', 'schemaVersion', 'dateCreated', 'dateModified', 'modifiedBy'
	Filters *[]Workitemfilter `json:"filters,omitempty"`

	// Attributes - List of entity attributes to be retrieved in the result.
	Attributes *[]string `json:"attributes,omitempty"`

	// After - The cursor that points to the end of the set of entities that has been returned.
	After *string `json:"after,omitempty"`

	// Sort - Sort
	Sort *Worktypequerysort `json:"sort,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Worktypequeryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Worktypequeryrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Worktypequeryrequest
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		VarSelect *string `json:"select,omitempty"`
		
		Filters *[]Workitemfilter `json:"filters,omitempty"`
		
		Attributes *[]string `json:"attributes,omitempty"`
		
		After *string `json:"after,omitempty"`
		
		Sort *Worktypequerysort `json:"sort,omitempty"`
		Alias
	}{ 
		PageSize: o.PageSize,
		
		VarSelect: o.VarSelect,
		
		Filters: o.Filters,
		
		Attributes: o.Attributes,
		
		After: o.After,
		
		Sort: o.Sort,
		Alias:    (Alias)(o),
	})
}

func (o *Worktypequeryrequest) UnmarshalJSON(b []byte) error {
	var WorktypequeryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WorktypequeryrequestMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := WorktypequeryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if VarSelect, ok := WorktypequeryrequestMap["select"].(string); ok {
		o.VarSelect = &VarSelect
	}
    
	if Filters, ok := WorktypequeryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Attributes, ok := WorktypequeryrequestMap["attributes"].([]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if After, ok := WorktypequeryrequestMap["after"].(string); ok {
		o.After = &After
	}
    
	if Sort, ok := WorktypequeryrequestMap["sort"].(map[string]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Worktypequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

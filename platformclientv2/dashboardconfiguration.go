package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dashboardconfiguration
type Dashboardconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of dashboard configuration.
	Name *string `json:"name,omitempty"`

	// Rows - The count of rows for the specific dashboard configuration.
	Rows *int `json:"rows,omitempty"`

	// Columns - The count of columns for the specific dashboard.
	Columns *int `json:"columns,omitempty"`

	// Widgets - List of widgets for dashboard configuration.
	Widgets *[]Widget `json:"widgets,omitempty"`

	// Favorite - The flag indicates if the dashboard is favorited by the user
	Favorite *bool `json:"favorite,omitempty"`

	// PublicDashboard - The flag to indicate if the dashboard is published by an user
	PublicDashboard *bool `json:"publicDashboard,omitempty"`

	// Restricted - The flag to indicate if the dashboard has any restricted data for that user
	Restricted *bool `json:"restricted,omitempty"`

	// LayoutType - The layout type of the dashboard
	LayoutType *string `json:"layoutType,omitempty"`

	// DateCreated - The created date of the dashboard. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The last modified date of the dashboard. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - The id of user who created the dashboard
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`

	// Shared - The flag to indicate if the dashboard is shared
	Shared *bool `json:"shared,omitempty"`

	// DashboardsSharedWith - The list of users and teams the dashboard is shared with
	DashboardsSharedWith *Dashboardssharedwith `json:"dashboardsSharedWith,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dashboardconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Dashboardconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Dashboardconfiguration
	
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
		
		Name *string `json:"name,omitempty"`
		
		Rows *int `json:"rows,omitempty"`
		
		Columns *int `json:"columns,omitempty"`
		
		Widgets *[]Widget `json:"widgets,omitempty"`
		
		Favorite *bool `json:"favorite,omitempty"`
		
		PublicDashboard *bool `json:"publicDashboard,omitempty"`
		
		Restricted *bool `json:"restricted,omitempty"`
		
		LayoutType *string `json:"layoutType,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		
		Shared *bool `json:"shared,omitempty"`
		
		DashboardsSharedWith *Dashboardssharedwith `json:"dashboardsSharedWith,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Rows: o.Rows,
		
		Columns: o.Columns,
		
		Widgets: o.Widgets,
		
		Favorite: o.Favorite,
		
		PublicDashboard: o.PublicDashboard,
		
		Restricted: o.Restricted,
		
		LayoutType: o.LayoutType,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		Shared: o.Shared,
		
		DashboardsSharedWith: o.DashboardsSharedWith,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dashboardconfiguration) UnmarshalJSON(b []byte) error {
	var DashboardconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &DashboardconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DashboardconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DashboardconfigurationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Rows, ok := DashboardconfigurationMap["rows"].(float64); ok {
		RowsInt := int(Rows)
		o.Rows = &RowsInt
	}
	
	if Columns, ok := DashboardconfigurationMap["columns"].(float64); ok {
		ColumnsInt := int(Columns)
		o.Columns = &ColumnsInt
	}
	
	if Widgets, ok := DashboardconfigurationMap["widgets"].([]interface{}); ok {
		WidgetsString, _ := json.Marshal(Widgets)
		json.Unmarshal(WidgetsString, &o.Widgets)
	}
	
	if Favorite, ok := DashboardconfigurationMap["favorite"].(bool); ok {
		o.Favorite = &Favorite
	}
    
	if PublicDashboard, ok := DashboardconfigurationMap["publicDashboard"].(bool); ok {
		o.PublicDashboard = &PublicDashboard
	}
    
	if Restricted, ok := DashboardconfigurationMap["restricted"].(bool); ok {
		o.Restricted = &Restricted
	}
    
	if LayoutType, ok := DashboardconfigurationMap["layoutType"].(string); ok {
		o.LayoutType = &LayoutType
	}
    
	if dateCreatedString, ok := DashboardconfigurationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DashboardconfigurationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := DashboardconfigurationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Shared, ok := DashboardconfigurationMap["shared"].(bool); ok {
		o.Shared = &Shared
	}
    
	if DashboardsSharedWith, ok := DashboardconfigurationMap["dashboardsSharedWith"].(map[string]interface{}); ok {
		DashboardsSharedWithString, _ := json.Marshal(DashboardsSharedWith)
		json.Unmarshal(DashboardsSharedWithString, &o.DashboardsSharedWith)
	}
	
	if SelfUri, ok := DashboardconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dashboardconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

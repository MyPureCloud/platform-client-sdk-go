package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Widget
type Widget struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Row - The row number for the specific dashboard widget configuration.
	Row *int `json:"row,omitempty"`

	// Column - The column number for the specific dashboard widget configuration.
	Column *int `json:"column,omitempty"`

	// Title - The title for the dashboard widget configuration.
	Title *string `json:"title,omitempty"`

	// VarType - The type of dashboard widget configuration.
	VarType *string `json:"type,omitempty"`

	// Metrics - The list of metrics for the dashboard widget configuration.
	Metrics *[]string `json:"metrics,omitempty"`

	// DisplayText - The display text for the dashboard widget configuration.
	DisplayText *string `json:"displayText,omitempty"`

	// DisplayTextColor - The color of the display text for the dashboard widget configuration in RGB hexadecimal format (for example \"#FF0000\" represents red).
	DisplayTextColor *string `json:"displayTextColor,omitempty"`

	// WebContentUrl - The external web URL for the dashboard widget configuration.
	WebContentUrl *string `json:"webContentUrl,omitempty"`

	// SplitFilters - Indicates each filter to be displayed individually.
	SplitFilters *bool `json:"splitFilters,omitempty"`

	// SplitByMediaType - Indicates that data for each media type should be shown individually.
	SplitByMediaType *bool `json:"splitByMediaType,omitempty"`

	// ShowLongest - Indicates the display be the longest time.
	ShowLongest *bool `json:"showLongest,omitempty"`

	// DisplayAsTable - Indicates the widget to be displayed as table.
	DisplayAsTable *bool `json:"displayAsTable,omitempty"`

	// ShowDuration - Indicates the display to include duration.
	ShowDuration *bool `json:"showDuration,omitempty"`

	// SortOrder - The sort order of the table.
	SortOrder *string `json:"sortOrder,omitempty"`

	// SortKey - The sort key of the table.
	SortKey *string `json:"sortKey,omitempty"`

	// EntityLimit - Indicates the limit of displayed entities.
	EntityLimit *int `json:"entityLimit,omitempty"`

	// DisplayAggregates - Indicates whether to display aggregate across all entity and media type combination.
	DisplayAggregates *bool `json:"displayAggregates,omitempty"`

	// IsFullWidth - Indicates whether a widget should take the full width of a dashboard or be shown only in a single slot.
	IsFullWidth *bool `json:"isFullWidth,omitempty"`

	// ShowPercentageChange - Indicates whether a widget should show the percentage diff between two values.
	ShowPercentageChange *bool `json:"showPercentageChange,omitempty"`

	// ShowProfilePicture - Indicates whether a widget should show the profile picture of an agent.
	ShowProfilePicture *bool `json:"showProfilePicture,omitempty"`

	// Filter - The filters to be applied for dashboard widget configuration
	Filter *Viewfilter `json:"filter,omitempty"`

	// Periods - The list of periods for the dashboard widget configuration
	Periods *[]string `json:"periods,omitempty"`

	// MediaTypes - The list of media types for the dashboard widget configuration
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// Warnings - List of warnings for dashboard widget configuration
	Warnings *[]Warning `json:"warnings,omitempty"`

	// ShowTimeInStatus - Indicates the show time in status of a widget configuration.
	ShowTimeInStatus *bool `json:"showTimeInStatus,omitempty"`

	// ShowOfflineAgents - Indicates to show offline agent widget.
	ShowOfflineAgents *bool `json:"showOfflineAgents,omitempty"`

	// SelectedStatuses - Indicates the selected statuses used to filter the agent widget in the dashboard.
	SelectedStatuses *[]string `json:"selectedStatuses,omitempty"`

	// SelectedSegmentTypes - Indicates the selected segment types used to filter the agent activity in the dashboard.
	SelectedSegmentTypes *[]string `json:"selectedSegmentTypes,omitempty"`

	// AgentInteractionSortOrder - The sort order of the interactions in the agent status widget.
	AgentInteractionSortOrder *string `json:"agentInteractionSortOrder,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Widget) SetField(field string, fieldValue interface{}) {
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

func (o Widget) MarshalJSON() ([]byte, error) {
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
	type Alias Widget
	
	return json.Marshal(&struct { 
		Row *int `json:"row,omitempty"`
		
		Column *int `json:"column,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		DisplayText *string `json:"displayText,omitempty"`
		
		DisplayTextColor *string `json:"displayTextColor,omitempty"`
		
		WebContentUrl *string `json:"webContentUrl,omitempty"`
		
		SplitFilters *bool `json:"splitFilters,omitempty"`
		
		SplitByMediaType *bool `json:"splitByMediaType,omitempty"`
		
		ShowLongest *bool `json:"showLongest,omitempty"`
		
		DisplayAsTable *bool `json:"displayAsTable,omitempty"`
		
		ShowDuration *bool `json:"showDuration,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortKey *string `json:"sortKey,omitempty"`
		
		EntityLimit *int `json:"entityLimit,omitempty"`
		
		DisplayAggregates *bool `json:"displayAggregates,omitempty"`
		
		IsFullWidth *bool `json:"isFullWidth,omitempty"`
		
		ShowPercentageChange *bool `json:"showPercentageChange,omitempty"`
		
		ShowProfilePicture *bool `json:"showProfilePicture,omitempty"`
		
		Filter *Viewfilter `json:"filter,omitempty"`
		
		Periods *[]string `json:"periods,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		Warnings *[]Warning `json:"warnings,omitempty"`
		
		ShowTimeInStatus *bool `json:"showTimeInStatus,omitempty"`
		
		ShowOfflineAgents *bool `json:"showOfflineAgents,omitempty"`
		
		SelectedStatuses *[]string `json:"selectedStatuses,omitempty"`
		
		SelectedSegmentTypes *[]string `json:"selectedSegmentTypes,omitempty"`
		
		AgentInteractionSortOrder *string `json:"agentInteractionSortOrder,omitempty"`
		Alias
	}{ 
		Row: o.Row,
		
		Column: o.Column,
		
		Title: o.Title,
		
		VarType: o.VarType,
		
		Metrics: o.Metrics,
		
		DisplayText: o.DisplayText,
		
		DisplayTextColor: o.DisplayTextColor,
		
		WebContentUrl: o.WebContentUrl,
		
		SplitFilters: o.SplitFilters,
		
		SplitByMediaType: o.SplitByMediaType,
		
		ShowLongest: o.ShowLongest,
		
		DisplayAsTable: o.DisplayAsTable,
		
		ShowDuration: o.ShowDuration,
		
		SortOrder: o.SortOrder,
		
		SortKey: o.SortKey,
		
		EntityLimit: o.EntityLimit,
		
		DisplayAggregates: o.DisplayAggregates,
		
		IsFullWidth: o.IsFullWidth,
		
		ShowPercentageChange: o.ShowPercentageChange,
		
		ShowProfilePicture: o.ShowProfilePicture,
		
		Filter: o.Filter,
		
		Periods: o.Periods,
		
		MediaTypes: o.MediaTypes,
		
		Warnings: o.Warnings,
		
		ShowTimeInStatus: o.ShowTimeInStatus,
		
		ShowOfflineAgents: o.ShowOfflineAgents,
		
		SelectedStatuses: o.SelectedStatuses,
		
		SelectedSegmentTypes: o.SelectedSegmentTypes,
		
		AgentInteractionSortOrder: o.AgentInteractionSortOrder,
		Alias:    (Alias)(o),
	})
}

func (o *Widget) UnmarshalJSON(b []byte) error {
	var WidgetMap map[string]interface{}
	err := json.Unmarshal(b, &WidgetMap)
	if err != nil {
		return err
	}
	
	if Row, ok := WidgetMap["row"].(float64); ok {
		RowInt := int(Row)
		o.Row = &RowInt
	}
	
	if Column, ok := WidgetMap["column"].(float64); ok {
		ColumnInt := int(Column)
		o.Column = &ColumnInt
	}
	
	if Title, ok := WidgetMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if VarType, ok := WidgetMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Metrics, ok := WidgetMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if DisplayText, ok := WidgetMap["displayText"].(string); ok {
		o.DisplayText = &DisplayText
	}
    
	if DisplayTextColor, ok := WidgetMap["displayTextColor"].(string); ok {
		o.DisplayTextColor = &DisplayTextColor
	}
    
	if WebContentUrl, ok := WidgetMap["webContentUrl"].(string); ok {
		o.WebContentUrl = &WebContentUrl
	}
    
	if SplitFilters, ok := WidgetMap["splitFilters"].(bool); ok {
		o.SplitFilters = &SplitFilters
	}
    
	if SplitByMediaType, ok := WidgetMap["splitByMediaType"].(bool); ok {
		o.SplitByMediaType = &SplitByMediaType
	}
    
	if ShowLongest, ok := WidgetMap["showLongest"].(bool); ok {
		o.ShowLongest = &ShowLongest
	}
    
	if DisplayAsTable, ok := WidgetMap["displayAsTable"].(bool); ok {
		o.DisplayAsTable = &DisplayAsTable
	}
    
	if ShowDuration, ok := WidgetMap["showDuration"].(bool); ok {
		o.ShowDuration = &ShowDuration
	}
    
	if SortOrder, ok := WidgetMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortKey, ok := WidgetMap["sortKey"].(string); ok {
		o.SortKey = &SortKey
	}
    
	if EntityLimit, ok := WidgetMap["entityLimit"].(float64); ok {
		EntityLimitInt := int(EntityLimit)
		o.EntityLimit = &EntityLimitInt
	}
	
	if DisplayAggregates, ok := WidgetMap["displayAggregates"].(bool); ok {
		o.DisplayAggregates = &DisplayAggregates
	}
    
	if IsFullWidth, ok := WidgetMap["isFullWidth"].(bool); ok {
		o.IsFullWidth = &IsFullWidth
	}
    
	if ShowPercentageChange, ok := WidgetMap["showPercentageChange"].(bool); ok {
		o.ShowPercentageChange = &ShowPercentageChange
	}
    
	if ShowProfilePicture, ok := WidgetMap["showProfilePicture"].(bool); ok {
		o.ShowProfilePicture = &ShowProfilePicture
	}
    
	if Filter, ok := WidgetMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Periods, ok := WidgetMap["periods"].([]interface{}); ok {
		PeriodsString, _ := json.Marshal(Periods)
		json.Unmarshal(PeriodsString, &o.Periods)
	}
	
	if MediaTypes, ok := WidgetMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if Warnings, ok := WidgetMap["warnings"].([]interface{}); ok {
		WarningsString, _ := json.Marshal(Warnings)
		json.Unmarshal(WarningsString, &o.Warnings)
	}
	
	if ShowTimeInStatus, ok := WidgetMap["showTimeInStatus"].(bool); ok {
		o.ShowTimeInStatus = &ShowTimeInStatus
	}
    
	if ShowOfflineAgents, ok := WidgetMap["showOfflineAgents"].(bool); ok {
		o.ShowOfflineAgents = &ShowOfflineAgents
	}
    
	if SelectedStatuses, ok := WidgetMap["selectedStatuses"].([]interface{}); ok {
		SelectedStatusesString, _ := json.Marshal(SelectedStatuses)
		json.Unmarshal(SelectedStatusesString, &o.SelectedStatuses)
	}
	
	if SelectedSegmentTypes, ok := WidgetMap["selectedSegmentTypes"].([]interface{}); ok {
		SelectedSegmentTypesString, _ := json.Marshal(SelectedSegmentTypes)
		json.Unmarshal(SelectedSegmentTypesString, &o.SelectedSegmentTypes)
	}
	
	if AgentInteractionSortOrder, ok := WidgetMap["agentInteractionSortOrder"].(string); ok {
		o.AgentInteractionSortOrder = &AgentInteractionSortOrder
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Widget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

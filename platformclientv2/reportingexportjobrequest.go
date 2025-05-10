package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingexportjobrequest
type Reportingexportjobrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The user supplied name of the export request
	Name *string `json:"name,omitempty"`

	// TimeZone - The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`

	// ExportFormat - The requested format of the exported data
	ExportFormat *string `json:"exportFormat,omitempty"`

	// Interval - The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Period - The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Period *string `json:"period,omitempty"`

	// ViewType - The type of view export job to be created
	ViewType *string `json:"viewType,omitempty"`

	// Filter - Filters to apply to create the view
	Filter *Viewfilter `json:"filter,omitempty"`

	// Read - Indicates if the request has been marked as read
	Read *bool `json:"read,omitempty"`

	// Locale - The locale used for localization of the exported data, i.e. en-US, es
	Locale *string `json:"locale,omitempty"`

	// HasFormatDurations - Indicates if durations are formatted in hh:mm:ss format instead of ms
	HasFormatDurations *bool `json:"hasFormatDurations,omitempty"`

	// HasSplitFilters - Indicates if filters will be split in aggregate detail exports
	HasSplitFilters *bool `json:"hasSplitFilters,omitempty"`

	// ExcludeEmptyRows - Excludes empty rows from the exports
	ExcludeEmptyRows *bool `json:"excludeEmptyRows,omitempty"`

	// HasSplitByMedia - Indicates if media type will be split in aggregate detail exports
	HasSplitByMedia *bool `json:"hasSplitByMedia,omitempty"`

	// HasSummaryRow - Indicates if summary row needs to be present in exports
	HasSummaryRow *bool `json:"hasSummaryRow,omitempty"`

	// CsvDelimiter - The user supplied csv delimiter string value either of type 'comma' or 'semicolon' permitted for the export request
	CsvDelimiter *string `json:"csvDelimiter,omitempty"`

	// SelectedColumns - The list of ordered selected columns from the export view by the user
	SelectedColumns *[]Selectedcolumns `json:"selectedColumns,omitempty"`

	// HasCustomParticipantAttributes - Indicates if custom participant attributes will be exported
	HasCustomParticipantAttributes *bool `json:"hasCustomParticipantAttributes,omitempty"`

	// RecipientEmails - The list of email recipients for the exports
	RecipientEmails *[]string `json:"recipientEmails,omitempty"`

	// IncludeDurationFormatInHeader - Indicates whether to include selected duration format to the column headers
	IncludeDurationFormatInHeader *bool `json:"includeDurationFormatInHeader,omitempty"`

	// DurationFormat - Indicates the duration format for the exports
	DurationFormat *string `json:"durationFormat,omitempty"`

	// ChartColumns - The list of columns for which chart is going to be displayed in export
	ChartColumns *[]Chartcolumn `json:"chartColumns,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingexportjobrequest) SetField(field string, fieldValue interface{}) {
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

func (o Reportingexportjobrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Reportingexportjobrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		ExportFormat *string `json:"exportFormat,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		Period *string `json:"period,omitempty"`
		
		ViewType *string `json:"viewType,omitempty"`
		
		Filter *Viewfilter `json:"filter,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		Locale *string `json:"locale,omitempty"`
		
		HasFormatDurations *bool `json:"hasFormatDurations,omitempty"`
		
		HasSplitFilters *bool `json:"hasSplitFilters,omitempty"`
		
		ExcludeEmptyRows *bool `json:"excludeEmptyRows,omitempty"`
		
		HasSplitByMedia *bool `json:"hasSplitByMedia,omitempty"`
		
		HasSummaryRow *bool `json:"hasSummaryRow,omitempty"`
		
		CsvDelimiter *string `json:"csvDelimiter,omitempty"`
		
		SelectedColumns *[]Selectedcolumns `json:"selectedColumns,omitempty"`
		
		HasCustomParticipantAttributes *bool `json:"hasCustomParticipantAttributes,omitempty"`
		
		RecipientEmails *[]string `json:"recipientEmails,omitempty"`
		
		IncludeDurationFormatInHeader *bool `json:"includeDurationFormatInHeader,omitempty"`
		
		DurationFormat *string `json:"durationFormat,omitempty"`
		
		ChartColumns *[]Chartcolumn `json:"chartColumns,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		TimeZone: o.TimeZone,
		
		ExportFormat: o.ExportFormat,
		
		Interval: o.Interval,
		
		Period: o.Period,
		
		ViewType: o.ViewType,
		
		Filter: o.Filter,
		
		Read: o.Read,
		
		Locale: o.Locale,
		
		HasFormatDurations: o.HasFormatDurations,
		
		HasSplitFilters: o.HasSplitFilters,
		
		ExcludeEmptyRows: o.ExcludeEmptyRows,
		
		HasSplitByMedia: o.HasSplitByMedia,
		
		HasSummaryRow: o.HasSummaryRow,
		
		CsvDelimiter: o.CsvDelimiter,
		
		SelectedColumns: o.SelectedColumns,
		
		HasCustomParticipantAttributes: o.HasCustomParticipantAttributes,
		
		RecipientEmails: o.RecipientEmails,
		
		IncludeDurationFormatInHeader: o.IncludeDurationFormatInHeader,
		
		DurationFormat: o.DurationFormat,
		
		ChartColumns: o.ChartColumns,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingexportjobrequest) UnmarshalJSON(b []byte) error {
	var ReportingexportjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingexportjobrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ReportingexportjobrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TimeZone, ok := ReportingexportjobrequestMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if ExportFormat, ok := ReportingexportjobrequestMap["exportFormat"].(string); ok {
		o.ExportFormat = &ExportFormat
	}
    
	if Interval, ok := ReportingexportjobrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Period, ok := ReportingexportjobrequestMap["period"].(string); ok {
		o.Period = &Period
	}
    
	if ViewType, ok := ReportingexportjobrequestMap["viewType"].(string); ok {
		o.ViewType = &ViewType
	}
    
	if Filter, ok := ReportingexportjobrequestMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Read, ok := ReportingexportjobrequestMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if Locale, ok := ReportingexportjobrequestMap["locale"].(string); ok {
		o.Locale = &Locale
	}
    
	if HasFormatDurations, ok := ReportingexportjobrequestMap["hasFormatDurations"].(bool); ok {
		o.HasFormatDurations = &HasFormatDurations
	}
    
	if HasSplitFilters, ok := ReportingexportjobrequestMap["hasSplitFilters"].(bool); ok {
		o.HasSplitFilters = &HasSplitFilters
	}
    
	if ExcludeEmptyRows, ok := ReportingexportjobrequestMap["excludeEmptyRows"].(bool); ok {
		o.ExcludeEmptyRows = &ExcludeEmptyRows
	}
    
	if HasSplitByMedia, ok := ReportingexportjobrequestMap["hasSplitByMedia"].(bool); ok {
		o.HasSplitByMedia = &HasSplitByMedia
	}
    
	if HasSummaryRow, ok := ReportingexportjobrequestMap["hasSummaryRow"].(bool); ok {
		o.HasSummaryRow = &HasSummaryRow
	}
    
	if CsvDelimiter, ok := ReportingexportjobrequestMap["csvDelimiter"].(string); ok {
		o.CsvDelimiter = &CsvDelimiter
	}
    
	if SelectedColumns, ok := ReportingexportjobrequestMap["selectedColumns"].([]interface{}); ok {
		SelectedColumnsString, _ := json.Marshal(SelectedColumns)
		json.Unmarshal(SelectedColumnsString, &o.SelectedColumns)
	}
	
	if HasCustomParticipantAttributes, ok := ReportingexportjobrequestMap["hasCustomParticipantAttributes"].(bool); ok {
		o.HasCustomParticipantAttributes = &HasCustomParticipantAttributes
	}
    
	if RecipientEmails, ok := ReportingexportjobrequestMap["recipientEmails"].([]interface{}); ok {
		RecipientEmailsString, _ := json.Marshal(RecipientEmails)
		json.Unmarshal(RecipientEmailsString, &o.RecipientEmails)
	}
	
	if IncludeDurationFormatInHeader, ok := ReportingexportjobrequestMap["includeDurationFormatInHeader"].(bool); ok {
		o.IncludeDurationFormatInHeader = &IncludeDurationFormatInHeader
	}
    
	if DurationFormat, ok := ReportingexportjobrequestMap["durationFormat"].(string); ok {
		o.DurationFormat = &DurationFormat
	}
    
	if ChartColumns, ok := ReportingexportjobrequestMap["chartColumns"].([]interface{}); ok {
		ChartColumnsString, _ := json.Marshal(ChartColumns)
		json.Unmarshal(ChartColumnsString, &o.ChartColumns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

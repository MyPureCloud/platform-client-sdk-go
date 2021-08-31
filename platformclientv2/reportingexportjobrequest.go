package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingexportjobrequest
type Reportingexportjobrequest struct { 
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


	// Locale - The locale use for localization of the exported data, i.e. en-us, es-mx  
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

}

func (o *Reportingexportjobrequest) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

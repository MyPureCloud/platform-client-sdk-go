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

func (u *Reportingexportjobrequest) MarshalJSON() ([]byte, error) {
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
		Name: u.Name,
		
		TimeZone: u.TimeZone,
		
		ExportFormat: u.ExportFormat,
		
		Interval: u.Interval,
		
		Period: u.Period,
		
		ViewType: u.ViewType,
		
		Filter: u.Filter,
		
		Read: u.Read,
		
		Locale: u.Locale,
		
		HasFormatDurations: u.HasFormatDurations,
		
		HasSplitFilters: u.HasSplitFilters,
		
		ExcludeEmptyRows: u.ExcludeEmptyRows,
		
		HasSplitByMedia: u.HasSplitByMedia,
		
		HasSummaryRow: u.HasSummaryRow,
		
		CsvDelimiter: u.CsvDelimiter,
		
		SelectedColumns: u.SelectedColumns,
		
		HasCustomParticipantAttributes: u.HasCustomParticipantAttributes,
		
		RecipientEmails: u.RecipientEmails,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

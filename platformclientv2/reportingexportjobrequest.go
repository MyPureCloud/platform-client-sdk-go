package platformclientv2
import (
	"encoding/json"
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


	// SelectedColumns - The list of ordered selected columns from the export view by the user
	SelectedColumns *[]Selectedcolumns `json:"selectedColumns,omitempty"`


	// HasCustomParticipantAttributes - Indicates if custom participant attributes will be exported
	HasCustomParticipantAttributes *bool `json:"hasCustomParticipantAttributes,omitempty"`


	// RecipientEmails - The list of email recipients for the exports
	RecipientEmails *[]string `json:"recipientEmails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportingexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

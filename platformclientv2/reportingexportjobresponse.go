package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reportingexportjobresponse
type Reportingexportjobresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// RunId - The unique run id of the export schedule execute
	RunId *string `json:"runId,omitempty"`


	// Status - The current status of the export request
	Status *string `json:"status,omitempty"`


	// TimeZone - The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`


	// ExportFormat - The requested format of the exported data
	ExportFormat *string `json:"exportFormat,omitempty"`


	// Interval - The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// DownloadUrl - The url to download the request if it's status is completed
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// ViewType - The type of view export job to be created
	ViewType *string `json:"viewType,omitempty"`


	// ExportErrorMessagesType - The error message in case the export request failed
	ExportErrorMessagesType *string `json:"exportErrorMessagesType,omitempty"`


	// Period - The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Period *string `json:"period,omitempty"`


	// Filter - Filters to apply to create the view
	Filter *Viewfilter `json:"filter,omitempty"`


	// Read - Indicates if the request has been marked as read
	Read *bool `json:"read,omitempty"`


	// CreatedDateTime - The created date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`


	// ModifiedDateTime - The last modified date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`


	// Locale - The locale use for localization of the exported data, i.e. en-us, es-mx  
	Locale *string `json:"locale,omitempty"`


	// PercentageComplete - The percentage of the job that has completed processing
	PercentageComplete *float64 `json:"percentageComplete,omitempty"`


	// HasFormatDurations - Indicates if durations are formatted in hh:mm:ss format instead of ms
	HasFormatDurations *bool `json:"hasFormatDurations,omitempty"`


	// HasSplitFilters - Indicates if filters will be split in aggregate detail exports
	HasSplitFilters *bool `json:"hasSplitFilters,omitempty"`


	// ExcludeEmptyRows - Excludes empty rows from the exports
	ExcludeEmptyRows *bool `json:"excludeEmptyRows,omitempty"`


	// HasSplitByMedia - Indicates if media type will be split in aggregate detail exports
	HasSplitByMedia *bool `json:"hasSplitByMedia,omitempty"`


	// SelectedColumns - The list of ordered selected columns from the export view by the user
	SelectedColumns *[]Selectedcolumns `json:"selectedColumns,omitempty"`


	// HasCustomParticipantAttributes - Indicates if custom participant attributes will be exported
	HasCustomParticipantAttributes *bool `json:"hasCustomParticipantAttributes,omitempty"`


	// RecipientEmails - The list of email recipients for the exports
	RecipientEmails *[]string `json:"recipientEmails,omitempty"`


	// EmailStatuses - The status of individual email addresses as a map
	EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`


	// EmailErrorDescription - The optional error message in case the export fail to email
	EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportingexportjobresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

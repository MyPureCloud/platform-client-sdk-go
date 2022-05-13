package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// CreatedDateTime - The created date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`


	// ModifiedDateTime - The last modified date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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


	// EmailStatuses - The status of individual email addresses as a map
	EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`


	// EmailErrorDescription - The optional error message in case the export fail to email
	EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Reportingexportjobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingexportjobresponse
	
	CreatedDateTime := new(string)
	if o.CreatedDateTime != nil {
		
		*CreatedDateTime = timeutil.Strftime(o.CreatedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDateTime = nil
	}
	
	ModifiedDateTime := new(string)
	if o.ModifiedDateTime != nil {
		
		*ModifiedDateTime = timeutil.Strftime(o.ModifiedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDateTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		ExportFormat *string `json:"exportFormat,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		ViewType *string `json:"viewType,omitempty"`
		
		ExportErrorMessagesType *string `json:"exportErrorMessagesType,omitempty"`
		
		Period *string `json:"period,omitempty"`
		
		Filter *Viewfilter `json:"filter,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		CreatedDateTime *string `json:"createdDateTime,omitempty"`
		
		ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`
		
		Locale *string `json:"locale,omitempty"`
		
		PercentageComplete *float64 `json:"percentageComplete,omitempty"`
		
		HasFormatDurations *bool `json:"hasFormatDurations,omitempty"`
		
		HasSplitFilters *bool `json:"hasSplitFilters,omitempty"`
		
		ExcludeEmptyRows *bool `json:"excludeEmptyRows,omitempty"`
		
		HasSplitByMedia *bool `json:"hasSplitByMedia,omitempty"`
		
		HasSummaryRow *bool `json:"hasSummaryRow,omitempty"`
		
		CsvDelimiter *string `json:"csvDelimiter,omitempty"`
		
		SelectedColumns *[]Selectedcolumns `json:"selectedColumns,omitempty"`
		
		HasCustomParticipantAttributes *bool `json:"hasCustomParticipantAttributes,omitempty"`
		
		RecipientEmails *[]string `json:"recipientEmails,omitempty"`
		
		EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`
		
		EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		RunId: o.RunId,
		
		Status: o.Status,
		
		TimeZone: o.TimeZone,
		
		ExportFormat: o.ExportFormat,
		
		Interval: o.Interval,
		
		DownloadUrl: o.DownloadUrl,
		
		ViewType: o.ViewType,
		
		ExportErrorMessagesType: o.ExportErrorMessagesType,
		
		Period: o.Period,
		
		Filter: o.Filter,
		
		Read: o.Read,
		
		CreatedDateTime: CreatedDateTime,
		
		ModifiedDateTime: ModifiedDateTime,
		
		Locale: o.Locale,
		
		PercentageComplete: o.PercentageComplete,
		
		HasFormatDurations: o.HasFormatDurations,
		
		HasSplitFilters: o.HasSplitFilters,
		
		ExcludeEmptyRows: o.ExcludeEmptyRows,
		
		HasSplitByMedia: o.HasSplitByMedia,
		
		HasSummaryRow: o.HasSummaryRow,
		
		CsvDelimiter: o.CsvDelimiter,
		
		SelectedColumns: o.SelectedColumns,
		
		HasCustomParticipantAttributes: o.HasCustomParticipantAttributes,
		
		RecipientEmails: o.RecipientEmails,
		
		EmailStatuses: o.EmailStatuses,
		
		EmailErrorDescription: o.EmailErrorDescription,
		
		Enabled: o.Enabled,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingexportjobresponse) UnmarshalJSON(b []byte) error {
	var ReportingexportjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingexportjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportingexportjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportingexportjobresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RunId, ok := ReportingexportjobresponseMap["runId"].(string); ok {
		o.RunId = &RunId
	}
    
	if Status, ok := ReportingexportjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if TimeZone, ok := ReportingexportjobresponseMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if ExportFormat, ok := ReportingexportjobresponseMap["exportFormat"].(string); ok {
		o.ExportFormat = &ExportFormat
	}
    
	if Interval, ok := ReportingexportjobresponseMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if DownloadUrl, ok := ReportingexportjobresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if ViewType, ok := ReportingexportjobresponseMap["viewType"].(string); ok {
		o.ViewType = &ViewType
	}
    
	if ExportErrorMessagesType, ok := ReportingexportjobresponseMap["exportErrorMessagesType"].(string); ok {
		o.ExportErrorMessagesType = &ExportErrorMessagesType
	}
    
	if Period, ok := ReportingexportjobresponseMap["period"].(string); ok {
		o.Period = &Period
	}
    
	if Filter, ok := ReportingexportjobresponseMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Read, ok := ReportingexportjobresponseMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if createdDateTimeString, ok := ReportingexportjobresponseMap["createdDateTime"].(string); ok {
		CreatedDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateTimeString)
		o.CreatedDateTime = &CreatedDateTime
	}
	
	if modifiedDateTimeString, ok := ReportingexportjobresponseMap["modifiedDateTime"].(string); ok {
		ModifiedDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateTimeString)
		o.ModifiedDateTime = &ModifiedDateTime
	}
	
	if Locale, ok := ReportingexportjobresponseMap["locale"].(string); ok {
		o.Locale = &Locale
	}
    
	if PercentageComplete, ok := ReportingexportjobresponseMap["percentageComplete"].(float64); ok {
		o.PercentageComplete = &PercentageComplete
	}
    
	if HasFormatDurations, ok := ReportingexportjobresponseMap["hasFormatDurations"].(bool); ok {
		o.HasFormatDurations = &HasFormatDurations
	}
    
	if HasSplitFilters, ok := ReportingexportjobresponseMap["hasSplitFilters"].(bool); ok {
		o.HasSplitFilters = &HasSplitFilters
	}
    
	if ExcludeEmptyRows, ok := ReportingexportjobresponseMap["excludeEmptyRows"].(bool); ok {
		o.ExcludeEmptyRows = &ExcludeEmptyRows
	}
    
	if HasSplitByMedia, ok := ReportingexportjobresponseMap["hasSplitByMedia"].(bool); ok {
		o.HasSplitByMedia = &HasSplitByMedia
	}
    
	if HasSummaryRow, ok := ReportingexportjobresponseMap["hasSummaryRow"].(bool); ok {
		o.HasSummaryRow = &HasSummaryRow
	}
    
	if CsvDelimiter, ok := ReportingexportjobresponseMap["csvDelimiter"].(string); ok {
		o.CsvDelimiter = &CsvDelimiter
	}
    
	if SelectedColumns, ok := ReportingexportjobresponseMap["selectedColumns"].([]interface{}); ok {
		SelectedColumnsString, _ := json.Marshal(SelectedColumns)
		json.Unmarshal(SelectedColumnsString, &o.SelectedColumns)
	}
	
	if HasCustomParticipantAttributes, ok := ReportingexportjobresponseMap["hasCustomParticipantAttributes"].(bool); ok {
		o.HasCustomParticipantAttributes = &HasCustomParticipantAttributes
	}
    
	if RecipientEmails, ok := ReportingexportjobresponseMap["recipientEmails"].([]interface{}); ok {
		RecipientEmailsString, _ := json.Marshal(RecipientEmails)
		json.Unmarshal(RecipientEmailsString, &o.RecipientEmails)
	}
	
	if EmailStatuses, ok := ReportingexportjobresponseMap["emailStatuses"].(map[string]interface{}); ok {
		EmailStatusesString, _ := json.Marshal(EmailStatuses)
		json.Unmarshal(EmailStatusesString, &o.EmailStatuses)
	}
	
	if EmailErrorDescription, ok := ReportingexportjobresponseMap["emailErrorDescription"].(string); ok {
		o.EmailErrorDescription = &EmailErrorDescription
	}
    
	if Enabled, ok := ReportingexportjobresponseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SelfUri, ok := ReportingexportjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingexportjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}

package platformclientv2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// APIResponse is a friendly interface for a response from the API
type APIResponse struct {
	Response      *http.Response      `json:"-"`
	HasBody       bool                `json:"hasBody,omitempty"`
	RawBody       []byte              `json:"rawBody,omitempty"`
	IsSuccess     bool                `json:"isSuccess,omitempty"`
	StatusCode    int                 `json:"statusCode,omitempty"`
	Status        string              `json:"status,omitempty"`
	Error         *APIError           `json:"error,omitempty"`
	ErrorMessage  string              `json:"errorMessage,omitempty"`
	CorrelationID string              `json:"correlationId,omitempty"`
	Header        map[string][]string `json:"header,omitempty"`
}

// String returns the JSON serialized object
func (r *APIResponse) String() string {
	s, _ := json.Marshal(r)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(s)), `\\u`, `\u`, -1))

	return str
}

// SetError returns the JSON serialized object
func (r *APIResponse) SetError(err *APIError) {
	if err == nil {
		return
	}
	r.Error = err
	r.ErrorMessage = fmt.Sprintf("API Error: %v - %v (%v)", r.StatusCode, err.Message, r.CorrelationID)
}

// APIError is the standard error body from the API
type APIError struct {
	Status            int                    `json:"status,omitempty"`
	Message           string                 `json:"message,omitempty"`
	MessageWithParams string                 `json:"messageWithParams,omitempty"`
	MessageParams     map[string]interface{} `json:"messageParams,omitempty"`
	Code              string                 `json:"code,omitempty"`
	ContextID         string                 `json:"contextId,omitempty"`
	Details           []Detail               `json:"details,omitempty"`
}

// String returns the JSON serialized object
func (r *APIError) String() string {
	s, _ := json.Marshal(r)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(s)), `\\u`, `\u`, -1))

	return str
}

// NewAPIResponse creates an APIResponse from a http Response
func NewAPIResponse(r *http.Response, body []byte) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	var apiError *APIError

	response.StatusCode = r.StatusCode
	response.Status = r.Status
	response.HasBody = len(body) > 0
	response.IsSuccess = r.StatusCode >= 200 && r.StatusCode < 300
	response.Header = r.Header
	if header, exists := response.Header["Inin-Correlation-Id"]; exists {
		response.CorrelationID = header[0]
	}

	if response.HasBody {
		// Set body
		response.RawBody = body

		// Handle error body
		if !response.IsSuccess {
			// Note: an error will always be returned from this block.
			// response.Error will only be set if the error is parsed from the HTTP response body

			// Parse response body as error
			err := json.Unmarshal(response.RawBody, &apiError)
			if err != nil {
				// Return response and json parse error
				return response, err
			}

			// Ensure context id is set
			if apiError.ContextID == "" {
				apiError.ContextID = response.CorrelationID
			}

			// Set response error on the response
			response.SetError(apiError)

			// Return response and api error message
			return response, errors.New(response.ErrorMessage)
		}
	}

	// Return response and unknown API error.
	// This shouldn't happen, but can if there isn't a response body for an unsuccessful request.
	if !response.IsSuccess {
		err := APIError{
			Message:   fmt.Sprintf("API Error: %v", r.Status),
			ContextID: response.CorrelationID,
		}
		response.SetError(&err)
		return response, errors.New(response.ErrorMessage)
	}

	// Return response
	return response, nil
}

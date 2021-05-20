package platformclientv2

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	traceLogger *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
)

type LoggingLevel int

const (
	LTrace LoggingLevel = iota
	LDebug
	LError
	LNone
)

type LoggingFormat int

const (
	JSON LoggingFormat = iota
	Text
)

func loggingFormatFromString(value string) *LoggingFormat {
	var logFormat LoggingFormat

	switch value {
	case "text":
		logFormat = Text
	case "json":
		logFormat = JSON
	default:
		return nil
	}

	return &logFormat
}

func loggingLevelFromString(value string) *LoggingLevel {
	var logLevel LoggingLevel

	switch value {
	case "trace":
		logLevel = LTrace
	case "debug":
		logLevel = LDebug
	case "error":
		logLevel = LError
	case "none":
		logLevel = LNone
	default:
		return nil
	}

	return &logLevel
}

type logStatement struct {
	Date            *time.Time  `json:"date,omitempty"`
	Level           string      `json:"level,omitempty"`
	Method          string      `json:"method,omitempty"`
	URL             string      `json:"url,omitempty"`
	RequestHeaders  http.Header `json:"requestHeaders,omitempty"`
	ResponseHeaders http.Header `json:"responseHeaders,omitempty"`
	CorrelationId   string      `json:"correlationId,omitempty"`
	StatusCode      int         `json:"statusCode,omitempty"`
	RequestBody     string      `json:"requestBody,omitempty"`
	ResponseBody    string      `json:"responseBody,omitempty"`
}

func (s *logStatement) string(format LoggingFormat, logRequestBody, logResponseBody bool) string {
	if len(s.RequestHeaders["Authorization"]) > 0 {
		s.RequestHeaders["Authorization"] = []string{"[REDACTED]"}
	}

	if !logRequestBody {
		s.RequestBody = ""
	}
	if !logResponseBody {
		s.ResponseBody = ""
	}

	if format == Text {
		return fmt.Sprintf(`
=== REQUEST ===%v%v%v%v
=== RESPONSE ===%v%v%v%v`,
			formatValue("URL", s.URL),
			formatValue("Method", s.Method),
			formatValue("Headers", formatHeaders(s.RequestHeaders)),
			formatValue("Body", s.RequestBody),

			formatValue("Status", fmt.Sprintf("%v", s.StatusCode)),
			formatValue("Headers", formatHeaders(s.ResponseHeaders)),
			formatValue("CorrelationId", s.CorrelationId),
			formatValue("Body", s.ResponseBody))
	}

	j, _ := json.Marshal(s)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return strings.TrimRight(string(pretty.Ugly([]byte(str))), "\n")
}

func (c *LoggingConfiguration) configureLogging() {
	var f *os.File
	if c.logFilePath != "" {
		f, _ = os.OpenFile(c.logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}

	var stdoutWrt io.Writer
	var stderrWrt io.Writer

	if f != nil && c.logToConsole { // Logging to console and file
		stdoutWrt = io.MultiWriter(f, os.Stdout)
		stderrWrt = io.MultiWriter(f, os.Stderr)
	} else if f == nil && c.logToConsole { // Logging to console
		stdoutWrt = io.MultiWriter(os.Stdout)
		stderrWrt = io.MultiWriter(os.Stderr)
	} else if f != nil && !c.logToConsole { // Logging to file
		stdoutWrt = io.MultiWriter(f)
		stderrWrt = io.MultiWriter(f)
	} else { // Cannot log to anything
		traceLogger = nil
		debugLogger = nil
		errorLogger = nil
		return
	}

	flags := 0
	tracePrefix := ""
	debugPrefix := ""
	errorPrefix := ""
	if c.logFormat == Text {
		flags = log.Ldate | log.Ltime
		tracePrefix = "TRACE: "
		debugPrefix = "DEBUG: "
		errorPrefix = "ERROR: "
	}

	traceLogger = log.New(stdoutWrt, tracePrefix, flags)
	debugLogger = log.New(stdoutWrt, debugPrefix, flags)
	errorLogger = log.New(stderrWrt, errorPrefix, flags)
}

func (c *LoggingConfiguration) trace(method, URL string, requestBody []byte, statusCode int, requestHeaders, responseHeaders http.Header) {
	now := time.Now()
	logStatement := &logStatement{
		Date:            &now,
		Level:           "trace",
		Method:          method,
		URL:             URL,
		RequestBody:     string(requestBody),
		CorrelationId:   getCorrelationId(responseHeaders),
		StatusCode:      statusCode,
		RequestHeaders:  requestHeaders,
		ResponseHeaders: responseHeaders,
	}

	c.log(traceLogger, LTrace, logStatement.string(c.logFormat, c.LogRequestBody, c.LogResponseBody))
}

func (c *LoggingConfiguration) debug(method, URL string, requestBody []byte, statusCode int, requestHeaders http.Header) {
	now := time.Now()
	logStatement := &logStatement{
		Date:           &now,
		Level:          "debug",
		Method:         method,
		URL:            URL,
		RequestBody:    string(requestBody),
		StatusCode:     statusCode,
		RequestHeaders: requestHeaders,
	}

	c.log(debugLogger, LDebug, logStatement.string(c.logFormat, c.LogRequestBody, c.LogResponseBody))
}

func (c *LoggingConfiguration) error(method, URL string, requestBody, responseBody []byte, statusCode int, requestHeaders, responseHeaders http.Header) {
	now := time.Now()
	logStatement := &logStatement{
		&now,
		"error",
		method,
		URL,
		requestHeaders,
		responseHeaders,
		getCorrelationId(responseHeaders),
		statusCode,
		string(requestBody),
		string(responseBody),
	}

	c.log(errorLogger, LError, logStatement.string(c.logFormat, c.LogRequestBody, c.LogResponseBody))
}

func (c *LoggingConfiguration) log(logger *log.Logger, logLevel LoggingLevel, v ...interface{}) {
	if logLevel >= c.LogLevel && logger != nil {
		logger.Println(v...)
	}
}

func getCorrelationId(headers http.Header) string {
	for key, values := range headers {
		if strings.ToLower(key) == "inin-correlation-id" {
			for _, value := range values {
				if value != "" {
					return value
				}
			}
		}
	}

	return ""
}

// Returns each header key value pair indented by one tab
func formatHeaders(headers http.Header) string {
	var result string
	for key, values := range headers {
		var valuesString string
		for _, value := range values {
			valuesString = fmt.Sprintf("%v, %v", valuesString, value)
		}
		valuesString = strings.TrimLeft(valuesString, ", ")
		result = fmt.Sprintf("%v\n\t%v: %v", result, key, valuesString)
	}
	return result
}

// Used to only print values that aren't empty
func formatValue(name, value string) string {
	if value != "" {
		return fmt.Sprintf("\n%v: %v", name, value)
	}
	return ""
}

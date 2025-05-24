package platformclientv2

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

func TestExampleUsersApi_GetUsers_OnProxy(t *testing.T) {

	config := GetDefaultConfiguration()
	config.LoggingConfiguration.LogLevel = LNone

	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")
	config.GateWayConfiguration = nil

	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	proxyconf := ProxyConfiguration{}
	config.ProxyConfiguration = &proxyconf
	config.ProxyConfiguration.Protocol = "http"
	config.ProxyConfiguration.Host = "localhost"
	config.ProxyConfiguration.Port = "4001"

	// Create an API instance using the default config
	usersAPI := NewUsersApiWithConfig(config)

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200

}

// With Proxy Agent
func TestExampleUsersApi_GetUsers_WithDefaultProxyAgent(t *testing.T) {

	config := GetDefaultConfiguration()
	config.LoggingConfiguration.LogLevel = LNone

	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")
	config.GateWayConfiguration = nil

	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	proxyAgent := &ProxyAgent{
		ProxyURL: "http://localhost:4001",
	}

	config.APIClient.SetProxyAgent(proxyAgent)

	// Create an API instance using the default config
	usersAPI := NewUsersApiWithConfig(config)

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200

}

// Testing for MTLS
func TestExampleUsersApi_GetUsers_CustomClient(t *testing.T) {

	config := GetDefaultConfiguration()
	config.LoggingConfiguration.LogLevel = LNone

	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com

	config.GateWayConfiguration = nil

	gatewayConfig := GateWayConfiguration{}
	gatewayConfig.Host = "localhost"
	gatewayConfig.Port = "4027"
	gatewayConfig.Protocol = "https"
	config.GateWayConfiguration = &gatewayConfig

	config.APIClient.SetMTLSCertificates("mtls-test/localhost.cert.pem", "mtls-test/localhost.key.pem", "mtls-test/ca-chain.cert.pem")

	config.APIClient.client.SetPreHook(PreHook)

	// Create an API instance using the default config
	usersAPI := NewUsersApiWithConfig(config)

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200
}

// PreHook implements the certificate revocation check
func PreHook(logger retryablehttp.Logger, req *http.Request, retry int) {
	config := GetDefaultConfiguration()

	//Extract certificate from request
	certificate, err := getCertificateFromConfig(config)
	if err != nil {
		logger.Printf("[ERROR] Certificate extraction failed: %v", err)
		return
	}
	issuerCertificate, err := getIssuerCertificate()
	if err != nil {
		logger.Printf("[ERROR] Issuer Certificate extraction failed: %v", err)
		return
	}

	// Check certificate status
	isValid, err := validateCertificate(issuerCertificate, certificate)
	if err != nil {
		logger.Printf("[ERROR] Certificate validation failed: %v", err)
		return
	}

	if !isValid {
		logger.Printf("[ERROR] Certificate is revoked")
	}

	logger.Printf("[INFO] Certificate validation successful")
}

// getCertificateFromConfig extracts the certificate from the configuration
func getCertificateFromConfig(config *Configuration) (*x509.Certificate, error) {
	if config.MTLSConfiguration == nil {
		return nil, fmt.Errorf("MTLS Configuration is required for certificate extraction.")
	}

	if config.MTLSConfiguration.CertFile == nil {
		return nil, fmt.Errorf("No Certificate found")
	}

	pemData := config.MTLSConfiguration.CertFile

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	return cert, nil
}

func getIssuerCertificate() (*x509.Certificate, error) {
	// Load issuer certificate from file
	issuerCertPath := "mtls-test/ca-chain.cert.pem"
	issuerCert, err := loadCACertificate(issuerCertPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load CA Certificate")
	}

	return issuerCert, nil
}

// loadCACertificate loads the CA certificate from file
func loadCACertificate(caCertPath string) (*x509.Certificate, error) {
	pemData, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA certificate: %v", err)
	}

	return cert, nil
}

// validateCertificate checks if the certificate is valid and not revoked
func validateCertificate(caCert *x509.Certificate, cert *x509.Certificate) (bool, error) {
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	// Verify certificate
	opts := x509.VerifyOptions{
		Roots:         certPool,
		CurrentTime:   cert.NotBefore,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		Intermediates: x509.NewCertPool(),
	}

	_, err := cert.Verify(opts)
	if err != nil {
		return false, fmt.Errorf("certificate verification failed: %v", err)
	}

	// Perform CRL validation
	if len(cert.CRLDistributionPoints) > 0 {
		isRevoked, err := checkCRL(cert)
		if err != nil {
			return false, err
		}
		if isRevoked {
			return false, nil
		}
	}

	return true, nil
}

// checkCRL performs a basic CRL check
func checkCRL(cert *x509.Certificate) (bool, error) {
	for _, crlURL := range cert.CRLDistributionPoints {
		resp, err := http.Get(crlURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		crlBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		crl, err := x509.ParseCRL(crlBytes)
		if err != nil {
			continue
		}

		for _, revokedCert := range crl.TBSCertList.RevokedCertificates {
			if cert.SerialNumber.Cmp(revokedCert.SerialNumber) == 0 {
				return true, nil // Certificate is revoked
			}
		}
	}

	return false, nil // Certificate is not revoked
}

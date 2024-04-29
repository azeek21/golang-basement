package security

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
)

func getCert(cert, key string) (tls.Certificate, error) {
	if len(cert) < 1 || len(key) < 1 {
		return tls.Certificate{}, errors.New("certificate or key file name invalid")
	}
	loadedCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return loadedCert, err
	}
	return loadedCert, nil
}

func LoadServerCert(certFile, keyFile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	cert, err := getCert(certFile, keyFile)
	return func(cri *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}

}
func LoadClientCert(certFile, keyFile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	cert, err := getCert(certFile, keyFile)
	return func(cri *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}

}
func VerifyCertChains(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	return nil
}

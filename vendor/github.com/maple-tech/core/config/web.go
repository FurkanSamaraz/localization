package config

//OptionsWeb holds info for the web serving components
type OptionsWeb struct {
	//Host is the hostname the http server will listen on
	Host string `json:"host"`

	//Port is the port the http server will listen on, to be combined with host in [host]:[port] fashion
	Port int `json:"port"`

	//HTTPS boolean if true, then use TLS
	HTTPS bool `json:"https"`

	//IDP enabled
	IDP bool `json:"idp"`

	//TLSCert The TLS Certificate file
	TLSCert string `json:"sslCert"`

	//TLSKey The TLS Certifacte key file
	TLSKey string `json:"sslKey"`

	//TLSChain The RootCA certificates to load
	TLSChain string `json:"sslChain"`

	//ReadTimeout seconds max timeout on reading from client before bailing, timed over WHOLE request (headers too)
	ReadTimeout int `json:"read-timeout"`

	//WriteTimeout seconds max timeout on writing to client before bailing, timed over WHOLE response (headers too)
	WriteTimeout int `json:"write-timeout"`

	//IdleTimeout seconds max time client can remain alive idle before we cut them loose
	IdleTimeout int `json:"idle-timeout"`

	//ShutDownGrace seconds time for the graceful shutdown period. During this new connections are blocked, and existing are allowed to complete
	ShutdownGrace int `json:"shutdown-grace"`

	// SAMLIDPMetadatURL saml identity providers list with key provider name and provider metadata url
	SAMLIDPsMetadata map[string]string `json:"saml-idps"`

	// SAMLSPCertificate file path
	SAMLSPCertificate string `json:"saml-sp-certificate"`

	// SAMLSPKey file path
	SAMLSPKey string `json:"saml-sp-key"`

	// SAMLSPKey file path
	SAMLSPRootURL string `json:"saml-sp-root-url"`

	// the of where the service is hosted
	DeploymentURL string `json:"deployment-url"`
}

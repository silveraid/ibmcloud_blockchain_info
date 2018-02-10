package credreader

type Client struct {
	Organization string    `json:"organization"`
}

type Credentials struct {
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	Client         Client                     `json:"client"`
	Channels       map[string]Channel         `json:"channels"`
	XNetworkId     string                     `json:"x-networkId"`
	XType          string                     `json:"x-type"`
	Organizations  map[string]Organization    `json:"organizations"`
	Orderers       map[string]Orderer         `json:"orderers"`
	Peers          map[string]Peer            `json:"peers"`
	CAs            map[string]CA              `json:"certificateAuthorities"`
}

type Organization struct {
	MSPID                   string      `json:"mspid"`
	Peers                   []string    `json:"peers"`
	CertificateAuthorities  []string    `json:"certificateAuthorities,omitempty"`
	SignedCert              Cert        `json:"signedCert"`
	XUploadedSignedCerts    []Cert      `json:"x-uploadedSignedCerts"`
}

type Cert struct {
	Pem    string    `json:"pem"`
	XName  string    `json:"x-name"`
}

type TLSCACert struct {
	Pem string    `json:"pem"`
}

type Channel struct {
	Orderers  []string                     `json:"orderers"`
	Peers     map[string]PeerAttributes    `json:"peers"`
}

type PeerAttributes struct {
	EndorsingPeer   bool    `json:"endorsingPeer"`
	ChaincodeQuery  bool    `json:"chaincodeQuery"`
	LedgerQuery     bool    `json:"ledgerQuery"`
	EventSource     bool    `json:"eventSource"`
}

type Peer struct {
	URL            string         `json:"url"`
	EventURL       string         `json:"eventUrl"`
	GRPCOptions    GRPCOptions    `json:"grpcOptions"`
	TLSCACerts     TLSCACert      `json:"tlsCACerts"`
	XMSPID         string         `json:"x-mspid"`
	XLedgerDBType  string         `json:"x-ledgerDbType"`
}

type Orderer struct {
	URL          string         `json:"url"`
	GRPCOptions  GRPCOptions    `json:"grpcOptions"`
	TLSCACerts   TLSCACert      `json:"tlsCACerts"`
}

type GRPCOptions struct {
	SSLTargetNameOverride      *string    `json:"ssl-target-name-override"`
	GRPCHttp2KeepaliveTime     int64      `json:"grpc.http2.keepalive_time"`
	GRPCKeepaliveTimeMs        int64      `json:"grpc.keepalive_time_ms"`
	GRPCHttp2KeepaliveTimeout  int64      `json:"grpc.http2.keepalive_timeout"`
	GRPCKeepaliveTimeoutMs     int64      `json:"grpc.keepalive_timeout_ms"`
}

type CA struct {
	URL          string           `json:"url"`
	CAName       string           `json:"caName"`
	XMSPID       string           `json:"x-mspid"`
	HttpOptions  HttpOptions      `json:"httpOptions"`
	TLSCACerts   TLSCACert        `json:"tlsCACerts"`
	Registrar    []CARegistrar    `json:"registrar"`
}

type HttpOptions struct {
	Verify bool    `json:"verify"`
}

type CARegistrar struct {
	EnrollId      string    `json:"enrollId"`
	EnrollSecret  string    `json:"enrollSecret"`
}

package credreader

import (
	"fmt"
)

// Function CredReader unmarshalls the untyped credentials data from
// Bluemix (IBM Cloud) API and returns it in a struct
func CredReader(input map[string]interface{}) Credentials {

	var c Credentials

	//
	//  strings
	//

	c.Name = input["name"].(string)
	c.Description = input["description"].(string)
	c.Client = Client{input["client"].(map[string]interface{})["organization"].(string)}
	c.XNetworkId = input["x-networkId"].(string)
	c.XType = input["x-type"].(string)

	//
	//  channels map
	//

	for chName, chAttrs := range input["channels"].(map[string]interface{}) {

		var ch Channel

		for attrName, attrData := range chAttrs.(map[string]interface{}) {

			switch attrName {
			// list of strings
			case "orderers":
				orderers := make([]string, len(attrData.([]interface{})))
				for i, v := range attrData.([]interface{}) {
					orderers[i] = fmt.Sprint(v)
				}
				ch.Orderers = orderers
			// map of map of peer attributes
			case "peers":
				for peerName, peerData := range attrData.(map[string]interface{}) {
					var peerAttr PeerAttributes
					for k, v := range peerData.(map[string]interface{}) {
						switch k {
						case "chaincodeQuery":
							peerAttr.ChaincodeQuery = v.(bool)
						case "endorsingPeer":
							peerAttr.EndorsingPeer = v.(bool)
						case "eventSource":
							peerAttr.EventSource = v.(bool)
						case "ledgerQuery":
							peerAttr.LedgerQuery = v.(bool)
						}
					}

					if ch.Peers == nil {
						ch.Peers = make(map[string]PeerAttributes)
					}

					ch.Peers[peerName] = peerAttr
				}
			}
		}

		if c.Channels == nil {
			c.Channels = make(map[string]Channel)
		}

		c.Channels[chName] = ch
	}

	//
	//  organizations map
	//

	for orgName, orgAttrs := range input["organizations"].(map[string]interface{}) {

		var org Organization

		for orgAttrName, orgAttr := range orgAttrs.(map[string]interface{}) {

			switch orgAttrName {
			// string
			case "mspid":
				org.MSPID = orgAttr.(string)
			// list of strings
			case "peers":
				peers := make([]string, len(orgAttr.([]interface{})))
				for i, v := range orgAttr.([]interface{}) {
					peers[i] = fmt.Sprint(v)
				}
				org.Peers = peers
			// map of struct
			case "signedCert":
				var cert Cert
				for k, v := range orgAttr.(map[string]interface{}) {
					switch k {
					case "pem":
						cert.Pem = v.(string)
					case "x-name":
						cert.XName = v.(string)
					}
				}
				org.SignedCert = cert
			// list of strings
			case "certificateAuthorities":
				cas := make([]string, len(orgAttr.([]interface{})))
				for i, v := range orgAttr.([]interface{}) {
					cas[i] = fmt.Sprint(v)
				}
				org.CertificateAuthorities = cas
			// array of map of struct
			case "x-uploadedSignedCerts":
				var certs []Cert
				for _, p := range orgAttr.([]interface{}) {
					var cert Cert
					for k, v := range p.(map[string]interface{}) {
						switch k {
						case "pem":
							cert.Pem = v.(string)
						case "x-name":
							cert.XName = v.(string)
						}
					}
					certs = append(certs, cert)
				}
				org.XUploadedSignedCerts = certs
			}
		}

		if c.Organizations == nil {
			c.Organizations = make(map[string]Organization)
		}

		// to copy the original behaviour
		if org.XUploadedSignedCerts == nil {
			org.XUploadedSignedCerts = []Cert{}
		}

		c.Organizations[orgName] = org
	}

	//
	//  orderers map
	//

	for ordName, ordAttrs := range input["orderers"].(map[string]interface{}) {

		var orderer Orderer

		for ordAttrName, ordAttr := range ordAttrs.(map[string]interface{}) {

			switch ordAttrName {
			// string
			case "url":
				orderer.URL = ordAttr.(string)
			// map of strings
			case "grpcOptions":
				var grpcOptions GRPCOptions
				for k, v := range ordAttr.(map[string]interface{}) {
					switch k {
					case "grpc.http2.keepalive_time":
						grpcOptions.GRPCHttp2KeepaliveTime = int64(v.(float64))
					case "grpc.keepalive_time_ms":
						grpcOptions.GRPCKeepaliveTimeMs = int64(v.(float64))
					case "grpc.http2.keepalive_timeout":
						grpcOptions.GRPCHttp2KeepaliveTimeout = int64(v.(float64))
					case "grpc.keepalive_timeout_ms":
						grpcOptions.GRPCKeepaliveTimeoutMs = int64(v.(float64))
					case "ssl-target-name-override":
						// TODO: Not sure how to parse this, need an example
						//grpcOptions.SSLTargetNameOverride =
					}
				}
				orderer.GRPCOptions = grpcOptions
			// map of string
			case "tlsCACerts":
				var tlsCACert TLSCACert
				for k, v := range ordAttr.(map[string]interface{}) {
					if k == "pem" {
						tlsCACert.Pem = v.(string)
					}
				}
				orderer.TLSCACerts = tlsCACert
			}
		}

		if c.Orderers == nil {
			c.Orderers = make(map[string]Orderer)
		}

		c.Orderers[ordName] = orderer
	}

	//
	//  peers map
	//

	for peerName, peerAttrs := range input["peers"].(map[string]interface{}) {

		var peer Peer

		for peerAttrName, peerAttr := range peerAttrs.(map[string]interface{}) {
			switch peerAttrName {
			// string
			case "url":
				peer.URL = peerAttr.(string)
			// string
			case "eventUrl":
				peer.EventURL = peerAttr.(string)
			// map of strings
			case "grpcOptions":
				var grpcOptions GRPCOptions
				for k, v := range peerAttr.(map[string]interface{}) {
					switch k {
					case "grpc.http2.keepalive_time":
						grpcOptions.GRPCHttp2KeepaliveTime = int64(v.(float64))
					case "grpc.keepalive_time_ms":
						grpcOptions.GRPCKeepaliveTimeMs = int64(v.(float64))
					case "grpc.http2.keepalive_timeout":
						grpcOptions.GRPCHttp2KeepaliveTimeout = int64(v.(float64))
					case "grpc.keepalive_timeout_ms":
						grpcOptions.GRPCKeepaliveTimeoutMs = int64(v.(float64))
					case "ssl-target-name-override":
						// TODO: Not sure how to parse this, need an example
						//grpcOptions.SSLTargetNameOverride =
					}
				}
				peer.GRPCOptions = grpcOptions
			// map of string
			case "tlsCACerts":
				var tlsCACert TLSCACert
				for k, v := range peerAttr.(map[string]interface{}) {
					if k == "pem" {
						tlsCACert.Pem = v.(string)
					}
				}
				peer.TLSCACerts = tlsCACert
			// string
			case "x-mspid":
				peer.XMSPID = peerAttr.(string)
			// string
			case "x-ledgerDbType":
				peer.XLedgerDBType = peerAttr.(string)
			}
		}

		if c.Peers == nil {
			c.Peers = make(map[string]Peer)
		}

		c.Peers[peerName] = peer
	}

	//
	//  certificateAuthorities map
	//

	for caName, caAttrs := range input["certificateAuthorities"].(map[string]interface{}) {

		var ca CA

		for caAttrName, caAttr := range caAttrs.(map[string]interface{}) {
			switch caAttrName {
			// string
			case "url":
				ca.URL = caAttr.(string)
			// map of strings
			case "httpOptions":
				var httpOptions HttpOptions
				for k, v := range caAttr.(map[string]interface{}) {
					if k == "verify" {
						httpOptions.Verify = v.(bool)
					}
				}
				ca.HttpOptions = httpOptions
			// map of strings
			case "tlsCACerts":
				var tlsCACert TLSCACert
				for k, v := range caAttr.(map[string]interface{}) {
					if k == "pem" {
						tlsCACert.Pem = v.(string)
					}
				}
				ca.TLSCACerts = tlsCACert
			// slice of map of strings
			case "registrar":
				var caRegistrars []CARegistrar
				for _, data := range caAttr.([]interface{}) {
					var caRegistrar CARegistrar
					for k, v := range data.(map[string]interface{}) {
						switch k {
						case "enrollId":
							caRegistrar.EnrollId = v.(string)
						case "enrollSecret":
							caRegistrar.EnrollSecret = v.(string)
						}
					}
					caRegistrars = append(caRegistrars, caRegistrar)
				}
				ca.Registrar = caRegistrars
			// string
			case "caName":
				ca.CAName = caAttr.(string)
			// string
			case "x-mspid":
				ca.XMSPID = caAttr.(string)
			}
		}

		if c.CAs == nil {
			c.CAs = make(map[string]CA)
		}

		c.CAs[caName] = ca
	}

	return c
}

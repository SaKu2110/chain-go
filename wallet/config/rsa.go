package config

import(
	"fmt"
	"errors"
	"io/ioutil"
	"crypto/rsa"
	"crypto/x509"

	"encoding/pem"
)

func ReadRsaPrivateKey(keyPath string) (*rsa.PrivateKey, error){
    var flag bool
    var key *rsa.PrivateKey

	bytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
    if block == nil {
        return nil, errors.New("invalid private key data")
    }
    switch block.Type {
    case "RSA PRIVATE KEY":
        key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
        if err != nil {
            return nil, err
        }
    case "PRIVATE KEY":
        keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
        if err != nil {
            return nil, err
        }
        key, flag = keyInterface.(*rsa.PrivateKey)
        if flag == false {
            return nil, errors.New("not RSA private key")
        }
    default:
        return nil, fmt.Errorf("invalid private key type : %s", block.Type)
    }

    key.Precompute()

    err = key.Validate()
    if err != nil {
        return nil, err
    }

    return key, nil
}
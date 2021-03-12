package config

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"flag"

	// "grpc_server/logutil"
	// "grpc_server/ratelimit"

	"os"

	"gopkg.in/yaml.v2"
)

var MonitorIfaces = []string{}

var ConfigVar = Config{}

var (
	KeyPair      *tls.Certificate
	CertPool     *x509.CertPool
	JwtPublicKey *ecdsa.PublicKey
	// Limiter        *ratelimit.RequestLimit
	EnableInfluxdb = flag.Bool("enable-influxdb", false, "Enable InfluxDb")
)

// func InitRateLimit(limit int) {
// 	Limiter = ratelimit.NewRequestLimit(limit)
// }

type DatabaseCfg struct {
	DbHost         string `yaml:"db_host"`
	DbType         string `yaml:"db_type"`
	DbName         string `yaml:"db_name"`
	DbUser         string `yaml:"db_user"`
	DbUserPassword string `yaml:"db_userpasswd"`
}

type Config struct {
	DatabaseConfig DatabaseCfg `yaml:"database"`
	// AuthConfig      AuthCfg        `yaml:"auth"`
	// TlsConfig       TlsCfg         `yaml:"tls"`
	// EthernetsConfig []*EthernetCfg `yaml:"ethers"`
	// InternalConfig  InternalCfg    `yaml:"internal"`
	// DiskQuotaConfig DiskQuotaCfg   `yaml:"quota"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) error {

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&ConfigVar); err != nil {
		return err
	}

	return nil
}

func GetDbUrl() map[string]string {
	return map[string]string{
		"Host":     ConfigVar.DatabaseConfig.DbHost,
		"User":     ConfigVar.DatabaseConfig.DbUser,
		"Name":     ConfigVar.DatabaseConfig.DbName,
		"Password": ConfigVar.DatabaseConfig.DbUserPassword,
		"Type":     ConfigVar.DatabaseConfig.DbType,
	}
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/slackhq/nebula"
	"github.com/slyngdk/nebula-provisioner/protocol"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"time"
)

// Build A version string that can be set with
//
//     -ldflags "-X main.Build=SOMEVERSION"
//
// at compile-time.
var Build string

var (
	l          *logrus.Logger
	configPath string
	config     *nebula.Config

	rootCmd = &cobra.Command{}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Use = os.Args[0]
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "agent.yml", "Path to either a file or directory to load configuration from")

	rootCmd.AddCommand(enrollCmd, exportCmd)
}

func initConfig() {
	l = logrus.New()
	l.Out = os.Stdout
	config = nebula.NewConfig(l)
	if configPath != "" {
		err := config.Load(configPath)
		if err != nil {
			l.WithError(err).Printf("failed to load config")
			os.Exit(1)
		}
	}
}

func main() {
	Execute()
}

type agentClient struct {
	l      *logrus.Logger
	conn   *grpc.ClientConn
	client protocol.AgentServiceClient
}

func NewClient(l *logrus.Logger, config *nebula.Config) (*agentClient, error) {

	var opts []grpc.DialOption

	cert := config.GetString("pki.cert", "agent.crt")
	key := config.GetString("pki.key", "agent.key")

	certExists, _ := fileExists(cert)
	keyExists, _ := fileExists(key)

	var keyPair tls.Certificate
	var err error

	if certExists && keyExists {
		keyPair, err = tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return nil, fmt.Errorf("unable to load keypair: %v", err)
		}
	} else if !certExists && !keyExists {
		l.Info("Generating new keypair")
		if err = generateAgentKeyPair(cert, key); err != nil {
			return nil, err
		}
		keyPair, err = tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return nil, fmt.Errorf("unable to load keypair: %v", err)
		}
	} else {
		return nil, fmt.Errorf("unable to load keypair: missing a part")
	}

	var caCertPool *x509.CertPool
	if config.IsSet("pki.ca") {
		ca := config.GetString("pki.ca", "ca.crt")
		srvcert, err := ioutil.ReadFile(ca)
		if err != nil {
			return nil, fmt.Errorf("unable to load server cert pool: %v", err)
		}
		caCertPool = x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(srvcert); !ok {
			return nil, fmt.Errorf("unable to append server cert to ca pool")
		}
	} else {
		caCertPool, err = x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("unable to load system cert pool: %v", err)
		}
	}

	ta := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair},
		RootCAs:      caCertPool,
	})

	opts = append(opts, grpc.WithTransportCredentials(ta))

	addr := config.GetString("server", "127.0.0.1:51150")
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	client := protocol.NewAgentServiceClient(conn)

	return &agentClient{l, conn, client}, nil
}

func (c agentClient) Close() error {
	return c.conn.Close()
}

func fileExists(filepath string) (bool, os.FileInfo) {

	fileinfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false, fileinfo
	}
	// Return false if the fileinfo says the file path is a directory.
	return !fileinfo.IsDir(), fileinfo
}

func generateAgentKeyPair(cert, key string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("cannot generate RSA key: %s", err)
	}

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = ioutil.WriteFile(key, pem.EncodeToMemory(privateKeyBlock), 0600)
	if err != nil {
		return fmt.Errorf("error while writing %s: %s", key, err)
	}
	serial, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	tml := x509.Certificate{
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName:   "New Name",
			Organization: []string{"New Org."},
		},
		BasicConstraintsValid: true,
	}
	publicKeyBytes, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("certificate cannot be created: %s", err)
	}

	publicKeyBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: publicKeyBytes,
	}

	err = ioutil.WriteFile(cert, pem.EncodeToMemory(publicKeyBlock), 0600)
	if err != nil {
		return fmt.Errorf("error while writing %s: %s", cert, err)
	}

	return nil
}

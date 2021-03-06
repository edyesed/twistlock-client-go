package main

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Init  InitConfig  `yaml:"init"`
	Saml  SamlConfig  `yaml:"saml"`
	Proxy ProxyConfig `yaml:"proxy"`
}

type InitConfig struct {
	Enabled bool `yaml:"enabled"`
}

type SamlConfig struct {
	Enabled    bool   `yaml:"enabled"`
	Type       string `yaml:"type"`
	URL        string `yaml:"url"`
	Cert       string `yaml:"cert"`
	Issuer     string `yaml:"issuer"`
	Audience   string `yaml:"audience"`
	ConsoleURL string `yaml:"consoleUrl"`
	TenantID   string `yaml:"tenantId"`
	AppID      string `yaml:"appId"`
	AppSecret  string `yaml:"appSecret"`
}

type ProxyConfig struct {
	HttpProxy string `yaml:"httpProxy"`
	NoProxy   string `yaml:"noProxy"`
	CaCert    string `yaml:"caCert"`
	ProxyUser string `yaml:"proxyUser"`
	ProxyPass string `yaml:"proxyPass"`
}

//var ValidSamlTypes = []string{"adfs", "azure", "gsuite", "okta", "ping", "shibboleth"}

func LoadConfig(path string) (*Config, error) {
	var c Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read error: ")
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, errors.Wrap(err, "parse error: ")
	}
	return &c, nil
}

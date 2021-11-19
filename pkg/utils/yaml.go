package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SecurityPolicyDefinition struct {
	NirikshanVersion string `yaml:"nirikshanVersion"`
	SiteConfigs      []struct {
		SiteData struct {
			SiteName         string   `yaml:"siteName"`
			ForwardingURL    string   `yaml:"forwardingUrl"`
			BlockedOs        string   `yaml:"blockedOs"`
			BlockedBrowser   string   `yaml:"blockedBrowser"`
			BlockedDevice    string   `yaml:"blockedDevice"`
			BlockedOSVersion string   `yaml:"blockedOSVersion"`
			BlockedLocations string   `yaml:"blockedLocations"`
			BlockedIPs       []string `yaml:"blockedIPs"`
		} `yaml:"siteData"`
	} `yaml:"siteConfigs"`
}

func (s *SecurityPolicyDefinition) GetConf() (*SecurityPolicyDefinition, error) {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Error("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		log.Error("Unmarshal: %v", err)
		return nil, err
	}
	return s, nil
}

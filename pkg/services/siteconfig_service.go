package services

import (
	"nirikshan-backend/pkg/entities"
)

type siteConfigService interface {
	GetSiteData(siteName string) (*entities.SiteConfigs, error)
	CreateSiteData(configs *entities.SiteConfigs) error
}

func (a applicationService) CreateSiteData(configs *entities.SiteConfigs) error {
	return a.siteConfigRepository.CreateSiteData(configs)
}

func (a applicationService) GetSiteData(siteName string) (*entities.SiteConfigs, error) {
	return a.siteConfigRepository.GetSiteData(siteName)
}

package service

import (
	"github.com/Jeffail/gabs"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ignw/cisco-aci-go-sdk/src/models"
)

// TODO: validate these settings are correct
const AP_RESOURCE_NAME_PREFIX = "ap"
const AP_OBJECT_CLASS = "fvAp"

var appProfileServiceInstance *AppProfileService

type AppProfileService struct {
	ResourceService
}

func GetAppProfileService(client *Client) *AppProfileService {
	if appProfileServiceInstance == nil {
		appProfileServiceInstance = &AppProfileService{ResourceService{
			ObjectClass:        AP_OBJECT_CLASS,
			ResourceNamePrefix: AP_RESOURCE_NAME_PREFIX,
			HasParent:          true,
		}}
	}
	return appProfileServiceInstance
}

/* New creates a new AppProfile with the appropriate default values */
func (aps AppProfileService) New(name string, description string) *models.AppProfile {

	a := models.AppProfile{models.ResourceAttributes{
		Name:         name,
		Description:  description,
		Status:       "created, modified",
		ObjectClass:  AP_OBJECT_CLASS,
		ResourceName: aps.getResourceName(name),
	},
		nil,
		nil,
	}
	//Do any additional construction logic here.
	return &a
}

func (aps AppProfileService) Save(t *models.AppProfile) error {

	err := aps.ResourceService.Save(t)
	if err != nil {
		return err
	}

	return nil

}

func (aps AppProfileService) Get(domainName string) (*models.AppProfile, error) {

	data, err := aps.ResourceService.Get(domainName)

	if err != nil {
		return nil, err
	}

	newAppProfile, err := aps.fromJSON(data)

	if err != nil {
		return nil, err
	}

	return newAppProfile, nil
}

func (aps AppProfileService) GetById(id string) (*models.AppProfile, error) {

	data, err := aps.ResourceService.GetById(id)

	if err != nil {
		return nil, err
	}

	return aps.fromJSON(data)
}

func (aps AppProfileService) GetByName(name string) ([]*models.AppProfile, error) {

	data, err := aps.ResourceService.GetByName(name)

	if err != nil {
		return nil, err
	}

	return aps.fromDataArray(data)
}

func (aps AppProfileService) GetAll() ([]*models.AppProfile, error) {

	data, err := aps.ResourceService.GetAll()
	if err != nil {
		return nil, err
	}

	return aps.fromDataArray(data)
}

func (aps AppProfileService) fromDataArray(data []*gabs.Container) ([]*models.AppProfile, error) {
	var appProfiles []*models.AppProfile
	var err, errors error

	// For each appProfile in the payload
	for _, fvAppProfile := range data {

		newAppProfile, err := aps.fromJSON(fvAppProfile)

		if err != nil {
			errors = multierror.Append(errors, err)
		} else {
			appProfiles = append(appProfiles, newAppProfile)

		}
	}

	return appProfiles, err
}

func (aps AppProfileService) fromJSON(data *gabs.Container) (*models.AppProfile, error) {
	resourceAttributes, err := aps.fromJSONToAttributes(aps.ObjectClass, data)

	if err != nil {
		return nil, err
	}

	// TODO: process child collections

	return &models.AppProfile{
		resourceAttributes,
		nil,
		nil,
	}, nil
}
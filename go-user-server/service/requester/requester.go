package requester

import (
	"os"
)

var (
	serviceURLOfHDHS string
	apiBaseURLOfHDSS string
)

func Init() {
	serviceURLOfHDHS = os.Getenv("HDHLM_WEBSERVICE_API_URL") // TODO: change to HDHS_SERVICE_URL
	apiBaseURLOfHDSS = os.Getenv("HDIWBS_API_BASE_URL")      // TODO: change to HDSS_API_BASE_URL
}

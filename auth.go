package main

const (
	OVH_APP_KEY_EU    = ""
	OVH_APP_SECRET_EU = ""
	OVH_APP_KEY_CA    = ""
	OVH_APP_SECRET_CA = ""
)

// getAppKey returns application key
func getAppKey(region string) string {
	if region == "ca" {
		return OVH_APP_KEY_CA
	}
	return OVH_APP_KEY_EU
}

// getAppSecret returns application sercret
func getAppSecret(region string) string {
	if region == "ca" {
		return OVH_APP_SECRET_CA
	}
	return OVH_APP_SECRET_EU
}

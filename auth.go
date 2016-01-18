package main

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

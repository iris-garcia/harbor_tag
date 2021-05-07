package tag

const (
	DEV_REGEX     = `^v[0-9]+\.[0-9]+.[0-9]+-dev\.[0-9]+$`
	STAGING_REGEX = `^v[0-9]+\.[0-9]+.[0-9]+-rc\.[0-9]+$`
	PROD_REGEX    = `^v[0-9]+\.[0-9]+.[0-9]+$`
)

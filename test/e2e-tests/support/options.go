package support

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	S3BucketNameEnvKey      = "S3_BUCKET"
	TargetImageTagsEnvKey   = "TARGET_IMAGE_TAGS_JSON"
	ClusterNamespaceEnvKey  = "NAMESPACE"
	GitSelfSignedCertEnvKey = "GIT_SELF_SIGNED_CERT"
	S3SelfSignedCertEnvKey  = "S3_SELF_SIGNED_CERT"
)

var (
	options *Options = nil
)

type Options struct {
	S3BucketName             string   // required
	ClusterNamespace         string   // required
	TargetImageTagReferences []string // required
	GitSelfSignedCert        string   // optional
	S3SelfSignedCert         string   // optional
}

func GetOptions() (*Options, error) {
	if options == nil {
		o, err := setOptions()
		if err != nil {
			return nil, err
		}

		options = o
	}

	return options, nil
}

func setOptions() (*Options, error) {
	if options != nil {
		return options, nil
	}

	var options = &Options{}

	if options.S3BucketName = os.Getenv(S3BucketNameEnvKey); options.S3BucketName == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", S3BucketNameEnvKey)
	}

	var err error
	if options.TargetImageTagReferences, err = parseImageTagsJSON(os.Getenv(TargetImageTagsEnvKey)); err != nil {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests: %w", TargetImageTagsEnvKey, err)
	}

	if options.ClusterNamespace = os.Getenv(ClusterNamespaceEnvKey); options.ClusterNamespace == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", ClusterNamespaceEnvKey)
	}

	if options.GitSelfSignedCert = os.Getenv(GitSelfSignedCertEnvKey); options.GitSelfSignedCert == "" {
		fmt.Printf("\noptional env variable %v not set, set it to use self-signed certs with git", GitSelfSignedCertEnvKey)
	}

	if options.S3SelfSignedCert = os.Getenv(S3SelfSignedCertEnvKey); options.S3SelfSignedCert == "" {
		fmt.Printf("\noptional env variable %v not set, set it to use self-signed certs with S3", S3SelfSignedCertEnvKey)
	}

	return options, nil
}

func parseImageTagsJSON(imageTagsJSON string) ([]string, error) {
	var imageTags []string

	err := json.Unmarshal([]byte(imageTagsJSON), &imageTags)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall image tags JSON string (%v) into []string: %w", imageTagsJSON, err)
	}

	return imageTags, nil
}

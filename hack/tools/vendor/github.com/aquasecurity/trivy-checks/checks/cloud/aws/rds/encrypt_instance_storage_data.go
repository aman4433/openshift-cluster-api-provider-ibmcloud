package rds

import (
	"github.com/aquasecurity/trivy-checks/pkg/rules"
	"github.com/aquasecurity/trivy/pkg/iac/providers"
	"github.com/aquasecurity/trivy/pkg/iac/scan"
	"github.com/aquasecurity/trivy/pkg/iac/severity"
	"github.com/aquasecurity/trivy/pkg/iac/state"
)

var CheckEncryptInstanceStorageData = rules.Register(
	scan.Rule{
		AVDID:      "AVD-AWS-0080",
		Provider:   providers.AWSProvider,
		Service:    "rds",
		ShortCode:  "encrypt-instance-storage-data",
		Summary:    "RDS encryption has not been enabled at a DB Instance level.",
		Impact:     "Data can be read from RDS instances if compromised",
		Resolution: "Enable encryption for RDS instances",
		Explanation: `Encryption should be enabled for an RDS Database instances. 

When enabling encryption by setting the kms_key_id.`,
		Links: []string{
			"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html",
		},
		Terraform: &scan.EngineMetadata{
			GoodExamples:        terraformEncryptInstanceStorageDataGoodExamples,
			BadExamples:         terraformEncryptInstanceStorageDataBadExamples,
			Links:               terraformEncryptInstanceStorageDataLinks,
			RemediationMarkdown: terraformEncryptInstanceStorageDataRemediationMarkdown,
		},
		CloudFormation: &scan.EngineMetadata{
			GoodExamples:        cloudFormationEncryptInstanceStorageDataGoodExamples,
			BadExamples:         cloudFormationEncryptInstanceStorageDataBadExamples,
			Links:               cloudFormationEncryptInstanceStorageDataLinks,
			RemediationMarkdown: cloudFormationEncryptInstanceStorageDataRemediationMarkdown,
		},
		Severity:   severity.High,
		Deprecated: true,
	},
	func(s *state.State) (results scan.Results) {
		for _, instance := range s.AWS.RDS.Instances {
			if !instance.ReplicationSourceARN.IsEmpty() {
				continue
			}
			if instance.Encryption.EncryptStorage.IsFalse() {
				results.Add(
					"Instance does not have storage encryption enabled.",
					instance.Encryption.EncryptStorage,
				)
			} else {
				results.AddPassed(&instance)
			}
		}
		return
	},
)
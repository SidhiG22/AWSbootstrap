package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformBootstrap(t *testing.T) {
	t.Parallel()
	testTerraformPath := "./"
	terraformOptions := configTerraformOptions(t, testTerraformPath)

	defer test_structure.RunTestStage(t, "teardown", func() {
		msg, err := terraform.DestroyE(t, terraformOptions)
		logger.Log(t, msg)
		if err != nil {
			logger.Log(t, "terraform destroy failed. Cleanup of resoures may be required")
		}
	})

	test_structure.RunTestStage(t, "deploy", func() {
		test_structure.SaveTerraformOptions(t, testTerraformPath, terraformOptions)
		logger.Log(t, "Running InitAndApply steps...")
		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		logger.Log(t, "Running validation steps...")

		logger.Log(t, "++++++++++++++++++++")
		logger.Log(t, "Before assertions...")
		logger.Log(t, "++++++++++++++++++++")

		bucket_id := terraform.Output(t, terraformOptions, "bucket_id")
		bucket_region := terraform.Output(t, terraformOptions, "bucket_region")

		aws.AssertS3BucketExistsE(t, bucket_region, bucket_id)

		logger.Log(t, "+++++++++++++++++")
		logger.Log(t, "After assertions.")
		logger.Log(t, "+++++++++++++++++")
	})
}

func configTerraformOptions(t *testing.T, path string) *terraform.Options {
	logger.Log(t, "Configuring Terraform options...")
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: path,
		// retry on known errors
		RetryableTerraformErrors: map[string]string{
			"net/http: TLS handshake timeout": "https://github.com/hashicorp/terraform/issues/22456",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 3 * time.Second,
	}
	return terraformOptions
}

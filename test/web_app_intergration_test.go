package test

import (
	"fmt"
	"testing"
	"time"
	"strings"
	"crypto/tls"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
)

const dbExampleDir = "../examples/mysql_database"
const webAppExampleDir = "../examples/web_app"

//Start up web app with real db
func TestIntegrationWebApp(t *testing.T) {
	//1. Make this test case parallel which means it will not block other test cases
	t.Parallel()
	//2. Deploy database
	defer test_structure.RunTestStage(t, "destroy_db", func() { destroyDb(t, dbExampleDir) })
	test_structure.RunTestStage(t, "deploy_db", func() { deployDb(t, dbExampleDir) })
	//3. Deploy web app
	defer test_structure.RunTestStage(t, "destroy_web_app", func() { destroyWebApp(t, webAppExampleDir) })
	test_structure.RunTestStage(t, "deploy_web_app", func() { deployWebApp(t, webAppExampleDir) })

	//4. Validate
	webAppOpts := test_structure.LoadTerraformOptions(t, webAppExampleDir)

	public_ip := terraform.OutputRequired(t, webAppOpts, "public_ip")
	listening_port := terraform.OutputRequired(t, webAppOpts, "listening_port")
	url := fmt.Sprintf("http://%s:%s", public_ip, listening_port)

	maxRetries := 10
	timeBetweenRetries := 10 * time.Second

	config := &tls.Config{}
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		config,
		maxRetries,
		timeBetweenRetries,
		func(status int, body string) bool {
			return status == 200 &&
				strings.Contains(body,
					fmt.Sprintf("(%s,%s) with (%s,%s)", webAppOpts.Vars["db_address"], webAppOpts.Vars["db_port"], webAppOpts.Vars["db_name"], webAppOpts.Vars["db_password"]),
				)
		})
}

func destroyWebApp(t *testing.T, webAppDir string) {
	webAppOpts := test_structure.LoadTerraformOptions(t, webAppDir)
	defer terraform.Destroy(t, webAppOpts)
}

func deployWebApp(t *testing.T, webAppDir string) {
	webAppOpts := createWebAppOpts(t, dbExampleDir, webAppDir)

	// Save data to disk so that other test stages executed at a later
	// time can read the data back in
	test_structure.SaveTerraformOptions(t, webAppDir, webAppOpts)

	terraform.InitAndApply(t, webAppOpts)
}

func createWebAppOpts(t *testing.T, dbDir string, webAppDir string) *terraform.Options {
	dbOpts := test_structure.LoadTerraformOptions(t, dbDir)
	db_address := terraform.OutputRequired(t, dbOpts, "address")
	db_port := terraform.OutputRequired(t, dbOpts, "port")
	return &terraform.Options{
		TerraformDir: webAppDir,

		Vars: map[string]interface{}{
			"security_group_name": fmt.Sprintf("testSecurityGroupName%s", random.UniqueId()),
			"ami":                 "ami-0fc20dd1da406780b",
			"instance_type":       "t2.micro",
			"db_address":          db_address,
			"db_port":             db_port,
			"db_name":             dbOpts.Vars["db_username"],
			"db_password":         dbOpts.Vars["db_password"],
		},

		// Retry up to 3 times, with 5 seconds between retries
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	}
}

func destroyDb(t *testing.T, dbDir string) {
	dbOpts := test_structure.LoadTerraformOptions(t, dbDir)
	defer terraform.Destroy(t, dbOpts)
}

func deployDb(t *testing.T, dbDir string) {
	dbOpts := createDbOpts(dbDir)

	// Save data to disk so that other test stages executed at a later
	// time can read the data back in
	test_structure.SaveTerraformOptions(t, dbDir, dbOpts)

	terraform.InitAndApply(t, dbOpts)
}

func createDbOpts(terraformDir string) *terraform.Options {
	return &terraform.Options{
		TerraformDir: terraformDir,

		Vars: map[string]interface{}{
			"db_name":        fmt.Sprintf("testDBName%s", random.UniqueId()),
			"db_username":     fmt.Sprintf("testDBUserName%s", random.UniqueId()),
			"db_password": "password",
		},

		// Retry up to 3 times, with 5 seconds between retries
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	}
}

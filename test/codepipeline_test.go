package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestPipelineInfrastructure(t *testing.T) {
    opts := &terraform.Options{
        TerraformDir: "..",
    }

    defer terraform.Destroy(t, opts)
    terraform.InitAndApply(t, opts)

    id := terraform.Output(t, opts, "codepipeline_id")
    assert.NotEmpty(t, id)
}

package rules

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

var PresetRules = map[string][]tflint.Rule{
	"all": {
		NewMcafModuleUsageRule(),
		NewMcafResourceUsageRule(),
	},
	"recommended": {
		NewMcafModuleUsageRule(),
		NewMcafResourceUsageRule(),
	},
}

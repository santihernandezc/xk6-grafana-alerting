// Package grafana-alerting contains the xk6-grafana-alerting extension.
package grafana_alerting

import (
	"fmt"
	"time"

	"github.com/grafana/grafana-openapi-client-go/models"
	"github.com/grafana/sobek"
	"go.k6.io/k6/js/modules"

	"github.com/grafana/alerting/testing/alerting-gen/pkg/execute"
)

type rootModule struct{}

func (*rootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &module{vu}
}

type module struct {
	vu modules.VU
}

func (m *module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]any{
			"GenerateGroups": m.generateGroups,
		},
	}
}

// TODO: how to parse in the incoming config options?
func parseConfig(rawConfig sobek.Value, runtime *sobek.Runtime) (execute.Config, error) {
	parsedConfig := execute.Config{}
	if rawConfig == nil || sobek.IsUndefined(rawConfig) {
		return parsedConfig, fmt.Errorf("generateGroups requires a configuration object")
	}
	converted := rawConfig.ToObject(runtime)
	if val := converted.Get("alertRuleCount"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.NumAlerting = int(val.ToInteger())
	}
	if val := converted.Get("recordingRuleCount"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.NumRecording = int(val.ToInteger())
	}
	if val := converted.Get("queryDatasource"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.QueryDS = val.String()
	}
	if val := converted.Get("writeDatasource"); val != nil && !sobek.IsUndefined(val) {

		parsedConfig.WriteDS = val.String()
	}
	if val := converted.Get("rulesPerGroup"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.RulesPerGroup = int(val.ToInteger())
	}
	if val := converted.Get("groupsPerFolder"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.GroupsPerFolder = int(val.ToInteger())
	}
	if val := converted.Get("seed"); val != nil && !sobek.IsUndefined(val) {
		parsedConfig.Seed = int64(val.ToInteger())
	} else {
		parsedConfig.Seed = time.Now().UnixNano()
	}
	// TODO: implement parsing of upload options
	return parsedConfig, nil
}

// FIXME: make this export with proper camel case instead of snake case from reflection
type GenerateGroupsOutput struct {
	Groups      []*models.AlertRuleGroup `json:"groups"`
	InputConfig execute.Config           `json:"inputConfig"`
}

func (m *module) generateGroups(rawConfig sobek.Value) *sobek.Object {
	runtime := m.vu.Runtime()
	config, err := parseConfig(rawConfig, runtime)
	if err != nil {
		panic(err)
	}
	// What type do we return? The data is an array of rules which can be json encoded
	groups, err := execute.Run(config, true)
	if err != nil {
		panic(err)
	}
	result := &GenerateGroupsOutput{
		Groups:      groups,
		InputConfig: config,
	}
	returnVal := runtime.ToValue(result)
	return returnVal.ToObject(runtime)
}

var _ modules.Module = (*rootModule)(nil)

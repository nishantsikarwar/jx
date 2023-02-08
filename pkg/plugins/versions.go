package plugins

import (
	jenkinsv1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx-helpers/v3/pkg/extensions"
)

const (
	// AdminVersion the version of the jx admin plugin
	AdminVersion = "0.2.1"

	// ApplicationVersion the version of the jx application plugin
	ApplicationVersion = "0.3.1"

	// ChangelogVersion the version of the jx changelog plugin
	ChangelogVersion = "0.9.5"

	// GitOpsVersion the version of the jx gitops plugin
	GitOpsVersion = "0.12.0"

	// HealthVersion the version of the jx health plugin
	HealthVersion = "0.0.77"

	// OctantVersion the default version of octant to use
	OctantVersion = "0.23.0"

	// OctantJXVersion the default version of octant-jx plugin to use
	OctantJXVersion = "0.0.44"

	// PipelineVersion the version of the jx pipeline plugin
	PipelineVersion = "0.5.9"

	// PreviewVersion the version of the jx preview plugin
	PreviewVersion = "0.0.232"

	// ProjectVersion the version of the jx project plugin
	ProjectVersion = "0.2.56"

	// PromoteVersion the version of the jx promote plugin
	PromoteVersion = "0.6.2"

	// SecretVersion the version of the jx secret plugin
	SecretVersion = "0.3.2"

	// TestVersion the version of the jx test plugin
	TestVersion = "0.0.51"

	// VerifyVersion the version of the jx verify plugin
	VerifyVersion = "0.3.2"
)

var (
	// Plugins default plugins
	Plugins = []jenkinsv1.Plugin{
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "admin", AdminVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "application", ApplicationVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "changelog", ChangelogVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "gitops", GitOpsVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "health", HealthVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "pipeline", PipelineVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "preview", PreviewVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "project", ProjectVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "promote", PromoteVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "secret", SecretVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "test", TestVersion),
		extensions.CreateJXPlugin(jenkinsxPluginsOrganisation, "verify", VerifyVersion),
	}

	// PluginMap a map of plugin names like `jx-gitops` to the Plugin object
	PluginMap = map[string]*jenkinsv1.Plugin{}
)

func init() {
	for i := range Plugins {
		plugin := &Plugins[i]
		PluginMap[plugin.Spec.Name] = plugin
	}
}

package cmd

const (

	// CLI version
	cliVersion = "v0.01"

	// Default username
	defaultUsername = "admin"

	// Default installation mode
	defaultMode = "namespace"

	// Default platform name
	defaultPlatform = "others"

	// Label of subscriber agent being deployed
	agentLabel = "app=subscriber"

	// Agent type is "external" for agents connected via kuberactl
	agentType = "external"

	// Default namespace for agent installation
	defaultNs = "litmus"

	// Default service account used for agent installation
	defaultSA = "litmus"
)
package cli

// CLI is the main command dispatcher.
var CLI struct {
	Debug  bool          `help:"Enable the debug mode."`
	Manage ManageCommand `cmd:"manage" help:"Start the bookmark manager."`
}

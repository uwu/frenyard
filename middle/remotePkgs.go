package middle

import (
	"github.com/CCDirectLink/CCUpdaterCLI"
	"github.com/CCDirectLink/CCUpdaterCLI/remote"
)

// FakeError should be enabled to prevent internet access by CCUpdaterUI.
const FakeError bool = false

// InternetConnectionWarning is true if the last GetRemotePackages() call actually resulted in error.
var InternetConnectionWarning bool = true

// GetRemotePackages retrieves remote packages from the server. (The CCUpdaterCLI-level cache semantics still apply.)
func GetRemotePackages() map[string]ccmodupdater.RemotePackage {
	InternetConnectionWarning = true
	if !FakeError {
		remote, err := remote.GetRemotePackages()
		if err == nil {
			InternetConnectionWarning = false
			return remote
		}
	}
	return map[string]ccmodupdater.RemotePackage{}
}

// GetLatestOf returns the latest of two possibly-nil packages (returning nil if both are nil)
func GetLatestOf(local ccmodupdater.Package, remote ccmodupdater.Package) ccmodupdater.Package {
	if local != nil {
		if remote != nil {
			if remote.Metadata().Version().GreaterThan(local.Metadata().Version()) {
				return remote
			}
		}
		return local
	}
	return remote
}

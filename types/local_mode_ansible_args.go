package types

// LocalModeAnsibleArgs is used by the local provisioner to feed Ansible with correct
// connection establishement information.
type LocalModeAnsibleArgs struct {
	Username        string
	Port            int
	PemFile         string
	KnownHostsFile  string
	BastionUsername string
	BastionHost     string
	BastionPort     int
	BastionPemFile  string
}

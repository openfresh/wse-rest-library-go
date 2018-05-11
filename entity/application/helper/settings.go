package helper

// Settings holds settings
type Settings struct {
	debug          bool
	host           string
	serverInstance string
	vhostInstance  string
	username       string
	password       string
	useDigest      bool
}

func NewSettings(
	debug bool,
	host string,
	serverInstance string,
	vhostInstance string,
	username string,
	password string,
	useDigest bool,
) *Settings {
	s := new(Settings)
	s.debug = debug
	s.host = host
	s.serverInstance = serverInstance
	s.vhostInstance = vhostInstance
	s.username = username
	s.password = password
	s.useDigest = useDigest
	return s
}

func NewDefaultSettings() *Settings {
	return NewSettings(
		false,
		"http://localhost:8087/v2",
		"_defaultServer_",
		"_defaultVHost_",
		"admin",
		"admin",
		false,
	)
}

// IsDebug get Debug.
func (s *Settings) IsDebug() bool {
	return s.debug
}

// SetDebug set debug.
func (s *Settings) SetDebug(debug bool) {
	s.debug = debug
}

// Host get host.
func (s *Settings) Host() string {
	return s.host
}

// SetHost set host.
func (s *Settings) SetHost(host string) {
	s.host = host
}

// ServerInstance get serverInstance.
func (s *Settings) ServerInstance() string {
	return s.serverInstance
}

// SetServerInstance set serverInstance.
func (s *Settings) SetServerInstance(serverInstance string) {
	s.serverInstance = serverInstance
}

// VhostInstance get vhostInstance.
func (s *Settings) VhostInstance() string {
	return s.vhostInstance
}

// SetVhostInstance set vhostInstance.
func (s *Settings) SetVhostInstance(vhostInstance string) {
	s.vhostInstance = vhostInstance
}

// Username get username.
func (s *Settings) Username() string {
	return s.username
}

// SetUsername set username.
func (s *Settings) SetUsername(username string) {
	s.username = username
}

// Password get password.
func (s *Settings) Password() string {
	return s.password
}

// SetPassword set password.
func (s *Settings) SetPassword(password string) {
	s.password = password
}

// IsUseDigest get useDigest.
func (s *Settings) IsUseDigest() bool {
	return s.useDigest
}

// SetUseDigest set debug.
func (s *Settings) SetUseDigest(useDigest bool) {
	s.useDigest = useDigest
}

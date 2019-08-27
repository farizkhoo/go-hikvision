package hikvision

const (
	defaultPort int    = 80
	defaultUser string = "admin"
)

// IPCamera describes an IPCamera
type IPCamera struct {
	ip       string
	port     int
	user     string
	password string
	prefix   string
}

// NewIPCamera returns a new instance of IPCamera
func NewIPCamera(ip, pass string) *IPCamera {
	return &IPCamera{
		ip:       ip,
		port:     defaultPort,
		user:     defaultUser,
		password: pass,
		prefix:   "ISAPI",
	}
}

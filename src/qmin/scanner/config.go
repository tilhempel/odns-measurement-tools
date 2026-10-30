package qmin_scanner

import "github.com/ilyakaznacheev/cleanenv"

type Cfg_db struct {
	// base Domain for which this server acts as Authoritve DNS Server
	BaseURL string `yaml:"BaseURL"`
	// Port on which to listen
	// DNS is 53
	Port int `yaml:"Port"`
	//TCP or UDP
	Protocol string `yaml:"Protocol"`
	// upper limit for random number of id Label in request
	// choose high to avoid two requests gettign the same id label
	RandMax int `yaml:"RandMax"`
	// how many label the request should be deep
	// used to test qmin pattern
	// > 2 to meaningfully test qmin
	LabelDepth int `yaml:"LabelDepth"`
	// how many IPs/Resolver to test in one go
	// depends on available compute power
	BatchSize int `yaml:"BatchSize"`
	// how often to test each Resolver
	Rounds int `yaml:"Rounds"`
	// normal timeout for a DNS request (in ms)
	Timeout int `yaml:"Timeout"`
	// timeout for trying again
	RetryTimeout int    `yaml:"RetryTimeout"`
	OutputDir    string `yaml:"OutputDir"`

	// the interface name of the interface used for scanning
	Iface_name string `yaml:"iface_name"`
	// the interface's IP address
	Iface_ip string `yaml:"iface_ip"`
	// the interface's IPv& address
	Iface_ip6 string `yaml:"iface_ip6"`
	// the scan's destination port, typically 53 for DNS
	Dst_port uint16 `yaml:"dst_port" env-default:"53"`
	// dns query type (only UDP)
	Dns_query_type string `yaml:"dns_query_type" env-default:"TXT"`
	// enable dnssec (only udp)
	Dnssec_enabled bool `yaml:"dnssec_enabled" env-default:"false"`
	// enable edns0 (only UDP)
	EDNS0_enabled bool `yaml:"edns0_enabled" env-default:"false"`
	// ENDS0 buffer size
	EDNS0_buffer_size int `yaml:"edns0_buffer_size" env-default:"4096"`
	// Log all dns resource records
	Log_dnsrecs bool `yaml:"log_dnsrecs" env-default:"false"`
	// list of ips or subnets in CIDR notation to exclude from the scan
	Excl_ips_fname string `yaml:"exclude_ips_fname"`
	// maximum pps to *send*
	Pkts_per_sec int `yaml:"pkts_per_sec"`
	// verbosity -> 0: OFF, 1: ERR, 2: WARN, 3: INFO, 4: DEBUG, 5: VERBOSE, 6: ALL
	Verbosity int    `yaml:"verbosity" env-default:"3"`
	Port_min  uint16 `yaml:"port_min"`
	Port_max  uint16 `yaml:"port_max"`
}

var Cfg Cfg_db

func Load_config(config_path string) {
	err := cleanenv.ReadConfig(config_path, &Cfg)
	if err != nil {
		panic(err)
	}
}

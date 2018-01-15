package main
// 定义命令行参数对应的变量
import (
	"os"
	"path/filepath"

	"github.com/CliffYuan/docker1.2.0/opts"
	flag "github.com/CliffYuan/docker1.2.0/pkg/mflag"
)

var (
	dockerCertPath = os.Getenv("DOCKER_CERT_PATH")
)
// 初始化部分参数
func init() {
	if dockerCertPath == "" {
		dockerCertPath = filepath.Join(os.Getenv("HOME"), ".docker")
	}
}

var (
	flVersion     = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
	flDaemon      = flag.Bool([]string{"d", "-daemon"}, false, "Enable daemon mode")     //是否以守护进程方式启动(依次对应:类型,使用标签，默认值，帮助信息)
	flDebug       = flag.Bool([]string{"D", "-debug"}, false, "Enable debug mode")
	flSocketGroup = flag.String([]string{"G", "-group"}, "docker", "Group to assign the unix socket specified by -H when running in daemon mode\nuse '' (the empty string) to disable setting of a group")
	flEnableCors  = flag.Bool([]string{"#api-enable-cors", "-api-enable-cors"}, false, "Enable CORS headers in the remote API")
	flTls         = flag.Bool([]string{"-tls"}, false, "Use TLS; implied by tls-verify flags")
	flTlsVerify   = flag.Bool([]string{"-tlsverify"}, false, "Use TLS and verify the remote (daemon: verify client, client: verify daemon)")

	// these are initialized in init() below since their default values depend on dockerCertPath which isn't fully initialized until init() runs
	flCa    *string
	flCert  *string
	flKey   *string
	flHosts []string
)

// 初始化部分参数
func init() {
	flCa = flag.String([]string{"-tlscacert"}, filepath.Join(dockerCertPath, defaultCaFile), "Trust only remotes providing a certificate signed by the CA given here")
	flCert = flag.String([]string{"-tlscert"}, filepath.Join(dockerCertPath, defaultCertFile), "Path to TLS certificate file")
	flKey = flag.String([]string{"-tlskey"}, filepath.Join(dockerCertPath, defaultKeyFile), "Path to TLS key file")
	opts.HostListVar(&flHosts, []string{"H", "-host"}, "The socket(s) to bind to in daemon mode\nspecified using one or more tcp://host:port, unix:///path/to/socket, fd://* or fd://socketfd.")
}
///////////////////////////////////////////////////
//知识积累:
// 在同一个文件中(或者同一个package下的不同文件可以定义多个同名的init方法)
// 在一个包被导入后，先手初始化其中的变量，然后按照代码(或文件)中的顺序依
//次执行包中的init方法,按照导入的顺序递归调用
/////////////////////////////////////////////

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/CliffYuan/docker1.2.0/api"
	"github.com/CliffYuan/docker1.2.0/api/client"
	"github.com/CliffYuan/docker1.2.0/dockerversion"
	flag "github.com/CliffYuan/docker1.2.0/pkg/mflag"
	"github.com/CliffYuan/docker1.2.0/reexec"
	"github.com/CliffYuan/docker1.2.0/utils"
)

const (
	defaultCaFile   = "ca.pem"
	defaultKeyFile  = "key.pem"
	defaultCertFile = "cert.pem"
)
// 对应命令行工具docker
func main() {
      // 检查是否有Initializer注册过,若存在(会调用它的初始化方法进行初始化)直接返回
	if reexec.Init() {
		return
	}
      // 根据用户输入，解析命令行参数
	flag.Parse()
	// FIXME: validate daemon flags here
        //根据参数解析的结果，设置(输出)用户想要的信息
	if *flVersion {
           // 对应docker -v 单纯查看docker版本信息
		showVersion()
		return
	}
	if *flDebug {
            // 设置调制模式标识(环境变量)
		os.Setenv("DEBUG", "1")
	}
         // 如果没有明确指定host(请求/daemon的地址),首先从环境变量获取,都没有指定,则设为默认的unix socket地址
	if len(flHosts) == 0 {
		defaultHost := os.Getenv("DOCKER_HOST")
		if defaultHost == "" || *flDaemon {
			// If we do not have a host, default to unix socket
			defaultHost = fmt.Sprintf("unix://%s", api.DEFAULTUNIXSOCKET)
		}
                // 并检查默认的地址是否有效
		if _, err := api.ValidateHost(defaultHost); err != nil {
			log.Fatal(err)
		}
		flHosts = append(flHosts, defaultHost)
	}
       // 若用户指定以docker daemon启动，则mianDaemon()作为宿主机守护进程启动,并直接结束
	if *flDaemon {
		mainDaemon()
		return
	}
        //用户只能指定一个请求监听地址,否则提示错误信息
	if len(flHosts) > 1 {
		log.Fatal("Please specify only one -H")
	}
        // 根据得到的监听地址按照求协议和地址分开存储
	protoAddrParts := strings.SplitN(flHosts[0], "://", 2)
	// 创建client对象变量和TLS配置变量
	var (
		cli       *client.DockerCli
		tlsConfig tls.Config
	)
	// 首先将跳过加密验证设为真
	tlsConfig.InsecureSkipVerify = true

	// 如果用户指定打开TLS加密验证，则获取授信的ca文件信息Rootcas，以上信息正确则设置启动TLS标识(将此前的跳过标识设为假)
	// If we should verify the server, we need to load a trusted ca
	if *flTlsVerify {
		*flTls = true
		certPool := x509.NewCertPool()
		file, err := ioutil.ReadFile(*flCa)
		if err != nil {
			log.Fatalf("Couldn't read ca cert %s: %s", *flCa, err)
		}
		certPool.AppendCertsFromPEM(file)
		tlsConfig.RootCAs = certPool
		tlsConfig.InsecureSkipVerify = false
	}
        //若启动了TLS加密验证，需要根据加载并发送客户端的证书
	// If tls is enabled, try to load and send client certificates
	if *flTls || *flTlsVerify {
		_, errCert := os.Stat(*flCert)
		_, errKey := os.Stat(*flKey)
		if errCert == nil && errKey == nil {
			*flTls = true
			cert, err := tls.LoadX509KeyPair(*flCert, *flKey)
			if err != nil {
				log.Fatalf("Couldn't load X509 key pair: %s. Key encrypted?", err)
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}
	// 根据先前得到的配置信息，调用clinet.NewDockerCLi()创建client实例
	if *flTls || *flTlsVerify {
		cli = client.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, protoAddrParts[0], protoAddrParts[1], &tlsConfig)
	} else {
		cli = client.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, protoAddrParts[0], protoAddrParts[1], nil)
	}
        //利用创建完成的client实例，解析先前得到的用户指定的命令参数
	if err := cli.Cmd(flag.Args()...); err != nil {
		if sterr, ok := err.(*utils.StatusError); ok {
			if sterr.Status != "" {
				log.Println(sterr.Status)
			}
			os.Exit(sterr.StatusCode)
		}
		log.Fatal(err)
	}
}

func showVersion() {
	fmt.Printf("Docker version %s, build %s\n", dockerversion.VERSION, dockerversion.GITCOMMIT)
}

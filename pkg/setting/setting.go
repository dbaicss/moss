package setting

import (
	"log"
	"time"

	"flag"
	"github.com/go-ini/ini"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

type App struct {
	PageSize  		int
	PrefixUrl 		string
	LogSavePath 	string
	LogSaveName 	string
	LogFileExt  	string
	TimeFormat  	string
	Kubecfg     	*kubernetes.Clientset
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Kubeconf struct {
	Kubecfg      	string
	KubeClientSet   *kubernetes.Clientset
}

var KubeSetting = &Kubeconf{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("kubernetes", KubeSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	kubeClients()
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
//初始化k8s clientSet
func kubeClients() {
	var (
		err error
		kubeconfig *string
		clientset *kubernetes.Clientset
	)
	/*
	//linux下获取kubeconfig路径
	if home := homeDir(); home != "" {
		kubeconfig = flag.String(home, filepath.Join(home, KubeSetting.Kubecfg, "config.dev"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String(home, "", "absolute path to the kubeconfig file")
	}
	*/
	//win下获取kubeconfig路径
	kubeconfig = flag.String("kubeconfig",KubeSetting.Kubecfg, "kube config path")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	KubeSetting.KubeClientSet = clientset
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
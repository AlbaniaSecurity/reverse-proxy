package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	config    = configRead()
	args      = *&os.Args
	targetUrl *url.URL
	target    = config.Target
	webPort   = config.WebInterfacePort
	proxyPort = config.ProxyPort
	kepAlive  = config.Kepalive
	banner    = "\bClient ⥂  Reverse-Proxy  ⥂ *Server\n"
)

func main() {

	// verify args
	if len(os.Args) < 2 {
		fmt.Println("[err] no params ,try --help for help")
		os.Exit(0)
	} else {
		args = os.Args
	}

	// for help
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		help()
		os.Exit(0)
	}

	// start reverse proxy with local config (config.json)
	if args[1] == "--listen" && args[2] == "--localconfig" || args[1] == "-l" && args[2] == "-lc" {
		proxy()
	}

	// set options and listen
	if args[1] == "--listen" && args[2] == "--set" || args[1] == "-l" && args[2] == "-s" || args[1] == "-s" && args[2] == "-l" {
		var isokay bool
		for {

			// default configs
			fmt.Println("\n[!] set mode, avoid using spaces")
			target = inputStr("[+] Set Target -> ")
			webPort = inputInt("[+] Set Web Port -> ")
			proxyPort = inputInt("[+] Set Proxy Port -> ")

			// keepAlive
			fmt.Println("\n[+] Enable KeepAliveConnection ? default: y ")
			chose := inputStr("(y/n) ")
			if chose == "y" || chose == "yes" || chose == "Y" || chose == "Yes" {
				kepAlive = true
			} else {
				kepAlive = false
			}

			// is okay ?
			fmt.Println("\n:: Using Config ::")
			fmt.Println("Target: ", target)
			fmt.Println("Wep Port: ", webPort)
			fmt.Println("Proxy Port: ", proxyPort)
			fmt.Println("Keep-a-Live Connection: ", kepAlive)

			okay := inputStr("\nOkay ? (y/n) ")
			if okay == "y" || chose == "yes" || chose == "Y" || chose == "Yes" {
				isokay = true
				fmt.Println("\033[2J", banner)
				break
			} else {
				fmt.Println("\033[2J", banner)
			}
		}
		if isokay {
			proxy()
		}
	}
}

// help banner
func help() {
	fmt.Println("\033[2J", banner)

	fmt.Println(":: Help Menu ::")
	fmt.Println(":: Use: ./proxy --listen --localconfig | to simple mode ")
	fmt.Println("\nArgs: ")
	fmt.Println(" --listen, -l       | listen proxy server ")
	fmt.Println(" --set, -s          | set options for listen")
	fmt.Println(" --localconfig, -lc | listen using local config\n")
}

// input
func inputStr(text string) string {
	fmt.Print(text)
	var temp string
	fmt.Scanln(&temp)
	return temp
}

func inputInt(text string) int {
	fmt.Print(text)
	var temp int
	fmt.Scanln(&temp)
	return temp
}

// handle error
func he(e error) {
	if e != nil {
		panic(e)
	}
}

// open file
type Config struct {
	Target           string `json:"target"`
	Kepalive         bool   `json:"kepalive"`
	WebInterfacePort int    `json:"web-interface-port"`
	ProxyPort        int    `json:"proxy-port"`
}

func configRead() Config {
	filename := "config.json"

	f, err := ioutil.ReadFile(filename)
	he(err)

	var obj Config
	json.Unmarshal(f, &obj)

	return obj
}

// web server
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	he(err)
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// web server
func webserver() {
	file := http.FileServer(http.Dir("./web"))
	http.Handle("/", file)

	http.HandleFunc("/addr", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, fmt.Sprintf("%s:%d", GetOutboundIP(), proxyPort))
	})

	http.ListenAndServe(":"+strconv.Itoa(webPort), nil)
}

// reverse proxy
func proxy() {
	fmt.Println("\033[2J", banner)
	go webserver()
	var err error
	targetUrl, err = url.Parse(target) // proxy url
	he(err)

	handler := http.NewServeMux()
	handler.HandleFunc("/", proxyHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s", GetOutboundIP()) + ":" + strconv.Itoa(proxyPort),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second, // caso queira uma proxy mais rapida deixe o timeout menor, porem se sua placa de rede for fraca deixe no padrao de 15 mesmo
		IdleTimeout:  15 * time.Second,
	}
	server.SetKeepAlivesEnabled(kepAlive)

	fmt.Println("[*] Web Interface     -> http://127.0.0.1:" + strconv.Itoa(webPort))
	fmt.Printf("[*] Proxy Listening on -> %s\n", server.Addr)

	log.Fatalln(server.ListenAndServe()) // print if error

}

/*
	handle sockets requests tcp/udp for addr:port proxy
*/
func proxyHandler(writer http.ResponseWriter, request *http.Request) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	request.Host = targetUrl.Host
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(writer, request)
}

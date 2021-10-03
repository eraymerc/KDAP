package sunucu

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

//local ip---------------------------------
func isFlagUp(f net.Flags) bool {
	if f&(1<<uint(0)) == 1 {
		return true
	}
	return false
}
func isFlagLoopback(f net.Flags) bool {
	if f&(1<<uint(2)) == 4 {
		return true
	}
	return false
}

//yerel ip listesi
func YerelIP() string {
	out := ""
	ifaces, err := net.Interfaces()
	if err != nil {
		out = "hata!"
		return out
	}

	for _, ifc := range ifaces {
		addrs, err := ifc.Addrs()
		if err != nil {
			fmt.Println(err)
		}

		if ifc.Flags != 0 {

			if isFlagUp(ifc.Flags) {

				if !isFlagLoopback(ifc.Flags) {
					out += ifc.Name + "\n"
					for _, addr := range addrs {
						out += "\t" + addr.String() + "\n"
					}
				}
			}
		}
	}
	return out
}

//---------------------------------------
//public ip
func GenelIP() string {
	apiUrl := "https://api.ipify.org?format=text"
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "api servisine bağlanılamadı :("
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "api yanıtı okunamadı :("
	}
	return string(ip)
}

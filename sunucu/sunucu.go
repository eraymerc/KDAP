package sunucu

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var server Server

func (server *Server) KlasorMu(path string) bool {
	fileInfo, err := os.Stat(server.kokDizin + path)
	if err != nil {
		// error handling
	}

	if fileInfo.IsDir() {
		return true
	}
	return false
}

func (server *Server) KlasorListe(path string) {
	list, err := filepath.Glob(server.kokDizin + path + "\\*")
	if err != nil {
		return //dosyalar listelenemedi
	}
	for i, item := range list {

		list[i] = strings.Replace(item, server.kokDizin, "", -1)

	}
	server.files = list
}

type Server struct {
	kokDizin string
	port     int
	sifre    string
	isKlasor bool
	files    []string
	listener net.Listener
}

func (server *Server) listenPort() {
	//listen port

	l, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))
	if err != nil {
		server.Kapat()
		server.listenPort()
		return
	}
	server.listener = l
}

func Baslat(filename string, port int, sifre string, ishttp bool, isKlasor bool) {

	var files []string
	var listener net.Listener
	server = Server{filename + "\\", port, sifre, isKlasor, files, listener}

	if isKlasor {
		server.KlasorListe("")
	}
	server.listenPort()
	//sunucu
	if ishttp {
		server.HTTPServer()
	} else {
		//tcp
	}
}

func (server *Server) Kapat() {

	server.listener.Close()
}

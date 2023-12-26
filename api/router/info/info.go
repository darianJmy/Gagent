package info

import (
	"bytes"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/emicklei/go-restful/v3"

	"gagent/pkg/types"
)

func GetCPUInfo(req *restful.Request, resp *restful.Response) {
	var r = types.Resp{
		Code: 200,
		Msg: types.CPU{
			Number: runtime.NumCPU(),
		},
	}

	resp.WriteHeaderAndJson(200, r, "application/json")
}

func GetNetWorkInfo(req *restful.Request, resp *restful.Response) {
	var network []types.Network
	interfaces, err := net.Interfaces()
	if err != nil {
		resp.WriteHeaderAndJson(200, types.Resp{Code: 400}, "application/json")
		return
	}

	for _, v := range interfaces {
		if v.Name != "lo" {
			network = append(network, types.Network{Name: v.Name})
		}
	}

	resp.WriteHeaderAndJson(200, types.Resp{Code: 200, Msg: network}, "application/json")
}

func GetSystemUUID(req *restful.Request, resp *restful.Response) {
	cmd := exec.Command("bash", "-c", "dmidecode -s system-uuid")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		resp.WriteHeaderAndJson(200, types.Resp{Code: 400, Msg: err}, "application/json")
		return
	}

	resp.WriteHeaderAndJson(200, types.Resp{Code: 200, Msg: types.System{UUID: strings.TrimSpace(stdout.String())}}, "application/json")
}

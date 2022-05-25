package parse

import (
	"fmt"
	"nacs/common"
	"nacs/utils"
	"strings"
)

func Parse(InputInfo *common.InputInfoStruct, RunningInfo *common.RunningInfoStruct) {
	if !strings.Contains(InputInfo.OutputFileName, "/") {
		RunningInfo.OutputFileName = utils.GetPath() + InputInfo.OutputFileName
	} else {
		RunningInfo.OutputFileName = InputInfo.OutputFileName
	}

	RunningInfo.Silent = InputInfo.Silent
	RunningInfo.NoSave = InputInfo.NoSave
	RunningInfo.LogLevel = InputInfo.LogLevel

	var err error
	RunningInfo.Hosts, err = ParseIP(InputInfo.Host, InputInfo.HostFile, InputInfo.SkipHost)
	if err != nil {
		fmt.Println("len(hosts)==0", err)
		return
	}
	RunningInfo.NoProbe = InputInfo.NoProbe

	if InputInfo.PortsOnly != "" {
		common.RunningInfo.Ports = OnlyPorts(InputInfo.PortsOnly)
	} else if InputInfo.PortsAdd != "" {
		common.RunningInfo.Ports = AddPorts(InputInfo.PortsAdd)
	} else {
		common.RunningInfo.Ports = common.DefaultPorts
	}

	RunningInfo.Thread = InputInfo.Thread

	ProxyParse()

	RunningInfo.DiscoverTimeout = InputInfo.Timeout
	RunningInfo.DiscoverMode = InputInfo.DiscoverMode
	RunningInfo.DiscoverType = InputInfo.DiscoverType
	RunningInfo.OutJson = InputInfo.OutJson

	ParseReversePlatform(InputInfo.CeyeKey, InputInfo.CeyeDomain)

	RunningInfo.NoPoc = InputInfo.NoPoc
	RunningInfo.NoBrute = InputInfo.NoBrute

	RunningInfo.PocTimeout = InputInfo.PocTimeout
	RunningInfo.PocRate = InputInfo.PocRate
	RunningInfo.PocThread = InputInfo.PocThread

	RunningInfo.FscanPocPath = InputInfo.FscanPocPath
	RunningInfo.NucleiPocPath = InputInfo.NucleiPocPath
	RunningInfo.PocDebug = InputInfo.PocDebug

	RunningInfo.Command = InputInfo.Command
	RunningInfo.SSHKey = InputInfo.SSHKey

	ParseUser(InputInfo)
	ParsePass(InputInfo)

	RunningInfo.BruteTimeout = InputInfo.BruteTimeout
	RunningInfo.BruteDebug = InputInfo.BruteDebug
	RunningInfo.BruteSocks5Proxy = InputInfo.BruteSocks5Proxy

	RunningInfo.RedisShell = InputInfo.RedisShell
	RunningInfo.RedisFile = InputInfo.RedisFile

	RunningInfo.PocSocks5Proxy = InputInfo.BruteSocks5Proxy

	RunningInfo.NoNuclei = InputInfo.NoNuclei
	RunningInfo.BruteThread = InputInfo.BruteThread
}
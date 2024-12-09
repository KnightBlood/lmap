/*
 *     lmap (LinuxHub's Nmap) is the nmap next generation pro plus max.
 *     Copyright (C) <2021>  <LinuxHub-Group>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package lmap

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const OUTPUT_IP_PER_LINE = 3

func CheckIP(subnets, excludes []string, isVerbose bool) {
	checkerGroup := &sync.WaitGroup{}
	t := time.Now()

	var allHosts []HostInfo
	for _, subnet := range subnets {
		hosts, _ := GetAllIPsFromCIDR(subnet, excludes)
		allHosts = append(allHosts, hosts...)
	}

	for index := range allHosts {
		checkerGroup.Add(1)

		go func(index int) {
			defer checkerGroup.Done()
			allHosts[index].isUsed = Ping(allHosts[index].host)
			if isVerbose {
				if allHosts[index].isUsed {
					fmt.Println("已使用IP：", allHosts[index].host.String())
				} else {
					fmt.Println("未使用IP：", allHosts[index].host.String())
				}
			}
		}(index)
	}

	checkerGroup.Wait()
	elapsed := time.Since(t)
	_, _ = fmt.Fprintln(os.Stderr, "IP扫描完成,耗时", elapsed)
	fmt.Println("已使用IP：")
	printIPList(allHosts, true)
	fmt.Println("未使用IP：")
	printIPList(allHosts, false)
}

func printIPList(hosts []HostInfo, boolFilter bool) {

	position := 1

	for _, hostInfo := range hosts {
		if boolFilter == hostInfo.isUsed {
			fmt.Print(hostInfo.host.String())
			if position%OUTPUT_IP_PER_LINE == 0 {
				fmt.Println()
			} else {
				fmt.Print(", ")
			}
			position++
		}
	}
	fmt.Println()
}

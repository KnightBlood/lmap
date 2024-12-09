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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LinuxHub-Group/lmap/pkg/lmap"
)

func main() {
	isVerbose := false
	flag.BoolVar(&isVerbose, "v", false, "be verbose")

	subnets := multiArg{}
	flag.Var(&subnets, "subnet", "network to scan (can be specified multiple times)")

	excludes := multiArg{}
	flag.Var(&excludes, "exclude", "IP or subnet to exclude (can be specified multiple times)")

	flag.Parse()

	if len(subnets) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "使用方法：%s [-v] -subnet <网络号>/<CIDR> [-subnet ...] [-exclude <IP>/<CIDR>] [-exclude ...]\n", os.Args[0])
		os.Exit(-1)
	}

	lmap.CheckIP(subnets, excludes, isVerbose)
}

type multiArg []string

func (m *multiArg) String() string {
	return fmt.Sprintf("%v", *m)
}

func (m *multiArg) Set(value string) error {
	*m = append(*m, value)
	return nil
}

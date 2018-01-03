/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 15:35:31 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 21:00:24 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"os/exec"
)

/*
**** PUBLIC ********************************************************************
*/

func getAddr() string {
	var ret string

	addrs, err := net.InterfaceAddrs()
	handleErr(err)

	for _, value := range addrs {
		ip, _, err := net.ParseCIDR(value.String())
		handleErr(err)
		if ip.To4() != nil && ip.IsLoopback() == false {
			ret = ip.String()
		}
	}
	return ret
}

func getGuid() string {
	cmd := exec.Command("sh", "-c", "/usr/bin/base64 /dev/random | /usr/bin/head -c 64")
	out, _ := cmd.Output()
	return (string(out))
}

func getHash(str string) uint32 {
	ret := uint32(0)
	for _, value := range str {
		ret += uint32(value)
		ret += (ret << 10)
		ret ^= (ret >> 6)
	}
	ret += (ret << 3)
	ret ^= (ret >> 11)
	ret += (ret << 15)
	ret = (ret % UINT_MAX)
	return ret
}

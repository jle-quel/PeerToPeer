/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:13 by Jefferso          #+#    #+#             */
/*   Updated: 2017/12/27 17:53:06 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os/exec"
	"net"
)

const USHRT_MAX = 65535

type s_broadcast struct {
	guid string
	hash uint16
	addr string
}

/* ************************************************************************** */

func getGuid() string {
	cmd := exec.Command("/bin/sh", "-c", "/usr/bin/base64 /dev/random | /usr/bin/head -c 64")
	out, _ := cmd.Output()
	return (string(out))
}

// CONFLICT
func getHash(str string) uint16 {
	var ret uint16

	ret = 0
	for _, value := range str {
		ret += uint16(value)
		ret += (ret << 10)
		ret ^= (ret >> 6)
	}
	ret += (ret << 3)
	ret ^= (ret >> 11)
	ret += (ret << 15)
	ret = (ret % USHRT_MAX)
	return ret
}

func getAddr() string {
	var ret string

	addrs, err := net.InterfaceAddrs()
	errHandler(err)

	for _, value := range addrs {
		ip, _, err := net.ParseCIDR(value.String())
		errHandler(err)
		if ip.To4() != nil && ip.IsLoopback() == false {
			ret = ip.String()
		}
	}

	return ret
}

func getHeader() s_broadcast {
	guid := getGuid()
	header := s_broadcast{guid, getHash(guid), getAddr()}
	return header
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:13 by Jefferso          #+#    #+#             */
/*   Updated: 2017/12/27 13:48:03 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os/exec"
	"os"
	"net"
	"fmt"
)

const USHRT_MAX = 65535

type s_broadcast struct {
	guid string
	hash uint16
	addr string
}

/* ************************************************************************** */

func getGuid() string {
	cmd := exec.Command("sh", "-c", "/usr/bin/base64 /dev/random | /usr/bin/head -c 64")
	out, _ := cmd.Output()
	return (string(out))
}

func getHash(str string) uint16 {
	var ret	uint16

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

func getAddr() {
	host, err := os.Hostname()
	handleErr(err)
	addrs, err := net.LookupIP(host)
	handleErr(err)

	fmt.Println(host)
	for _, value := range addrs {
		fmt.Println(value)
	}
}

func getHeader() s_broadcast {
	guid := getGuid()
	header := s_broadcast{guid, getHash(guid), "lol"}
	return header
}

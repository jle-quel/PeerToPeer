/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   header.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:59:51 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/14 14:59:03 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"time"
	"os/exec"
)

/*
**** PRIVATE *******************************************************************
*/

func getId() string {
	cmd := exec.Command("/bin/sh", "-c", "/usr/bin/base64 /dev/urandom | /usr/bin/head -c 64")
	ret, err := cmd.Output()
	handleErr(err)
	return string(ret)
}

func getTimestamp() uint32 {
	return uint32(time.Now().Unix())
}

func getAddr() string {
	var addr string

	addrs, err := net.InterfaceAddrs()
	handleErr(err)
	for _, value := range addrs {
		ip, _, err := net.ParseCIDR(value.String())
		handleErr(err)
		if ip.IsLoopback() == false && ip.To4() != nil {
			addr = ip.String()
		}
	}
	return addr
}

/*
**** PUBLIC ********************************************************************
*/

func initHeader() func() header {
	id := getId()
	return func() header {
		return header{id, getTimestamp(), getAddr()}
	}
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/02 18:14:17 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
	"os"
)

const HEADER_SIZE = 10
const PORT42 = ":8888"

/*
**** PRIVATE *******************************************************************
*/

func debug(buf []byte, size int, addr string) {
}

func server() {

	addr, err := net.ResolveUDPAddr("udp4", PORT42)
	handleErr(err)
	conn, err := net.ListenUDP("udp", addr)
	handleErr(err)
	defer conn.Close()

	buf := make([]byte, HEADER_SIZE)
	size, ip, err := conn.ReadFrom(buf)
	handleErr(err)

	fmt.Println("Peer found", size, string(buf))
	fmt.Println("Ip", ip)
}

/*
**** PUBLIC ********************************************************************
*/

func main() {

	switch len(os.Args) {
	case 2:
		fmt.Println("Broadcasting ...")
		broadcast()
	default:
		fmt.Println("Waiting for peers ...")
		server()
	}
}

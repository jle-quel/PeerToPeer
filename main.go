/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/04 12:06:26 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"strings"
)

const UINT_MAX =  4294967295
const BROADCAST_ADDR = "255.255.255.255:8888"
const HEADER_SIZE = 10
const BROADCAST_PORT = ":8888"
const BUF_SIZE = 1024

func mainServer() {
	fmt.Println("Listening for peers ...\n")
	ch := make(chan t_map)
	go server(ch)

	for {
		fmt.Println(<- ch)
	}
}

func mainBroadcast() {
	fmt.Println("Broadcasting ...")
	broadcast(s_header{getAddr(), os.Args[1], getHash(os.Args[1])}.Encode())
}

/*
**** PUBLIC ********************************************************************
*/

func main() {
	switch len(os.Args) {
	case 2:
		// Temporary
		switch strings.Compare(os.Args[1], "root") {
		case 0:
			mainServer()
		default:
			mainBroadcast()
		}
	default:
		fmt.Println("Usage: ./p2p [username]")
	}
}s

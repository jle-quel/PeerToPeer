/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/03 16:45:26 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

const UINT_MAX =  4294967295
const BROADCAST_ADDR = "255.255.255.255:8888"
const HEADER_SIZE = 10
const BROADCAST_PORT = ":8888"
const BUF_SIZE = 1024

/*
**** PUBLIC ********************************************************************
*/

func main() {

	switch len(os.Args) {
	case 2:
		fmt.Println("Broadcasting ...\n")
		broadcast()
	default:
		fmt.Println("Waiting for peers ...\n")
		server()
	}
}

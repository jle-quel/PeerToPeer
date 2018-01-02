/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tools.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 21:36:30 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/02 18:14:01 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"net"
)

const UINT_MAX =  4294967295

/*
**** PUBLIC ********************************************************************
*/

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error ->", err)
		os.Exit(1)
	}
}

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

// CONFLICT
// func getHash(str string) uint32 {
// 	var ret uint32
//
// 	ret = 0
// 	for _, value := range str {
// 		ret += uint32(value)
// 		ret += (ret << 10)
// 		ret ^= (ret >> 6)
// 	}
// 	ret += (ret << 3)
// 	ret ^= (ret >> 11)
// 	ret += (ret << 15)
// 	ret = (ret % UINT_MAX)
// 	return ret
// }

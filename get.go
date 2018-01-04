/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 15:35:31 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 12:42:23 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
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

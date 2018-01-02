/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/02 13:30:49 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/02 20:18:07 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
	"fmt"
	"os"
)

type s_header struct {
	addr string
	username string
	hash uint32
}

func broadcast() {
	addr, err := net.ResolveUDPAddr("udp4", BROADCAST_ADDR)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)

	new := s_header{getAddr(), os.Args[1], getHash(os.Args[1])}
	fmt.Println(new)

	conn.Close()
}

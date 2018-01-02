/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/02 13:30:49 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/02 18:14:31 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"net"
)

func broadcast() {
	addr, err := net.ResolveUDPAddr("udp4", PORT42)
	conn, err := net.DialUDP("udp", nil, addr)
	handleErr(err)

	conn.Write([]byte(getAddr()))
	conn.Close()
}

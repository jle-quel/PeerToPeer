/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   send.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/14 14:55:19 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/16 13:29:05 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"time"
)

/*
**** PUBLIC ********************************************************************
*/

func (self header) Broadcast() {
	conn := initUDPConn()
	_, err := conn.Write(self.Encode())
	handleErr(err)
	conn.Close()
}

func (self header) Bootstrap(ip string) {
	time.Sleep(1 * time.Second)
	conn := initTCPConn(ip)
	_, err := conn.Write(self.Encode())
	handleErr(err)
	conn.Close()
}

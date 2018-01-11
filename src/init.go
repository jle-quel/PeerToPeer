/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/09 14:45:09 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/09 15:54:01 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
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

func getMaj() uint32 {
	return uint32(time.Now().Unix())
}

/*
**** PUBLIC ********************************************************************
*/

func initHeader() func() s_header {
	id := getId()

	return func() s_header {
		return s_header{id, getMaj()}
	}
}

func initRoutingTable() func(peer s_header) []s_header {
	new := make([]s_header, 0)

	return func(peer s_header) []s_header {
		new = append(new, peer)
		return new
	}
}

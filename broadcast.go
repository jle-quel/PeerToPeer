/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   broadcast.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 10:41:07 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 19:36:57 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func broadcast(b []byte) {
	fmt.Println("Broadcasting...")

	conn := initUDPConn()
	_, err := conn.Write(b)
	handleErr(err)
	conn.Close()
}

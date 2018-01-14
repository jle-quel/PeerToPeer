/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/14 20:19:21 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/*
**** PUBLIC ********************************************************************
*/

func main() {
	ch := make(chan t_map)
	getHeader := initHeader()

	getHeader().Broadcast()
	go UDPServer(getHeader, ch)
	TCPServer(ch)
}

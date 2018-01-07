/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/04 18:58:05 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func core() {
	fmt.Println("Boot up")
	broadcast(s_header{getAddr(), os.Args[1]}.Encode())
	// server()
}

/*
**** PUBLIC ********************************************************************
*/

func main() {
	switch {
	case len(os.Args) == 2:
		switch {
		case len(os.Args[1]) > 10:
			fmt.Println("XYZ: username must be <= 10char")
		default:
			core()
		}
	default:
		fmt.Println("XYZ: ./p2p [username]")
	}
}

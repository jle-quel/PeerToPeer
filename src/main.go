/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 11:52:54 by jle-quel         ###   ########.fr       */
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

// func decode(code []byte) header {
// 	var buf header
// 	b := bytes.NewReader(code)
// 	err := json.NewDecoder(b).Decode(&buf)
// 	handleErr(err)
// 	return buf
// }

/*
**** PUBLIC ********************************************************************
*/

func main() {
	// Peer Discovery
	getHeader := initHeader()
	broadcast(getHeader().Encode())
	UDPServer()

}

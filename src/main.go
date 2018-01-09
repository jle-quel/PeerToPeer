/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 21:58:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/09 12:29:47 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "time"

func main() {
	peer := getHeader()

	// Peer Discovery
	//peerDiscovery(self)
	//listenUDP()

	// Peer Bootstrap
	// peerBootstrap(peer().Encode().Bytes(), "192.168.0.11")
	// listenTCP()
}

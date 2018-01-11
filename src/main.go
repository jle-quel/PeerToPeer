/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 21:58:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/11 12:49:26 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func main() {
	// getPeer := initHeader()

	// Peer Discovery
	//peerDiscovery(self)
	//listenUDP()

	// Peer Bootstrap
	// peerBootstrap(getPeer().Encode().Bytes(), "192.168.0.11")
	listenTCP()
}

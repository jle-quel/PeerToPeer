/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 21:58:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/08 22:18:47 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import _ "fmt"

func main() {
	// Peer Discovery
	PeerDiscovery(s_header{getId(), getMaj()}.Encode().Bytes())
	// go listenUDP()
	// go listenTCP()
}
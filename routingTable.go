/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routingTable.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 18:34:40 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 19:30:38 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type s_routing struct {
	Addr string
	Username string
}

type routingTable map[uint32]s_routing

/*
**** PUBLIC ********************************************************************
*/

func tableClosure(peer s_header) func() routingTable {

	table := make(routingTable)
	return func() routingTable {
		table[peer.Hash] = s_routing{peer.Addr, peer.Username}
		return table
	}
}

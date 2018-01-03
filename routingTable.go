/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routingTable.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 18:34:40 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/03 20:56:45 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type s_routing struct {
	Addr string
	Username string
}

type t_map map[uint32]s_routing

/*
**** PUBLIC ********************************************************************
*/

func routingTable() func(peer s_header) t_map {
	table := make(t_map)
	return func(peer s_header) t_map {
		table[peer.Hash] = s_routing{peer.Addr, peer.Username}
		return table
	}
}

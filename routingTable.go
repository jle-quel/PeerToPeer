/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routingTable.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/03 18:34:40 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 13:53:37 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type t_map map[string]string

/*
**** PUBLIC ********************************************************************
*/

func routingTable() func(peer s_header) t_map {
	table := make(t_map)
	return func(peer s_header) t_map {
		table[peer.Username] = peer.Addr
		return table
	}
}

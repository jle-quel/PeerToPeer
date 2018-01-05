/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   routingTable.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 14:31:18 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/05 14:40:27 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

/*
**** PUBLIC ********************************************************************
*/

func tableClosure() func(addr string, guid string) t_map {
	routingTable := make(t_map)
	return func(addr string, guid string) t_map {
		routingTable[guid] = addr
		return routingTable
	}
}

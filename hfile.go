/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   hfile.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:40:05 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 20:00:06 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

/*
**** STRUCTURES ****************************************************************
*/

type header struct {
	Id string
	Timestamp uint32
	Addr string
}

type routingTable struct {
	Addr string
}

/*
***** TYPES ********************************************************************
*/

type t_map map[string]routingTable

/*
**** CONSTANTES ****************************************************************
*/

const BROADCAST_ADDR = "255.255.255.255:8888"
const BROADCAST_PORT = ":8888"
const HEADER_SIZE = 119
const TCP_PORT = ":8889"

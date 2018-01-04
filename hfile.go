/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   hfile.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/04 16:27:42 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/04 17:21:50 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type s_header struct {
	Addr string
	Username string
}

type t_map map[string]string

type t_bytes []byte

const UINT_MAX =  4294967295
const BROADCAST_ADDR = "10.13.255.255:8888"//"255.255.255.255:8888"
const BROADCAST_PORT = ":8888"
const HEADER_SIZE = 46

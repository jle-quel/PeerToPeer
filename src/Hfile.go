/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Hfile.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 22:06:01 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/09 15:06:49 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type s_header struct {
	Guid string
	Maj uint32
}

type t_byte []byte

const BROADCAST_ADDR = "255.255.255.255"
const BROADCAST_PORT = ":8888"
const BOOTSTRAP_PORT = ":8889"
const HEADERSIZE = 93

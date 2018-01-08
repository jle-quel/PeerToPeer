/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Hfile.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 22:06:01 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/08 22:27:53 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type s_header struct {
	Guid string
	Maj uint32
}

type t_bytes []byte

const BROADCAST_ADDR = "255.255.255.255"
const BROADCAST_PORT = ":8888"

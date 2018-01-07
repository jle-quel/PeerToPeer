/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   hfile.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 11:38:19 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/07 21:56:41 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

const BROADCAST_ADDR = "10.13.255.255:8888"
const RT_PORT = ":8888"
const DATA_PORT = ":8889"

type s_header struct {
	Guid string
	N uint32
}

type t_map map[string]string

type t_bytes []byte

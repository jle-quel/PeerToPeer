/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   json.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/04 11:19:29 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/04 19:26:17 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"encoding/json"
	"bytes"
)

/*
**** PUBLIC ********************************************************************
*/

func (self s_header) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(self)
	handleErr(err)
	return buf
}

func (self t_bytes) DecodeHeader() s_header {
	var peer s_header

	b := bytes.NewReader(self)
	err := json.NewDecoder(b).Decode(&peer)
	handleErr(err)
	return peer
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   json.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/04 11:19:29 by Jefferso          #+#    #+#             */
/*   Updated: 2018/01/04 11:54:19 by Jefferson        ###   ########.fr       */
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
	json.NewEncoder(buf).Encode(self)
	return buf
}

func (self t_map) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(self)
	return buf
}

func (self t_bytes) DecodeHeader() s_header {
	var peer s_header

	b := bytes.NewReader(self)
	json.NewDecoder(b).Decode(&peer)
	return peer
}

func (self t_bytes) DecodeTable() t_map {
	var table t_map

	b := bytes.NewReader(self)
	json.NewDecoder(b).Decode(&table)
	return table
}

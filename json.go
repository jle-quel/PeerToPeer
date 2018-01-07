/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   json.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/07 21:50:20 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/07 21:50:56 by jle-quel         ###   ########.fr       */
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

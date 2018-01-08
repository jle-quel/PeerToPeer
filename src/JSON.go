/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   JSON.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/08 22:14:03 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/08 22:16:10 by jle-quel         ###   ########.fr       */
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

func (this s_header) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(this)
	handleErr(err)
	return buf
}

func (this t_bytes) Decode() s_header {
	var peer s_header

	b := bytes.NewReader(this)
	err := json.NewDecoder(b).Decode(&peer)
	handleErr(err)
	return peer
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   json.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 19:35:42 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/13 19:37:44 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bytes"
	"encoding/json"
)

func decode(peer []byte) header {
	var ret header

	b := bytes.NewReader(peer)
	err := json.NewDecoder(b).Decode(&ret)
	handleErr(err)
	return ret
}

func (self header) Encode() []byte {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(self)
	handleErr(err)
	return b.Bytes()
}

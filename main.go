/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: Jefferso <Jefferso@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/26 20:55:08 by Jefferso          #+#    #+#             */
/*   Updated: 2017/12/27 13:22:50 by Jefferson        ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

/* ************************************************************************** */

func	main() {
	var header	s_broadcast

	// CONFLICT WITH THE HASH
	header = getHeader()
	fmt.Println("Guid ->", header.guid)
	fmt.Println("Hash ->", header.hash)
	fmt.Println("Addr ->", header.addr)
	fmt.Printf("\n")
	getAddr()
}

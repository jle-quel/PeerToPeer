/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   signal.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/15 16:17:13 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 16:58:34 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"
	"os/signal"
)

func handleSignal(getHeader func()header ) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c
	getHeader().Broadcast()
	os.Exit(42)
}

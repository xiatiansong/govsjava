package com.lumi.largedata.govsjava.netty;

import io.netty.bootstrap.Bootstrap;
import io.netty.buffer.Unpooled;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.DatagramPacket;
import io.netty.channel.socket.nio.NioDatagramChannel;
import io.netty.util.CharsetUtil;

import java.net.InetSocketAddress;

/**
 * A UDP broadcast client that asks for a quote of the moment (QOTM) to {@link QuoteOfTheMomentServer}.
 *
 * Inspired by <a href="http://docs.oracle.com/javase/tutorial/networking/datagrams/clientServer.html">the official
 * Java tutorial</a>.
 */
public final class NettyUDPClient {

	static final int PORT = Integer.parseInt(System.getProperty("port", "8080"));

	public static void main(String[] args) throws Exception {

		EventLoopGroup group = new NioEventLoopGroup();
		try {
			Bootstrap b = new Bootstrap();
			b.group(group).channel(NioDatagramChannel.class).option(ChannelOption.SO_BROADCAST, true)
					.handler(new NettyUDPClientHandler());

			Channel ch = b.bind(0).sync().channel();

			long t1 = System.currentTimeMillis();
			System.out.println(t1);
			// Broadcast the QOTM request to port 8080.
			for (int i = 0; i < 500000; i++) {
				ChannelFuture cf = ch.writeAndFlush(
						new DatagramPacket(Unpooled.copiedBuffer("QOTM?", CharsetUtil.UTF_8), new InetSocketAddress(
								"127.0.0.1", PORT))).sync();
			}
			long t2 = System.currentTimeMillis();
			System.out.println(t2-t1);
			

			// QuoteOfTheMomentClientHandler will close the DatagramChannel when a
			// response is received.  If the channel is not closed within 5 seconds,
			// print an error message and quit.
			if (!ch.closeFuture().await(5000)) {
				System.err.println("QOTM request timed out.");
			}
		} finally {
			group.shutdownGracefully();
		}
	}
}
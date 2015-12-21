package com.lumi.largedata.govsjava.netty;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

import java.net.InetSocketAddress;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class NettyHttpServer {

	private static Log log = LogFactory.getLog(NettyHttpServer.class);

	public void start(int port) throws Exception {
		NioEventLoopGroup group = new NioEventLoopGroup();
		try {
			ServerBootstrap b = new ServerBootstrap();
			b.group(group).channel(NioServerSocketChannel.class).localAddress(new InetSocketAddress(port))
					.childHandler(new ChannelInitializer<SocketChannel>() {
						@Override
						public void initChannel(SocketChannel ch) throws Exception {
							ch.pipeline().addLast(new HttpServerInboundHandler());
						}
					});

			ChannelFuture f = b.bind().sync();
			System.out
					.println(NettyHttpServer.class.getName() + " started and listen on " + f.channel().localAddress());
			f.channel().closeFuture().sync();
		} finally {
			group.shutdownGracefully().sync();
		}
	}

	public static void main(String[] args) throws Exception {
		NettyHttpServer server = new NettyHttpServer();
		log.info("Http Server listening on 8080 ...");
		server.start(8080);
	}
}

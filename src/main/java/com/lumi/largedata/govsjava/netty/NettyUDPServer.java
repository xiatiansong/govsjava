package com.lumi.largedata.govsjava.netty;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.DatagramPacket;
import io.netty.channel.socket.nio.NioDatagramChannel;
import io.netty.handler.codec.MessageToMessageDecoder;

import java.nio.charset.Charset;
import java.util.List;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/**
 * @author TinyZ on 2015/6/8.
 */
public class NettyUDPServer {

	private static Log log = LogFactory.getLog(NettyHttpServer.class);

	public void start(int port) throws Exception {
		final NioEventLoopGroup nioEventLoopGroup = new NioEventLoopGroup();
		try {
			Bootstrap bootstrap = new Bootstrap();
			bootstrap.channel(NioDatagramChannel.class);
			bootstrap.group(nioEventLoopGroup);
			bootstrap.handler(new ChannelInitializer<NioDatagramChannel>() {
				@Override
				public void channelActive(ChannelHandlerContext ctx) throws Exception {
					super.channelActive(ctx);
				}
				@Override
				protected void initChannel(NioDatagramChannel ch) throws Exception {
					ChannelPipeline cp = ch.pipeline();
					cp.addLast("framer", new MessageToMessageDecoder<DatagramPacket>() {
						@Override
						protected void decode(ChannelHandlerContext ctx, DatagramPacket msg, List<Object> out)
								throws Exception {
							out.add(msg.content().toString(Charset.forName("UTF-8")));
						}
					}).addLast("handler", new NettyUdpHandler());
				}
			});
			// 监听端口
			ChannelFuture sync = bootstrap.bind(port).sync();
			sync.channel().closeFuture().sync();
		} finally {
			nioEventLoopGroup.shutdownGracefully();
		}
	}
	
	public static void main(String[] args) throws Exception {
		NettyUDPServer server = new NettyUDPServer();
		log.info("Http Server listening on 8080 ...");
		server.start(8080);
	}
}
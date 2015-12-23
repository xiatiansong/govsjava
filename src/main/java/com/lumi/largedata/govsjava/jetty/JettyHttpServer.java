package com.lumi.largedata.govsjava.jetty;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.eclipse.jetty.server.Connector;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Request;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.AbstractHandler;
import org.eclipse.jetty.server.handler.ContextHandler;
import org.eclipse.jetty.server.nio.SelectChannelConnector;
import org.eclipse.jetty.util.thread.QueuedThreadPool;

public class JettyHttpServer {

	public static void main(String[] args) throws Exception {
		QueuedThreadPool threadPool = new QueuedThreadPool();
		threadPool.setMaxThreads(11);
		threadPool.setMinThreads(4);
		Server server = new Server();
		server.setThreadPool(threadPool);

		Connector connector = new SelectChannelConnector();
		connector.setPort(8081);
		server.setConnectors(new Connector[] { connector });

		ContextHandler context = new ContextHandler();
		context.setContextPath("/rwer");
		context.setResourceBase(".");
		context.setClassLoader(Thread.currentThread().getContextClassLoader());
		server.setHandler(context);

		Handler handler = new HelloHandler();
		context.setHandler(handler);

		server.start();
		server.join();
	}

	public static class HelloHandler extends AbstractHandler {
		@Override
		public void handle(String arg0, Request request1, HttpServletRequest request, HttpServletResponse response)
				throws IOException, ServletException {
			response.getWriter().println("<h1>Hellodfgdfgdfgfd OneContext</h1>");
		}
	}
}
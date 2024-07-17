package io.ncraft.armory.auth.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.63.0)",
    comments = "Source: armory/auth/v1/authing.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class AuthingGrpc {

  private AuthingGrpc() {}

  public static final java.lang.String SERVICE_NAME = "armory.auth.v1.Authing";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateAccountRequest,
      io.ncraft.armory.auth.Account> getCreateAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_account",
      requestType = io.ncraft.armory.auth.v1.CreateAccountRequest.class,
      responseType = io.ncraft.armory.auth.Account.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateAccountRequest,
      io.ncraft.armory.auth.Account> getCreateAccountMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateAccountRequest, io.ncraft.armory.auth.Account> getCreateAccountMethod;
    if ((getCreateAccountMethod = AuthingGrpc.getCreateAccountMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getCreateAccountMethod = AuthingGrpc.getCreateAccountMethod) == null) {
          AuthingGrpc.getCreateAccountMethod = getCreateAccountMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.CreateAccountRequest, io.ncraft.armory.auth.Account>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_account"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.CreateAccountRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.Account.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("create_account"))
              .build();
        }
      }
    }
    return getCreateAccountMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AuthingStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthingStub>() {
        @java.lang.Override
        public AuthingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthingStub(channel, callOptions);
        }
      };
    return AuthingStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AuthingBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthingBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthingBlockingStub>() {
        @java.lang.Override
        public AuthingBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthingBlockingStub(channel, callOptions);
        }
      };
    return AuthingBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AuthingFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthingFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthingFutureStub>() {
        @java.lang.Override
        public AuthingFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthingFutureStub(channel, callOptions);
        }
      };
    return AuthingFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createAccount(io.ncraft.armory.auth.v1.CreateAccountRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.Account> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAccountMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Authing.
   */
  public static abstract class AuthingImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AuthingGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Authing.
   */
  public static final class AuthingStub
      extends io.grpc.stub.AbstractAsyncStub<AuthingStub> {
    private AuthingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthingStub(channel, callOptions);
    }

    /**
     */
    public void createAccount(io.ncraft.armory.auth.v1.CreateAccountRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.Account> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Authing.
   */
  public static final class AuthingBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AuthingBlockingStub> {
    private AuthingBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthingBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthingBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.ncraft.armory.auth.Account createAccount(io.ncraft.armory.auth.v1.CreateAccountRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAccountMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Authing.
   */
  public static final class AuthingFutureStub
      extends io.grpc.stub.AbstractFutureStub<AuthingFutureStub> {
    private AuthingFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthingFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthingFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.Account> createAccount(
        io.ncraft.armory.auth.v1.CreateAccountRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAccountMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_ACCOUNT = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_ACCOUNT:
          serviceImpl.createAccount((io.ncraft.armory.auth.v1.CreateAccountRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.Account>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getCreateAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.CreateAccountRequest,
              io.ncraft.armory.auth.Account>(
                service, METHODID_CREATE_ACCOUNT)))
        .build();
  }

  private static abstract class AuthingBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AuthingBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.ncraft.armory.auth.v1.AuthingProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Authing");
    }
  }

  private static final class AuthingFileDescriptorSupplier
      extends AuthingBaseDescriptorSupplier {
    AuthingFileDescriptorSupplier() {}
  }

  private static final class AuthingMethodDescriptorSupplier
      extends AuthingBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AuthingMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AuthingGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AuthingFileDescriptorSupplier())
              .addMethod(getCreateAccountMethod())
              .build();
        }
      }
    }
    return result;
  }
}

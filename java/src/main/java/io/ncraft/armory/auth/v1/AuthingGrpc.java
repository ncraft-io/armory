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
  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateUserRequest,
      io.ncraft.armory.auth.User> getCreateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_user",
      requestType = io.ncraft.armory.auth.v1.CreateUserRequest.class,
      responseType = io.ncraft.armory.auth.User.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateUserRequest,
      io.ncraft.armory.auth.User> getCreateUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.CreateUserRequest, io.ncraft.armory.auth.User> getCreateUserMethod;
    if ((getCreateUserMethod = AuthingGrpc.getCreateUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getCreateUserMethod = AuthingGrpc.getCreateUserMethod) == null) {
          AuthingGrpc.getCreateUserMethod = getCreateUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.CreateUserRequest, io.ncraft.armory.auth.User>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.CreateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.User.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("create_user"))
              .build();
        }
      }
    }
    return getCreateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.BatchCreateUsersRequest,
      io.ncraft.armory.auth.v1.BatchCreateUsersResponse> getBatchCreateUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_users",
      requestType = io.ncraft.armory.auth.v1.BatchCreateUsersRequest.class,
      responseType = io.ncraft.armory.auth.v1.BatchCreateUsersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.BatchCreateUsersRequest,
      io.ncraft.armory.auth.v1.BatchCreateUsersResponse> getBatchCreateUsersMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.BatchCreateUsersRequest, io.ncraft.armory.auth.v1.BatchCreateUsersResponse> getBatchCreateUsersMethod;
    if ((getBatchCreateUsersMethod = AuthingGrpc.getBatchCreateUsersMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getBatchCreateUsersMethod = AuthingGrpc.getBatchCreateUsersMethod) == null) {
          AuthingGrpc.getBatchCreateUsersMethod = getBatchCreateUsersMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.BatchCreateUsersRequest, io.ncraft.armory.auth.v1.BatchCreateUsersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_users"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.BatchCreateUsersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.BatchCreateUsersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("batch_create_users"))
              .build();
        }
      }
    }
    return getBatchCreateUsersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdateUserRequest,
      org.mojolang.mojo.core.Null> getUpdateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_user",
      requestType = io.ncraft.armory.auth.v1.UpdateUserRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdateUserRequest,
      org.mojolang.mojo.core.Null> getUpdateUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdateUserRequest, org.mojolang.mojo.core.Null> getUpdateUserMethod;
    if ((getUpdateUserMethod = AuthingGrpc.getUpdateUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getUpdateUserMethod = AuthingGrpc.getUpdateUserMethod) == null) {
          AuthingGrpc.getUpdateUserMethod = getUpdateUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.UpdateUserRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.UpdateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("update_user"))
              .build();
        }
      }
    }
    return getUpdateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ActiveUserRequest,
      io.ncraft.armory.auth.LogonUser> getActiveUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "active_user",
      requestType = io.ncraft.armory.auth.v1.ActiveUserRequest.class,
      responseType = io.ncraft.armory.auth.LogonUser.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ActiveUserRequest,
      io.ncraft.armory.auth.LogonUser> getActiveUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ActiveUserRequest, io.ncraft.armory.auth.LogonUser> getActiveUserMethod;
    if ((getActiveUserMethod = AuthingGrpc.getActiveUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getActiveUserMethod = AuthingGrpc.getActiveUserMethod) == null) {
          AuthingGrpc.getActiveUserMethod = getActiveUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.ActiveUserRequest, io.ncraft.armory.auth.LogonUser>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "active_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.ActiveUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.LogonUser.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("active_user"))
              .build();
        }
      }
    }
    return getActiveUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.GetUserRequest,
      io.ncraft.armory.auth.User> getGetUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_user",
      requestType = io.ncraft.armory.auth.v1.GetUserRequest.class,
      responseType = io.ncraft.armory.auth.User.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.GetUserRequest,
      io.ncraft.armory.auth.User> getGetUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.GetUserRequest, io.ncraft.armory.auth.User> getGetUserMethod;
    if ((getGetUserMethod = AuthingGrpc.getGetUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getGetUserMethod = AuthingGrpc.getGetUserMethod) == null) {
          AuthingGrpc.getGetUserMethod = getGetUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.GetUserRequest, io.ncraft.armory.auth.User>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.GetUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.User.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("get_user"))
              .build();
        }
      }
    }
    return getGetUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ListUserRequest,
      io.ncraft.armory.auth.v1.ListUserResponse> getListUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "list_user",
      requestType = io.ncraft.armory.auth.v1.ListUserRequest.class,
      responseType = io.ncraft.armory.auth.v1.ListUserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ListUserRequest,
      io.ncraft.armory.auth.v1.ListUserResponse> getListUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.ListUserRequest, io.ncraft.armory.auth.v1.ListUserResponse> getListUserMethod;
    if ((getListUserMethod = AuthingGrpc.getListUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getListUserMethod = AuthingGrpc.getListUserMethod) == null) {
          AuthingGrpc.getListUserMethod = getListUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.ListUserRequest, io.ncraft.armory.auth.v1.ListUserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "list_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.ListUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.ListUserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("list_user"))
              .build();
        }
      }
    }
    return getListUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.DeleteUserRequest,
      org.mojolang.mojo.core.Null> getDeleteUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete_user",
      requestType = io.ncraft.armory.auth.v1.DeleteUserRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.DeleteUserRequest,
      org.mojolang.mojo.core.Null> getDeleteUserMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.DeleteUserRequest, org.mojolang.mojo.core.Null> getDeleteUserMethod;
    if ((getDeleteUserMethod = AuthingGrpc.getDeleteUserMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getDeleteUserMethod = AuthingGrpc.getDeleteUserMethod) == null) {
          AuthingGrpc.getDeleteUserMethod = getDeleteUserMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.DeleteUserRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete_user"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.DeleteUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("delete_user"))
              .build();
        }
      }
    }
    return getDeleteUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdatePasswordRequest,
      org.mojolang.mojo.core.Null> getUpdatePasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_password",
      requestType = io.ncraft.armory.auth.v1.UpdatePasswordRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdatePasswordRequest,
      org.mojolang.mojo.core.Null> getUpdatePasswordMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.UpdatePasswordRequest, org.mojolang.mojo.core.Null> getUpdatePasswordMethod;
    if ((getUpdatePasswordMethod = AuthingGrpc.getUpdatePasswordMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getUpdatePasswordMethod = AuthingGrpc.getUpdatePasswordMethod) == null) {
          AuthingGrpc.getUpdatePasswordMethod = getUpdatePasswordMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.UpdatePasswordRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_password"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.UpdatePasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("update_password"))
              .build();
        }
      }
    }
    return getUpdatePasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LoginRequest,
      io.ncraft.armory.auth.LogonUser> getLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "login",
      requestType = io.ncraft.armory.auth.v1.LoginRequest.class,
      responseType = io.ncraft.armory.auth.LogonUser.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LoginRequest,
      io.ncraft.armory.auth.LogonUser> getLoginMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LoginRequest, io.ncraft.armory.auth.LogonUser> getLoginMethod;
    if ((getLoginMethod = AuthingGrpc.getLoginMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getLoginMethod = AuthingGrpc.getLoginMethod) == null) {
          AuthingGrpc.getLoginMethod = getLoginMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.LoginRequest, io.ncraft.armory.auth.LogonUser>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "login"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.LoginRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.LogonUser.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("login"))
              .build();
        }
      }
    }
    return getLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LogoutRequest,
      org.mojolang.mojo.core.Null> getLogoutMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "logout",
      requestType = io.ncraft.armory.auth.v1.LogoutRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LogoutRequest,
      org.mojolang.mojo.core.Null> getLogoutMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.auth.v1.LogoutRequest, org.mojolang.mojo.core.Null> getLogoutMethod;
    if ((getLogoutMethod = AuthingGrpc.getLogoutMethod) == null) {
      synchronized (AuthingGrpc.class) {
        if ((getLogoutMethod = AuthingGrpc.getLogoutMethod) == null) {
          AuthingGrpc.getLogoutMethod = getLogoutMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.auth.v1.LogoutRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "logout"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.auth.v1.LogoutRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new AuthingMethodDescriptorSupplier("logout"))
              .build();
        }
      }
    }
    return getLogoutMethod;
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
    default void createUser(io.ncraft.armory.auth.v1.CreateUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateUserMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateUsers(io.ncraft.armory.auth.v1.BatchCreateUsersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.BatchCreateUsersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateUsersMethod(), responseObserver);
    }

    /**
     */
    default void updateUser(io.ncraft.armory.auth.v1.UpdateUserRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateUserMethod(), responseObserver);
    }

    /**
     */
    default void activeUser(io.ncraft.armory.auth.v1.ActiveUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getActiveUserMethod(), responseObserver);
    }

    /**
     */
    default void getUser(io.ncraft.armory.auth.v1.GetUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserMethod(), responseObserver);
    }

    /**
     */
    default void listUser(io.ncraft.armory.auth.v1.ListUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.ListUserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListUserMethod(), responseObserver);
    }

    /**
     */
    default void deleteUser(io.ncraft.armory.auth.v1.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteUserMethod(), responseObserver);
    }

    /**
     */
    default void updatePassword(io.ncraft.armory.auth.v1.UpdatePasswordRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePasswordMethod(), responseObserver);
    }

    /**
     */
    default void login(io.ncraft.armory.auth.v1.LoginRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLoginMethod(), responseObserver);
    }

    /**
     */
    default void logout(io.ncraft.armory.auth.v1.LogoutRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLogoutMethod(), responseObserver);
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
    public void createUser(io.ncraft.armory.auth.v1.CreateUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateUsers(io.ncraft.armory.auth.v1.BatchCreateUsersRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.BatchCreateUsersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateUsersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateUser(io.ncraft.armory.auth.v1.UpdateUserRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void activeUser(io.ncraft.armory.auth.v1.ActiveUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getActiveUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUser(io.ncraft.armory.auth.v1.GetUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listUser(io.ncraft.armory.auth.v1.ListUserRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.ListUserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteUser(io.ncraft.armory.auth.v1.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePassword(io.ncraft.armory.auth.v1.UpdatePasswordRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void login(io.ncraft.armory.auth.v1.LoginRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void logout(io.ncraft.armory.auth.v1.LogoutRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request, responseObserver);
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
    public io.ncraft.armory.auth.User createUser(io.ncraft.armory.auth.v1.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.auth.v1.BatchCreateUsersResponse batchCreateUsers(io.ncraft.armory.auth.v1.BatchCreateUsersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateUsersMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateUser(io.ncraft.armory.auth.v1.UpdateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.auth.LogonUser activeUser(io.ncraft.armory.auth.v1.ActiveUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getActiveUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.auth.User getUser(io.ncraft.armory.auth.v1.GetUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.auth.v1.ListUserResponse listUser(io.ncraft.armory.auth.v1.ListUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null deleteUser(io.ncraft.armory.auth.v1.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updatePassword(io.ncraft.armory.auth.v1.UpdatePasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.auth.LogonUser login(io.ncraft.armory.auth.v1.LoginRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLoginMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null logout(io.ncraft.armory.auth.v1.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLogoutMethod(), getCallOptions(), request);
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
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.User> createUser(
        io.ncraft.armory.auth.v1.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.v1.BatchCreateUsersResponse> batchCreateUsers(
        io.ncraft.armory.auth.v1.BatchCreateUsersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateUsersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateUser(
        io.ncraft.armory.auth.v1.UpdateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.LogonUser> activeUser(
        io.ncraft.armory.auth.v1.ActiveUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getActiveUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.User> getUser(
        io.ncraft.armory.auth.v1.GetUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.v1.ListUserResponse> listUser(
        io.ncraft.armory.auth.v1.ListUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> deleteUser(
        io.ncraft.armory.auth.v1.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updatePassword(
        io.ncraft.armory.auth.v1.UpdatePasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.auth.LogonUser> login(
        io.ncraft.armory.auth.v1.LoginRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> logout(
        io.ncraft.armory.auth.v1.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_USER = 0;
  private static final int METHODID_BATCH_CREATE_USERS = 1;
  private static final int METHODID_UPDATE_USER = 2;
  private static final int METHODID_ACTIVE_USER = 3;
  private static final int METHODID_GET_USER = 4;
  private static final int METHODID_LIST_USER = 5;
  private static final int METHODID_DELETE_USER = 6;
  private static final int METHODID_UPDATE_PASSWORD = 7;
  private static final int METHODID_LOGIN = 8;
  private static final int METHODID_LOGOUT = 9;

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
        case METHODID_CREATE_USER:
          serviceImpl.createUser((io.ncraft.armory.auth.v1.CreateUserRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_USERS:
          serviceImpl.batchCreateUsers((io.ncraft.armory.auth.v1.BatchCreateUsersRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.BatchCreateUsersResponse>) responseObserver);
          break;
        case METHODID_UPDATE_USER:
          serviceImpl.updateUser((io.ncraft.armory.auth.v1.UpdateUserRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_ACTIVE_USER:
          serviceImpl.activeUser((io.ncraft.armory.auth.v1.ActiveUserRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser>) responseObserver);
          break;
        case METHODID_GET_USER:
          serviceImpl.getUser((io.ncraft.armory.auth.v1.GetUserRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.User>) responseObserver);
          break;
        case METHODID_LIST_USER:
          serviceImpl.listUser((io.ncraft.armory.auth.v1.ListUserRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.v1.ListUserResponse>) responseObserver);
          break;
        case METHODID_DELETE_USER:
          serviceImpl.deleteUser((io.ncraft.armory.auth.v1.DeleteUserRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_UPDATE_PASSWORD:
          serviceImpl.updatePassword((io.ncraft.armory.auth.v1.UpdatePasswordRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_LOGIN:
          serviceImpl.login((io.ncraft.armory.auth.v1.LoginRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.auth.LogonUser>) responseObserver);
          break;
        case METHODID_LOGOUT:
          serviceImpl.logout((io.ncraft.armory.auth.v1.LogoutRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
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
          getCreateUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.CreateUserRequest,
              io.ncraft.armory.auth.User>(
                service, METHODID_CREATE_USER)))
        .addMethod(
          getBatchCreateUsersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.BatchCreateUsersRequest,
              io.ncraft.armory.auth.v1.BatchCreateUsersResponse>(
                service, METHODID_BATCH_CREATE_USERS)))
        .addMethod(
          getUpdateUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.UpdateUserRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_USER)))
        .addMethod(
          getActiveUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.ActiveUserRequest,
              io.ncraft.armory.auth.LogonUser>(
                service, METHODID_ACTIVE_USER)))
        .addMethod(
          getGetUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.GetUserRequest,
              io.ncraft.armory.auth.User>(
                service, METHODID_GET_USER)))
        .addMethod(
          getListUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.ListUserRequest,
              io.ncraft.armory.auth.v1.ListUserResponse>(
                service, METHODID_LIST_USER)))
        .addMethod(
          getDeleteUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.DeleteUserRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_DELETE_USER)))
        .addMethod(
          getUpdatePasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.UpdatePasswordRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_PASSWORD)))
        .addMethod(
          getLoginMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.LoginRequest,
              io.ncraft.armory.auth.LogonUser>(
                service, METHODID_LOGIN)))
        .addMethod(
          getLogoutMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.auth.v1.LogoutRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_LOGOUT)))
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
              .addMethod(getCreateUserMethod())
              .addMethod(getBatchCreateUsersMethod())
              .addMethod(getUpdateUserMethod())
              .addMethod(getActiveUserMethod())
              .addMethod(getGetUserMethod())
              .addMethod(getListUserMethod())
              .addMethod(getDeleteUserMethod())
              .addMethod(getUpdatePasswordMethod())
              .addMethod(getLoginMethod())
              .addMethod(getLogoutMethod())
              .build();
        }
      }
    }
    return result;
  }
}

package io.ncraft.armory.file.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.63.0)",
    comments = "Source: armory/file/v1/file.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class FileGrpc {

  private FileGrpc() {}

  public static final java.lang.String SERVICE_NAME = "armory.file.v1.File";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.GetFileRequest,
      io.ncraft.armory.file.BinaryFile> getGetFileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_file",
      requestType = io.ncraft.armory.file.v1.GetFileRequest.class,
      responseType = io.ncraft.armory.file.BinaryFile.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.GetFileRequest,
      io.ncraft.armory.file.BinaryFile> getGetFileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.GetFileRequest, io.ncraft.armory.file.BinaryFile> getGetFileMethod;
    if ((getGetFileMethod = FileGrpc.getGetFileMethod) == null) {
      synchronized (FileGrpc.class) {
        if ((getGetFileMethod = FileGrpc.getGetFileMethod) == null) {
          FileGrpc.getGetFileMethod = getGetFileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.file.v1.GetFileRequest, io.ncraft.armory.file.BinaryFile>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_file"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.v1.GetFileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.BinaryFile.getDefaultInstance()))
              .setSchemaDescriptor(new FileMethodDescriptorSupplier("get_file"))
              .build();
        }
      }
    }
    return getGetFileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.CreateFileRequest,
      io.ncraft.armory.file.BinaryFile> getCreateFileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_file",
      requestType = io.ncraft.armory.file.v1.CreateFileRequest.class,
      responseType = io.ncraft.armory.file.BinaryFile.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.CreateFileRequest,
      io.ncraft.armory.file.BinaryFile> getCreateFileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.CreateFileRequest, io.ncraft.armory.file.BinaryFile> getCreateFileMethod;
    if ((getCreateFileMethod = FileGrpc.getCreateFileMethod) == null) {
      synchronized (FileGrpc.class) {
        if ((getCreateFileMethod = FileGrpc.getCreateFileMethod) == null) {
          FileGrpc.getCreateFileMethod = getCreateFileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.file.v1.CreateFileRequest, io.ncraft.armory.file.BinaryFile>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_file"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.v1.CreateFileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.BinaryFile.getDefaultInstance()))
              .setSchemaDescriptor(new FileMethodDescriptorSupplier("create_file"))
              .build();
        }
      }
    }
    return getCreateFileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.BatchCreateFileRequest,
      io.ncraft.armory.file.v1.BatchCreateFileResponse> getBatchCreateFileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_file",
      requestType = io.ncraft.armory.file.v1.BatchCreateFileRequest.class,
      responseType = io.ncraft.armory.file.v1.BatchCreateFileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.BatchCreateFileRequest,
      io.ncraft.armory.file.v1.BatchCreateFileResponse> getBatchCreateFileMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.file.v1.BatchCreateFileRequest, io.ncraft.armory.file.v1.BatchCreateFileResponse> getBatchCreateFileMethod;
    if ((getBatchCreateFileMethod = FileGrpc.getBatchCreateFileMethod) == null) {
      synchronized (FileGrpc.class) {
        if ((getBatchCreateFileMethod = FileGrpc.getBatchCreateFileMethod) == null) {
          FileGrpc.getBatchCreateFileMethod = getBatchCreateFileMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.file.v1.BatchCreateFileRequest, io.ncraft.armory.file.v1.BatchCreateFileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_file"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.v1.BatchCreateFileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.file.v1.BatchCreateFileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new FileMethodDescriptorSupplier("batch_create_file"))
              .build();
        }
      }
    }
    return getBatchCreateFileMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FileStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileStub>() {
        @java.lang.Override
        public FileStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileStub(channel, callOptions);
        }
      };
    return FileStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FileBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileBlockingStub>() {
        @java.lang.Override
        public FileBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileBlockingStub(channel, callOptions);
        }
      };
    return FileBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FileFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileFutureStub>() {
        @java.lang.Override
        public FileFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileFutureStub(channel, callOptions);
        }
      };
    return FileFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void getFile(io.ncraft.armory.file.v1.GetFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetFileMethod(), responseObserver);
    }

    /**
     */
    default void createFile(io.ncraft.armory.file.v1.CreateFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateFileMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateFile(io.ncraft.armory.file.v1.BatchCreateFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.v1.BatchCreateFileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateFileMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service File.
   */
  public static abstract class FileImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return FileGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service File.
   */
  public static final class FileStub
      extends io.grpc.stub.AbstractAsyncStub<FileStub> {
    private FileStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileStub(channel, callOptions);
    }

    /**
     */
    public void getFile(io.ncraft.armory.file.v1.GetFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetFileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createFile(io.ncraft.armory.file.v1.CreateFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateFileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateFile(io.ncraft.armory.file.v1.BatchCreateFileRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.file.v1.BatchCreateFileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateFileMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service File.
   */
  public static final class FileBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<FileBlockingStub> {
    private FileBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.ncraft.armory.file.BinaryFile getFile(io.ncraft.armory.file.v1.GetFileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetFileMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.file.BinaryFile createFile(io.ncraft.armory.file.v1.CreateFileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateFileMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.file.v1.BatchCreateFileResponse batchCreateFile(io.ncraft.armory.file.v1.BatchCreateFileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateFileMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service File.
   */
  public static final class FileFutureStub
      extends io.grpc.stub.AbstractFutureStub<FileFutureStub> {
    private FileFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.file.BinaryFile> getFile(
        io.ncraft.armory.file.v1.GetFileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetFileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.file.BinaryFile> createFile(
        io.ncraft.armory.file.v1.CreateFileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateFileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.file.v1.BatchCreateFileResponse> batchCreateFile(
        io.ncraft.armory.file.v1.BatchCreateFileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateFileMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_FILE = 0;
  private static final int METHODID_CREATE_FILE = 1;
  private static final int METHODID_BATCH_CREATE_FILE = 2;

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
        case METHODID_GET_FILE:
          serviceImpl.getFile((io.ncraft.armory.file.v1.GetFileRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile>) responseObserver);
          break;
        case METHODID_CREATE_FILE:
          serviceImpl.createFile((io.ncraft.armory.file.v1.CreateFileRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.file.BinaryFile>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_FILE:
          serviceImpl.batchCreateFile((io.ncraft.armory.file.v1.BatchCreateFileRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.file.v1.BatchCreateFileResponse>) responseObserver);
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
          getGetFileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.file.v1.GetFileRequest,
              io.ncraft.armory.file.BinaryFile>(
                service, METHODID_GET_FILE)))
        .addMethod(
          getCreateFileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.file.v1.CreateFileRequest,
              io.ncraft.armory.file.BinaryFile>(
                service, METHODID_CREATE_FILE)))
        .addMethod(
          getBatchCreateFileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.file.v1.BatchCreateFileRequest,
              io.ncraft.armory.file.v1.BatchCreateFileResponse>(
                service, METHODID_BATCH_CREATE_FILE)))
        .build();
  }

  private static abstract class FileBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    FileBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.ncraft.armory.file.v1.FileProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("File");
    }
  }

  private static final class FileFileDescriptorSupplier
      extends FileBaseDescriptorSupplier {
    FileFileDescriptorSupplier() {}
  }

  private static final class FileMethodDescriptorSupplier
      extends FileBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    FileMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (FileGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FileFileDescriptorSupplier())
              .addMethod(getGetFileMethod())
              .addMethod(getCreateFileMethod())
              .addMethod(getBatchCreateFileMethod())
              .build();
        }
      }
    }
    return result;
  }
}

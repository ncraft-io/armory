package io.ncraft.armory.unitable.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.63.0)",
    comments = "Source: armory/unitable/v1/unitable.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UnitableGrpc {

  private UnitableGrpc() {}

  public static final java.lang.String SERVICE_NAME = "armory.unitable.v1.Unitable";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateTableRequest,
      io.ncraft.armory.unitable.Table> getCreateTableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_table",
      requestType = io.ncraft.armory.unitable.v1.CreateTableRequest.class,
      responseType = io.ncraft.armory.unitable.Table.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateTableRequest,
      io.ncraft.armory.unitable.Table> getCreateTableMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateTableRequest, io.ncraft.armory.unitable.Table> getCreateTableMethod;
    if ((getCreateTableMethod = UnitableGrpc.getCreateTableMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getCreateTableMethod = UnitableGrpc.getCreateTableMethod) == null) {
          UnitableGrpc.getCreateTableMethod = getCreateTableMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.CreateTableRequest, io.ncraft.armory.unitable.Table>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_table"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.CreateTableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.Table.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("create_table"))
              .build();
        }
      }
    }
    return getCreateTableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateTableRequest,
      org.mojolang.mojo.core.Null> getUpdateTableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_table",
      requestType = io.ncraft.armory.unitable.v1.UpdateTableRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateTableRequest,
      org.mojolang.mojo.core.Null> getUpdateTableMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateTableRequest, org.mojolang.mojo.core.Null> getUpdateTableMethod;
    if ((getUpdateTableMethod = UnitableGrpc.getUpdateTableMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getUpdateTableMethod = UnitableGrpc.getUpdateTableMethod) == null) {
          UnitableGrpc.getUpdateTableMethod = getUpdateTableMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.UpdateTableRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_table"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.UpdateTableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("update_table"))
              .build();
        }
      }
    }
    return getUpdateTableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetTableRequest,
      io.ncraft.armory.unitable.Table> getGetTableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_table",
      requestType = io.ncraft.armory.unitable.v1.GetTableRequest.class,
      responseType = io.ncraft.armory.unitable.Table.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetTableRequest,
      io.ncraft.armory.unitable.Table> getGetTableMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetTableRequest, io.ncraft.armory.unitable.Table> getGetTableMethod;
    if ((getGetTableMethod = UnitableGrpc.getGetTableMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getGetTableMethod = UnitableGrpc.getGetTableMethod) == null) {
          UnitableGrpc.getGetTableMethod = getGetTableMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.GetTableRequest, io.ncraft.armory.unitable.Table>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_table"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.GetTableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.Table.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("get_table"))
              .build();
        }
      }
    }
    return getGetTableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListTablesRequest,
      io.ncraft.armory.unitable.v1.ListTablesResponse> getListTablesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "list_tables",
      requestType = io.ncraft.armory.unitable.v1.ListTablesRequest.class,
      responseType = io.ncraft.armory.unitable.v1.ListTablesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListTablesRequest,
      io.ncraft.armory.unitable.v1.ListTablesResponse> getListTablesMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListTablesRequest, io.ncraft.armory.unitable.v1.ListTablesResponse> getListTablesMethod;
    if ((getListTablesMethod = UnitableGrpc.getListTablesMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getListTablesMethod = UnitableGrpc.getListTablesMethod) == null) {
          UnitableGrpc.getListTablesMethod = getListTablesMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.ListTablesRequest, io.ncraft.armory.unitable.v1.ListTablesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "list_tables"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListTablesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListTablesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("list_tables"))
              .build();
        }
      }
    }
    return getListTablesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteTableRequest,
      org.mojolang.mojo.core.Null> getDeleteTableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete_table",
      requestType = io.ncraft.armory.unitable.v1.DeleteTableRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteTableRequest,
      org.mojolang.mojo.core.Null> getDeleteTableMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteTableRequest, org.mojolang.mojo.core.Null> getDeleteTableMethod;
    if ((getDeleteTableMethod = UnitableGrpc.getDeleteTableMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getDeleteTableMethod = UnitableGrpc.getDeleteTableMethod) == null) {
          UnitableGrpc.getDeleteTableMethod = getDeleteTableMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.DeleteTableRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete_table"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.DeleteTableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("delete_table"))
              .build();
        }
      }
    }
    return getDeleteTableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.SyncTableRequest,
      io.ncraft.armory.unitable.Table> getSyncTableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "sync_table",
      requestType = io.ncraft.armory.unitable.v1.SyncTableRequest.class,
      responseType = io.ncraft.armory.unitable.Table.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.SyncTableRequest,
      io.ncraft.armory.unitable.Table> getSyncTableMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.SyncTableRequest, io.ncraft.armory.unitable.Table> getSyncTableMethod;
    if ((getSyncTableMethod = UnitableGrpc.getSyncTableMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getSyncTableMethod = UnitableGrpc.getSyncTableMethod) == null) {
          UnitableGrpc.getSyncTableMethod = getSyncTableMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.SyncTableRequest, io.ncraft.armory.unitable.Table>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "sync_table"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.SyncTableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.Table.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("sync_table"))
              .build();
        }
      }
    }
    return getSyncTableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateColumnRequest,
      io.ncraft.armory.unitable.Column> getCreateColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_column",
      requestType = io.ncraft.armory.unitable.v1.CreateColumnRequest.class,
      responseType = io.ncraft.armory.unitable.Column.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateColumnRequest,
      io.ncraft.armory.unitable.Column> getCreateColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateColumnRequest, io.ncraft.armory.unitable.Column> getCreateColumnMethod;
    if ((getCreateColumnMethod = UnitableGrpc.getCreateColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getCreateColumnMethod = UnitableGrpc.getCreateColumnMethod) == null) {
          UnitableGrpc.getCreateColumnMethod = getCreateColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.CreateColumnRequest, io.ncraft.armory.unitable.Column>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.CreateColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.Column.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("create_column"))
              .build();
        }
      }
    }
    return getCreateColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateColumnRequest,
      org.mojolang.mojo.core.Null> getUpdateColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_column",
      requestType = io.ncraft.armory.unitable.v1.UpdateColumnRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateColumnRequest,
      org.mojolang.mojo.core.Null> getUpdateColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateColumnRequest, org.mojolang.mojo.core.Null> getUpdateColumnMethod;
    if ((getUpdateColumnMethod = UnitableGrpc.getUpdateColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getUpdateColumnMethod = UnitableGrpc.getUpdateColumnMethod) == null) {
          UnitableGrpc.getUpdateColumnMethod = getUpdateColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.UpdateColumnRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.UpdateColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("update_column"))
              .build();
        }
      }
    }
    return getUpdateColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetColumnRequest,
      io.ncraft.armory.unitable.Column> getGetColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_column",
      requestType = io.ncraft.armory.unitable.v1.GetColumnRequest.class,
      responseType = io.ncraft.armory.unitable.Column.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetColumnRequest,
      io.ncraft.armory.unitable.Column> getGetColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetColumnRequest, io.ncraft.armory.unitable.Column> getGetColumnMethod;
    if ((getGetColumnMethod = UnitableGrpc.getGetColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getGetColumnMethod = UnitableGrpc.getGetColumnMethod) == null) {
          UnitableGrpc.getGetColumnMethod = getGetColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.GetColumnRequest, io.ncraft.armory.unitable.Column>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.GetColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.Column.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("get_column"))
              .build();
        }
      }
    }
    return getGetColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteColumnRequest,
      org.mojolang.mojo.core.Null> getDeleteColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete_column",
      requestType = io.ncraft.armory.unitable.v1.DeleteColumnRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteColumnRequest,
      org.mojolang.mojo.core.Null> getDeleteColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteColumnRequest, org.mojolang.mojo.core.Null> getDeleteColumnMethod;
    if ((getDeleteColumnMethod = UnitableGrpc.getDeleteColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getDeleteColumnMethod = UnitableGrpc.getDeleteColumnMethod) == null) {
          UnitableGrpc.getDeleteColumnMethod = getDeleteColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.DeleteColumnRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.DeleteColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("delete_column"))
              .build();
        }
      }
    }
    return getDeleteColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListColumnsRequest,
      io.ncraft.armory.unitable.v1.ListColumnsResponse> getListColumnsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "list_columns",
      requestType = io.ncraft.armory.unitable.v1.ListColumnsRequest.class,
      responseType = io.ncraft.armory.unitable.v1.ListColumnsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListColumnsRequest,
      io.ncraft.armory.unitable.v1.ListColumnsResponse> getListColumnsMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListColumnsRequest, io.ncraft.armory.unitable.v1.ListColumnsResponse> getListColumnsMethod;
    if ((getListColumnsMethod = UnitableGrpc.getListColumnsMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getListColumnsMethod = UnitableGrpc.getListColumnsMethod) == null) {
          UnitableGrpc.getListColumnsMethod = getListColumnsMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.ListColumnsRequest, io.ncraft.armory.unitable.v1.ListColumnsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "list_columns"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListColumnsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListColumnsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("list_columns"))
              .build();
        }
      }
    }
    return getListColumnsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest,
      org.mojolang.mojo.core.Null> getBatchCreateColumnsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_columns",
      requestType = io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest,
      org.mojolang.mojo.core.Null> getBatchCreateColumnsMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest, org.mojolang.mojo.core.Null> getBatchCreateColumnsMethod;
    if ((getBatchCreateColumnsMethod = UnitableGrpc.getBatchCreateColumnsMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchCreateColumnsMethod = UnitableGrpc.getBatchCreateColumnsMethod) == null) {
          UnitableGrpc.getBatchCreateColumnsMethod = getBatchCreateColumnsMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_columns"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_create_columns"))
              .build();
        }
      }
    }
    return getBatchCreateColumnsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_update_column",
      requestType = io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest, org.mojolang.mojo.core.Null> getBatchUpdateColumnMethod;
    if ((getBatchUpdateColumnMethod = UnitableGrpc.getBatchUpdateColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchUpdateColumnMethod = UnitableGrpc.getBatchUpdateColumnMethod) == null) {
          UnitableGrpc.getBatchUpdateColumnMethod = getBatchUpdateColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_update_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_update_column"))
              .build();
        }
      }
    }
    return getBatchUpdateColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest,
      org.mojolang.mojo.core.Null> getBatchDeleteColumnMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_delete_column",
      requestType = io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest,
      org.mojolang.mojo.core.Null> getBatchDeleteColumnMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest, org.mojolang.mojo.core.Null> getBatchDeleteColumnMethod;
    if ((getBatchDeleteColumnMethod = UnitableGrpc.getBatchDeleteColumnMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchDeleteColumnMethod = UnitableGrpc.getBatchDeleteColumnMethod) == null) {
          UnitableGrpc.getBatchDeleteColumnMethod = getBatchDeleteColumnMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_delete_column"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_delete_column"))
              .build();
        }
      }
    }
    return getBatchDeleteColumnMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateRowRequest,
      org.mojolang.mojo.core.Object> getCreateRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "create_row",
      requestType = io.ncraft.armory.unitable.v1.CreateRowRequest.class,
      responseType = org.mojolang.mojo.core.Object.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateRowRequest,
      org.mojolang.mojo.core.Object> getCreateRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.CreateRowRequest, org.mojolang.mojo.core.Object> getCreateRowMethod;
    if ((getCreateRowMethod = UnitableGrpc.getCreateRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getCreateRowMethod = UnitableGrpc.getCreateRowMethod) == null) {
          UnitableGrpc.getCreateRowMethod = getCreateRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.CreateRowRequest, org.mojolang.mojo.core.Object>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "create_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.CreateRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Object.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("create_row"))
              .build();
        }
      }
    }
    return getCreateRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateRowRequest,
      org.mojolang.mojo.core.Null> getUpdateRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "update_row",
      requestType = io.ncraft.armory.unitable.v1.UpdateRowRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateRowRequest,
      org.mojolang.mojo.core.Null> getUpdateRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.UpdateRowRequest, org.mojolang.mojo.core.Null> getUpdateRowMethod;
    if ((getUpdateRowMethod = UnitableGrpc.getUpdateRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getUpdateRowMethod = UnitableGrpc.getUpdateRowMethod) == null) {
          UnitableGrpc.getUpdateRowMethod = getUpdateRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.UpdateRowRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "update_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.UpdateRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("update_row"))
              .build();
        }
      }
    }
    return getUpdateRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetRowRequest,
      org.mojolang.mojo.core.Object> getGetRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "get_row",
      requestType = io.ncraft.armory.unitable.v1.GetRowRequest.class,
      responseType = org.mojolang.mojo.core.Object.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetRowRequest,
      org.mojolang.mojo.core.Object> getGetRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.GetRowRequest, org.mojolang.mojo.core.Object> getGetRowMethod;
    if ((getGetRowMethod = UnitableGrpc.getGetRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getGetRowMethod = UnitableGrpc.getGetRowMethod) == null) {
          UnitableGrpc.getGetRowMethod = getGetRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.GetRowRequest, org.mojolang.mojo.core.Object>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "get_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.GetRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Object.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("get_row"))
              .build();
        }
      }
    }
    return getGetRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteRowRequest,
      org.mojolang.mojo.core.Null> getDeleteRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "delete_row",
      requestType = io.ncraft.armory.unitable.v1.DeleteRowRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteRowRequest,
      org.mojolang.mojo.core.Null> getDeleteRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.DeleteRowRequest, org.mojolang.mojo.core.Null> getDeleteRowMethod;
    if ((getDeleteRowMethod = UnitableGrpc.getDeleteRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getDeleteRowMethod = UnitableGrpc.getDeleteRowMethod) == null) {
          UnitableGrpc.getDeleteRowMethod = getDeleteRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.DeleteRowRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "delete_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.DeleteRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("delete_row"))
              .build();
        }
      }
    }
    return getDeleteRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListRowRequest,
      io.ncraft.armory.unitable.v1.ListRowResponse> getListRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "list_row",
      requestType = io.ncraft.armory.unitable.v1.ListRowRequest.class,
      responseType = io.ncraft.armory.unitable.v1.ListRowResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListRowRequest,
      io.ncraft.armory.unitable.v1.ListRowResponse> getListRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ListRowRequest, io.ncraft.armory.unitable.v1.ListRowResponse> getListRowMethod;
    if ((getListRowMethod = UnitableGrpc.getListRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getListRowMethod = UnitableGrpc.getListRowMethod) == null) {
          UnitableGrpc.getListRowMethod = getListRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.ListRowRequest, io.ncraft.armory.unitable.v1.ListRowResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "list_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ListRowResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("list_row"))
              .build();
        }
      }
    }
    return getListRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ExportRowRequest,
      io.ncraft.armory.unitable.v1.ExportRowResponse> getExportRowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "export_row",
      requestType = io.ncraft.armory.unitable.v1.ExportRowRequest.class,
      responseType = io.ncraft.armory.unitable.v1.ExportRowResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ExportRowRequest,
      io.ncraft.armory.unitable.v1.ExportRowResponse> getExportRowMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.ExportRowRequest, io.ncraft.armory.unitable.v1.ExportRowResponse> getExportRowMethod;
    if ((getExportRowMethod = UnitableGrpc.getExportRowMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getExportRowMethod = UnitableGrpc.getExportRowMethod) == null) {
          UnitableGrpc.getExportRowMethod = getExportRowMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.ExportRowRequest, io.ncraft.armory.unitable.v1.ExportRowResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "export_row"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ExportRowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.ExportRowResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("export_row"))
              .build();
        }
      }
    }
    return getExportRowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateRowsRequest,
      org.mojolang.mojo.core.Null> getBatchCreateRowsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_create_rows",
      requestType = io.ncraft.armory.unitable.v1.BatchCreateRowsRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateRowsRequest,
      org.mojolang.mojo.core.Null> getBatchCreateRowsMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchCreateRowsRequest, org.mojolang.mojo.core.Null> getBatchCreateRowsMethod;
    if ((getBatchCreateRowsMethod = UnitableGrpc.getBatchCreateRowsMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchCreateRowsMethod = UnitableGrpc.getBatchCreateRowsMethod) == null) {
          UnitableGrpc.getBatchCreateRowsMethod = getBatchCreateRowsMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchCreateRowsRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_create_rows"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchCreateRowsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_create_rows"))
              .build();
        }
      }
    }
    return getBatchCreateRowsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateRowsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_update_rows",
      requestType = io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest,
      org.mojolang.mojo.core.Null> getBatchUpdateRowsMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest, org.mojolang.mojo.core.Null> getBatchUpdateRowsMethod;
    if ((getBatchUpdateRowsMethod = UnitableGrpc.getBatchUpdateRowsMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchUpdateRowsMethod = UnitableGrpc.getBatchUpdateRowsMethod) == null) {
          UnitableGrpc.getBatchUpdateRowsMethod = getBatchUpdateRowsMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_update_rows"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_update_rows"))
              .build();
        }
      }
    }
    return getBatchUpdateRowsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest,
      org.mojolang.mojo.core.Null> getBatchDeleteRowsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "batch_delete_rows",
      requestType = io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest.class,
      responseType = org.mojolang.mojo.core.Null.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest,
      org.mojolang.mojo.core.Null> getBatchDeleteRowsMethod() {
    io.grpc.MethodDescriptor<io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest, org.mojolang.mojo.core.Null> getBatchDeleteRowsMethod;
    if ((getBatchDeleteRowsMethod = UnitableGrpc.getBatchDeleteRowsMethod) == null) {
      synchronized (UnitableGrpc.class) {
        if ((getBatchDeleteRowsMethod = UnitableGrpc.getBatchDeleteRowsMethod) == null) {
          UnitableGrpc.getBatchDeleteRowsMethod = getBatchDeleteRowsMethod =
              io.grpc.MethodDescriptor.<io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest, org.mojolang.mojo.core.Null>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "batch_delete_rows"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  org.mojolang.mojo.core.Null.getDefaultInstance()))
              .setSchemaDescriptor(new UnitableMethodDescriptorSupplier("batch_delete_rows"))
              .build();
        }
      }
    }
    return getBatchDeleteRowsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UnitableStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UnitableStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UnitableStub>() {
        @java.lang.Override
        public UnitableStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UnitableStub(channel, callOptions);
        }
      };
    return UnitableStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UnitableBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UnitableBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UnitableBlockingStub>() {
        @java.lang.Override
        public UnitableBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UnitableBlockingStub(channel, callOptions);
        }
      };
    return UnitableBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UnitableFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UnitableFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UnitableFutureStub>() {
        @java.lang.Override
        public UnitableFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UnitableFutureStub(channel, callOptions);
        }
      };
    return UnitableFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createTable(io.ncraft.armory.unitable.v1.CreateTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTableMethod(), responseObserver);
    }

    /**
     */
    default void updateTable(io.ncraft.armory.unitable.v1.UpdateTableRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateTableMethod(), responseObserver);
    }

    /**
     */
    default void getTable(io.ncraft.armory.unitable.v1.GetTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTableMethod(), responseObserver);
    }

    /**
     */
    default void listTables(io.ncraft.armory.unitable.v1.ListTablesRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListTablesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListTablesMethod(), responseObserver);
    }

    /**
     */
    default void deleteTable(io.ncraft.armory.unitable.v1.DeleteTableRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteTableMethod(), responseObserver);
    }

    /**
     */
    default void syncTable(io.ncraft.armory.unitable.v1.SyncTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSyncTableMethod(), responseObserver);
    }

    /**
     */
    default void createColumn(io.ncraft.armory.unitable.v1.CreateColumnRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateColumnMethod(), responseObserver);
    }

    /**
     */
    default void updateColumn(io.ncraft.armory.unitable.v1.UpdateColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateColumnMethod(), responseObserver);
    }

    /**
     */
    default void getColumn(io.ncraft.armory.unitable.v1.GetColumnRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetColumnMethod(), responseObserver);
    }

    /**
     */
    default void deleteColumn(io.ncraft.armory.unitable.v1.DeleteColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteColumnMethod(), responseObserver);
    }

    /**
     */
    default void listColumns(io.ncraft.armory.unitable.v1.ListColumnsRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListColumnsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListColumnsMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateColumns(io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateColumnsMethod(), responseObserver);
    }

    /**
     */
    default void batchUpdateColumn(io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchUpdateColumnMethod(), responseObserver);
    }

    /**
     */
    default void batchDeleteColumn(io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchDeleteColumnMethod(), responseObserver);
    }

    /**
     */
    default void createRow(io.ncraft.armory.unitable.v1.CreateRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateRowMethod(), responseObserver);
    }

    /**
     */
    default void updateRow(io.ncraft.armory.unitable.v1.UpdateRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateRowMethod(), responseObserver);
    }

    /**
     */
    default void getRow(io.ncraft.armory.unitable.v1.GetRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetRowMethod(), responseObserver);
    }

    /**
     */
    default void deleteRow(io.ncraft.armory.unitable.v1.DeleteRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteRowMethod(), responseObserver);
    }

    /**
     */
    default void listRow(io.ncraft.armory.unitable.v1.ListRowRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListRowResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListRowMethod(), responseObserver);
    }

    /**
     */
    default void exportRow(io.ncraft.armory.unitable.v1.ExportRowRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ExportRowResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getExportRowMethod(), responseObserver);
    }

    /**
     */
    default void batchCreateRows(io.ncraft.armory.unitable.v1.BatchCreateRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchCreateRowsMethod(), responseObserver);
    }

    /**
     */
    default void batchUpdateRows(io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchUpdateRowsMethod(), responseObserver);
    }

    /**
     */
    default void batchDeleteRows(io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBatchDeleteRowsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Unitable.
   */
  public static abstract class UnitableImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return UnitableGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Unitable.
   */
  public static final class UnitableStub
      extends io.grpc.stub.AbstractAsyncStub<UnitableStub> {
    private UnitableStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UnitableStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UnitableStub(channel, callOptions);
    }

    /**
     */
    public void createTable(io.ncraft.armory.unitable.v1.CreateTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateTableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateTable(io.ncraft.armory.unitable.v1.UpdateTableRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateTableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTable(io.ncraft.armory.unitable.v1.GetTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listTables(io.ncraft.armory.unitable.v1.ListTablesRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListTablesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListTablesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteTable(io.ncraft.armory.unitable.v1.DeleteTableRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteTableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void syncTable(io.ncraft.armory.unitable.v1.SyncTableRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSyncTableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createColumn(io.ncraft.armory.unitable.v1.CreateColumnRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateColumn(io.ncraft.armory.unitable.v1.UpdateColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getColumn(io.ncraft.armory.unitable.v1.GetColumnRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteColumn(io.ncraft.armory.unitable.v1.DeleteColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listColumns(io.ncraft.armory.unitable.v1.ListColumnsRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListColumnsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListColumnsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateColumns(io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateColumnsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchUpdateColumn(io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchUpdateColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchDeleteColumn(io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchDeleteColumnMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createRow(io.ncraft.armory.unitable.v1.CreateRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateRow(io.ncraft.armory.unitable.v1.UpdateRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getRow(io.ncraft.armory.unitable.v1.GetRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteRow(io.ncraft.armory.unitable.v1.DeleteRowRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listRow(io.ncraft.armory.unitable.v1.ListRowRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListRowResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void exportRow(io.ncraft.armory.unitable.v1.ExportRowRequest request,
        io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ExportRowResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getExportRowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchCreateRows(io.ncraft.armory.unitable.v1.BatchCreateRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchCreateRowsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchUpdateRows(io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchUpdateRowsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void batchDeleteRows(io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest request,
        io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBatchDeleteRowsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Unitable.
   */
  public static final class UnitableBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<UnitableBlockingStub> {
    private UnitableBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UnitableBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UnitableBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.ncraft.armory.unitable.Table createTable(io.ncraft.armory.unitable.v1.CreateTableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateTableMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateTable(io.ncraft.armory.unitable.v1.UpdateTableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateTableMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.Table getTable(io.ncraft.armory.unitable.v1.GetTableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTableMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.v1.ListTablesResponse listTables(io.ncraft.armory.unitable.v1.ListTablesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListTablesMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null deleteTable(io.ncraft.armory.unitable.v1.DeleteTableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteTableMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.Table syncTable(io.ncraft.armory.unitable.v1.SyncTableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSyncTableMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.Column createColumn(io.ncraft.armory.unitable.v1.CreateColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateColumn(io.ncraft.armory.unitable.v1.UpdateColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.Column getColumn(io.ncraft.armory.unitable.v1.GetColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null deleteColumn(io.ncraft.armory.unitable.v1.DeleteColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.v1.ListColumnsResponse listColumns(io.ncraft.armory.unitable.v1.ListColumnsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListColumnsMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchCreateColumns(io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateColumnsMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchUpdateColumn(io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchUpdateColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchDeleteColumn(io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchDeleteColumnMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Object createRow(io.ncraft.armory.unitable.v1.CreateRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null updateRow(io.ncraft.armory.unitable.v1.UpdateRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Object getRow(io.ncraft.armory.unitable.v1.GetRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null deleteRow(io.ncraft.armory.unitable.v1.DeleteRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.v1.ListRowResponse listRow(io.ncraft.armory.unitable.v1.ListRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.ncraft.armory.unitable.v1.ExportRowResponse exportRow(io.ncraft.armory.unitable.v1.ExportRowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getExportRowMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchCreateRows(io.ncraft.armory.unitable.v1.BatchCreateRowsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchCreateRowsMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchUpdateRows(io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchUpdateRowsMethod(), getCallOptions(), request);
    }

    /**
     */
    public org.mojolang.mojo.core.Null batchDeleteRows(io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBatchDeleteRowsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Unitable.
   */
  public static final class UnitableFutureStub
      extends io.grpc.stub.AbstractFutureStub<UnitableFutureStub> {
    private UnitableFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UnitableFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UnitableFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.Table> createTable(
        io.ncraft.armory.unitable.v1.CreateTableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateTableMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateTable(
        io.ncraft.armory.unitable.v1.UpdateTableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateTableMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.Table> getTable(
        io.ncraft.armory.unitable.v1.GetTableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTableMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.v1.ListTablesResponse> listTables(
        io.ncraft.armory.unitable.v1.ListTablesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListTablesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> deleteTable(
        io.ncraft.armory.unitable.v1.DeleteTableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteTableMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.Table> syncTable(
        io.ncraft.armory.unitable.v1.SyncTableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSyncTableMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.Column> createColumn(
        io.ncraft.armory.unitable.v1.CreateColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateColumn(
        io.ncraft.armory.unitable.v1.UpdateColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.Column> getColumn(
        io.ncraft.armory.unitable.v1.GetColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> deleteColumn(
        io.ncraft.armory.unitable.v1.DeleteColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.v1.ListColumnsResponse> listColumns(
        io.ncraft.armory.unitable.v1.ListColumnsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListColumnsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchCreateColumns(
        io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateColumnsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchUpdateColumn(
        io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchUpdateColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchDeleteColumn(
        io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchDeleteColumnMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Object> createRow(
        io.ncraft.armory.unitable.v1.CreateRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> updateRow(
        io.ncraft.armory.unitable.v1.UpdateRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Object> getRow(
        io.ncraft.armory.unitable.v1.GetRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> deleteRow(
        io.ncraft.armory.unitable.v1.DeleteRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.v1.ListRowResponse> listRow(
        io.ncraft.armory.unitable.v1.ListRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.ncraft.armory.unitable.v1.ExportRowResponse> exportRow(
        io.ncraft.armory.unitable.v1.ExportRowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getExportRowMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchCreateRows(
        io.ncraft.armory.unitable.v1.BatchCreateRowsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchCreateRowsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchUpdateRows(
        io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchUpdateRowsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<org.mojolang.mojo.core.Null> batchDeleteRows(
        io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBatchDeleteRowsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_TABLE = 0;
  private static final int METHODID_UPDATE_TABLE = 1;
  private static final int METHODID_GET_TABLE = 2;
  private static final int METHODID_LIST_TABLES = 3;
  private static final int METHODID_DELETE_TABLE = 4;
  private static final int METHODID_SYNC_TABLE = 5;
  private static final int METHODID_CREATE_COLUMN = 6;
  private static final int METHODID_UPDATE_COLUMN = 7;
  private static final int METHODID_GET_COLUMN = 8;
  private static final int METHODID_DELETE_COLUMN = 9;
  private static final int METHODID_LIST_COLUMNS = 10;
  private static final int METHODID_BATCH_CREATE_COLUMNS = 11;
  private static final int METHODID_BATCH_UPDATE_COLUMN = 12;
  private static final int METHODID_BATCH_DELETE_COLUMN = 13;
  private static final int METHODID_CREATE_ROW = 14;
  private static final int METHODID_UPDATE_ROW = 15;
  private static final int METHODID_GET_ROW = 16;
  private static final int METHODID_DELETE_ROW = 17;
  private static final int METHODID_LIST_ROW = 18;
  private static final int METHODID_EXPORT_ROW = 19;
  private static final int METHODID_BATCH_CREATE_ROWS = 20;
  private static final int METHODID_BATCH_UPDATE_ROWS = 21;
  private static final int METHODID_BATCH_DELETE_ROWS = 22;

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
        case METHODID_CREATE_TABLE:
          serviceImpl.createTable((io.ncraft.armory.unitable.v1.CreateTableRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table>) responseObserver);
          break;
        case METHODID_UPDATE_TABLE:
          serviceImpl.updateTable((io.ncraft.armory.unitable.v1.UpdateTableRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_GET_TABLE:
          serviceImpl.getTable((io.ncraft.armory.unitable.v1.GetTableRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table>) responseObserver);
          break;
        case METHODID_LIST_TABLES:
          serviceImpl.listTables((io.ncraft.armory.unitable.v1.ListTablesRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListTablesResponse>) responseObserver);
          break;
        case METHODID_DELETE_TABLE:
          serviceImpl.deleteTable((io.ncraft.armory.unitable.v1.DeleteTableRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_SYNC_TABLE:
          serviceImpl.syncTable((io.ncraft.armory.unitable.v1.SyncTableRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Table>) responseObserver);
          break;
        case METHODID_CREATE_COLUMN:
          serviceImpl.createColumn((io.ncraft.armory.unitable.v1.CreateColumnRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column>) responseObserver);
          break;
        case METHODID_UPDATE_COLUMN:
          serviceImpl.updateColumn((io.ncraft.armory.unitable.v1.UpdateColumnRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_GET_COLUMN:
          serviceImpl.getColumn((io.ncraft.armory.unitable.v1.GetColumnRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.Column>) responseObserver);
          break;
        case METHODID_DELETE_COLUMN:
          serviceImpl.deleteColumn((io.ncraft.armory.unitable.v1.DeleteColumnRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_LIST_COLUMNS:
          serviceImpl.listColumns((io.ncraft.armory.unitable.v1.ListColumnsRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListColumnsResponse>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_COLUMNS:
          serviceImpl.batchCreateColumns((io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_BATCH_UPDATE_COLUMN:
          serviceImpl.batchUpdateColumn((io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_BATCH_DELETE_COLUMN:
          serviceImpl.batchDeleteColumn((io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_CREATE_ROW:
          serviceImpl.createRow((io.ncraft.armory.unitable.v1.CreateRowRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object>) responseObserver);
          break;
        case METHODID_UPDATE_ROW:
          serviceImpl.updateRow((io.ncraft.armory.unitable.v1.UpdateRowRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_GET_ROW:
          serviceImpl.getRow((io.ncraft.armory.unitable.v1.GetRowRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Object>) responseObserver);
          break;
        case METHODID_DELETE_ROW:
          serviceImpl.deleteRow((io.ncraft.armory.unitable.v1.DeleteRowRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_LIST_ROW:
          serviceImpl.listRow((io.ncraft.armory.unitable.v1.ListRowRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ListRowResponse>) responseObserver);
          break;
        case METHODID_EXPORT_ROW:
          serviceImpl.exportRow((io.ncraft.armory.unitable.v1.ExportRowRequest) request,
              (io.grpc.stub.StreamObserver<io.ncraft.armory.unitable.v1.ExportRowResponse>) responseObserver);
          break;
        case METHODID_BATCH_CREATE_ROWS:
          serviceImpl.batchCreateRows((io.ncraft.armory.unitable.v1.BatchCreateRowsRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_BATCH_UPDATE_ROWS:
          serviceImpl.batchUpdateRows((io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest) request,
              (io.grpc.stub.StreamObserver<org.mojolang.mojo.core.Null>) responseObserver);
          break;
        case METHODID_BATCH_DELETE_ROWS:
          serviceImpl.batchDeleteRows((io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest) request,
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
          getCreateTableMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.CreateTableRequest,
              io.ncraft.armory.unitable.Table>(
                service, METHODID_CREATE_TABLE)))
        .addMethod(
          getUpdateTableMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.UpdateTableRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_TABLE)))
        .addMethod(
          getGetTableMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.GetTableRequest,
              io.ncraft.armory.unitable.Table>(
                service, METHODID_GET_TABLE)))
        .addMethod(
          getListTablesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.ListTablesRequest,
              io.ncraft.armory.unitable.v1.ListTablesResponse>(
                service, METHODID_LIST_TABLES)))
        .addMethod(
          getDeleteTableMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.DeleteTableRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_DELETE_TABLE)))
        .addMethod(
          getSyncTableMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.SyncTableRequest,
              io.ncraft.armory.unitable.Table>(
                service, METHODID_SYNC_TABLE)))
        .addMethod(
          getCreateColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.CreateColumnRequest,
              io.ncraft.armory.unitable.Column>(
                service, METHODID_CREATE_COLUMN)))
        .addMethod(
          getUpdateColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.UpdateColumnRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_COLUMN)))
        .addMethod(
          getGetColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.GetColumnRequest,
              io.ncraft.armory.unitable.Column>(
                service, METHODID_GET_COLUMN)))
        .addMethod(
          getDeleteColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.DeleteColumnRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_DELETE_COLUMN)))
        .addMethod(
          getListColumnsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.ListColumnsRequest,
              io.ncraft.armory.unitable.v1.ListColumnsResponse>(
                service, METHODID_LIST_COLUMNS)))
        .addMethod(
          getBatchCreateColumnsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchCreateColumnsRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_CREATE_COLUMNS)))
        .addMethod(
          getBatchUpdateColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_UPDATE_COLUMN)))
        .addMethod(
          getBatchDeleteColumnMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchDeleteColumnRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_DELETE_COLUMN)))
        .addMethod(
          getCreateRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.CreateRowRequest,
              org.mojolang.mojo.core.Object>(
                service, METHODID_CREATE_ROW)))
        .addMethod(
          getUpdateRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.UpdateRowRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_UPDATE_ROW)))
        .addMethod(
          getGetRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.GetRowRequest,
              org.mojolang.mojo.core.Object>(
                service, METHODID_GET_ROW)))
        .addMethod(
          getDeleteRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.DeleteRowRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_DELETE_ROW)))
        .addMethod(
          getListRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.ListRowRequest,
              io.ncraft.armory.unitable.v1.ListRowResponse>(
                service, METHODID_LIST_ROW)))
        .addMethod(
          getExportRowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.ExportRowRequest,
              io.ncraft.armory.unitable.v1.ExportRowResponse>(
                service, METHODID_EXPORT_ROW)))
        .addMethod(
          getBatchCreateRowsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchCreateRowsRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_CREATE_ROWS)))
        .addMethod(
          getBatchUpdateRowsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_UPDATE_ROWS)))
        .addMethod(
          getBatchDeleteRowsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.ncraft.armory.unitable.v1.BatchDeleteRowsRequest,
              org.mojolang.mojo.core.Null>(
                service, METHODID_BATCH_DELETE_ROWS)))
        .build();
  }

  private static abstract class UnitableBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UnitableBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.ncraft.armory.unitable.v1.UnitableProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Unitable");
    }
  }

  private static final class UnitableFileDescriptorSupplier
      extends UnitableBaseDescriptorSupplier {
    UnitableFileDescriptorSupplier() {}
  }

  private static final class UnitableMethodDescriptorSupplier
      extends UnitableBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    UnitableMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (UnitableGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UnitableFileDescriptorSupplier())
              .addMethod(getCreateTableMethod())
              .addMethod(getUpdateTableMethod())
              .addMethod(getGetTableMethod())
              .addMethod(getListTablesMethod())
              .addMethod(getDeleteTableMethod())
              .addMethod(getSyncTableMethod())
              .addMethod(getCreateColumnMethod())
              .addMethod(getUpdateColumnMethod())
              .addMethod(getGetColumnMethod())
              .addMethod(getDeleteColumnMethod())
              .addMethod(getListColumnsMethod())
              .addMethod(getBatchCreateColumnsMethod())
              .addMethod(getBatchUpdateColumnMethod())
              .addMethod(getBatchDeleteColumnMethod())
              .addMethod(getCreateRowMethod())
              .addMethod(getUpdateRowMethod())
              .addMethod(getGetRowMethod())
              .addMethod(getDeleteRowMethod())
              .addMethod(getListRowMethod())
              .addMethod(getExportRowMethod())
              .addMethod(getBatchCreateRowsMethod())
              .addMethod(getBatchUpdateRowsMethod())
              .addMethod(getBatchDeleteRowsMethod())
              .build();
        }
      }
    }
    return result;
  }
}

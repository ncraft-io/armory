// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/db_query_config.proto

package io.ncraft.armory.unitable;

public final class DbQueryConfigProto {
  private DbQueryConfigProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_armory_unitable_DbQueryConfig_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_armory_unitable_DbQueryConfig_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n%armory/unitable/db_query_config.proto\022" +
      "\017armory.unitable\032\036armory/unitable/db_que" +
      "ry.proto\":\n\rDbQueryConfig\022)\n\007queries\030\001 \003" +
      "(\0132\030.armory.unitable.DbQueryBn\n\031io.ncraf" +
      "t.armory.unitableB\022DbQueryConfigProtoP\001Z" +
      ";github.com/ncraft-io/armory/go/pkg/armo" +
      "ry/unitable;unitableb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.ncraft.armory.unitable.DbQueryProto.getDescriptor(),
        });
    internal_static_armory_unitable_DbQueryConfig_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_armory_unitable_DbQueryConfig_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_armory_unitable_DbQueryConfig_descriptor,
        new java.lang.String[] { "Queries", });
    io.ncraft.armory.unitable.DbQueryProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}

// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/db_query.proto

package io.ncraft.armory.unitable;

public final class DbQueryProto {
  private DbQueryProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_armory_unitable_DbQuery_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_armory_unitable_DbQuery_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\036armory/unitable/db_query.proto\022\017armory" +
      ".unitable\032\034armory/unitable/column.proto\032" +
      "\017mojo/mojo.proto\"\207\001\n\007DbQuery\022\n\n\002id\030\001 \001(\t" +
      "\022\014\n\004name\030\002 \001(\t\022\013\n\003sql\030\003 \001(\t\022\022\n\nparameter" +
      "s\030\004 \003(\t\022\r\n\005table\030\005 \001(\t\0222\n\007columns\030\017 \003(\0132" +
      "\027.armory.unitable.ColumnB\010\332\317$\000\350\325$\001Bh\n\031io" +
      ".ncraft.armory.unitableB\014DbQueryProtoP\001Z" +
      ";github.com/ncraft-io/armory/go/pkg/armo" +
      "ry/unitable;unitableb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.ncraft.armory.unitable.ColumnProto.getDescriptor(),
          org.mojolang.mojo.MojoProtos.getDescriptor(),
        });
    internal_static_armory_unitable_DbQuery_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_armory_unitable_DbQuery_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_armory_unitable_DbQuery_descriptor,
        new java.lang.String[] { "Id", "Name", "Sql", "Parameters", "Table", "Columns", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(org.mojolang.mojo.MojoProtos.dbJson);
    registry.add(org.mojolang.mojo.MojoProtos.reference);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    io.ncraft.armory.unitable.ColumnProto.getDescriptor();
    org.mojolang.mojo.MojoProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}

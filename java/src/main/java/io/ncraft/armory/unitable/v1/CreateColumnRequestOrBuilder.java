// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

public interface CreateColumnRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:armory.unitable.v1.CreateColumnRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string database = 1;</code>
   * @return The database.
   */
  java.lang.String getDatabase();
  /**
   * <code>string database = 1;</code>
   * @return The bytes for database.
   */
  com.google.protobuf.ByteString
      getDatabaseBytes();

  /**
   * <code>string table = 2;</code>
   * @return The table.
   */
  java.lang.String getTable();
  /**
   * <code>string table = 2;</code>
   * @return The bytes for table.
   */
  com.google.protobuf.ByteString
      getTableBytes();

  /**
   * <code>.armory.unitable.Column column = 3;</code>
   * @return Whether the column field is set.
   */
  boolean hasColumn();
  /**
   * <code>.armory.unitable.Column column = 3;</code>
   * @return The column.
   */
  io.ncraft.armory.unitable.Column getColumn();
  /**
   * <code>.armory.unitable.Column column = 3;</code>
   */
  io.ncraft.armory.unitable.ColumnOrBuilder getColumnOrBuilder();
}
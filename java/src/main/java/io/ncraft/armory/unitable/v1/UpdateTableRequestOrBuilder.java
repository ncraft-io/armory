// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

public interface UpdateTableRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:armory.unitable.v1.UpdateTableRequest)
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
   * <code>.armory.unitable.Table table = 2;</code>
   * @return Whether the table field is set.
   */
  boolean hasTable();
  /**
   * <code>.armory.unitable.Table table = 2;</code>
   * @return The table.
   */
  io.ncraft.armory.unitable.Table getTable();
  /**
   * <code>.armory.unitable.Table table = 2;</code>
   */
  io.ncraft.armory.unitable.TableOrBuilder getTableOrBuilder();

  /**
   * <code>string id = 3;</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 3;</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();
}
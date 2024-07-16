// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/table.proto

package io.ncraft.armory.unitable;

public interface TableOrBuilder extends
    // @@protoc_insertion_point(interface_extends:armory.unitable.Table)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string id = 1;</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 1;</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <code>string name = 2 [(.mojo.db_index) = ""];</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 2 [(.mojo.db_index) = ""];</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string display_name = 3;</code>
   * @return The displayName.
   */
  java.lang.String getDisplayName();
  /**
   * <code>string display_name = 3;</code>
   * @return The bytes for displayName.
   */
  com.google.protobuf.ByteString
      getDisplayNameBytes();

  /**
   * <code>string export_name = 4;</code>
   * @return The exportName.
   */
  java.lang.String getExportName();
  /**
   * <code>string export_name = 4;</code>
   * @return The bytes for exportName.
   */
  com.google.protobuf.ByteString
      getExportNameBytes();

  /**
   * <code>string tenant = 5 [(.mojo.db_index) = ""];</code>
   * @return The tenant.
   */
  java.lang.String getTenant();
  /**
   * <code>string tenant = 5 [(.mojo.db_index) = ""];</code>
   * @return The bytes for tenant.
   */
  com.google.protobuf.ByteString
      getTenantBytes();

  /**
   * <code>string database = 10 [(.mojo.db_index) = ""];</code>
   * @return The database.
   */
  java.lang.String getDatabase();
  /**
   * <code>string database = 10 [(.mojo.db_index) = ""];</code>
   * @return The bytes for database.
   */
  com.google.protobuf.ByteString
      getDatabaseBytes();

  /**
   * <code>repeated .armory.unitable.Column columns = 15 [(.mojo.reference) = ""];</code>
   */
  java.util.List<io.ncraft.armory.unitable.Column> 
      getColumnsList();
  /**
   * <code>repeated .armory.unitable.Column columns = 15 [(.mojo.reference) = ""];</code>
   */
  io.ncraft.armory.unitable.Column getColumns(int index);
  /**
   * <code>repeated .armory.unitable.Column columns = 15 [(.mojo.reference) = ""];</code>
   */
  int getColumnsCount();
  /**
   * <code>repeated .armory.unitable.Column columns = 15 [(.mojo.reference) = ""];</code>
   */
  java.util.List<? extends io.ncraft.armory.unitable.ColumnOrBuilder> 
      getColumnsOrBuilderList();
  /**
   * <code>repeated .armory.unitable.Column columns = 15 [(.mojo.reference) = ""];</code>
   */
  io.ncraft.armory.unitable.ColumnOrBuilder getColumnsOrBuilder(
      int index);

  /**
   * <code>.mojo.core.Timestamp create_time = 100;</code>
   * @return Whether the createTime field is set.
   */
  boolean hasCreateTime();
  /**
   * <code>.mojo.core.Timestamp create_time = 100;</code>
   * @return The createTime.
   */
  org.mojolang.mojo.core.Timestamp getCreateTime();
  /**
   * <code>.mojo.core.Timestamp create_time = 100;</code>
   */
  org.mojolang.mojo.core.TimestampOrBuilder getCreateTimeOrBuilder();

  /**
   * <code>.mojo.core.Timestamp update_time = 101;</code>
   * @return Whether the updateTime field is set.
   */
  boolean hasUpdateTime();
  /**
   * <code>.mojo.core.Timestamp update_time = 101;</code>
   * @return The updateTime.
   */
  org.mojolang.mojo.core.Timestamp getUpdateTime();
  /**
   * <code>.mojo.core.Timestamp update_time = 101;</code>
   */
  org.mojolang.mojo.core.TimestampOrBuilder getUpdateTimeOrBuilder();
}
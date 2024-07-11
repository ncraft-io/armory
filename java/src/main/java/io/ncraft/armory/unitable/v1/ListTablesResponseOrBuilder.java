// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

public interface ListTablesResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:armory.unitable.v1.ListTablesResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .armory.unitable.Table tables = 1;</code>
   */
  java.util.List<io.ncraft.armory.unitable.Table> 
      getTablesList();
  /**
   * <code>repeated .armory.unitable.Table tables = 1;</code>
   */
  io.ncraft.armory.unitable.Table getTables(int index);
  /**
   * <code>repeated .armory.unitable.Table tables = 1;</code>
   */
  int getTablesCount();
  /**
   * <code>repeated .armory.unitable.Table tables = 1;</code>
   */
  java.util.List<? extends io.ncraft.armory.unitable.TableOrBuilder> 
      getTablesOrBuilderList();
  /**
   * <code>repeated .armory.unitable.Table tables = 1;</code>
   */
  io.ncraft.armory.unitable.TableOrBuilder getTablesOrBuilder(
      int index);

  /**
   * <code>int32 total_count = 2000;</code>
   * @return The totalCount.
   */
  int getTotalCount();

  /**
   * <code>string next_page_token = 2001;</code>
   * @return The nextPageToken.
   */
  java.lang.String getNextPageToken();
  /**
   * <code>string next_page_token = 2001;</code>
   * @return The bytes for nextPageToken.
   */
  com.google.protobuf.ByteString
      getNextPageTokenBytes();
}

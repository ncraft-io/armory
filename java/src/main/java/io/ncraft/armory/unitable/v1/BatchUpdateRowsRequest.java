// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

/**
 * Protobuf type {@code armory.unitable.v1.BatchUpdateRowsRequest}
 */
public final class BatchUpdateRowsRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:armory.unitable.v1.BatchUpdateRowsRequest)
    BatchUpdateRowsRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use BatchUpdateRowsRequest.newBuilder() to construct.
  private BatchUpdateRowsRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private BatchUpdateRowsRequest() {
    database_ = "";
    table_ = "";
    rows_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new BatchUpdateRowsRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateRowsRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateRowsRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.class, io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.Builder.class);
  }

  public static final int DATABASE_FIELD_NUMBER = 1;
  private volatile java.lang.Object database_;
  /**
   * <code>string database = 1;</code>
   * @return The database.
   */
  @java.lang.Override
  public java.lang.String getDatabase() {
    java.lang.Object ref = database_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      database_ = s;
      return s;
    }
  }
  /**
   * <code>string database = 1;</code>
   * @return The bytes for database.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getDatabaseBytes() {
    java.lang.Object ref = database_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      database_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TABLE_FIELD_NUMBER = 2;
  private volatile java.lang.Object table_;
  /**
   * <code>string table = 2;</code>
   * @return The table.
   */
  @java.lang.Override
  public java.lang.String getTable() {
    java.lang.Object ref = table_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      table_ = s;
      return s;
    }
  }
  /**
   * <code>string table = 2;</code>
   * @return The bytes for table.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getTableBytes() {
    java.lang.Object ref = table_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      table_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int ROWS_FIELD_NUMBER = 4;
  private java.util.List<org.mojolang.mojo.core.Object> rows_;
  /**
   * <code>repeated .mojo.core.Object rows = 4;</code>
   */
  @java.lang.Override
  public java.util.List<org.mojolang.mojo.core.Object> getRowsList() {
    return rows_;
  }
  /**
   * <code>repeated .mojo.core.Object rows = 4;</code>
   */
  @java.lang.Override
  public java.util.List<? extends org.mojolang.mojo.core.ObjectOrBuilder> 
      getRowsOrBuilderList() {
    return rows_;
  }
  /**
   * <code>repeated .mojo.core.Object rows = 4;</code>
   */
  @java.lang.Override
  public int getRowsCount() {
    return rows_.size();
  }
  /**
   * <code>repeated .mojo.core.Object rows = 4;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Object getRows(int index) {
    return rows_.get(index);
  }
  /**
   * <code>repeated .mojo.core.Object rows = 4;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.ObjectOrBuilder getRowsOrBuilder(
      int index) {
    return rows_.get(index);
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(database_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, database_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(table_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, table_);
    }
    for (int i = 0; i < rows_.size(); i++) {
      output.writeMessage(4, rows_.get(i));
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(database_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, database_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(table_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, table_);
    }
    for (int i = 0; i < rows_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(4, rows_.get(i));
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest)) {
      return super.equals(obj);
    }
    io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest other = (io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest) obj;

    if (!getDatabase()
        .equals(other.getDatabase())) return false;
    if (!getTable()
        .equals(other.getTable())) return false;
    if (!getRowsList()
        .equals(other.getRowsList())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + DATABASE_FIELD_NUMBER;
    hash = (53 * hash) + getDatabase().hashCode();
    hash = (37 * hash) + TABLE_FIELD_NUMBER;
    hash = (53 * hash) + getTable().hashCode();
    if (getRowsCount() > 0) {
      hash = (37 * hash) + ROWS_FIELD_NUMBER;
      hash = (53 * hash) + getRowsList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code armory.unitable.v1.BatchUpdateRowsRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:armory.unitable.v1.BatchUpdateRowsRequest)
      io.ncraft.armory.unitable.v1.BatchUpdateRowsRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateRowsRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateRowsRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.class, io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.Builder.class);
    }

    // Construct using io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      database_ = "";

      table_ = "";

      if (rowsBuilder_ == null) {
        rows_ = java.util.Collections.emptyList();
      } else {
        rows_ = null;
        rowsBuilder_.clear();
      }
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateRowsRequest_descriptor;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest getDefaultInstanceForType() {
      return io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest build() {
      io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest buildPartial() {
      io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest result = new io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest(this);
      int from_bitField0_ = bitField0_;
      result.database_ = database_;
      result.table_ = table_;
      if (rowsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          rows_ = java.util.Collections.unmodifiableList(rows_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.rows_ = rows_;
      } else {
        result.rows_ = rowsBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest) {
        return mergeFrom((io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest other) {
      if (other == io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest.getDefaultInstance()) return this;
      if (!other.getDatabase().isEmpty()) {
        database_ = other.database_;
        onChanged();
      }
      if (!other.getTable().isEmpty()) {
        table_ = other.table_;
        onChanged();
      }
      if (rowsBuilder_ == null) {
        if (!other.rows_.isEmpty()) {
          if (rows_.isEmpty()) {
            rows_ = other.rows_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureRowsIsMutable();
            rows_.addAll(other.rows_);
          }
          onChanged();
        }
      } else {
        if (!other.rows_.isEmpty()) {
          if (rowsBuilder_.isEmpty()) {
            rowsBuilder_.dispose();
            rowsBuilder_ = null;
            rows_ = other.rows_;
            bitField0_ = (bitField0_ & ~0x00000001);
            rowsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getRowsFieldBuilder() : null;
          } else {
            rowsBuilder_.addAllMessages(other.rows_);
          }
        }
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              database_ = input.readStringRequireUtf8();

              break;
            } // case 10
            case 18: {
              table_ = input.readStringRequireUtf8();

              break;
            } // case 18
            case 34: {
              org.mojolang.mojo.core.Object m =
                  input.readMessage(
                      org.mojolang.mojo.core.Object.parser(),
                      extensionRegistry);
              if (rowsBuilder_ == null) {
                ensureRowsIsMutable();
                rows_.add(m);
              } else {
                rowsBuilder_.addMessage(m);
              }
              break;
            } // case 34
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private java.lang.Object database_ = "";
    /**
     * <code>string database = 1;</code>
     * @return The database.
     */
    public java.lang.String getDatabase() {
      java.lang.Object ref = database_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        database_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string database = 1;</code>
     * @return The bytes for database.
     */
    public com.google.protobuf.ByteString
        getDatabaseBytes() {
      java.lang.Object ref = database_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        database_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string database = 1;</code>
     * @param value The database to set.
     * @return This builder for chaining.
     */
    public Builder setDatabase(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      database_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string database = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearDatabase() {
      
      database_ = getDefaultInstance().getDatabase();
      onChanged();
      return this;
    }
    /**
     * <code>string database = 1;</code>
     * @param value The bytes for database to set.
     * @return This builder for chaining.
     */
    public Builder setDatabaseBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      database_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object table_ = "";
    /**
     * <code>string table = 2;</code>
     * @return The table.
     */
    public java.lang.String getTable() {
      java.lang.Object ref = table_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        table_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string table = 2;</code>
     * @return The bytes for table.
     */
    public com.google.protobuf.ByteString
        getTableBytes() {
      java.lang.Object ref = table_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        table_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string table = 2;</code>
     * @param value The table to set.
     * @return This builder for chaining.
     */
    public Builder setTable(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      table_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string table = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearTable() {
      
      table_ = getDefaultInstance().getTable();
      onChanged();
      return this;
    }
    /**
     * <code>string table = 2;</code>
     * @param value The bytes for table to set.
     * @return This builder for chaining.
     */
    public Builder setTableBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      table_ = value;
      onChanged();
      return this;
    }

    private java.util.List<org.mojolang.mojo.core.Object> rows_ =
      java.util.Collections.emptyList();
    private void ensureRowsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        rows_ = new java.util.ArrayList<org.mojolang.mojo.core.Object>(rows_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder> rowsBuilder_;

    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Object> getRowsList() {
      if (rowsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(rows_);
      } else {
        return rowsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public int getRowsCount() {
      if (rowsBuilder_ == null) {
        return rows_.size();
      } else {
        return rowsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public org.mojolang.mojo.core.Object getRows(int index) {
      if (rowsBuilder_ == null) {
        return rows_.get(index);
      } else {
        return rowsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder setRows(
        int index, org.mojolang.mojo.core.Object value) {
      if (rowsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureRowsIsMutable();
        rows_.set(index, value);
        onChanged();
      } else {
        rowsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder setRows(
        int index, org.mojolang.mojo.core.Object.Builder builderForValue) {
      if (rowsBuilder_ == null) {
        ensureRowsIsMutable();
        rows_.set(index, builderForValue.build());
        onChanged();
      } else {
        rowsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder addRows(org.mojolang.mojo.core.Object value) {
      if (rowsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureRowsIsMutable();
        rows_.add(value);
        onChanged();
      } else {
        rowsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder addRows(
        int index, org.mojolang.mojo.core.Object value) {
      if (rowsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureRowsIsMutable();
        rows_.add(index, value);
        onChanged();
      } else {
        rowsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder addRows(
        org.mojolang.mojo.core.Object.Builder builderForValue) {
      if (rowsBuilder_ == null) {
        ensureRowsIsMutable();
        rows_.add(builderForValue.build());
        onChanged();
      } else {
        rowsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder addRows(
        int index, org.mojolang.mojo.core.Object.Builder builderForValue) {
      if (rowsBuilder_ == null) {
        ensureRowsIsMutable();
        rows_.add(index, builderForValue.build());
        onChanged();
      } else {
        rowsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder addAllRows(
        java.lang.Iterable<? extends org.mojolang.mojo.core.Object> values) {
      if (rowsBuilder_ == null) {
        ensureRowsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, rows_);
        onChanged();
      } else {
        rowsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder clearRows() {
      if (rowsBuilder_ == null) {
        rows_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        rowsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public Builder removeRows(int index) {
      if (rowsBuilder_ == null) {
        ensureRowsIsMutable();
        rows_.remove(index);
        onChanged();
      } else {
        rowsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public org.mojolang.mojo.core.Object.Builder getRowsBuilder(
        int index) {
      return getRowsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public org.mojolang.mojo.core.ObjectOrBuilder getRowsOrBuilder(
        int index) {
      if (rowsBuilder_ == null) {
        return rows_.get(index);  } else {
        return rowsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public java.util.List<? extends org.mojolang.mojo.core.ObjectOrBuilder> 
         getRowsOrBuilderList() {
      if (rowsBuilder_ != null) {
        return rowsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(rows_);
      }
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public org.mojolang.mojo.core.Object.Builder addRowsBuilder() {
      return getRowsFieldBuilder().addBuilder(
          org.mojolang.mojo.core.Object.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public org.mojolang.mojo.core.Object.Builder addRowsBuilder(
        int index) {
      return getRowsFieldBuilder().addBuilder(
          index, org.mojolang.mojo.core.Object.getDefaultInstance());
    }
    /**
     * <code>repeated .mojo.core.Object rows = 4;</code>
     */
    public java.util.List<org.mojolang.mojo.core.Object.Builder> 
         getRowsBuilderList() {
      return getRowsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder> 
        getRowsFieldBuilder() {
      if (rowsBuilder_ == null) {
        rowsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder>(
                rows_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        rows_ = null;
      }
      return rowsBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:armory.unitable.v1.BatchUpdateRowsRequest)
  }

  // @@protoc_insertion_point(class_scope:armory.unitable.v1.BatchUpdateRowsRequest)
  private static final io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest();
  }

  public static io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<BatchUpdateRowsRequest>
      PARSER = new com.google.protobuf.AbstractParser<BatchUpdateRowsRequest>() {
    @java.lang.Override
    public BatchUpdateRowsRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<BatchUpdateRowsRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<BatchUpdateRowsRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.ncraft.armory.unitable.v1.BatchUpdateRowsRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}


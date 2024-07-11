// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

/**
 * Protobuf type {@code armory.unitable.v1.BatchUpdateColumnRequest}
 */
public final class BatchUpdateColumnRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:armory.unitable.v1.BatchUpdateColumnRequest)
    BatchUpdateColumnRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use BatchUpdateColumnRequest.newBuilder() to construct.
  private BatchUpdateColumnRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private BatchUpdateColumnRequest() {
    database_ = "";
    table_ = "";
    columns_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new BatchUpdateColumnRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateColumnRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateColumnRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.class, io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.Builder.class);
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

  public static final int COLUMNS_FIELD_NUMBER = 3;
  private java.util.List<io.ncraft.armory.unitable.Column> columns_;
  /**
   * <code>repeated .armory.unitable.Column columns = 3;</code>
   */
  @java.lang.Override
  public java.util.List<io.ncraft.armory.unitable.Column> getColumnsList() {
    return columns_;
  }
  /**
   * <code>repeated .armory.unitable.Column columns = 3;</code>
   */
  @java.lang.Override
  public java.util.List<? extends io.ncraft.armory.unitable.ColumnOrBuilder> 
      getColumnsOrBuilderList() {
    return columns_;
  }
  /**
   * <code>repeated .armory.unitable.Column columns = 3;</code>
   */
  @java.lang.Override
  public int getColumnsCount() {
    return columns_.size();
  }
  /**
   * <code>repeated .armory.unitable.Column columns = 3;</code>
   */
  @java.lang.Override
  public io.ncraft.armory.unitable.Column getColumns(int index) {
    return columns_.get(index);
  }
  /**
   * <code>repeated .armory.unitable.Column columns = 3;</code>
   */
  @java.lang.Override
  public io.ncraft.armory.unitable.ColumnOrBuilder getColumnsOrBuilder(
      int index) {
    return columns_.get(index);
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
    for (int i = 0; i < columns_.size(); i++) {
      output.writeMessage(3, columns_.get(i));
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
    for (int i = 0; i < columns_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(3, columns_.get(i));
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
    if (!(obj instanceof io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest)) {
      return super.equals(obj);
    }
    io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest other = (io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest) obj;

    if (!getDatabase()
        .equals(other.getDatabase())) return false;
    if (!getTable()
        .equals(other.getTable())) return false;
    if (!getColumnsList()
        .equals(other.getColumnsList())) return false;
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
    if (getColumnsCount() > 0) {
      hash = (37 * hash) + COLUMNS_FIELD_NUMBER;
      hash = (53 * hash) + getColumnsList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest parseFrom(
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
  public static Builder newBuilder(io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest prototype) {
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
   * Protobuf type {@code armory.unitable.v1.BatchUpdateColumnRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:armory.unitable.v1.BatchUpdateColumnRequest)
      io.ncraft.armory.unitable.v1.BatchUpdateColumnRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateColumnRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateColumnRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.class, io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.Builder.class);
    }

    // Construct using io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.newBuilder()
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

      if (columnsBuilder_ == null) {
        columns_ = java.util.Collections.emptyList();
      } else {
        columns_ = null;
        columnsBuilder_.clear();
      }
      bitField0_ = (bitField0_ & ~0x00000001);
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_BatchUpdateColumnRequest_descriptor;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest getDefaultInstanceForType() {
      return io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest build() {
      io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest buildPartial() {
      io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest result = new io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest(this);
      int from_bitField0_ = bitField0_;
      result.database_ = database_;
      result.table_ = table_;
      if (columnsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          columns_ = java.util.Collections.unmodifiableList(columns_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.columns_ = columns_;
      } else {
        result.columns_ = columnsBuilder_.build();
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
      if (other instanceof io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest) {
        return mergeFrom((io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest other) {
      if (other == io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest.getDefaultInstance()) return this;
      if (!other.getDatabase().isEmpty()) {
        database_ = other.database_;
        onChanged();
      }
      if (!other.getTable().isEmpty()) {
        table_ = other.table_;
        onChanged();
      }
      if (columnsBuilder_ == null) {
        if (!other.columns_.isEmpty()) {
          if (columns_.isEmpty()) {
            columns_ = other.columns_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureColumnsIsMutable();
            columns_.addAll(other.columns_);
          }
          onChanged();
        }
      } else {
        if (!other.columns_.isEmpty()) {
          if (columnsBuilder_.isEmpty()) {
            columnsBuilder_.dispose();
            columnsBuilder_ = null;
            columns_ = other.columns_;
            bitField0_ = (bitField0_ & ~0x00000001);
            columnsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getColumnsFieldBuilder() : null;
          } else {
            columnsBuilder_.addAllMessages(other.columns_);
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
            case 26: {
              io.ncraft.armory.unitable.Column m =
                  input.readMessage(
                      io.ncraft.armory.unitable.Column.parser(),
                      extensionRegistry);
              if (columnsBuilder_ == null) {
                ensureColumnsIsMutable();
                columns_.add(m);
              } else {
                columnsBuilder_.addMessage(m);
              }
              break;
            } // case 26
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

    private java.util.List<io.ncraft.armory.unitable.Column> columns_ =
      java.util.Collections.emptyList();
    private void ensureColumnsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        columns_ = new java.util.ArrayList<io.ncraft.armory.unitable.Column>(columns_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        io.ncraft.armory.unitable.Column, io.ncraft.armory.unitable.Column.Builder, io.ncraft.armory.unitable.ColumnOrBuilder> columnsBuilder_;

    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public java.util.List<io.ncraft.armory.unitable.Column> getColumnsList() {
      if (columnsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(columns_);
      } else {
        return columnsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public int getColumnsCount() {
      if (columnsBuilder_ == null) {
        return columns_.size();
      } else {
        return columnsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public io.ncraft.armory.unitable.Column getColumns(int index) {
      if (columnsBuilder_ == null) {
        return columns_.get(index);
      } else {
        return columnsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder setColumns(
        int index, io.ncraft.armory.unitable.Column value) {
      if (columnsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureColumnsIsMutable();
        columns_.set(index, value);
        onChanged();
      } else {
        columnsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder setColumns(
        int index, io.ncraft.armory.unitable.Column.Builder builderForValue) {
      if (columnsBuilder_ == null) {
        ensureColumnsIsMutable();
        columns_.set(index, builderForValue.build());
        onChanged();
      } else {
        columnsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder addColumns(io.ncraft.armory.unitable.Column value) {
      if (columnsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureColumnsIsMutable();
        columns_.add(value);
        onChanged();
      } else {
        columnsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder addColumns(
        int index, io.ncraft.armory.unitable.Column value) {
      if (columnsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureColumnsIsMutable();
        columns_.add(index, value);
        onChanged();
      } else {
        columnsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder addColumns(
        io.ncraft.armory.unitable.Column.Builder builderForValue) {
      if (columnsBuilder_ == null) {
        ensureColumnsIsMutable();
        columns_.add(builderForValue.build());
        onChanged();
      } else {
        columnsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder addColumns(
        int index, io.ncraft.armory.unitable.Column.Builder builderForValue) {
      if (columnsBuilder_ == null) {
        ensureColumnsIsMutable();
        columns_.add(index, builderForValue.build());
        onChanged();
      } else {
        columnsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder addAllColumns(
        java.lang.Iterable<? extends io.ncraft.armory.unitable.Column> values) {
      if (columnsBuilder_ == null) {
        ensureColumnsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, columns_);
        onChanged();
      } else {
        columnsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder clearColumns() {
      if (columnsBuilder_ == null) {
        columns_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        columnsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public Builder removeColumns(int index) {
      if (columnsBuilder_ == null) {
        ensureColumnsIsMutable();
        columns_.remove(index);
        onChanged();
      } else {
        columnsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public io.ncraft.armory.unitable.Column.Builder getColumnsBuilder(
        int index) {
      return getColumnsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public io.ncraft.armory.unitable.ColumnOrBuilder getColumnsOrBuilder(
        int index) {
      if (columnsBuilder_ == null) {
        return columns_.get(index);  } else {
        return columnsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public java.util.List<? extends io.ncraft.armory.unitable.ColumnOrBuilder> 
         getColumnsOrBuilderList() {
      if (columnsBuilder_ != null) {
        return columnsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(columns_);
      }
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public io.ncraft.armory.unitable.Column.Builder addColumnsBuilder() {
      return getColumnsFieldBuilder().addBuilder(
          io.ncraft.armory.unitable.Column.getDefaultInstance());
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public io.ncraft.armory.unitable.Column.Builder addColumnsBuilder(
        int index) {
      return getColumnsFieldBuilder().addBuilder(
          index, io.ncraft.armory.unitable.Column.getDefaultInstance());
    }
    /**
     * <code>repeated .armory.unitable.Column columns = 3;</code>
     */
    public java.util.List<io.ncraft.armory.unitable.Column.Builder> 
         getColumnsBuilderList() {
      return getColumnsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        io.ncraft.armory.unitable.Column, io.ncraft.armory.unitable.Column.Builder, io.ncraft.armory.unitable.ColumnOrBuilder> 
        getColumnsFieldBuilder() {
      if (columnsBuilder_ == null) {
        columnsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            io.ncraft.armory.unitable.Column, io.ncraft.armory.unitable.Column.Builder, io.ncraft.armory.unitable.ColumnOrBuilder>(
                columns_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        columns_ = null;
      }
      return columnsBuilder_;
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


    // @@protoc_insertion_point(builder_scope:armory.unitable.v1.BatchUpdateColumnRequest)
  }

  // @@protoc_insertion_point(class_scope:armory.unitable.v1.BatchUpdateColumnRequest)
  private static final io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest();
  }

  public static io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<BatchUpdateColumnRequest>
      PARSER = new com.google.protobuf.AbstractParser<BatchUpdateColumnRequest>() {
    @java.lang.Override
    public BatchUpdateColumnRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<BatchUpdateColumnRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<BatchUpdateColumnRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.ncraft.armory.unitable.v1.BatchUpdateColumnRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}


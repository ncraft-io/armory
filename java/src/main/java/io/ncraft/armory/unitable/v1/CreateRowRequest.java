// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: armory/unitable/v1/unitable.proto

package io.ncraft.armory.unitable.v1;

/**
 * Protobuf type {@code armory.unitable.v1.CreateRowRequest}
 */
public final class CreateRowRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:armory.unitable.v1.CreateRowRequest)
    CreateRowRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CreateRowRequest.newBuilder() to construct.
  private CreateRowRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CreateRowRequest() {
    database_ = "";
    table_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CreateRowRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_CreateRowRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_CreateRowRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.ncraft.armory.unitable.v1.CreateRowRequest.class, io.ncraft.armory.unitable.v1.CreateRowRequest.Builder.class);
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

  public static final int ROW_FIELD_NUMBER = 3;
  private org.mojolang.mojo.core.Object row_;
  /**
   * <code>.mojo.core.Object row = 3;</code>
   * @return Whether the row field is set.
   */
  @java.lang.Override
  public boolean hasRow() {
    return row_ != null;
  }
  /**
   * <code>.mojo.core.Object row = 3;</code>
   * @return The row.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Object getRow() {
    return row_ == null ? org.mojolang.mojo.core.Object.getDefaultInstance() : row_;
  }
  /**
   * <code>.mojo.core.Object row = 3;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.ObjectOrBuilder getRowOrBuilder() {
    return getRow();
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
    if (row_ != null) {
      output.writeMessage(3, getRow());
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
    if (row_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(3, getRow());
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
    if (!(obj instanceof io.ncraft.armory.unitable.v1.CreateRowRequest)) {
      return super.equals(obj);
    }
    io.ncraft.armory.unitable.v1.CreateRowRequest other = (io.ncraft.armory.unitable.v1.CreateRowRequest) obj;

    if (!getDatabase()
        .equals(other.getDatabase())) return false;
    if (!getTable()
        .equals(other.getTable())) return false;
    if (hasRow() != other.hasRow()) return false;
    if (hasRow()) {
      if (!getRow()
          .equals(other.getRow())) return false;
    }
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
    if (hasRow()) {
      hash = (37 * hash) + ROW_FIELD_NUMBER;
      hash = (53 * hash) + getRow().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.ncraft.armory.unitable.v1.CreateRowRequest parseFrom(
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
  public static Builder newBuilder(io.ncraft.armory.unitable.v1.CreateRowRequest prototype) {
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
   * Protobuf type {@code armory.unitable.v1.CreateRowRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:armory.unitable.v1.CreateRowRequest)
      io.ncraft.armory.unitable.v1.CreateRowRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_CreateRowRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_CreateRowRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.ncraft.armory.unitable.v1.CreateRowRequest.class, io.ncraft.armory.unitable.v1.CreateRowRequest.Builder.class);
    }

    // Construct using io.ncraft.armory.unitable.v1.CreateRowRequest.newBuilder()
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

      if (rowBuilder_ == null) {
        row_ = null;
      } else {
        row_ = null;
        rowBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.ncraft.armory.unitable.v1.UnitableProto.internal_static_armory_unitable_v1_CreateRowRequest_descriptor;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.CreateRowRequest getDefaultInstanceForType() {
      return io.ncraft.armory.unitable.v1.CreateRowRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.CreateRowRequest build() {
      io.ncraft.armory.unitable.v1.CreateRowRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.ncraft.armory.unitable.v1.CreateRowRequest buildPartial() {
      io.ncraft.armory.unitable.v1.CreateRowRequest result = new io.ncraft.armory.unitable.v1.CreateRowRequest(this);
      result.database_ = database_;
      result.table_ = table_;
      if (rowBuilder_ == null) {
        result.row_ = row_;
      } else {
        result.row_ = rowBuilder_.build();
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
      if (other instanceof io.ncraft.armory.unitable.v1.CreateRowRequest) {
        return mergeFrom((io.ncraft.armory.unitable.v1.CreateRowRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.ncraft.armory.unitable.v1.CreateRowRequest other) {
      if (other == io.ncraft.armory.unitable.v1.CreateRowRequest.getDefaultInstance()) return this;
      if (!other.getDatabase().isEmpty()) {
        database_ = other.database_;
        onChanged();
      }
      if (!other.getTable().isEmpty()) {
        table_ = other.table_;
        onChanged();
      }
      if (other.hasRow()) {
        mergeRow(other.getRow());
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
              input.readMessage(
                  getRowFieldBuilder().getBuilder(),
                  extensionRegistry);

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

    private org.mojolang.mojo.core.Object row_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder> rowBuilder_;
    /**
     * <code>.mojo.core.Object row = 3;</code>
     * @return Whether the row field is set.
     */
    public boolean hasRow() {
      return rowBuilder_ != null || row_ != null;
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     * @return The row.
     */
    public org.mojolang.mojo.core.Object getRow() {
      if (rowBuilder_ == null) {
        return row_ == null ? org.mojolang.mojo.core.Object.getDefaultInstance() : row_;
      } else {
        return rowBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public Builder setRow(org.mojolang.mojo.core.Object value) {
      if (rowBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        row_ = value;
        onChanged();
      } else {
        rowBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public Builder setRow(
        org.mojolang.mojo.core.Object.Builder builderForValue) {
      if (rowBuilder_ == null) {
        row_ = builderForValue.build();
        onChanged();
      } else {
        rowBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public Builder mergeRow(org.mojolang.mojo.core.Object value) {
      if (rowBuilder_ == null) {
        if (row_ != null) {
          row_ =
            org.mojolang.mojo.core.Object.newBuilder(row_).mergeFrom(value).buildPartial();
        } else {
          row_ = value;
        }
        onChanged();
      } else {
        rowBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public Builder clearRow() {
      if (rowBuilder_ == null) {
        row_ = null;
        onChanged();
      } else {
        row_ = null;
        rowBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public org.mojolang.mojo.core.Object.Builder getRowBuilder() {
      
      onChanged();
      return getRowFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    public org.mojolang.mojo.core.ObjectOrBuilder getRowOrBuilder() {
      if (rowBuilder_ != null) {
        return rowBuilder_.getMessageOrBuilder();
      } else {
        return row_ == null ?
            org.mojolang.mojo.core.Object.getDefaultInstance() : row_;
      }
    }
    /**
     * <code>.mojo.core.Object row = 3;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder> 
        getRowFieldBuilder() {
      if (rowBuilder_ == null) {
        rowBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Object, org.mojolang.mojo.core.Object.Builder, org.mojolang.mojo.core.ObjectOrBuilder>(
                getRow(),
                getParentForChildren(),
                isClean());
        row_ = null;
      }
      return rowBuilder_;
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


    // @@protoc_insertion_point(builder_scope:armory.unitable.v1.CreateRowRequest)
  }

  // @@protoc_insertion_point(class_scope:armory.unitable.v1.CreateRowRequest)
  private static final io.ncraft.armory.unitable.v1.CreateRowRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.ncraft.armory.unitable.v1.CreateRowRequest();
  }

  public static io.ncraft.armory.unitable.v1.CreateRowRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CreateRowRequest>
      PARSER = new com.google.protobuf.AbstractParser<CreateRowRequest>() {
    @java.lang.Override
    public CreateRowRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<CreateRowRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CreateRowRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.ncraft.armory.unitable.v1.CreateRowRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}


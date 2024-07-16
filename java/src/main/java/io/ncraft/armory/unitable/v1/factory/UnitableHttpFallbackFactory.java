// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.

package io.ncraft.armory.unitable.v1.factory;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.openfeign.FallbackFactory;
import org.springframework.stereotype.Component;
import java.util.List;
import java.util.Map;
import org.mojolang.mojo.core.Null;
import org.mojolang.mojo.core.Result;
import org.mojolang.mojo.core.ErrorException;
import org.mojolang.mojo.core.ErrorCodes;
import org.mojolang.mojo.core.Pagination;
import io.ncraft.armory.unitable.*;
import io.ncraft.armory.unitable.v1.UnitableHttp;

@Component
public class UnitableHttpFallbackFactory implements FallbackFactory<UnitableHttp> {
    private static final Logger log = LoggerFactory.getLogger(UnitableHttpFallbackFactory.class);

    @Override
    public UnitableHttp create(Throwable throwable) {
        return new UnitableHttp() {
            
            @Override
            public Result<Table> createTable(String database, Table table) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to create_table."));
            }
            
            @Override
            public Result<Null> updateTable(String database, Table table, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to update_table."));
            }
            
            @Override
            public Result<Table> getTable(String database, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to get_table."));
            }
            
            @Override
            public Pagination<Table> listTables(String database, int pageSize, String pageToken, int skip, String filter, String order, String fieldMask, boolean unique) {
                 return Pagination.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to list_tables."));
            }
            
            @Override
            public Result<Null> deleteTable(String database, String id, boolean force) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to delete_table."));
            }
            
            @Override
            public Result<Table> syncTable(String database, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to sync_table."));
            }
            
            @Override
            public Result<Column> createColumn(String database, String table, Column column) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to create_column."));
            }
            
            @Override
            public Result<Null> updateColumn(String database, String table, Column column, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to update_column."));
            }
            
            @Override
            public Result<Column> getColumn(String database, String table, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to get_column."));
            }
            
            @Override
            public Result<Null> deleteColumn(String database, String table, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to delete_column."));
            }
            
            @Override
            public Pagination<Column> listColumns(String database, String table, int pageSize, String pageToken, int skip, String filter, String order, String fieldMask, boolean unique) {
                 return Pagination.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to list_columns."));
            }
            
            @Override
            public Result<Null> batchCreateColumns(String database, String table, List<Column> columns) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_create_columns."));
            }
            
            @Override
            public Result<Null> batchUpdateColumn(String database, String table, List<Column> columns) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_update_column."));
            }
            
            @Override
            public Result<Null> batchDeleteColumn(String database, String table, List<String> ids) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_delete_column."));
            }
            
            @Override
            public Result<org.mojolang.mojo.core.Object> createRow(String database, String table, org.mojolang.mojo.core.Object row) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to create_row."));
            }
            
            @Override
            public Result<Null> updateRow(String database, String table, String id, org.mojolang.mojo.core.Object row) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to update_row."));
            }
            
            @Override
            public Result<org.mojolang.mojo.core.Object> getRow(String database, String table, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to get_row."));
            }
            
            @Override
            public Result<Null> deleteRow(String database, String table, String id) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to delete_row."));
            }
            
            @Override
            public Pagination<org.mojolang.mojo.core.Object> listRow(String database, String table, int pageSize, String pageToken, int skip, String filter, String order, String fieldMask, boolean unique) {
                 return Pagination.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to list_row."));
            }
            
            @Override
            public Pagination<org.mojolang.mojo.core.Object> exportRow(String database, String table, int pageSize, String pageToken, int skip, String filter, String order, String fieldMask, boolean unique) {
                 return Pagination.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to export_row."));
            }
            
            @Override
            public Result<Null> batchCreateRows(String database, String table, List<org.mojolang.mojo.core.Object> rows) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_create_rows."));
            }
            
            @Override
            public Result<Null> batchUpdateRows(String database, String table, List<org.mojolang.mojo.core.Object> rows) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_update_rows."));
            }
            
            @Override
            public Result<Null> batchDeleteRows(String database, String table, List<String> ids) {
                 return Result.fail(new ErrorException(ErrorCodes.INTERNAL_ERROR, "failed to batch_delete_rows."));
            }
            
        };
    }
}
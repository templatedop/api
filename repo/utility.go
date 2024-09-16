package repo

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/templatedop/api/db"
)

type SQLValue string

func execReturn[T any](ctx context.Context, db *db.DB, sql string, args []any, scanFn pgx.RowToFunc[T]) (T, error) {
	var result T
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	collectedRow, err := pgx.CollectOneRow(rows, scanFn)
	if err != nil {
		return result, err
	}

	return collectedRow, nil
}

func UpdateReturning[T any](ctx context.Context, db *db.DB, query sq.UpdateBuilder, scanFn pgx.RowToFunc[T]) (T, error) {
	var result T
	sql, args, err := query.ToSql()
	if err != nil {
		return result, err
	}
	collectedrows, err := execReturn(ctx, db, sql, args, scanFn)
	if err != nil {
		return result, err
	}
	return collectedrows, nil

}

func execinsert(ctx context.Context, db *db.DB, sql string, args []any) (pgconn.CommandTag, error) {

	rows, err := db.Exec(ctx, sql, args...)

	if err != nil {
		return rows, err
	}

	return rows, err
}

func execupdate(ctx context.Context, db *db.DB, sql string, args []any) (pgconn.CommandTag, error) {

	rows, err := db.Exec(ctx, sql, args...)

	if err != nil {
		return rows, err
	}

	return rows, err
}

func execdelete(ctx context.Context, db *db.DB, sql string, args []any) (pgconn.CommandTag, error) {

	rows, err := db.Exec(ctx, sql, args...)

	if err != nil {
		return rows, err
	}

	return rows, err
}

func exec(ctx context.Context, db *db.DB, sql string, args []any) (pgconn.CommandTag, error) {

	rows, err := db.Query(ctx, sql, args...)

	if err != nil {
		return pgconn.CommandTag{}, err
	}
	defer rows.Close()
	return rows.CommandTag(), rows.Err()
}

func Update(ctx context.Context, db *db.DB, query sq.UpdateBuilder) (pgconn.CommandTag, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return execupdate(ctx, db, sql, args)
}

func Delete(ctx context.Context, db *db.DB, query sq.DeleteBuilder) (pgconn.CommandTag, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return execdelete(ctx, db, sql, args)
}

func Insert(ctx context.Context, db *db.DB, query sq.InsertBuilder) (pgconn.CommandTag, error) {
	sql, args, err := query.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return execinsert(ctx, db, sql, args)
}

func ExecRow(ctx context.Context, db *db.DB, sql string, args ...any) (pgconn.CommandTag, error) {
	ct, err := exec(ctx, db, sql, args)
	if err != nil {
		return ct, err
	}
	rowsAffected := ct.RowsAffected()
	if rowsAffected == 0 {
		return ct, pgx.ErrNoRows
	}
	return ct, nil
}

func SelectOneOK[T any](ctx context.Context, db *db.DB, builder sq.SelectBuilder, scanFn pgx.RowToFunc[T]) (T, bool, error) {

	var zero T
	sql, args, err := builder.ToSql()
	if err != nil {
		return zero, false, err
	}
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			return zero, false, nil
		}
		return zero, false, err
	}
	defer rows.Close()
	collectedRow, b, err := CollectOneRowOK(rows, scanFn)
	if err != nil {
		return zero, false, err
	}

	return collectedRow, b, nil
}

func SelectOne[T any](ctx context.Context, db *db.DB, builder sq.SelectBuilder, scanFn pgx.RowToFunc[T]) (T, error) {
	var zero T
	sql, args, err := builder.ToSql()
	if err != nil {
		return zero, err
	}
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return zero, err
	}
	defer rows.Close()

	collectedRow, err := pgx.CollectOneRow(rows, scanFn)
	if err != nil {
		return zero, err
	}

	return collectedRow, nil
}

func InsertReturning[T any](ctx context.Context, db *db.DB, builder sq.InsertBuilder, scanFn pgx.RowToFunc[T]) (T, error) {
	var zero T
	sql, args, err := builder.ToSql()
	if err != nil {
		return zero, err
	}
	collectedRow, err := execReturn(ctx, db, sql, args, scanFn)
	if err != nil {
		return zero, err
	}
	return collectedRow, nil

}

func SelectRows[T any](ctx context.Context, db *db.DB, builder sq.SelectBuilder, scanFn pgx.RowToFunc[T]) ([]T, error) {

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	collectedRows, err := pgx.CollectRows(rows, scanFn)
	if err != nil {
		return nil, err
	}

	return collectedRows, nil
}

func SelectRowsOK[T any](ctx context.Context, db *db.DB, builder sq.SelectBuilder, scanFn pgx.RowToFunc[T]) ([]T, bool, error) {
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, false, err
	}
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	defer rows.Close()
	collectedRows, b, err := CollectRowsOK(rows, scanFn)

	if err != nil {
		return nil, false, err
	}

	return collectedRows, b, nil
}

func SelectRowsTag[T any](ctx context.Context, db *db.DB, builder sq.SelectBuilder, tag string) ([]T, error) {

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	collectedRows, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (T, error) {
		return RowToStructByTag[T](row, tag)
	})
	if err != nil {
		return nil, err
	}

	return collectedRows, nil
}

func RowToStructByTag[T any](row pgx.CollectableRow, tag string) (T, error) {

	var value T
	err := row.Scan(&tagStructRowScanner{ptrToStruct: &value, lax: true, tag: tag})
	return value, err
}

type tagStructRowScanner struct {
	ptrToStruct any
	lax         bool
	tag         string
}

func (ts *tagStructRowScanner) ScanRow(rows pgx.Rows) error {

	dst := ts.ptrToStruct
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		return fmt.Errorf("dst not a pointer")
	}

	dstElemValue := dstValue.Elem()

	scanTargets, err := ts.appendScanTargets(dstElemValue, nil, rows.FieldDescriptions(), ts.tag)

	if err != nil {

		return err
	}

	for i, t := range scanTargets {

		if t == nil {
			return fmt.Errorf("struct doesn't have corresponding field to match returned column %s", rows.FieldDescriptions()[i].Name)
		}
	}

	return rows.Scan(scanTargets...)
}

func (rs *tagStructRowScanner) appendScanTargets(dstElemValue reflect.Value, scanTargets []any, fldDescs []pgconn.FieldDescription, tagkey string) ([]any, error) {
	var err error
	dstElemType := dstElemValue.Type()

	if scanTargets == nil {
		scanTargets = make([]any, len(fldDescs))
	}

	for i := 0; i < dstElemType.NumField(); i++ {
		sf := dstElemType.Field(i)

		if sf.PkgPath != "" && !sf.Anonymous {

			// Field is unexported, skip it.
			continue
		}

		if sf.Anonymous && sf.Type.Kind() == reflect.Struct {

			scanTargets, err = rs.appendScanTargets(dstElemValue.Field(i), scanTargets, fldDescs, tagkey)
			if err != nil {
				return nil, err
			}
		} else {

			dbTag, dbTagPresent := sf.Tag.Lookup(tagkey)
			if dbTagPresent {

				dbTag = strings.Split(dbTag, ",")[0]
			}
			if dbTag == "-" {

				// Field is ignored, skip it.
				continue
			}

			colName := dbTag

			if !dbTagPresent {

				colName = sf.Name
			}

			fpos := fieldPosByName(fldDescs, colName)

			if fpos == -1 {
				if rs.lax {

					continue
				}
				return nil, fmt.Errorf("cannot find field %s in returned row", colName)
			}
			if fpos >= len(scanTargets) && !rs.lax {
				return nil, fmt.Errorf("cannot find field %s in returned row", colName)
			}

			scanTargets[fpos] = dstElemValue.Field(i).Addr().Interface()
		}
	}

	return scanTargets, err
}

func fieldPosByName(fldDescs []pgconn.FieldDescription, field string) (i int) {
	i = -1
	for i, desc := range fldDescs {
		if strings.EqualFold(desc.Name, field) {
			return i
		}
	}
	return
}

func StructToSetMap(article interface{}) map[string]interface{} {

	setMap := make(map[string]interface{})

	val := reflect.ValueOf(article).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get("json")

		// Skip fields without the "db" tag
		if tag == "" {
			continue
		}

		// Check if the value is the zero value for its type
		switch val.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val.Field(i).Int() == 0 {
				continue
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val.Field(i).Uint() == 0 {
				continue
			}
		case reflect.Float32, reflect.Float64:
			if val.Field(i).Float() == 0 {
				continue
			}
		case reflect.String:
			if val.Field(i).String() == "" {
				continue
			}
		case reflect.Bool:
			if !val.Field(i).Bool() {
				continue
			}

		case reflect.Struct:
			if val.Field(i).Type() == reflect.TypeOf(time.Time{}) && val.Field(i).Interface().(time.Time).IsZero() {
				continue
			}

		default:
			// Handle other types as needed
		}

		setMap[tag] = val.Field(i).Interface()
	}

	return setMap
}
func QueueExecRow(batch *pgx.Batch, builder sq.Sqlizer) error {
	var qErr error

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	batch.Queue(sql, args...)
	batch.Queue(sql, args...).Exec(func(ct pgconn.CommandTag) error {
		rowsAffected := ct.RowsAffected()
		if rowsAffected == 0 {
			qErr = pgx.ErrNoRows
			return nil
		}
		return nil
	})

	return qErr
}

func QueueReturn[T any](batch *pgx.Batch, builder sq.Sqlizer, scanFn pgx.RowToFunc[T], result *[]T) error {

	var qErr error

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	batch.Queue(sql, args...).Query(func(rows pgx.Rows) error {
		collectedRows, err := pgx.CollectRows(rows, scanFn)
		if err != nil {
			qErr = err
			return nil
		}
		*result = collectedRows
		return nil
	})

	return qErr
}

func QueueReturnRow[T any](batch *pgx.Batch, builder sq.Sqlizer, scanFn pgx.RowToFunc[T], result *T) error {
	var qErr error

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	batch.Queue(sql, args...).Query(func(rows pgx.Rows) error {
		collectedRow, err := pgx.CollectOneRow(rows, scanFn)
		if err != nil {
			qErr = err
			return nil
		}

		*result = collectedRow
		return nil
	})

	return qErr
}

func TxReturnRow[T any](ctx context.Context, tx pgx.Tx, builder sq.Sqlizer, scanFn pgx.RowToFunc[T], result *T) error {

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	collectedRow, err := pgx.CollectOneRow(rows, scanFn)
	if err != nil {
		return err
	}
	*result = collectedRow
	return nil
}

func TxRows[T any](ctx context.Context, tx pgx.Tx, builder sq.Sqlizer, scanFn pgx.RowToFunc[T], result *[]T) error {

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	collectedRows, err := pgx.CollectRows(rows, scanFn)
	if err != nil {
		return err
	}

	*result = collectedRows
	return nil
}

func TxExec(ctx context.Context, tx pgx.Tx, builder sq.Sqlizer) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func GenerateMapFromStruct(instance interface{}, tag string) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.Indirect(reflect.ValueOf(instance))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tag)
		if tag != "" {
			result[tag] = val.Field(i).Interface()
		}
	}
	return result
}

func GenerateColumnsFromStruct(instance interface{}, tag string) []string {
	var columns []string

	val := reflect.Indirect(reflect.ValueOf(instance))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tag)
		if tag != "" {
			columns = append(columns, tag)
		}
	}

	return columns
}

func CollectRowsOK[T any](rows pgx.Rows, fn pgx.RowToFunc[T]) ([]T, bool, error) {
	var value []T
	var err error
	value, err = pgx.CollectRows(rows, fn)
	if err != nil {
		if err == pgx.ErrNoRows {
			return value, false, nil
		}
		return value, false, err
	}
	return value, true, nil
}
func CollectOneRowOK[T any](rows pgx.Rows, fn pgx.RowToFunc[T]) (T, bool, error) {
	var value T
	var err error
	value, err = pgx.CollectOneRow(rows, fn)
	if err != nil {
		if err == pgx.ErrNoRows {
			return value, false, nil
		}
		return value, false, err
	}
	return value, true, nil
}

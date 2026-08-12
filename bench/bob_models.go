package bench

import (
	"context"
	"io"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/bob/expr"
	"github.com/stephenafamo/scan"
)

// These types mirror the model, setter, and reflection-free scanner generated
// by Bob for the models table. Keeping them here makes the benchmark build
// reproducible without requiring a running database during code generation.
type bobModel struct {
	ID      uint32 `db:"id,pk"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int32  `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

type bobModelSlice []*bobModel

type bobModelSetter struct {
	ID      omit.Val[uint32] `db:"id,pk"`
	Name    omit.Val[string] `db:"name"`
	Title   omit.Val[string] `db:"title"`
	Fax     omit.Val[string] `db:"fax"`
	Web     omit.Val[string] `db:"web"`
	Age     omit.Val[int32]  `db:"age"`
	Right   omit.Val[bool]   `db:"right"`
	Counter omit.Val[int64]  `db:"counter"`
}

var bobModels = psql.NewTablex[*bobModel, bobModelSlice, *bobModelSetter](
	"",
	"models",
	expr.ColsForStruct[bobModel]("models"),
	bobModelScanMapper,
)

func newBobModelSetter(m *Model) *bobModelSetter {
	return &bobModelSetter{
		Name:    omit.From(m.Name),
		Title:   omit.From(m.Title),
		Fax:     omit.From(m.Fax),
		Web:     omit.From(m.Web),
		Age:     omit.From(int32(m.Age)),
		Right:   omit.From(m.Right),
		Counter: omit.From(m.Counter),
	}
}

func (s bobModelSetter) SetColumns() []string {
	columns := make([]string, 0, 8)
	if s.ID.IsValue() {
		columns = append(columns, "id")
	}
	if s.Name.IsValue() {
		columns = append(columns, "name")
	}
	if s.Title.IsValue() {
		columns = append(columns, "title")
	}
	if s.Fax.IsValue() {
		columns = append(columns, "fax")
	}
	if s.Web.IsValue() {
		columns = append(columns, "web")
	}
	if s.Age.IsValue() {
		columns = append(columns, "age")
	}
	if s.Right.IsValue() {
		columns = append(columns, "right")
	}
	if s.Counter.IsValue() {
		columns = append(columns, "counter")
	}

	return columns
}

func (s *bobModelSetter) Apply(query *dialect.InsertQuery) {
	query.AppendHooks(func(ctx context.Context, exec bob.Executor) (context.Context, error) {
		return bobModels.BeforeInsertHooks.RunHooks(ctx, exec, s)
	})

	query.AppendValues(
		bobInsertValue(&s.ID),
		bobInsertValue(&s.Name),
		bobInsertValue(&s.Title),
		bobInsertValue(&s.Fax),
		bobInsertValue(&s.Web),
		bobInsertValue(&s.Age),
		bobInsertValue(&s.Right),
		bobInsertValue(&s.Counter),
	)
}

func bobInsertValue[T any](value *omit.Val[T]) bob.Expression {
	return bob.ExpressionFunc(func(ctx context.Context, w io.StringWriter, d bob.Dialect, start int) ([]any, error) {
		if value.IsUnset() {
			return psql.Raw("DEFAULT").WriteSQL(ctx, w, d, start)
		}

		return psql.Arg(value.MustGet()).WriteSQL(ctx, w, d, start)
	})
}

func (s bobModelSetter) UpdateMod() bob.Mod[*dialect.UpdateQuery] {
	return um.Set(s.Expressions()...)
}

func (s bobModelSetter) Expressions(prefix ...string) []bob.Expression {
	expressions := make([]bob.Expression, 0, 8)
	if s.ID.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "id", s.ID))
	}
	if s.Name.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "name", s.Name))
	}
	if s.Title.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "title", s.Title))
	}
	if s.Fax.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "fax", s.Fax))
	}
	if s.Web.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "web", s.Web))
	}
	if s.Age.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "age", s.Age))
	}
	if s.Right.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "right", s.Right))
	}
	if s.Counter.IsValue() {
		expressions = append(expressions, bobUpdateExpression(prefix, "counter", s.Counter))
	}

	return expressions
}

func bobUpdateExpression[T any](prefix []string, column string, value omit.Val[T]) bob.Expression {
	return expr.Join{
		Sep: " = ",
		Exprs: []bob.Expression{
			psql.Quote(append(prefix, column)...),
			psql.Arg(value),
		},
	}
}

func bobModelScanMapper(_ context.Context, columns []string) (scan.BeforeFunc, func(any) (*bobModel, error)) {
	type target struct {
		index       int
		destination func(*bobModel) any
	}

	targets := make([]target, 0, 8)
	for index, column := range columns {
		switch column {
		case "id":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.ID }})
		case "name":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Name }})
		case "title":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Title }})
		case "fax":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Fax }})
		case "web":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Web }})
		case "age":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Age }})
		case "right":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Right }})
		case "counter":
			targets = append(targets, target{index, func(m *bobModel) any { return &m.Counter }})
		}
	}

	return func(row *scan.Row) (any, error) {
			model := new(bobModel)
			for _, target := range targets {
				row.ScheduleScanByIndex(target.index, target.destination(model))
			}

			return model, nil
		}, func(value any) (*bobModel, error) {
			return value.(*bobModel), nil
		}
}

func (m *bobModel) AfterQueryHook(ctx context.Context, exec bob.Executor, queryType bob.QueryType) error {
	var err error
	switch queryType {
	case bob.QueryTypeSelect:
		_, err = bobModels.AfterSelectHooks.RunHooks(ctx, exec, bobModelSlice{m})
	case bob.QueryTypeInsert:
		_, err = bobModels.AfterInsertHooks.RunHooks(ctx, exec, bobModelSlice{m})
	case bob.QueryTypeUpdate:
		_, err = bobModels.AfterUpdateHooks.RunHooks(ctx, exec, bobModelSlice{m})
	case bob.QueryTypeDelete:
		_, err = bobModels.AfterDeleteHooks.RunHooks(ctx, exec, bobModelSlice{m})
	case bob.QueryTypeMerge:
		_, err = bobModels.AfterMergeHooks.RunHooks(ctx, exec, bobModelSlice{m})
	}

	return err
}

func (models bobModelSlice) AfterQueryHook(ctx context.Context, exec bob.Executor, queryType bob.QueryType) error {
	var err error
	switch queryType {
	case bob.QueryTypeSelect:
		_, err = bobModels.AfterSelectHooks.RunHooks(ctx, exec, models)
	case bob.QueryTypeInsert:
		_, err = bobModels.AfterInsertHooks.RunHooks(ctx, exec, models)
	case bob.QueryTypeUpdate:
		_, err = bobModels.AfterUpdateHooks.RunHooks(ctx, exec, models)
	case bob.QueryTypeDelete:
		_, err = bobModels.AfterDeleteHooks.RunHooks(ctx, exec, models)
	case bob.QueryTypeMerge:
		_, err = bobModels.AfterMergeHooks.RunHooks(ctx, exec, models)
	}

	return err
}

func (m *bobModel) primaryKeyValues() bob.Expression {
	return psql.Arg(m.ID)
}

func (m *bobModel) primaryKeyEquals() dialect.Expression {
	return psql.Quote("models", "id").EQ(
		bob.ExpressionFunc(func(ctx context.Context, w io.StringWriter, d bob.Dialect, start int) ([]any, error) {
			return m.primaryKeyValues().WriteSQL(ctx, w, d, start)
		}),
	)
}

func (m *bobModel) Update(ctx context.Context, exec bob.Executor, setter *bobModelSetter) error {
	updated, err := bobModels.Update(
		setter.UpdateMod(),
		um.Where(m.primaryKeyEquals()),
	).One(ctx, exec)
	if err != nil {
		return err
	}

	*m = *updated
	return nil
}

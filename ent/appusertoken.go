// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/boshangad/v1/ent/app"
	"github.com/boshangad/v1/ent/appuser"
	"github.com/boshangad/v1/ent/appusertoken"
	"github.com/boshangad/v1/ent/user"
	"github.com/google/uuid"
)

// AppUserToken is the model entity for the AppUserToken schema.
type AppUserToken struct {
	config `json:"-"`
	// ID of the ent.
	// ID
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// AppID holds the value of the "app_id" field.
	// 应用
	AppID uint64 `json:"app_id,omitempty"`
	// AppUserID holds the value of the "app_user_id" field.
	// 应用用户
	AppUserID uint64 `json:"app_user_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	// 用户
	UserID uint64 `json:"user_id,omitempty"`
	// ClientVersion holds the value of the "client_version" field.
	// 客户端版本
	ClientVersion string `json:"client_version,omitempty"`
	// UUID holds the value of the "uuid" field.
	// 设备唯一标识
	UUID uuid.UUID `json:"uuid,omitempty"`
	// IP holds the value of the "ip" field.
	// IP地址
	IP string `json:"ip,omitempty"`
	// ExpireTime holds the value of the "expire_time" field.
	// 失效时间
	ExpireTime uint64 `json:"expire_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppUserTokenQuery when eager-loading is set.
	Edges AppUserTokenEdges `json:"edges"`
}

// AppUserTokenEdges holds the relations/edges for other nodes in the graph.
type AppUserTokenEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// AppUser holds the value of the appUser edge.
	AppUser *AppUser `json:"appUser,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppUserTokenEdges) AppOrErr() (*App, error) {
	if e.loadedTypes[0] {
		if e.App == nil {
			// The edge app was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.App, nil
	}
	return nil, &NotLoadedError{edge: "app"}
}

// AppUserOrErr returns the AppUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppUserTokenEdges) AppUserOrErr() (*AppUser, error) {
	if e.loadedTypes[1] {
		if e.AppUser == nil {
			// The edge appUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: appuser.Label}
		}
		return e.AppUser, nil
	}
	return nil, &NotLoadedError{edge: "appUser"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppUserTokenEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUserToken) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appusertoken.FieldID, appusertoken.FieldCreateTime, appusertoken.FieldAppID, appusertoken.FieldAppUserID, appusertoken.FieldUserID, appusertoken.FieldExpireTime:
			values[i] = new(sql.NullInt64)
		case appusertoken.FieldClientVersion, appusertoken.FieldIP:
			values[i] = new(sql.NullString)
		case appusertoken.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUserToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUserToken fields.
func (aut *AppUserToken) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appusertoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aut.ID = uint64(value.Int64)
		case appusertoken.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				aut.CreateTime = value.Int64
			}
		case appusertoken.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				aut.AppID = uint64(value.Int64)
			}
		case appusertoken.FieldAppUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_user_id", values[i])
			} else if value.Valid {
				aut.AppUserID = uint64(value.Int64)
			}
		case appusertoken.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				aut.UserID = uint64(value.Int64)
			}
		case appusertoken.FieldClientVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_version", values[i])
			} else if value.Valid {
				aut.ClientVersion = value.String
			}
		case appusertoken.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				aut.UUID = *value
			}
		case appusertoken.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				aut.IP = value.String
			}
		case appusertoken.FieldExpireTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expire_time", values[i])
			} else if value.Valid {
				aut.ExpireTime = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryApp queries the "app" edge of the AppUserToken entity.
func (aut *AppUserToken) QueryApp() *AppQuery {
	return (&AppUserTokenClient{config: aut.config}).QueryApp(aut)
}

// QueryAppUser queries the "appUser" edge of the AppUserToken entity.
func (aut *AppUserToken) QueryAppUser() *AppUserQuery {
	return (&AppUserTokenClient{config: aut.config}).QueryAppUser(aut)
}

// QueryUser queries the "user" edge of the AppUserToken entity.
func (aut *AppUserToken) QueryUser() *UserQuery {
	return (&AppUserTokenClient{config: aut.config}).QueryUser(aut)
}

// Update returns a builder for updating this AppUserToken.
// Note that you need to call AppUserToken.Unwrap() before calling this method if this AppUserToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (aut *AppUserToken) Update() *AppUserTokenUpdateOne {
	return (&AppUserTokenClient{config: aut.config}).UpdateOne(aut)
}

// Unwrap unwraps the AppUserToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aut *AppUserToken) Unwrap() *AppUserToken {
	tx, ok := aut.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUserToken is not a transactional entity")
	}
	aut.config.driver = tx.drv
	return aut
}

// String implements the fmt.Stringer.
func (aut *AppUserToken) String() string {
	var builder strings.Builder
	builder.WriteString("AppUserToken(")
	builder.WriteString(fmt.Sprintf("id=%v", aut.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", aut.CreateTime))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", aut.AppID))
	builder.WriteString(", app_user_id=")
	builder.WriteString(fmt.Sprintf("%v", aut.AppUserID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", aut.UserID))
	builder.WriteString(", client_version=")
	builder.WriteString(aut.ClientVersion)
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", aut.UUID))
	builder.WriteString(", ip=")
	builder.WriteString(aut.IP)
	builder.WriteString(", expire_time=")
	builder.WriteString(fmt.Sprintf("%v", aut.ExpireTime))
	builder.WriteByte(')')
	return builder.String()
}

// AppUserTokens is a parsable slice of AppUserToken.
type AppUserTokens []*AppUserToken

func (aut AppUserTokens) config(cfg config) {
	for _i := range aut {
		aut[_i].config = cfg
	}
}

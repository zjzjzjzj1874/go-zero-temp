package svc

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github/zjzjzjzj1874/go-zero-temp/access/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	DB     *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		panic(fmt.Sprintf("connect mysql failed, error: %v", err))
	}

	svcCtx := &ServiceContext{
		Config: c,
		DB:     db,
	}

	// 初始化权限数据
	initPermissions(svcCtx)

	return svcCtx
}

// initPermissions 初始化权限数据和管理员角色
func initPermissions(svcCtx *ServiceContext) {
	// 1. 初始化权限数据
	var routes []rest.Route
	routes = append(routes,
		rest.Route{Method: http.MethodGet, Path: "/access/liveness", Handler: nil},
		rest.Route{Method: http.MethodGet, Path: "/access/swagger", Handler: nil},
		rest.Route{Method: http.MethodPost, Path: "/access/roles", Handler: nil},
		rest.Route{Method: http.MethodPut, Path: "/access/roles/:id", Handler: nil},
		rest.Route{Method: http.MethodDelete, Path: "/access/roles/:id", Handler: nil},
		rest.Route{Method: http.MethodGet, Path: "/access/roles", Handler: nil},
		rest.Route{Method: http.MethodPost, Path: "/access/roles/:roleId/permissions", Handler: nil},
		rest.Route{Method: http.MethodGet, Path: "/access/roles/:roleId/permissions", Handler: nil},
		rest.Route{Method: http.MethodPost, Path: "/access/permissions", Handler: nil},
		rest.Route{Method: http.MethodPut, Path: "/access/permissions/:id", Handler: nil},
		rest.Route{Method: http.MethodDelete, Path: "/access/permissions/:id", Handler: nil},
		rest.Route{Method: http.MethodGet, Path: "/access/permissions", Handler: nil},
	)

	// 2. 遍历路由，将权限数据写入数据库
	for _, route := range routes {
		// 使用upsert语句，如果存在则更新，不存在则插入
		query := `INSERT INTO permission (name, type, description, resource, action, status) 
			VALUES (?, 'api', ?, ?, ?, 1) 
			ON DUPLICATE KEY UPDATE 
			name = VALUES(name), 
			description = VALUES(description), 
			status = VALUES(status)`

		// 生成权限名称和描述
		name := fmt.Sprintf("%s %s", route.Method, route.Path)
		description := fmt.Sprintf("Access to %s %s", route.Method, route.Path)

		// 执行upsert
		_, err := svcCtx.DB.Exec(query, name, description, route.Path, strings.ToUpper(route.Method))
		if err != nil {
			fmt.Printf("init permission failed, error: %v\n", err)
		}
	}

	// 3. 初始化管理员角色
	roleQuery := `INSERT INTO role (name, description, status) 
		VALUES ('admin', 'System Administrator', 1) 
		ON DUPLICATE KEY UPDATE 
		description = VALUES(description), 
		status = VALUES(status)`

	result, err := svcCtx.DB.Exec(roleQuery)
	if err != nil {
		fmt.Printf("init admin role failed, error: %v\n", err)
		return
	}

	// 4. 为管理员角色分配所有权限
	roleID, err := result.LastInsertId()
	if err != nil {
		// 如果是更新操作，需要查询admin角色的ID
		var id int64
		err = svcCtx.DB.QueryRow("SELECT id FROM role WHERE name = 'admin'").Scan(&id)
		if err != nil {
			fmt.Printf("get admin role id failed, error: %v\n", err)
			return
		}
		roleID = id
	}

	// 5. 分配所有权限给管理员角色
	_, err = svcCtx.DB.Exec(`INSERT INTO role_permission (role_id, permission_id) 
		SELECT ?, id FROM permission 
		ON DUPLICATE KEY UPDATE role_id = VALUES(role_id)`, roleID)
	if err != nil {
		fmt.Printf("assign permissions to admin role failed, error: %v\n", err)
	}
}

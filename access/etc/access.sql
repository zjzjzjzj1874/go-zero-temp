-- 角色表
CREATE TABLE IF NOT EXISTS `role` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name` varchar(50) NOT NULL COMMENT '角色名称',
    `description` varchar(200) DEFAULT NULL COMMENT '角色描述',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0-禁用 1-启用',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 权限表
CREATE TABLE IF NOT EXISTS `permission` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
    `name` varchar(50) NOT NULL COMMENT '权限名称',
    `type` varchar(20) NOT NULL COMMENT '权限类型：api-接口权限',
    `description` varchar(200) DEFAULT NULL COMMENT '权限描述',
    `resource` varchar(200) NOT NULL COMMENT '资源标识（如API路径）',
    `action` varchar(20) NOT NULL COMMENT '操作类型（如GET/POST等）',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0-禁用 1-启用',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_resource_action` (`resource`, `action`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

-- 角色权限关联表
CREATE TABLE IF NOT EXISTS `role_permission` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `role_id` bigint(20) NOT NULL COMMENT '角色ID',
    `permission_id` bigint(20) NOT NULL COMMENT '权限ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_role_permission` (`role_id`, `permission_id`),
    KEY `idx_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关联表';

-- 用户角色关联表
CREATE TABLE IF NOT EXISTS `user_role` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `role_id` bigint(20) NOT NULL COMMENT '角色ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_role` (`user_id`, `role_id`),
    KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';
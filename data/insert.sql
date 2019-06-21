BEGIN;
INSERT INTO `users` VALUES (1, 'admin', 'admin', '93a57b286d7f77fdce1c8e17f5c2dfb6459c739b058c85b168cdd1df599e1f35', '1118447772383584256', 'admin@mshk.top', 0, 'adminnote', 0, '2019-04-12 13:52:57', '2019-04-12 14:40:38');
COMMIT;


-- ----------------------------
-- Records of columns
-- ----------------------------
BEGIN;
INSERT INTO `columns` VALUES (1, '系统设置', '', 0, 100, 1, 'fa fa-gear', '2019-04-22 00:07:47', '2019-06-22 00:30:00');
INSERT INTO `columns` VALUES (2, '栏目管理', '/v1/admin/columns', 1, 20, 1, '', '2019-04-22 00:08:35', '2019-04-22 00:08:35');
INSERT INTO `columns` VALUES (3, '用户相关设置', '', 0, 90, 1, 'fa fa-user', '2019-04-22 00:09:04', '2019-04-22 00:09:04');
INSERT INTO `columns` VALUES (4, '用户管理', '/v1/admin/users', 3, 0, 1, 'fa fa-child', '2019-04-22 00:09:59', '2019-04-22 00:09:59');
INSERT INTO `columns` VALUES (5, '只能查看自己和自己创建的用户列表', '', 4, 0, 0, '', '2019-04-22 00:11:37', '2019-04-22 00:11:37');
INSERT INTO `columns` VALUES (6, '显示查看用户所属用户组按钮', '', 4, 0, 0, '', '2019-04-22 00:12:10', '2019-04-22 00:12:10');
INSERT INTO `columns` VALUES (7, '查看所有用户列表', '', 4, 0, 0, '', '2019-04-22 00:12:27', '2019-04-22 00:12:27');
INSERT INTO `columns` VALUES (8, '用户组管理', '/v1/admin/usersgroup', 3, 0, 1, 'fa fa-group', '2019-04-22 00:13:33', '2019-04-22 00:13:33');
INSERT INTO `columns` VALUES (9, '显示配置权限按钮', '', 8, 0, 0, '', '2019-04-22 00:14:02', '2019-04-22 00:14:02');
INSERT INTO `columns` VALUES (10, '显示删除和批量删除用户组列表', '', 8, 0, 0, '', '2019-04-22 00:14:38', '2019-04-22 00:14:38');
INSERT INTO `columns` VALUES (11, '查看所有用户组列表', '', 8, 0, 0, '', '2019-04-22 00:15:34', '2019-04-22 00:15:34');
INSERT INTO `columns` VALUES (12, '登录历史记录', '/v1/admin/historyuserlogin', 3, 0, 1, '', '2019-04-22 00:16:05', '2019-04-22 00:16:05');
INSERT INTO `columns` VALUES (13, '任务计划', '', 0, 0, 1, 'fa fa-calendar-plus-o', '2019-06-22 00:30:59', '2019-06-22 00:30:59');
INSERT INTO `columns` VALUES (14, '订单管理', '', 0, 0, 1, 'fa fa-th-list', '2019-06-22 00:31:17', '2019-06-22 00:31:17');
INSERT INTO `columns` VALUES (15, '委托管理', '', 0, 0, 1, 'fa fa-flask', '2019-06-22 00:31:34', '2019-06-22 00:31:34');
INSERT INTO `columns` VALUES (16, '分析报表', '', 0, 0, 1, 'fa fa fa-bar-chart-o', '2019-06-22 00:32:00', '2019-06-22 00:35:14');
INSERT INTO `columns` VALUES (17, '用户/交易所/交易对', '', 1, 0, 1, 'fa fa-object-ungroup', '2019-06-22 00:33:03', '2019-06-22 00:33:03');
INSERT INTO `columns` VALUES (18, '交易所管理', '', 1, 0, 1, 'fa fa-university', '2019-06-22 00:33:41', '2019-06-22 00:33:41');
INSERT INTO `columns` VALUES (19, '交易对管理', '', 1, 0, 1, 'nav-label', '2019-06-22 00:34:04', '2019-06-22 00:34:04');
COMMIT;

BEGIN;
INSERT INTO `users_group` VALUES (1, '普通用户组', 0, 0, '', '2019-06-16 02:54:37', '2019-06-16 02:54:37');
INSERT INTO `users_group` VALUES (2, '管理员', 0, 0, '', '2019-06-16 02:54:44', '2019-06-16 02:54:44');
INSERT INTO `users_group` VALUES (3, '超级管理员', 0, 0, '', '2019-06-16 02:54:54', '2019-06-16 02:54:54');
COMMIT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (1, 'g', '1', 'usergroup_3', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', 'usergroup_3', '1', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', 'usergroup_3', '2', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', 'usergroup_3', '17', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', 'usergroup_3', '18', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', 'usergroup_3', '19', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', 'usergroup_3', '3', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'p', 'usergroup_3', '4', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', 'usergroup_3', '5', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'usergroup_3', '6', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'usergroup_3', '7', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', 'usergroup_3', '8', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', 'usergroup_3', '9', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'usergroup_3', '10', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', 'usergroup_3', '11', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', 'usergroup_3', '12', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'p', 'usergroup_3', '13', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', 'usergroup_3', '14', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', 'usergroup_3', '15', '', 'GET', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', 'usergroup_3', '16', '', 'GET', '', '');
COMMIT;
BEGIN;
INSERT INTO `users` VALUES (1, 'admin', 'admin', '93a57b286d7f77fdce1c8e17f5c2dfb6459c739b058c85b168cdd1df599e1f35', '1118447772383584256', 'admin@mshk.top', 0, 'adminnote', 0, '2019-04-12 13:52:57', '2019-04-12 14:40:38');
COMMIT;


BEGIN;
INSERT INTO `columns` VALUES (1, '系统设置', '', 0, 100, 1, '', '2019-04-22 00:07:47', '2019-04-22 00:10:11');
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
COMMIT;

BEGIN;
INSERT INTO `users_group` VALUES (1, '普通用户组', 0, 0, '', '2019-06-16 02:54:37', '2019-06-16 02:54:37');
INSERT INTO `users_group` VALUES (2, '管理员', 0, 0, '', '2019-06-16 02:54:44', '2019-06-16 02:54:44');
INSERT INTO `users_group` VALUES (3, '超级管理员', 0, 0, '', '2019-06-16 02:54:54', '2019-06-16 02:54:54');
COMMIT;
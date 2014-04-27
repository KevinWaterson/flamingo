CREATE TABLE users(
	    id SERIAL PRIMARY KEY NOT NULL,
	    email VARCHAR(20) NOT NULL,
	    password VARCHAR(50) NOT NULL
);
INSERT INTO users(email, password) VALUES ('admin@golang.pro', 'ce7d2196c0756cd0fc46217e69f10b6a');
INSERT INTO users(email, password) VALUES ('moderator@golang.pro', 'ce7d2196c0756cd0fc46217e69f10b6a');
INSERT INTO users(email, password) VALUES ('user@golang.pro', 'ce7d2196c0756cd0fc46217e69f10b6a');
INSERT INTO users(email, password) VALUES ('picker@golang.pro', 'ce7d2196c0756cd0fc46217e69f10b6a');

CREATE TABLE user_group(
	id SERIAL  PRIMARY KEY NOT NULL,
        name VARCHAR(50) NOT NULL,
        description VARCHAR(100) NOT NULL
);
INSERT INTO user_group( id, name, description ) VALUES ( 1, 'admin', 'For admins only' );
INSERT INTO user_group( id, name, description ) VALUES ( 2, 'user', 'For general users' );
INSERT INTO user_group( id, name, description ) VALUES ( 3, 'moderator', 'For site moderators' );
INSERT INTO user_group( id, name, description ) VALUES ( 4, 'picker', 'Warehouse Picker' );

CREATE TABLE user_group_link(
	id SERIAL PRIMARY KEY NOT NULL,
        user_id INTEGER NOT NULL,
        group_id INTEGER NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(group_id) REFERENCES user_group(id)
);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 1, 1, 1 );
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 2, 1, 2 );
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 3, 1, 3 );
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 4, 1, 4);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 5, 2, 3);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 6, 3, 4);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 7, 4, 1);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 8, 4, 2);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 9, 4, 3);
INSERT INTO user_group_link(id, user_id, group_id) VALUES ( 10,4, 4);


CREATE TABLE group_permission(
	id SERIAL PRIMARY KEY NOT NULL,
        name    TEXT NOT NULL,
        value   TEXT NOT NULL
);
INSERT INTO group_permission(id, name, value) VALUES (1, 'Create', 'Create');
INSERT INTO group_permission(id, name, value) VALUES (2, 'Read', 'Read');
INSERT INTO group_permission(id, name, value) VALUES (3, 'Update', 'Update');
INSERT INTO group_permission(id, name, value) VALUES (4, 'Delete', 'Delete');

CREATE TABLE acl_object(
	id SERIAL PRIMARY KEY NOT NULL,
        filename TEXT NOT NULL,
        function_name TEXT NOT NULL
);
INSERT INTO acl_object(id, filename, function_name) VALUES (1, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (2, 'main.go', '/reports/sales/missingOrders');
INSERT INTO acl_object(id, filename, function_name) VALUES (3, 'main.go', '/admin/user/groups');
INSERT INTO acl_object(id, filename, function_name) VALUES (4, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (5, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (6, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (7, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (8, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (9, 'main.go', '/sales/reports/monthly');
INSERT INTO acl_object(id, filename, function_name) VALUES (10, 'main.go', '/sales/reports/monthly');

-- this is a registry of permissions a group has on objects
-- if a user is in a group that is part of that group, all is well
CREATE TABLE group_permission_object_link(
	id SERIAL PRIMARY KEY NOT NULL,
        group_id INTEGER NOT NULL,
        permission_id INTEGER NOT NULL,
        object_id INTEGER NOT NULL,
        FOREIGN KEY(group_id) REFERENCES user_group(id),
        FOREIGN KEY(permission_id) REFERENCES group_permission(id),
        FOREIGN KEY(object_id) REFERENCES acl_object(id)
);
-- a group must have READ permission on 
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 1, 1, 1, 1); -- group id 1 (admin) has permission_id 1 (CREATE) privs on object_id 1
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 2, 1, 2, 1); -- group id 1 (admin) has permission_id 2 (READ) privs on object_id 1
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 3, 1, 3, 1); -- group id 1 (admin) has permission_id 3 (UPDATE) privs on object_id 1
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 4, 1, 4, 1); -- group id 1 (admin) has permission_id 4 (Delete) privs on object_id 1
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 5, 3, 2, 1); -- group id 3 (acounts) has permission_id 2 (READ) privs on object_id 1

INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 6, 1, 1, 2); -- group id 1 (admin) has permission_id 1 (CREATE) privs on object_id 2 
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 7, 1, 2, 2); -- group id 1 (admin) has permission_id 2 (READ) privs on object_id 2
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 8, 1, 3, 2); -- group id 1 (admin) has permission_id 3 (UPDATE) privs on object_id 2
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 9, 1, 4, 2); -- group id 1 (admin) has permission_id 4 (Delete) privs on object_id 2

INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 10, 1, 1, 3); -- group id 1 (admin) has permission_id 1 (CREATE) privs on object_id 3
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 11, 1, 2, 3); -- group id 1 (admin) has permission_id 2 (READ) privs on object_id 3
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 12, 1, 3, 3); -- group id 1 (admin) has permission_id 3 (UPDATE) privs on object_id 3
INSERT INTO group_permission_object_link(id, group_id, permission_id, object_id) VALUES ( 13, 1, 4, 3); -- group id 1 (admin) has permission_id 4 (Delete) privs on object_id 3




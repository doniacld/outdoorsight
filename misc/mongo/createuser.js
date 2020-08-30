// give all permissions to developer to connect
db.createUser({
    user: 'os',
    pwd: 'os',
    roles: [
        {role: "userAdminAnyDatabase", db: "admin"},
        {role: "readWriteAnyDatabase", db: "admin"},
        {role: "dbAdminAnyDatabase", db: "admin"}, {role: "clusterAdmin", db: "admin"}]
});

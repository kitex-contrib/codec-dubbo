namespace go user

struct Request {
    1: required string Name,
}(JavaClassName="org.apache.dubbo.Request")

struct User {
    1: required string ID,
    2: required string Name,
    3: required i32 Age,
}(JavaClassName="org.apache.dubbo.User")

service UserService {
    User GetUser(1: Request req)
}(JavaClassName="org.apache.dubbo.UserProvider")
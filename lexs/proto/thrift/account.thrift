struct RegisterReq {
    1: string phone;
    2: string pwd;
    3: string type;
}

struct RegisterResp {
    1: string id;
    2: string phone;
}

struct LoginReq {
    1: string phone;
    2: string pwd;
    3: string type;
}

struct LoginResp {
    1: string id;
    2: string phone;
}

service AccountSevice {
    RegisterResp registe(1:RegisterReq req),
    LoginResp login(1:LoginReq req),
}


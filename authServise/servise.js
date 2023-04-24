const grpc = require('@grpc/grpc-js');
const PROTO_PATH = './auth.proto';
var protoLoader = require('@grpc/proto-loader');

const IP = '127.0.0.1';
const PORT = '50051';

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};

const packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const authProto = grpc.loadPackageDefinition(packageDefinition);

const server = new grpc.Server();

const logins = new Map([
    ['Алексей', 1],
    ['Артемий', 2],
    ['Владимир', 3],
]);

server.addService(authProto.Auth.service, {
    IsAuth: (call, callback) => {
        const AuthResp = { reply: logins.get(call.request.login) };
        callback(null, AuthResp);
    },
});

server.bindAsync(
    `${IP}:${PORT}`,
    grpc.ServerCredentials.createInsecure(),
    (error, port) => {
        console.log(`Server running at http:${IP}:${PORT}`);
        server.start();
    }
);

const PROTO_PATH = './grpchw.proto';
const grpc = require('grpc');
const olapProto = grpc.load(PROTO_PATH);

let client = new olapProto.grpchw.Olap('127.0.0.1:50051',
                                       grpc.credentials.createInsecure());

client.getMessage({name: 'nodejs'}, (error, response) => {
  if (error)
    console.log(`Error: ${error}`);
  else
    console.log(response);
});

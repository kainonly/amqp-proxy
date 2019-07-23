import { MongoClient } from 'mongodb';

const Mongo = (uri: string): Promise<MongoClient> => new Promise((resove, reject) => {
  const client = new MongoClient(uri, { useNewUrlParser: true });
  client.connect((err: any) => err ? reject(err) : resove(client));
});

export { Mongo };
